from cffi import FFI

import evt_exception


class libevt:
    lib = None
    ffi = None
    abi = None


def check_lib_init():
    if(libevt.lib == None):
        evt_exception.evt_exception_raiser(-15)


def init_lib():
    if(libevt.lib == None):
        libevt.ffi = FFI()
        libevt.ffi.cdef("""
                typedef struct {
                    size_t  sz;
                    char    buf[0];
                } evt_data_t;

                #define EVT_OK                       0
                #define EVT_INTERNAL_ERROR          -1
                #define EVT_INVALID_ARGUMENT        -2
                #define EVT_INVALID_PRIVATE_KEY     -3
                #define EVT_INVALID_PUBLIC_KEY      -4
                #define EVT_INVALID_SIGNATURE       -5
                #define EVT_INVALID_HASH            -6
                #define EVT_INVALID_ACTION          -7
                #define EVT_INVALID_BINARY          -8
                #define EVT_SIZE_NOT_EQUALS         -11
                #define EVT_DATA_NOT_EQUALS         -12

                int evt_free(void*);
                int evt_equals(evt_data_t* rhs, evt_data_t* lhs);


                typedef evt_data_t evt_bin_t;
                typedef evt_data_t evt_chain_id_t;
                typedef evt_data_t evt_public_key_t;
                typedef evt_data_t evt_private_key_t;
                typedef evt_data_t evt_signature_t;
                typedef evt_data_t evt_checksum_t;


                void* evt_abi();
                void evt_free_abi(void* abi);
                int evt_abi_json_to_bin(void* evt_abi, const char* action, const char* json, evt_bin_t** bin /* out */);
                int evt_abi_bin_to_json(void* evt_abi, const char* action, evt_bin_t* bin, char** json /* out */);
                int evt_trx_json_to_digest(void* evt_abi, const char* json, evt_chain_id_t* chain_id, evt_checksum_t** digest /* out */);
                int evt_chain_id_from_string(const char* str, evt_chain_id_t** chain_id /* out */);



                int evt_generate_new_pair(evt_public_key_t** pub_key /* out */, evt_private_key_t** priv_key /* out */);
                int evt_get_public_key(evt_private_key_t* priv_key, evt_public_key_t** pub_key /* out */);
                int evt_sign_hash(evt_private_key_t* priv_key, evt_checksum_t* hash, evt_signature_t** sign /* out */);
                int evt_recover(evt_signature_t* sign, evt_checksum_t* hash, evt_public_key_t** pub_key /* out */);
                int evt_hash(const char* buf, size_t sz, evt_checksum_t** hash /* out */);

                int evt_public_key_string(evt_public_key_t* pub_key, char** str /* out */);
                int evt_private_key_string(evt_private_key_t* priv_key, char** str /* out */);
                int evt_signature_string(evt_signature_t* sign, char** str /* out */);
                int evt_checksum_string(evt_checksum_t* hash, char** str /* out */);

                int evt_public_key_from_string(const char* str, evt_public_key_t** pub_key /* out */);
                int evt_private_key_from_string(const char* str, evt_private_key_t** priv_key /* out */);
                int evt_signature_from_string(const char* str, evt_signature_t** sign /* out */);
                int evt_checksum_from_string(const char* str, evt_checksum_t** hash /* out */);






                """)
        libevt.lib = libevt.ffi.dlopen('./lib/libevt.so')
        libevt.abi = libevt.lib.evt_abi()
