/** @file      mixnet.h
 *****************************************************************************
 * @author     Stefanos Chaliasos
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/
#ifndef MIXNET_H_
#define MIXNET_H_

#include <iostream>
#include <vector>
#include <map>
#include <sstream>
#include <algorithm> //has random_shuffle
#include <libff/algebra/fields/bigint.hpp>
#include <libff/algebra/curves/bn128/bn128_pp.hpp>
#include <libff/algebra/curves/public_params.hpp>
#include <fstream>
#include <string>
#include "types.hpp"
#include "crs.hpp"
#include "utils.hpp"
#include "prover.hpp"
#include "verifier.hpp"
#include "shuffle_util.hpp"
#include "nlohmann/json.hpp"

using namespace std;
using json = nlohmann::json;
using s_cipher = pair<string [4],string [4]>;
using s_ciphers = vector<s_cipher>;

typedef libff::bn128_pp curve;

json serialiaze_proof(Proof<curve> proof);

bool key_gen(long n, string public_file, string secret_file);

bool create_crs(long votes_number, string crs_file, string public_file);

bool generate_encoded_votes(string crs_file, string votes_file);

bool encrypt(string crs_file, string votes_file, string ciphertexts_file);

bool prove(string crs_file, string ciphertexts_file, string proofs_file);

bool verify(string crs_file, string ciphertexts_file, string proofs_file);

bool decrypt(string crs_file, string votes_file, string proofs_file,
             string decrypted_votes_file, string secret_file);

bool test_mixnet(long n);

#endif /* MIXNET_H_ */
