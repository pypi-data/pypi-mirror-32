/** @file      shuffle_util.hpp
 *****************************************************************************
 * @author     Janno Siim, Stefanos Chaliasos
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/
#ifndef SHUFFLE_UTIL_H_
#define SHUFFLE_UTIL_H_

#include <vector>
#include <libff/algebra/curves/public_params.hpp> // definitions for groups and
                                                  // other structures

using namespace std;

/*! Adds group or field elements together.
 *
 * \param vec vector
 * \tparam S structure type (G1, G2, GT, Fr)
 * \returns sum of vector
 */
template <typename S>
S vector_sum(vector<S> vec){
S sum = S::zero();
    for(size_t i = 0; i < vec.size(); i++){
          sum = sum + vec.at(i);
    }
       return sum;
}

/*! Generates n uniformly random and independent elements of either
 * G1, G2, GT or Fr.
 * For field elements generate_randomizers(long n, long k) is much faster.
 *
 * \param n number of field elements
 * \tparam S structure type (G1, G2, Gt or Fr)
 * \returns n fields elements
 */
template <typename S>
vector<S> generate_randomizers(long n){
      vector<S> randomizers;
      for(long i = 0; i < n; i++) randomizers.push_back(S::random_element());
      return randomizers;
}

/*! Generates a vector of n uniformly random field elements.
 *
 * \param n number of field elements
 * \param k number of bits in a field element
 * \tparam ppT curve type
 * \returns n random elements;
 */
template <typename ppT>
libff::Fr_vector<ppT> generate_randomizers(long n, long k){
      size_t k_in_bytes = k / 8;
      libff::Fr_vector<ppT> vec;
      FILE *fp = fopen("/dev/urandom", "r");
      for(long i = 0; i < n; i++){
          libff::bigint<libff::Fr<ppT>::num_limbs> r;
          size_t bytes_read = fread(r.data, 1, k_in_bytes , fp);
          assert(bytes_read == k_in_bytes);
          libff::Fr<ppT> f(r);
          vec.push_back(f);
      }
      fclose(fp);
      return vec;
}

/*! Generates k bit fields element.
 * \param k number of bits
 * \returns randomizer
 * \paramt curve type
 */
template <typename ppT>
libff::Fr<ppT> generate_randomizer(long k){
      size_t k_in_bytes = k / 8;
      FILE *fp = fopen("/dev/urandom", "r");

      libff::bigint<libff::Fr<ppT>::num_limbs> r;
      size_t bytes_read = fread(r.data, 1, k_in_bytes , fp);
      assert(bytes_read == k_in_bytes);
      libff::Fr<ppT> f(r);

      fclose(fp);

      return f;
}


#endif /* SHUFFLE_UTIL_H_ */
