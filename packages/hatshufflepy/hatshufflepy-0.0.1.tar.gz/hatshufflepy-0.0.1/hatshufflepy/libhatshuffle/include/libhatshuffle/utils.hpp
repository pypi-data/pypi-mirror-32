/** @file      utils.hpp
 *****************************************************************************
 * @author     Stefanos Chaliasos
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/
#ifndef UTILS_H_
#define UTILS_H_

#include <vector>
#include <map>
#include <string>
#include <libff/algebra/curves/bn128/bn128_pp.hpp>
#include "types.hpp"

using namespace std;
using s_cipher = pair<string [4],string [4]>;
using s_ciphers = vector<s_cipher>;

typedef libff::bn128_pp ppT; //Choose elliptic ppT (currently Barreto-Naehrig)

namespace utils {

/**
 * Print the name of a field and its values.
 *
 * \param name The name of the field to print
 * \param elements A vector of elements (G1 or G2)
 */
template <typename T>
void print_field(string name, vector<T> elements) {
    cout << name << endl;
    for (auto i = elements.begin(); i != elements.end(); ++i) {
        i->print();
    }
}

template <typename T>
void print_field(string name, T element) {
    cout << name << endl;
    element.print();
}

/**
 * Set affine coordinates.
 *
 * \param elements A vector of elements (G1 or G2)
 */
template <typename T>
void to_affine_coordinates(vector<T>* elements) {
    for (auto i = elements->begin(); i != elements->end(); ++i) {
        i->to_special();
    }
}

template <typename T>
void to_affine_coordinates(T *element) {
    element->to_special();
}

/**
 * Create a vector of long from 1 to n.
 *
 * \param n Number of votes to generate
 * \return votes A vector of integers (votes)
 */
vector<long> generate_votes(long n);

/**
 * Get the coorfinates of an ellipt curve point (G1 or G2).
 *
 * \param element Point (G1 or G2)
 * \return coordinates Vector of coordinates in format of
 *  [coord[0], coord[1]] or
 *  [coord[0].a_, coord[0].b_, coord[1].a_, coord[1].b_]
 */
vector <string> get_coordinates (libff::G1<ppT> element);

vector <string> get_coordinates (libff::G2<ppT> element);

vector <vector<string>> get_coordinates (vector<libff::G1<ppT>> list);

vector <vector<string>> get_coordinates (vector<libff::G2<ppT>> list);

/**
 * Create a G1 from coordinates.
 *
 * \param coord A vector of coordinates
 * \return a The G1 element
 */
libff::G1<ppT> set_coordinates_G1(vector<string> coord);

vector<libff::G1<ppT>> set_vector_coordinates_G1(vector<vector<string>> coords);

libff::G2<ppT> set_coordinates_G2(vector<string> coord);

vector<libff::G2<ppT>> set_vector_coordinates_G2(vector<vector<string>> coords);

/**
 * Create ElGamal_pair from a serialized ciphertext
 *
 * \param s_ciphertext A serialized ciphertext
 * \return ciphertext A ciphertext (ElGamal_pair)
 */
ElGamal_pair<ppT> deserialize_ElGamal_pair (
    vector<vector<string>> s_ciphertext);

void to_affine_ElGamal_pair (ElGamal_pair<ppT>* p);

ElGamal_vector<ppT> deserialize_ciphertexts(
    vector<vector<vector<string>>> s_ciphertexts);

void to_affine_ElGamal_vector (ElGamal_vector<ppT>* v);

/**
 * Serialize a ciphertext (ElGamal_pair)
 *
 * \param p An ElGamal_pair (ciphertext)
 * \return serialized_cipher A serialized ciphertext (ElGamal_pair). Format:
 *  pair<A[x1,y1,x2,y2],B[x1,y1,x2,y2]>
 */
pair<string [4], string [4]> serialize_ElGamal_pair (ElGamal_pair<ppT> p);

s_ciphers serialize_ciphertexts (ElGamal_vector<ppT> ciphertexts);

/**
 * Get string representation of Fr (bignum)
 *
 * \param num
 * \return str The string representation
 */
string Fr_toString (libff::Fr<ppT> num);

/**
 * Generate n random votes for testing purposes
 *
 * \param n Number of votes
 * \param h g2_sk
 * \return votes Vector of votes (strings)
 */
vector<string> generate_random_votes (long n, libff::G2<ppT> h);

/*! Generates ElGamal ciphertexts from votes.
 *
 * \param votes
 * \param h ElGamal public key
 * \returns ElGamal ciphertexts
 */
ElGamal_vector<ppT> get_ciphertexts(vector<libff::Fr<ppT>> votes,
                                    libff::G2<ppT> h);

/**
 * Deserialize votes
 *
 * \param d_voted Vector of votes (strings)
 * \return votes Vector of deserialized votes (Fr<ppT>
 */
vector<libff::Fr<ppT>> deserialize_votes (vector<string> d_votes);

/*! Generates a permutation vector containing numbers 0 to n - 1 in random order.
 * Current implementation uses default c++ random generator.
 * This is probably not suitable for crypto.
 *
 * \param n length of the permutation vector
 * \return permutation vector
 */
vector<long> generate_perm(long n);

/*! Generates n ElGamal ciphertexts. Encrypted messages are random.
 *
 * \param n number of ciphertexts
 * \param h ElGamal public key
 * \returns ElGamal ciphertexts
 */
ElGamal_vector<ppT> generate_ciphertexts(long n, libff::G2<ppT> h);

string get_key_for_ciphertext (libff::G2<ppT> ciphertext);

map<string, libff::Fr<ppT>> make_table(vector<libff::Fr<ppT>> votes);

vector<string> decrypt_ciphers (ElGamal_vector<ppT> ciphertexts,
                                libff::Fr<ppT> secret,
                                map<string, libff::Fr<ppT>> table);

} // utils

#endif /* UTILS_H_ */
