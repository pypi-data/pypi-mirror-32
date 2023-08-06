
#include "pastml.h"
#include "likelihood.h"

void swap(size_t i, size_t j, size_t* array) {
    size_t tmp = array[i];
    array[i] = array[j];
    array[j] = tmp;
}

void quick_sort_indices(size_t* array, const double* values, size_t start, size_t stop) {
    if ((start + 1) >= stop) {
        return;
    }
    size_t i = start, j = stop - 1, pivot_index = start;
    double pivot = values[array[pivot_index]];
    while (i < pivot_index || j > pivot_index) {
        while (i < pivot_index && values[array[i]] > pivot) {
            i++;
        }
        if (i < pivot_index) {
            swap(i, pivot_index, array);
            pivot_index = i;
            i++;
        }
        while (j > pivot_index && values[array[j]] <= pivot) {
            j--;
        }
        if (j > pivot_index) {
            swap(j, pivot_index, array);
            pivot_index = j;
            j--;
        }
    }
    quick_sort_indices(array, values, start, i);
    quick_sort_indices(array, values, i + 1, stop);
}

void sort_indices(size_t* indices, double* values, size_t n) {
    /* sort indices by the corresponding values in the values array, in decreasing order */
    quick_sort_indices(indices, values, 0, n);
}



void
order_marginal(Tree* tree, size_t num_annotations)
{
    /**
     * Recursively sets node's marginal field to marginal state probabilities in decreasing order,
     * and node's tmp_best field to the corresponding state indices.
     */
    size_t i;
    Node* nd;

    for (i = 0; i < tree->nb_nodes; i++) {
        nd = tree->nodes[i];

        // had to implement a quicksort instead of qsort with a comparator
        // as nested functions cause problems on MacOS with pypi
        sort_indices(nd->best_states, nd->result_probs, num_annotations);
    }
}

void calc_correct(Tree *tree, size_t n) {
    /**
     * Chooses an optimal number of non-zero probabilities to keep, and sets all of them to be equal.
     */

    size_t i, j, k, best_num_states;
    double smallest_correction, correction_i, equal_p_i;
    Node* nd;

    for (k = 0; k < tree->nb_nodes; k++) {
        nd = tree->nodes[k];
        smallest_correction = INFINITY;
        best_num_states = n;
        /* local fraction optimisation */
        for (i = 0; i < n; i++) {
            /* Calculate the squared difference between our probabilities
             * and choosing (i + 1) states with probability 1 / (i + 1) each:
             * correction_i = sum_{j <= i} (p_j - 1 / (i + 1))^2 + sum_{j > i} p_j^2 */
            correction_i = 0.;
            equal_p_i = 1.0 / ((double) i + 1.0);
            for (j = 0; j < n; j++) {
                if (j <= i) {
                    correction_i += pow(nd->result_probs[nd->best_states[j]] - equal_p_i, 2);
                } else {
                    correction_i += pow(nd->result_probs[nd->best_states[j]], 2);
                }
            }
            if (smallest_correction > correction_i) {
                smallest_correction = correction_i;
                best_num_states = i + 1;
            }
        }

        equal_p_i = 1.0 / ((double) best_num_states);
        for (i = 0; i < n; i++) {
            nd->result_probs[nd->best_states[i]] = (i < best_num_states) ? equal_p_i : 0.0;
        }
    }
}

void normalize_result_probabilities(Tree *tree, size_t n) {
    size_t i;
    for (i = 0; i < tree->nb_nodes; i++) {
        normalize(tree->nodes[i]->result_probs, n);
    }
}

void set_id_best_states(Tree *tree, size_t n) {
    size_t * indices = malloc(n * sizeof(size_t));
    size_t i;
    for (i = 0; i < n; i++) {
        indices[i] = i;
    }
    for (i = 0; i < tree->nb_nodes; i++) {
        memcpy(tree->nodes[i]->best_states, indices, n * sizeof(size_t));
    }
    free(indices);
}

void choose_likely_states(Tree *tree, size_t n) {
    /**
     * Chooses an optimal number of non-zero probabilities to keep, and sets all of them to be equal.
     */
    // order marginal probabilities
    order_marginal(tree, n);

    //choose the most likely states to keep among those with non-zero probabilities
    calc_correct(tree, n);
}


void choose_best_marginal_states(Tree *tree, size_t n) {
    /**
     * Chooses the state with the highest marginal probability for each node,
     * sets its result_probs to 1, and the others to 0.
     */
    size_t i, j;
    Node* nd;
    size_t best_state_i;
    double best_p, cur_p;

    for (i = 0; i < tree->nb_nodes; i++) {
        nd = tree->nodes[i];
        best_state_i = 0;
        best_p = 0.0;
        for (j = 0; j < n; j++) {
            cur_p = nd->result_probs[j];
            // set current prob to 0 (if it is the best we'll reset it to 1 below)
            nd->result_probs[j] = 0.0;
            if (best_p < cur_p) {
                best_p = cur_p;
                // set previous best to 0
                nd->result_probs[best_state_i] = 0.0;
                best_state_i = j;
                // set new best to 1
                nd->result_probs[best_state_i] = 1.0;
            }
        }
    }
}
