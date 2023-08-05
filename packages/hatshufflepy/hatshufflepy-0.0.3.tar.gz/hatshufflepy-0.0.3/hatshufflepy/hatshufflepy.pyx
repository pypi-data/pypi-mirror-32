cimport hatshufflepy


def hat_shuffle_key_gen(votes_number, public_file, secret_file):
    """Generate El Gamal key pair                                          
                                                                                        
    Args:                                                                          
        votes_number: Number of votes to mix                                          
        public_file: JSON file to save public key
        secret_file: File to save secret key
    """ 
    hatshufflepy.key_gen(votes_number, public_file, secret_file)


def hat_shuffle_create_crs(votes_number, crs_file, public_file):
    """Generate Common Reference string                                       
                                                                                        
    Args:                                                                          
        votes_number: Number of votes to mix 
        crs_file: JSON file to save crs
        public_file: Public Key file
    """ 
    hatshufflepy.create_crs(votes_number, crs_file, public_file)


def hat_shuffle_generate_encoded_votes(crs_file, votes_file):
    """Generate encoded votes for testing purposes
                                                                                        
    Args:                                                                          
        crs_file: JSON crs file
        votes_file: JSON file to save votes
    """ 
    hatshufflepy.generate_encoded_votes(crs_file, votes_file)


def hat_shuffle_encrypt(crs_file, votes_file, ciphertexts_file):
    """Encrypt Votes
                                                                                        
    Args:                                                                          
        crs_file: JSON crs file
        votes_file: JSON votes file
        ciphertexts_file: JSON file to save ciphertexts
    """ 
    hatshufflepy.encrypt(crs_file, votes_file, ciphertexts_file)


def hat_shuffle_prove(crs_file, ciphertexts_file, proofs_file):
    """Proving
                                                                                        
    Args:                                                                          
        crs_file: JSON crs file
        ciphertexts_file: JSON ciphertexts file 
        proofs_file: JSON file to save proofs
    """ 
    hatshufflepy.prove(crs_file, ciphertexts_file, proofs_file)


def hat_shuffle_verify(crs_file, ciphertexts_file, proofs_file):
    """Verifying
                                                                                        
    Args:                                                                          
        crs_file: JSON crs file
        ciphertexts_file: JSON ciphertexts file 
        proofs_file: JSON proofs file
    Returns:
        True if verification was successful, otherwise returns False
    """ 
    return hatshufflepy.verify(crs_file, ciphertexts_file, proofs_file)


def hat_shuffle_decrypt(crs_file, votes_file, proofs_file,
                        decrypted_votes_file, secret_file):
    """Decryption of shuffled ciphertexts
                                         
    Args:
        crs_file: JSON crs file
        votes_file: JSON votes file
        proofs_file: JSON proofs file
        secret_file: Secret key filee
        decrypted_votes_file: JSON file to save decrypted votes
    """ 
    hatshufflepy.decrypt(crs_file, votes_file, proofs_file,
                         decrypted_votes_file, secret_file)


def hat_shuffle_test_mixnet(n):
    hatshufflepy.test_mixnet(n)
