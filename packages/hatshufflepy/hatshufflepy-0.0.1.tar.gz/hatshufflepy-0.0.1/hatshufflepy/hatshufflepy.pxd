from libcpp cimport bool
from libcpp.string cimport string


cdef extern from "mixnet.h":
    bool key_gen(int votes_number, string public_file, string secret_file)


cdef extern from "mixnet.h":
    bool create_crs(int votes_number, string crs_file, string public_file)


cdef extern from "mixnet.h":
    bool generate_encoded_votes(string crs_file, string votes_file)


cdef extern from "mixnet.h":
    bool encrypt(string crs_file, string votes_file, string ciphertexts_file)


cdef extern from "mixnet.h":
    bool prove(string crs_file, string ciphertexts_file, string proofs_file)


cdef extern from "mixnet.h":
    bool verify(string crs_file, string ciphertexts_file, string proofs_file)


cdef extern from "mixnet.h":
    bool decrypt(string crs_file, string votes_file, string proofs_file,
                 string decrypted_votes_file, string secret_file)


cdef extern from "mixnet.h":
    bool test_mixnet(int n)
