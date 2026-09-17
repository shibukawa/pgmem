package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgp_cfb_decrypt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v7 != 0 {
		v8 = int32(_a_F_pgp_cfb_decrypt_0)
	} else {
		v8 = int32(_a_F_pgp_cfb_decrypt_1)
	}
	F_cfb_process(m, l0, l1, l2, l3, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_pgp_get_digest_code(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v202 int32
	_ = v202
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v292 int32
	_ = v292
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	v7 = int32(_a_F_pgp_get_digest_code_0)
	v8 = l0
	goto L4
L1:
	;
	return v320
L2:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)+4))
	v320 = v319
	goto L1
L3:
	;
	if v45 == int32(0) {
		v318 = int32(_a_F_pgp_get_digest_code_1)
		goto L2
	} else {
		goto L16
	}
L4:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v11 == v12 {
		v34 = v11
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v45 = int32(0)
	goto L3
L6:
	;
	v36 = int32(1)
	if v34 != 0 {
		v7 = v7 + v36
		v8 = v8 + v36
		goto L4
	} else {
		goto L15
	}
L7:
	;
	if base.Ui32((v11-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v22 = v11 | int32(32)
	goto L10
L9:
	;
	v22 = v11
	goto L10
L10:
	;
	if base.Ui32((v12-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v31 = v12 | int32(32)
	goto L13
L12:
	;
	v31 = v12
	goto L13
L13:
	;
	if v22 == v31 {
		v34 = v22
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v45 = v22 - v31
	goto L3
L15:
	;
	goto L5
L16:
	;
	v52 = int32(_a_F_pgp_get_digest_code_2)
	v53 = l0
	goto L18
L17:
	;
	if v90 == int32(0) {
		v318 = int32(_a_F_pgp_get_digest_code_3)
		goto L2
	} else {
		goto L30
	}
L18:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v56 == v57 {
		v79 = v56
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v90 = int32(0)
	goto L17
L20:
	;
	v81 = int32(1)
	if v79 != 0 {
		v52 = v52 + v81
		v53 = v53 + v81
		goto L18
	} else {
		goto L29
	}
L21:
	;
	if base.Ui32((v56-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v67 = v56 | int32(32)
	goto L24
L23:
	;
	v67 = v56
	goto L24
L24:
	;
	if base.Ui32((v57-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v76 = v57 | int32(32)
	goto L27
L26:
	;
	v76 = v57
	goto L27
L27:
	;
	if v67 == v76 {
		v79 = v67
		goto L20
	} else {
		goto L28
	}
L28:
	;
	v90 = v67 - v76
	goto L17
L29:
	;
	goto L19
L30:
	;
	v97 = int32(_a_F_pgp_get_digest_code_4)
	v98 = l0
	goto L32
L31:
	;
	if v135 == int32(0) {
		v318 = int32(_a_F_pgp_get_digest_code_5)
		goto L2
	} else {
		goto L44
	}
L32:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if v101 == v102 {
		v124 = v101
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v135 = int32(0)
	goto L31
L34:
	;
	v126 = int32(1)
	if v124 != 0 {
		v97 = v97 + v126
		v98 = v98 + v126
		goto L32
	} else {
		goto L43
	}
L35:
	;
	if base.Ui32((v101-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v112 = v101 | int32(32)
	goto L38
L37:
	;
	v112 = v101
	goto L38
L38:
	;
	if base.Ui32((v102-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v121 = v102 | int32(32)
	goto L41
L40:
	;
	v121 = v102
	goto L41
L41:
	;
	if v112 == v121 {
		v124 = v112
		goto L34
	} else {
		goto L42
	}
L42:
	;
	v135 = v112 - v121
	goto L31
L43:
	;
	goto L33
L44:
	;
	v142 = int32(_a_F_pgp_get_digest_code_6)
	v143 = l0
	goto L46
L45:
	;
	if v180 == int32(0) {
		v318 = int32(_a_F_pgp_get_digest_code_7)
		goto L2
	} else {
		goto L58
	}
L46:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	if v146 == v147 {
		v169 = v146
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v180 = int32(0)
	goto L45
L48:
	;
	v171 = int32(1)
	if v169 != 0 {
		v142 = v142 + v171
		v143 = v143 + v171
		goto L46
	} else {
		goto L57
	}
L49:
	;
	if base.Ui32((v146-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v157 = v146 | int32(32)
	goto L52
L51:
	;
	v157 = v146
	goto L52
L52:
	;
	if base.Ui32((v147-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v166 = v147 | int32(32)
	goto L55
L54:
	;
	v166 = v147
	goto L55
L55:
	;
	if v157 == v166 {
		v169 = v157
		goto L48
	} else {
		goto L56
	}
L56:
	;
	v180 = v157 - v166
	goto L45
L57:
	;
	goto L47
L58:
	;
	v187 = int32(_a_F_pgp_get_digest_code_8)
	v188 = l0
	goto L60
L59:
	;
	if v225 == int32(0) {
		v318 = int32(_a_F_pgp_get_digest_code_9)
		goto L2
	} else {
		goto L72
	}
L60:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if v191 == v192 {
		v214 = v191
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v225 = int32(0)
	goto L59
L62:
	;
	v216 = int32(1)
	if v214 != 0 {
		v187 = v187 + v216
		v188 = v188 + v216
		goto L60
	} else {
		goto L71
	}
L63:
	;
	if base.Ui32((v191-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v202 = v191 | int32(32)
	goto L66
L65:
	;
	v202 = v191
	goto L66
L66:
	;
	if base.Ui32((v192-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v211 = v192 | int32(32)
	goto L69
L68:
	;
	v211 = v192
	goto L69
L69:
	;
	if v202 == v211 {
		v214 = v202
		goto L62
	} else {
		goto L70
	}
L70:
	;
	v225 = v202 - v211
	goto L59
L71:
	;
	goto L61
L72:
	;
	v232 = int32(_a_F_pgp_get_digest_code_10)
	v233 = l0
	goto L74
L73:
	;
	if v270 == int32(0) {
		v318 = int32(_a_F_pgp_get_digest_code_11)
		goto L2
	} else {
		goto L86
	}
L74:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232))))
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
	if v236 == v237 {
		v259 = v236
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v270 = int32(0)
	goto L73
L76:
	;
	v261 = int32(1)
	if v259 != 0 {
		v232 = v232 + v261
		v233 = v233 + v261
		goto L74
	} else {
		goto L85
	}
L77:
	;
	if base.Ui32((v236-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v247 = v236 | int32(32)
	goto L80
L79:
	;
	v247 = v236
	goto L80
L80:
	;
	if base.Ui32((v237-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v256 = v237 | int32(32)
	goto L83
L82:
	;
	v256 = v237
	goto L83
L83:
	;
	if v247 == v256 {
		v259 = v247
		goto L76
	} else {
		goto L84
	}
L84:
	;
	v270 = v247 - v256
	goto L73
L85:
	;
	goto L75
L86:
	;
	v277 = int32(_a_F_pgp_get_digest_code_12)
	v278 = l0
	goto L88
L87:
	;
	if v315 != 0 {
		v320 = int32(-104)
		goto L1
	} else {
		goto L100
	}
L88:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277))))
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278))))
	if v281 == v282 {
		v304 = v281
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v315 = int32(0)
	goto L87
L90:
	;
	v306 = int32(1)
	if v304 != 0 {
		v277 = v277 + v306
		v278 = v278 + v306
		goto L88
	} else {
		goto L99
	}
L91:
	;
	if base.Ui32((v281-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v292 = v281 | int32(32)
	goto L94
L93:
	;
	v292 = v281
	goto L94
L94:
	;
	if base.Ui32((v282-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v301 = v282 | int32(32)
	goto L97
L96:
	;
	v301 = v282
	goto L97
L97:
	;
	if v292 == v301 {
		v304 = v292
		goto L90
	} else {
		goto L98
	}
L98:
	;
	v315 = v292 - v301
	goto L87
L99:
	;
	goto L89
L100:
	;
	v318 = int32(_a_F_pgp_get_digest_code_13)
	goto L2
}
func F_pgp_get_digest_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v3 = l0 - int32(1)
	if base.Ui32(v3) <= base.Ui32(int32(9)) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v3<<(uint(int32(2))%32))+uint32(_c_F_pgp_get_digest_name[0])))
		v12 = v10
	} else {
		v12 = int32(0)
	}
	return v12
}
func F_pgp_key_alloc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_palloc0(m, int32(52))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v3
		return int32(0)
	}
}
func F_pgp_s2k_read(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v13 = F_pullf_read_fixed(m, l0, int32(1), v8+int32(15))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 < int32(0) {
			v53 = v13
			m.G0 = v8 + int32(16)
			return v53
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v19)
			v24 = F_pullf_read_fixed(m, l0, int32(1), v8+int32(14))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				if v24 < int32(0) {
					v53 = v24
					m.G0 = v8 + int32(16)
					return v53
				} else {
					v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)))
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v28)
					v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
					switch v30 {
					case 0:
						v53 = v30
						m.G0 = v8 + int32(16)
						return v53
					case 1:
						v34 = F_pullf_read_fixed(m, l0, int32(8), l1+int32(2))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							v53 = v34
							m.G0 = v8 + int32(16)
							return v53
						}
					default:
						v53 = int32(-121)
						m.G0 = v8 + int32(16)
						return v53
					case 3:
						v39 = F_pullf_read_fixed(m, l0, int32(8), l1+int32(2))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							if v39 < int32(0) {
								v53 = v39
								m.G0 = v8 + int32(16)
								return v53
							} else {
								v46 = F_pullf_read_fixed(m, l0, int32(1), v8+int32(13))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									if v46 < int32(0) {
										v53 = v46
									} else {
										v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+13)))
										*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)) = uint8(v50)
										v53 = v39
									}
									m.G0 = v8 + int32(16)
									return v53
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pgp_set_compress_algo(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	if base.Ui32(l1) <= base.Ui32(int32(3)) {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = l1
		v8 = int32(0)
	} else {
		v8 = int32(-13)
	}
	return v8
}
func F_pgp_set_s2k_cipher_algo(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	v6 = F_pg_strcasecmp(m, int32(_a_F_pgp_set_s2k_cipher_algo_0), l1)
	mBase = m.M
	if v6 == int32(0) {
		v49 = int32(_a_F_pgp_set_s2k_cipher_algo_1)
		v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
		v51 = v50
	} else {
		v11 = F_pg_strcasecmp(m, int32(_a_F_pgp_set_s2k_cipher_algo_2), l1)
		mBase = m.M
		if v11 == int32(0) {
			v49 = int32(_a_F_pgp_set_s2k_cipher_algo_3)
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
			v51 = v50
		} else {
			v16 = F_pg_strcasecmp(m, int32(_a_F_pgp_set_s2k_cipher_algo_4), l1)
			mBase = m.M
			if v16 == int32(0) {
				v49 = int32(_a_F_pgp_set_s2k_cipher_algo_5)
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
				v51 = v50
			} else {
				v21 = F_pg_strcasecmp(m, int32(_a_F_pgp_set_s2k_cipher_algo_6), l1)
				mBase = m.M
				if v21 == int32(0) {
					v49 = int32(_a_F_pgp_set_s2k_cipher_algo_7)
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
					v51 = v50
				} else {
					v26 = F_pg_strcasecmp(m, int32(_a_F_pgp_set_s2k_cipher_algo_8), l1)
					mBase = m.M
					if v26 == int32(0) {
						v49 = int32(_a_F_pgp_set_s2k_cipher_algo_9)
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
						v51 = v50
					} else {
						v31 = F_pg_strcasecmp(m, int32(_a_F_pgp_set_s2k_cipher_algo_10), l1)
						mBase = m.M
						if v31 == int32(0) {
							v49 = int32(_a_F_pgp_set_s2k_cipher_algo_11)
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
							v51 = v50
						} else {
							v36 = F_pg_strcasecmp(m, int32(_a_F_pgp_set_s2k_cipher_algo_12), l1)
							mBase = m.M
							if v36 == int32(0) {
								v49 = int32(_a_F_pgp_set_s2k_cipher_algo_13)
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
								v51 = v50
							} else {
								v41 = F_pg_strcasecmp(m, int32(_a_F_pgp_set_s2k_cipher_algo_14), l1)
								mBase = m.M
								if v41 == int32(0) {
									v49 = int32(_a_F_pgp_set_s2k_cipher_algo_15)
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
									v51 = v50
								} else {
									v46 = F_pg_strcasecmp(m, int32(_a_F_pgp_set_s2k_cipher_algo_16), l1)
									mBase = m.M
									if v46 != 0 {
										v51 = int32(-103)
									} else {
										v49 = int32(_a_F_pgp_set_s2k_cipher_algo_17)
										v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
										v51 = v50
									}
								}
							}
						}
					}
				}
			}
		}
	}
	if v51 < int32(0) {
		return v51
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v51
		return int32(0)
	}
}
func F_pgp_set_symkey(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v15 int32
	_ = v15
	v4 = int32(0)
	if base.B2i32(l1 == v4)|base.B2i32(l2 <= v4) != 0 {
		v15 = int32(-13)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = l1
		v15 = int32(0)
	}
	return v15
}
