/* Copyright (C) 2015 Atsushi Togo */
/* All rights reserved. */

/* This file is part of phonopy. */

/* Redistribution and use in source and binary forms, with or without */
/* modification, are permitted provided that the following conditions */
/* are met: */

/* * Redistributions of source code must retain the above copyright */
/*   notice, this list of conditions and the following disclaimer. */

/* * Redistributions in binary form must reproduce the above copyright */
/*   notice, this list of conditions and the following disclaimer in */
/*   the documentation and/or other materials provided with the */
/*   distribution. */

/* * Neither the name of the phonopy project nor the names of its */
/*   contributors may be used to endorse or promote products derived */
/*   from this software without specific prior written permission. */

/* THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS */
/* "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT */
/* LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS */
/* FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE */
/* COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, */
/* INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, */
/* BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; */
/* LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER */
/* CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT */
/* LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN */
/* ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE */
/* POSSIBILITY OF SUCH DAMAGE. */

#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include <phonoc_array.h>
#include <phonoc_utils.h>
#include <phonon3_h/collision_matrix.h>

static void _get_collision_matrix(double *collision_matrix,
                                  const double *fc3_normal_squared,
                                  const int num_band0,
                                  const int num_band,
                                  const double *frequencies,
                                  const int *triplets,
                                  const int *triplets_map,
                                  const int num_gp,
                                  const int *stabilized_gp_map,
                                  const int *rot_BZ_grid_points,
                                  const int num_ir_gp,
                                  const int num_rot,
                                  const double *rotations_cartesian,
                                  const double *g,
                                  const double temperature,
                                  const double unit_conversion_factor,
                                  const double cutoff_frequency);
static void
_get_reducible_collision_matrix(double *collision_matrix,
                                const double *fc3_normal_squared,
                                const int num_band0,
                                const int num_band,
                                const double *frequencies,
                                const int *triplets,
                                const int *triplets_map,
                                const int num_gp,
                                const int *stabilized_gp_map,
                                const double *g,
                                const double temperature,
                                const double unit_conversion_factor,
                                const double cutoff_frequency);
static void get_inv_sinh(double *inv_sinh,
                         const int gp,
                         const int temperature,
                         const double *frequencies,
                         const int *triplets,
                         const int *triplets_map,
                         const int *stabilized_gp_map,
                         const int ti,
                         const int num_band,
                         const double cutoff_frequency);
static int *create_gp2tp_map(const int *triplets,
                             const int num_gp);

void col_get_collision_matrix(double *collision_matrix,
                              const Darray *fc3_normal_squared,
                              const double *frequencies,
                              const int *triplets,
                              const Iarray *triplets_map,
                              const int *stabilized_gp_map,
                              const Iarray *rot_BZ_grid_points,
                              const double *rotations_cartesian,
                              const double *g,
                              const double temperature,
                              const double unit_conversion_factor,
                              const double cutoff_frequency)
{
  int num_triplets, num_ir_gp, num_rot, num_gp, num_band0, num_band;

  num_triplets = fc3_normal_squared->dims[0];
  num_band0 = fc3_normal_squared->dims[1];
  num_band = fc3_normal_squared->dims[2];
  num_ir_gp = rot_BZ_grid_points->dims[0];
  num_rot = rot_BZ_grid_points->dims[1];
  num_gp = triplets_map->dims[0];

  _get_collision_matrix(
    collision_matrix,
    fc3_normal_squared->data,
    num_band0,
    num_band,
    frequencies,
    triplets,
    triplets_map->data,
    num_gp,
    stabilized_gp_map,
    rot_BZ_grid_points->data,
    num_ir_gp,
    num_rot,
    rotations_cartesian,
    g + 2 * num_triplets * num_band0 * num_band * num_band,
    temperature,
    unit_conversion_factor,
    cutoff_frequency);
}

void col_get_reducible_collision_matrix(double *collision_matrix,
                                        const Darray *fc3_normal_squared,
                                        const double *frequencies,
                                        const int *triplets,
                                        const Iarray *triplets_map,
                                        const int *stabilized_gp_map,
                                        const double *g,
                                        const double temperature,
                                        const double unit_conversion_factor,
                                        const double cutoff_frequency)
{
  int num_triplets, num_gp, num_band, num_band0;

  num_triplets = fc3_normal_squared->dims[0];
  num_band0 = fc3_normal_squared->dims[1];
  num_band = fc3_normal_squared->dims[2];
  num_gp = triplets_map->dims[0];

  _get_reducible_collision_matrix(
    collision_matrix,
    fc3_normal_squared->data,
    num_band0,
    num_band,
    frequencies,
    triplets,
    triplets_map->data,
    num_gp,
    stabilized_gp_map,
    g + 2 * num_triplets * num_band0 * num_band * num_band,
    temperature,
    unit_conversion_factor,
    cutoff_frequency);
}

