/** @file      crs.hpp
 *****************************************************************************
 * @author     Janno Siim, Stefanos Chaliasos
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/
#ifndef CRS_H_
#define CRS_H_

#include <vector>
#include "gmp.h"
#include <fstream>
#include <libff/algebra/fields/bigint.hpp>
#include <libff/algebra/curves/public_params.hpp>
#include <libff/algebra/scalar_multiplication/multiexp.hpp>
#include "utils.hpp"
#include "types.hpp"
#include "shuffle_util.hpp"
#include "nlohmann/json.hpp"


using namespace std;
using json = nlohmann::json;


/*! Common reference string (CRS) for the prover and the verifier.
 * It contains elements of groups G1, G2. Groups G1 and G2
 * are additive (and GT is multiplicative).
 * \brief Common reference string
 * \tparam ppt elliptic curve group
 */
template<typename ppT>
struct CRS {

    //! Number of ciphertexts
    long n;

    //-----Group G1 elements of CRS-----

    //! Elements [P_i(chi)]_1 for i = 1, ..., n
    vector<libff::G1<ppT>> g1_Pis;

    //! Element [rho]_1
    libff::G1<ppT> g1_rho;

    //! Elements [\hat{P}_i(chi)]_1 for i = 1, ..., n
    vector<libff::G1<ppT>> g1_Pi_hats;

    //! Element [\hat{rho}]_1
    libff::G1<ppT> g1_rho_hat;

    //! Element [P0(chi)]_1
    libff::G1<ppT> g1_P0;

    //! Elements [((P_i(chi) + P_0(chi)^2 - 1) / rho)]_1 for i = 1,..., n
    vector<libff::G1<ppT>> g1_Pi_longs;

    //! Element [sum P_i(chi)]_1 for i = 1,..., n
    libff::G1<ppT> g1_sum_Pi;

    //! Element [sum \hat{P}_i(chi)]_1 for i = 1,..., n
    libff::G1<ppT> g1_sum_Pi_hat;

    //! Elements
    // [((beta * P_i(chi) - \hat{beta} * P_i^{hat}(chi))]_1 for i = 1,..., n
    vector<libff::G1<ppT>> g1_Pi_longs2;

    //! Element [(beta * rho + \hat{beta} * \hat{rho})]_1
    libff::G1<ppT> g1_beta_rho;

    //-----Group G2 elements of CRS-----

    //! Element [sk]_2 - El Gamal pk
    libff::G2<ppT> g2_sk;

    //! Elements [P_i(chi)]_2 for i = 1, ..., n
    vector<libff::G2<ppT>> g2_Pis;

    //! Element [P0(chi)]_2
    libff::G2<ppT> g2_P0;

    //! Element [rho]_2
    libff::G2<ppT> g2_rho;

    //! Element [sum P_i(chi)]_2 for i = 1,..., n
    libff::G2<ppT> g2_sum_Pi;

    //! Element [beta]_2
    libff::G2<ppT> g2_beta;

    //! Element [\hat{beta}]_2
    libff::G2<ppT> g2_beta_hat;

