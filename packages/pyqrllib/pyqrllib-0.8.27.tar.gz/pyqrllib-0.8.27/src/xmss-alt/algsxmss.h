// Distributed under the MIT software license, see the accompanying
// file LICENSE or http://www.opensource.org/licenses/mit-license.php.
// Based on the public domain XMSS reference implementation
// by Andreas Hülsing and Joost Rijneveld

#include <cstddef>
#include "xmss_common.h"

int xmss_Genkeypair(eHashFunction hash_func,
                    xmss_params *params,
                    unsigned char *pk,
                    unsigned char *sk,
                    unsigned char *seed);

int xmss_updateSK(unsigned char *sk,
                  unsigned long k);

int xmss_Signmsg(eHashFunction hash_func,
                 xmss_params *params,
                 unsigned char *sk,
                 unsigned char *sig_msg,
                 unsigned char *msg,
                 size_t msglen);