static void _get_collision_matrix(double *collision_matrix,
                                  const double *fc3_normal_squared,
                                  const int num_band0,
                                  const int num_band,
                                  const double *frequencies,
                                  const int *triplets,
                                  const int *triplets_map,
                                  const int num_gp,
                                  const int *stabilized_gp_map,
                                  const int *rot_BZ_grid_points,
                                  const int num_ir_gp,
                                  const int num_rot,
                                  const double *rotations_cartesian,
                                  const double *g,
                                  const double temperature,
                                  const double unit_conversion_factor,
                                  const double cutoff_frequency)
{
  int i, j, k, l, m, n, ti, r_gp;
  int *gp2tp_map;
  double collision;
  double *inv_sinh;

  gp2tp_map = create_gp2tp_map(triplets_map, num_gp);

#pragma omp parallel for private(j, k, l, m, n, ti, r_gp, collision, inv_sinh)
  for (i = 0; i < num_ir_gp; i++) {
    inv_sinh = (double*)malloc(sizeof(double) * num_band);
    for (j = 0; j < num_rot; j++) {
      r_gp = rot_BZ_grid_points[i * num_rot + j];
      ti = gp2tp_map[triplets_map[r_gp]];
      get_inv_sinh(inv_sinh,
                   r_gp,
                   temperature,
                   frequencies,
                   triplets,
                   triplets_map,
                   stabilized_gp_map,
                   ti,
                   num_band,
                   cutoff_frequency);

      for (k = 0; k < num_band0; k++) {
        for (l = 0; l < num_band; l++) {
          collision = 0;
          for (m = 0; m < num_band; m++) {
            collision +=
              fc3_normal_squared[ti * num_band0 * num_band * num_band +
                                 k * num_band * num_band +
                                 l * num_band + m] *
              g[ti * num_band0 * num_band * num_band +
                k * num_band * num_band +
                l * num_band + m] *
              inv_sinh[m] * unit_conversion_factor;
          }
          for (m = 0; m < 3; m++) {
            for (n = 0; n < 3; n++) {
              collision_matrix[k * 3 * num_ir_gp * num_band * 3 +
                               m * num_ir_gp * num_band * 3 +
                               i * num_band * 3 + l * 3 + n] +=
                collision * rotations_cartesian[j * 9 + m * 3 + n];
            }
          }
        }
      }
    }
    free(inv_sinh);
    inv_sinh = NULL;
  }

  free(gp2tp_map);
  gp2tp_map = NULL;
}

static void
_get_reducible_collision_matrix(double *collision_matrix,
                                const double *fc3_normal_squared,
                                const int num_band0,
                                const int num_band,
                                const double *frequencies,
                                const int *triplets,
                                const int *triplets_map,
                                const int num_gp,
                                const int *stabilized_gp_map,
                                const double *g,
                                const double temperature,
                                const double unit_conversion_factor,
                                const double cutoff_frequency)
{
  int i, j, k, l, ti;
  int *gp2tp_map;
  double collision;
  double *inv_sinh;

  gp2tp_map = create_gp2tp_map(triplets_map, num_gp);

#pragma omp parallel for private(j, k, l, ti, collision, inv_sinh)
  for (i = 0; i < num_gp; i++) {
    inv_sinh = (double*)malloc(sizeof(double) * num_band);
    ti = gp2tp_map[triplets_map[i]];
    get_inv_sinh(inv_sinh,
                 i,
                 temperature,
                 frequencies,
                 triplets,
                 triplets_map,
                 stabilized_gp_map,
                 ti,
                 num_band,
                 cutoff_frequency);

    for (j = 0; j < num_band0; j++) {
      for (k = 0; k < num_band; k++) {
        collision = 0;
        for (l = 0; l < num_band; l++) {
          collision +=
            fc3_normal_squared[ti * num_band0 * num_band * num_band +
                               j * num_band * num_band +
                               k * num_band + l] *
            g[ti * num_band0 * num_band * num_band +
              j * num_band * num_band +
              k * num_band + l] *
            inv_sinh[l] * unit_conversion_factor;
        }
        collision_matrix[j * num_gp * num_band + i * num_band + k] += collision;
      }
    }

    free(inv_sinh);
    inv_sinh = NULL;
  }

  free(gp2tp_map);
  gp2tp_map = NULL;
}

static void get_inv_sinh(double *inv_sinh,
                         const int gp,
                         const int temperature,
                         const double *frequencies,
                         const int *triplets,
                         const int *triplets_map,
                         const int *stabilized_gp_map,
                         const int ti,
                         const int num_band,
                         const double cutoff_frequency)
{
  int i, gp2;
  double f;

  if (triplets_map[gp] == stabilized_gp_map[gp]) {
    gp2 = triplets[ti * 3 + 2];
  } else {
    gp2 = triplets[ti * 3 + 1];
  }
  for (i = 0; i < num_band; i++) {
    f = frequencies[gp2 * num_band + i];
    if (f > cutoff_frequency) {
      inv_sinh[i] = inv_sinh_occupation(f, temperature);
    } else {
      inv_sinh[i] = 0;
    }
  }
}

static int *create_gp2tp_map(const int *triplets_map,
                             const int num_gp)
{
  int i, max_i, count;
  int *gp2tp_map;

  max_i = 0;
  for (i = 0; i < num_gp; i++) {
    if (max_i < triplets_map[i]) {
      max_i = triplets_map[i];
    }
  }

  gp2tp_map = (int*)malloc(sizeof(int) * (max_i + 1));
  for (i = 0; i < max_i + 1; i++) {
    gp2tp_map[i] = 0;
  }

  count = 0;
  for (i = 0; i < num_gp; i++) {
    if (triplets_map[i] == i) {
      gp2tp_map[i] = count;
      count++;
    }
  }

  return gp2tp_map;
}