    CRS(long n, json &y): n {n} {
        ppT::init_public_params(); //initializes pairing parameters
        deserialize(y);
    }
    /*! Constructs CRS
     * \param n number of ciphertexts
     */
    CRS(long n, libff::G2<ppT> pk): n {n} {
        libff::enter_block("CRS generation", true);
        ppT::init_public_params(); //initializes pairing parameters

        //Random generates secret values
        libff::Fr<ppT> chi = libff::Fr<ppT>::random_element();
        libff::Fr<ppT> beta = libff::Fr<ppT>::random_element();
        libff::Fr<ppT> beta_hat = libff::Fr<ppT>::random_element();
        libff::Fr<ppT> rho = generate_nonzero();
        libff::Fr<ppT> rho_hat = generate_nonzero();

        // set g2_sk
        g2_sk = pk;

        libff::enter_block("Compute P_i(chi)-s", false);
        libff::Fr<ppT> P0 = lagrangian(n + 1, chi) - 1;
        libff::Fr_vector<ppT> Pis = generate_Pi(chi);
        libff::leave_block("Compute P_i(chi)-s", false);

        // Computes Z(chi)
        libff::Fr<ppT> Zchi = libff::Fr<ppT>::one();
        for (long i = 1; i < n + 2; i++) {
            Zchi *= chi - i;
        }

        libff::enter_block("Compute hat{P}_i(chi)-s", false);
        libff::Fr_vector<ppT> Pi_hats;
        libff::Fr<ppT> chi_n_1 = chi ^ (n + 1);
        libff::Fr<ppT> temp = chi_n_1;
        for (long i = 0; i < n; i++) {
            temp = temp * chi_n_1;
            Pi_hats.push_back(temp);
        }
        libff::leave_block("Compute hat{P}_i(chi)-s", false);

        libff::enter_block("Init G1 elements", false);
        libff::inhibit_profiling_info = true;
        init_G1_elems(chi, beta, beta_hat, rho, rho_hat, Pis, Pi_hats, P0);
        libff::inhibit_profiling_info = false;
        libff::leave_block("Init G1 elements", false);

        libff::enter_block("Init G2 elements", false);
        libff::inhibit_profiling_info = true;
        init_G2_elems(beta, beta_hat, rho, Pis, P0);
        libff::inhibit_profiling_info = false;
        libff::leave_block("Init G2 elements", false);

        libff::leave_block("CRS generation", true);
    }

    void print() {
        cout << "n: " << n << endl;
        cout << "G1" << endl;
        utils::print_field <libff::G1<ppT>> ("g1_Pis", g1_Pis);
        utils::print_field <libff::G1<ppT>> ("g1_rho", g1_rho);
        utils::print_field <libff::G1<ppT>> ("g1_Pi_hats", g1_Pi_hats);
        utils::print_field <libff::G1<ppT>> ("g1_P0", g1_P0);
        utils::print_field <libff::G1<ppT>> ("g1_Pi_longs", g1_Pi_longs);
        utils::print_field <libff::G1<ppT>> ("g1_sum_Pi", g1_sum_Pi);
        utils::print_field <libff::G1<ppT>> ("g1_sum_Pi_hat", g1_sum_Pi_hat);
        utils::print_field <libff::G1<ppT>> ("g1_Pi_longs2", g1_Pi_longs2);
        utils::print_field <libff::G1<ppT>> ("g1_beta_rho", g1_beta_rho);
        cout << "G2" << endl;
        utils::print_field <libff::G2<ppT>> ("g2_sk", g2_sk);
        utils::print_field <libff::G2<ppT>> ("g2_Pis", g2_Pis);
        utils::print_field <libff::G2<ppT>> ("g2_P0", g2_P0);
        utils::print_field <libff::G2<ppT>> ("g2_rho", g2_rho);
        utils::print_field <libff::G2<ppT>> ("g2_sum_Pi", g2_sum_Pi);
        utils::print_field <libff::G2<ppT>> ("g2_beta", g2_beta);
        utils::print_field <libff::G2<ppT>> ("g2_beta_hat", g2_beta_hat);
    }

    void to_affine() {
        utils::to_affine_coordinates(&g1_Pis);
        utils::to_affine_coordinates(&g1_rho);
        utils::to_affine_coordinates(&g1_rho_hat);
        utils::to_affine_coordinates(&g1_Pi_hats);
        utils::to_affine_coordinates(&g1_P0);
        utils::to_affine_coordinates(&g1_Pi_longs);
        utils::to_affine_coordinates(&g1_sum_Pi);
        utils::to_affine_coordinates(&g1_sum_Pi_hat);
        utils::to_affine_coordinates(&g1_Pi_longs2);
        utils::to_affine_coordinates(&g1_beta_rho);
        utils::to_affine_coordinates(&g2_sk);
        utils::to_affine_coordinates(&g2_Pis);
        utils::to_affine_coordinates(&g2_P0);
        utils::to_affine_coordinates(&g2_rho);
        utils::to_affine_coordinates(&g2_sum_Pi);
        utils::to_affine_coordinates(&g2_beta);
        utils::to_affine_coordinates(&g2_beta_hat);
    }

