/** @file      types.hpp
 *****************************************************************************
 * @author     Janno Siim, Stefanos Chaliasos
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/
#ifndef TYPES_H_
#define TYPES_H_

#include <vector>
#include <libff/algebra/curves/public_params.hpp> // definitions for groups and
                                                  // other structures

using namespace std;


/*! Structure to represent ElGamal ciphertexts. Uses group G2.
 *
 * \tparam elliptic curve group. Defines group G2
 */
template<typename ppT>
using ElGamal_pair = pair<libff::G2<ppT>, libff::G2<ppT>>;


/*! Vector of elgamal elements of group G2
 *
 * \tparam elliptic curve group. Defines group G2
 */
template<typename ppT>
using ElGamal_vector = vector<ElGamal_pair<ppT>>;

#endif /* TYPES_H_ */
