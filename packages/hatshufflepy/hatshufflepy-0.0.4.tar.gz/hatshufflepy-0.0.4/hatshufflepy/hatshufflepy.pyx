cimport hatshufflepy
from cpython.version cimport PY_MAJOR_VERSION


cdef unicode _ustring(s):
    if type(s) is unicode:
        # fast path for most common case(s)
        return <unicode>s
    elif PY_MAJOR_VERSION < 3 and isinstance(s, bytes):
        # only accept byte strings in Python 2.x, not in Py3
        return (<bytes>s).decode('ascii')
    elif isinstance(s, unicode):
        # an evil cast to <unicode> might work here in some(!) cases,
        # depending on what the further processing does.  to be safe,
        # we can always create a copy instead
        return unicode(s)
    else:
        raise TypeError(...)


def hat_shuffle_key_gen(votes_number, public_file, secret_file):
    """Generate El Gamal key pair                                          
                                                                                        
    Args:                                                                          
        votes_number: Number of votes to mix                                          
        public_file: JSON file to save public key
        secret_file: File to save secret key
    """ 
    public_file_ = public_file.encode('utf-8')
    secret_file_ = secret_file.encode('utf-8')
    hatshufflepy.key_gen(votes_number, public_file_, secret_file_)


def hat_shuffle_create_crs(votes_number, crs_file, public_file):
    """Generate Common Reference string                                       
                                                                                        
    Args:                                                                          
        votes_number: Number of votes to mix 
        crs_file: JSON file to save crs
        public_file: Public Key file
    """ 
    crs_file_ = crs_file.encode('utf-8')
    public_file_ = public_file.encode('utf-8')
    hatshufflepy.create_crs(votes_number, crs_file_, public_file_)


def hat_shuffle_generate_encoded_votes(crs_file, votes_file):
    """Generate encoded votes for testing purposes
                                                                                        
    Args:                                                                          
        crs_file: JSON crs file
        votes_file: JSON file to save votes
    """ 
    crs_file_ = crs_file.encode('utf-8')
    votes_file_ = votes_file.encode('utf-8')
    hatshufflepy.generate_encoded_votes(crs_file_, votes_file_)


def hat_shuffle_encrypt(crs_file, votes_file, ciphertexts_file):
    """Encrypt Votes
                                                                                        
    Args:                                                                          
        crs_file: JSON crs file
        votes_file: JSON votes file
        ciphertexts_file: JSON file to save ciphertexts
    """ 
    crs_file_ = crs_file.encode('utf-8')
    votes_file_ = votes_file.encode('utf-8')
    ciphertexts_file_ = ciphertexts_file.encode('utf-8')
    hatshufflepy.encrypt(crs_file_, votes_file_, ciphertexts_file_)


def hat_shuffle_prove(crs_file, ciphertexts_file, proofs_file):
    """Proving
                                                                                        
    Args:                                                                          
        crs_file: JSON crs file
        ciphertexts_file: JSON ciphertexts file 
        proofs_file: JSON file to save proofs
    """ 
    crs_file_ = crs_file.encode('utf-8')
    proofs_file_ = proofs_file.encode('utf-8')
    ciphertexts_file_ = ciphertexts_file.encode('utf-8')
    hatshufflepy.prove(crs_file_, ciphertexts_file_, proofs_file_)


def hat_shuffle_verify(crs_file, ciphertexts_file, proofs_file):
    """Verifying
                                                                                        
    Args:                                                                          
        crs_file: JSON crs file
        ciphertexts_file: JSON ciphertexts file 
        proofs_file: JSON proofs file
    Returns:
        True if verification was successful, otherwise returns False
    """ 
    crs_file_ = crs_file.encode('utf-8')
    proofs_file_ = proofs_file.encode('utf-8')
    ciphertexts_file_ = ciphertexts_file.encode('utf-8')
    return hatshufflepy.verify(crs_file_, ciphertexts_file_, proofs_file_)


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
    crs_file_ = crs_file.encode('utf-8')
    proofs_file_ = proofs_file.encode('utf-8')
    votes_file_ = votes_file.encode('utf-8')
    secret_file_ = secret_file.encode('utf-8')
    decrypted_votes_file_ = decrypted_votes_file.encode('utf-8')
    hatshufflepy.decrypt(crs_file_, votes_file_, proofs_file_,
                         decrypted_votes_file_, secret_file_)


def hat_shuffle_test_mixnet(n):
    hatshufflepy.test_mixnet(n)