    json serialize() {
        json crs_json = {
            {"n", n},
            {"g1_Pis", utils::get_coordinates(g1_Pis)},
            {"g1_rho", utils::get_coordinates(g1_rho)},
            {"g1_rho_hat", utils::get_coordinates(g1_rho_hat)},
            {"g1_Pi_hats", utils::get_coordinates(g1_Pi_hats)},
            {"g1_P0", utils::get_coordinates(g1_P0)},
            {"g1_Pi_longs", utils::get_coordinates(g1_Pi_longs)},
            {"g1_sum_Pi", utils::get_coordinates(g1_sum_Pi)},
            {"g1_sum_Pi_hat", utils::get_coordinates(g1_sum_Pi_hat)},
            {"g1_Pi_longs2", utils::get_coordinates(g1_Pi_longs2)},
            {"g1_beta_rho", utils::get_coordinates(g1_beta_rho)},
            {"g2_sk", utils::get_coordinates(g2_sk)},
            {"g2_Pis", utils::get_coordinates(g2_Pis)},
            {"g2_P0", utils::get_coordinates(g2_P0)},
            {"g2_rho", utils::get_coordinates(g2_rho)},
            {"g2_sum_Pi", utils::get_coordinates(g2_sum_Pi)},
            {"g2_beta", utils::get_coordinates(g2_beta)},
            {"g2_beta_hat", utils::get_coordinates(g2_beta_hat)},
        };
        return crs_json;
    }

    void deserialize (json &y) {
        g1_rho = utils::set_coordinates_G1 (y["g1_rho"]);
        g1_rho_hat = utils::set_coordinates_G1 (y["g1_rho_hat"]);
        g1_P0 = utils::set_coordinates_G1 (y["g1_P0"]);
        g1_beta_rho = utils::set_coordinates_G1 (y["g1_beta_rho"]);
        g1_sum_Pi = utils::set_coordinates_G1 (y["g1_sum_Pi"]);
        g1_sum_Pi_hat = utils::set_coordinates_G1 (y["g1_sum_Pi_hat"]);
        g1_Pi_hats = utils::set_vector_coordinates_G1 (y["g1_Pi_hats"]);
        g1_Pi_longs = utils::set_vector_coordinates_G1 (y["g1_Pi_longs"]);
        g1_Pi_longs2 = utils::set_vector_coordinates_G1 (y["g1_Pi_longs2"]);
        g1_Pis = utils::set_vector_coordinates_G1 (y["g1_Pis"]);
        g2_sk = utils::set_coordinates_G2 (y["g2_sk"]);
        g2_P0 = utils::set_coordinates_G2 (y["g2_P0"]);
        g2_beta = utils::set_coordinates_G2 (y["g2_beta"]);
        g2_beta_hat = utils::set_coordinates_G2 (y["g2_beta_hat"]);
        g2_rho = utils::set_coordinates_G2 (y["g2_rho"]);
        g2_sum_Pi = utils::set_coordinates_G2 (y["g2_sum_Pi"]);
        g2_Pis = utils::set_vector_coordinates_G2 (y["g2_Pis"]);
    }

private:
    /*! Generates CRS elements of group G1.
     * See the paper for the meaning of parameters.
     * \param Pis polynomials Pi evaluated at point chi
     * \param P0 polynomial P0 evaluated at point chi
     */
    void init_G1_elems(libff::Fr<ppT> chi, libff::Fr<ppT> beta,
                       libff::Fr<ppT> beta_hat, libff::Fr<ppT> rho,
                       libff::Fr<ppT> rho_hat, libff::Fr_vector<ppT> Pis,
                       libff::Fr_vector<ppT> Pi_hats, libff::Fr<ppT> P0) {

    libff::G1<ppT> g1 = libff::G1<ppT>::one();

    // Precomputation for fixed-base multi-exponentiation of g1.
    const size_t g1_exp_count = 4 * n + 6;
    size_t g1_window_size = libff::get_exp_window_size<libff::G1<ppT>>(
                                g1_exp_count);
    libff::window_table<libff::G1<ppT>> g1_table = libff::get_window_table(
                                                libff::Fr<ppT>::size_in_bits(),
                                                g1_window_size,
                                                g1);

    g1_Pis = batch_exp(libff::Fr<ppT>::size_in_bits(),
                       g1_window_size, g1_table, Pis);
    g1_rho = windowed_exp(libff::Fr<ppT>::size_in_bits(),
                          g1_window_size, g1_table, rho);

    g1_Pi_hats = batch_exp(libff::Fr<ppT>::size_in_bits(),
                           g1_window_size, g1_table, Pi_hats);
    g1_rho_hat = windowed_exp(libff::Fr<ppT>::size_in_bits(),
                              g1_window_size, g1_table, rho_hat);

    g1_P0 = windowed_exp(libff::Fr<ppT>::size_in_bits(),
                         g1_window_size, g1_table, P0);

    // field elements ((P_i(chi) + P0(chi))^2 - 1)/rho
    libff::Fr_vector<ppT> Pi_longs;
    libff::Fr<ppT> rho_inv = rho.inverse();
    for (long i = 0; i < n; i++) {
        Pi_longs.push_back((((Pis.at(i) + P0) ^ 2) - 1) * rho_inv);
    }

    g1_Pi_longs = batch_exp(libff::Fr<ppT>::size_in_bits(), g1_window_size,
                            g1_table, Pi_longs);

    libff::Fr<ppT> sum_Pi = vector_sum<libff::Fr<ppT>>(Pis);
    libff::Fr<ppT> sum_Pi_hat = vector_sum<libff::Fr<ppT>>(Pi_hats);

    g1_sum_Pi = windowed_exp(libff::Fr<ppT>::size_in_bits(), g1_window_size,
                             g1_table, sum_Pi);
    g1_sum_Pi_hat = windowed_exp(libff::Fr<ppT>::size_in_bits(),
                                 g1_window_size, g1_table, sum_Pi_hat);

    // field elements ((beta * P_i(chi) + \hat{beta} * \hat{P}_i(chi))
    libff::Fr_vector<ppT> Pi_longs2;
    for (long i = 0; i < n; i++) {
        Pi_longs2.push_back(beta * Pis.at(i) + beta_hat * Pi_hats.at(i));
    }

    g1_Pi_longs2 = batch_exp(libff::Fr<ppT>::size_in_bits(), g1_window_size,
                             g1_table, Pi_longs2);
    g1_beta_rho = windowed_exp(libff::Fr<ppT>::size_in_bits(), g1_window_size,
                               g1_table, beta * rho + beta_hat * rho_hat);
  }


    /*! Generates CRS elements of group G2.
    * See the paper for the meaning of parameters.
    *
    * \param sk ElGamal secret key
    */
    void init_G2_elems(libff::Fr<ppT> beta,
                       libff::Fr<ppT> beta_hat, libff::Fr<ppT> rho,
                       libff::Fr_vector<ppT> Pis, libff::Fr<ppT> P0) {
        libff::G2<ppT> g2 = libff::G2<ppT>::one();

        // Precomputation for fixed-based multi-exponentiation of g2.
        const size_t g2_exp_count = n + 6;
        size_t g2_window_size = libff::get_exp_window_size<libff::G2<ppT>>(
                                    g2_exp_count);
        libff::window_table<libff::G2<ppT>> g2_table = get_window_table(
                                                libff::Fr<ppT>::size_in_bits(),
                                                g2_window_size, g2);

        g2_Pis = batch_exp(libff::Fr<ppT>::size_in_bits(), g2_window_size,
                           g2_table, Pis);
        g2_rho = windowed_exp(libff::Fr<ppT>::size_in_bits(), g2_window_size,
                              g2_table, rho);
        g2_P0 = windowed_exp(libff::Fr<ppT>::size_in_bits(), g2_window_size,
                             g2_table, P0);

        libff::Fr<ppT> sum_Pi = vector_sum<libff::Fr<ppT>>(Pis);

        g2_sum_Pi = windowed_exp(libff::Fr<ppT>::size_in_bits(),
                                 g2_window_size,
                                 g2_table, sum_Pi);
        g2_beta = windowed_exp(libff::Fr<ppT>::size_in_bits(), g2_window_size,
                               g2_table, beta);
        g2_beta_hat = windowed_exp(libff::Fr<ppT>::size_in_bits(),
                                   g2_window_size,
                                   g2_table,
                                   beta_hat);
    }


    /*! Generates uniformly random nonzero field element.
     * \returns nonzero field element
     */
    libff::Fr<ppT> generate_nonzero() {
        libff::Fr<ppT> elem = libff::Fr<ppT>::random_element();
        while (elem == libff::Fr<ppT>::zero()) {
            elem = libff::Fr<ppT>::random_element();
        }

        return elem;
    }


    /*! Evaluates Lagrange basis polynomial l_i at point x.
     * Distinct points are 1, 2, ..., n + 1.
     * \param i polynomial index in range 1, 2, ..., n + 1
     * \param chi input value of the polynomial
     * \returns nonzero field element
     */
    libff::Fr<ppT> lagrangian(long i, libff::Fr<ppT> chi) {
        libff::Fr<ppT> numerator = libff::Fr<ppT>::one();
        libff::Fr<ppT> denominator = libff::Fr<ppT>::one();
        for (long j = 1; j < n + 2; j++) {
            if (i == j) continue;

            numerator = numerator * (chi - j);
            libff::Fr<ppT> elem = i - j;
            denominator = denominator * elem;
        }

        return numerator * denominator.invert();
    }


    /*! Computes denominators for Lagrange basis polynomials.
     * Uses distinct points 1, ...,k
     * \param k number of basis polynomials
     */
    libff::Fr_vector<ppT> compute_denominators(long k) {

        libff::Fr_vector<ppT> denominators;

        libff::Fr<ppT> temp = libff::Fr<ppT>::one();
        for (long i = 1; i <= k; i++) {
            if (i == 1) {
            for (long j = 2; j <= k; j++) {
                libff::Fr<ppT> elem = i - j;
                temp = temp * elem;
            }

            } else if (i == k) {
                libff::Fr<ppT> elem = libff::Fr<ppT>::one() - k;
                temp = elem * temp;
            } else {

                libff::Fr<ppT> inverse = i - 1 - k;
                inverse = inverse.invert();

                libff::Fr<ppT> elem = i - 1;

                temp = elem * temp * inverse;
            }
            denominators.push_back(temp);
        }

        return denominators;
    }


    /*! Computes vector of elements P_i(chi) for i = 1, ..., n.
     * Uses Lagrange basis polynomials with distinct points 1, ..., n + 1
     * \param chi point of evaluation
     */
    libff::Fr_vector<ppT> generate_Pi(libff::Fr<ppT> chi) {
        libff::Fr_vector<ppT> Pis;

        libff::Fr<ppT> prod = libff::Fr<ppT>::one();
        for (long j = 1; j < n + 2; j++) prod = prod * (chi - j);

        libff::Fr_vector<ppT> denoms = compute_denominators(n + 1);

        libff::Fr<ppT> missing_factor = chi - (n + 1);

        //l_{n + 1} (\chi)
        libff::Fr<ppT> l_n_1 = prod * (missing_factor * denoms.at(n)).inverse();

        libff::Fr<ppT> two = 2;
        for (long i = 1; i < n + 1; i++) {
            missing_factor = chi - i;
            libff::Fr<ppT> l_i = prod * (missing_factor * denoms.at(i - 1)).inverse();
            Pis.push_back(two * l_i + l_n_1);
        }

        return Pis;
    }
}; // CRS

#endif /* CRS_H_ */
