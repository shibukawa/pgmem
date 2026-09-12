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
		v8 = int32(6867)
	} else {
		v8 = int32(6868)
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
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v191 int32
	_ = v191
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v216 int32
	_ = v216
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v264 int32
	_ = v264
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v287 int32
	_ = v287
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v313 int32
	_ = v313
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	v9 = int32(557448)
	v10 = l0
	goto L4
L1:
	;
	return v342
L2:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v340)+4))
	v342 = v341
	goto L1
L3:
	;
	if v47 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L4:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v13 == v14 {
		v36 = v13
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v47 = int32(0)
	goto L3
L6:
	;
	v38 = int32(1)
	if v36 != 0 {
		v9 = v9 + v38
		v10 = v10 + v38
		goto L4
	} else {
		goto L15
	}
L7:
	;
	if base.Ui32((v13-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v24 = v13 | int32(32)
	goto L10
L9:
	;
	v24 = v13
	goto L10
L10:
	;
	if base.Ui32((v14-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v33 = v14 | int32(32)
	goto L13
L12:
	;
	v33 = v14
	goto L13
L13:
	;
	if v24 == v33 {
		v36 = v24
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v47 = v24 - v33
	goto L3
L15:
	;
	goto L5
L16:
	;
	v340 = int32(4394960)
	goto L2
L17:
	;
	goto L18
L18:
	;
	v57 = int32(561770)
	v58 = l0
	goto L20
L19:
	;
	if v95 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L20:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v61 == v62 {
		v84 = v61
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v95 = int32(0)
	goto L19
L22:
	;
	v86 = int32(1)
	if v84 != 0 {
		v57 = v57 + v86
		v58 = v58 + v86
		goto L20
	} else {
		goto L31
	}
L23:
	;
	if base.Ui32((v61-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v72 = v61 | int32(32)
	goto L26
L25:
	;
	v72 = v61
	goto L26
L26:
	;
	if base.Ui32((v62-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v81 = v62 | int32(32)
	goto L29
L28:
	;
	v81 = v62
	goto L29
L29:
	;
	if v72 == v81 {
		v84 = v72
		goto L22
	} else {
		goto L30
	}
L30:
	;
	v95 = v72 - v81
	goto L19
L31:
	;
	goto L21
L32:
	;
	v340 = int32(4394968)
	goto L2
L33:
	;
	goto L34
L34:
	;
	v105 = int32(562191)
	v106 = l0
	goto L36
L35:
	;
	if v143 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L36:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	if v109 == v110 {
		v132 = v109
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v143 = int32(0)
	goto L35
L38:
	;
	v134 = int32(1)
	if v132 != 0 {
		v105 = v105 + v134
		v106 = v106 + v134
		goto L36
	} else {
		goto L47
	}
L39:
	;
	if base.Ui32((v109-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v120 = v109 | int32(32)
	goto L42
L41:
	;
	v120 = v109
	goto L42
L42:
	;
	if base.Ui32((v110-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v129 = v110 | int32(32)
	goto L45
L44:
	;
	v129 = v110
	goto L45
L45:
	;
	if v120 == v129 {
		v132 = v120
		goto L38
	} else {
		goto L46
	}
L46:
	;
	v143 = v120 - v129
	goto L35
L47:
	;
	goto L37
L48:
	;
	v340 = int32(4394976)
	goto L2
L49:
	;
	goto L50
L50:
	;
	v153 = int32(563148)
	v154 = l0
	goto L52
L51:
	;
	if v191 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L52:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	if v157 == v158 {
		v180 = v157
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v191 = int32(0)
	goto L51
L54:
	;
	v182 = int32(1)
	if v180 != 0 {
		v153 = v153 + v182
		v154 = v154 + v182
		goto L52
	} else {
		goto L63
	}
L55:
	;
	if base.Ui32((v157-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v168 = v157 | int32(32)
	goto L58
L57:
	;
	v168 = v157
	goto L58
L58:
	;
	if base.Ui32((v158-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v177 = v158 | int32(32)
	goto L61
L60:
	;
	v177 = v158
	goto L61
L61:
	;
	if v168 == v177 {
		v180 = v168
		goto L54
	} else {
		goto L62
	}
L62:
	;
	v191 = v168 - v177
	goto L51
L63:
	;
	goto L53
L64:
	;
	v340 = int32(4394984)
	goto L2
L65:
	;
	goto L66
L66:
	;
	v201 = int32(557034)
	v202 = l0
	goto L68
L67:
	;
	if v239 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L68:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	if v205 == v206 {
		v228 = v205
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v239 = int32(0)
	goto L67
L70:
	;
	v230 = int32(1)
	if v228 != 0 {
		v201 = v201 + v230
		v202 = v202 + v230
		goto L68
	} else {
		goto L79
	}
L71:
	;
	if base.Ui32((v205-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v216 = v205 | int32(32)
	goto L74
L73:
	;
	v216 = v205
	goto L74
L74:
	;
	if base.Ui32((v206-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v225 = v206 | int32(32)
	goto L77
L76:
	;
	v225 = v206
	goto L77
L77:
	;
	if v216 == v225 {
		v228 = v216
		goto L70
	} else {
		goto L78
	}
L78:
	;
	v239 = v216 - v225
	goto L67
L79:
	;
	goto L69
L80:
	;
	v340 = int32(4394992)
	goto L2
L81:
	;
	goto L82
L82:
	;
	v249 = int32(558928)
	v250 = l0
	goto L84
L83:
	;
	if v287 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L84:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	if v253 == v254 {
		v276 = v253
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v287 = int32(0)
	goto L83
L86:
	;
	v278 = int32(1)
	if v276 != 0 {
		v249 = v249 + v278
		v250 = v250 + v278
		goto L84
	} else {
		goto L95
	}
L87:
	;
	if base.Ui32((v253-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v264 = v253 | int32(32)
	goto L90
L89:
	;
	v264 = v253
	goto L90
L90:
	;
	if base.Ui32((v254-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v273 = v254 | int32(32)
	goto L93
L92:
	;
	v273 = v254
	goto L93
L93:
	;
	if v264 == v273 {
		v276 = v264
		goto L86
	} else {
		goto L94
	}
L94:
	;
	v287 = v264 - v273
	goto L83
L95:
	;
	goto L85
L96:
	;
	v340 = int32(4395000)
	goto L2
L97:
	;
	goto L98
L98:
	;
	v298 = int32(561224)
	v299 = l0
	goto L100
L99:
	;
	if v336 != 0 {
		v342 = int32(-104)
		goto L1
	} else {
		goto L112
	}
L100:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298))))
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299))))
	if v302 == v303 {
		v325 = v302
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v336 = int32(0)
	goto L99
L102:
	;
	v327 = int32(1)
	if v325 != 0 {
		v298 = v298 + v327
		v299 = v299 + v327
		goto L100
	} else {
		goto L111
	}
L103:
	;
	if base.Ui32((v302-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v313 = v302 | int32(32)
	goto L106
L105:
	;
	v313 = v302
	goto L106
L106:
	;
	if base.Ui32((v303-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v322 = v303 | int32(32)
	goto L109
L108:
	;
	v322 = v303
	goto L109
L109:
	;
	if v313 == v322 {
		v325 = v313
		goto L102
	} else {
		goto L110
	}
L110:
	;
	v336 = v313 - v322
	goto L99
L111:
	;
	goto L101
L112:
	;
	v340 = int32(4395008)
	goto L2
}
func F_pgp_get_digest_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v4 = l0 - int32(1)
	if base.Ui32(v4) <= base.Ui32(int32(9)) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v4<<(uint(int32(2))%32))+uint32(_consts[1433])))
		v12 = v11
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
	var v9 int32
	_ = v9
	if base.Ui32(l1) <= base.Ui32(int32(3)) {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = l1
		v9 = int32(0)
	} else {
		v9 = int32(-13)
	}
	return v9
}
func F_pgp_set_s2k_cipher_algo(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	v8 = F_pg_strcasecmp(m, int32(171408), l1)
	mBase = m.M
	if v8 == int32(0) {
		v77 = int32(4395024)
		v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
		v79 = v78
	} else {
		v16 = F_pg_strcasecmp(m, int32(557283), l1)
		mBase = m.M
		if v16 == int32(0) {
			v77 = int32(4395044)
			v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
			v79 = v78
		} else {
			v24 = F_pg_strcasecmp(m, int32(340048), l1)
			mBase = m.M
			if v24 == int32(0) {
				v77 = int32(4395064)
				v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
				v79 = v78
			} else {
				v32 = F_pg_strcasecmp(m, int32(322871), l1)
				mBase = m.M
				if v32 == int32(0) {
					v77 = int32(4395084)
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
					v79 = v78
				} else {
					v40 = F_pg_strcasecmp(m, int32(171908), l1)
					mBase = m.M
					if v40 == int32(0) {
						v77 = int32(4395104)
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
						v79 = v78
					} else {
						v48 = F_pg_strcasecmp(m, int32(556478), l1)
						mBase = m.M
						if v48 == int32(0) {
							v77 = int32(4395124)
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
							v79 = v78
						} else {
							v56 = F_pg_strcasecmp(m, int32(561053), l1)
							mBase = m.M
							if v56 == int32(0) {
								v77 = int32(4395144)
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
								v79 = v78
							} else {
								v64 = F_pg_strcasecmp(m, int32(556999), l1)
								mBase = m.M
								if v64 == int32(0) {
									v77 = int32(4395164)
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
									v79 = v78
								} else {
									v73 = F_pg_strcasecmp(m, int32(322880), l1)
									mBase = m.M
									if v73 != 0 {
										v79 = int32(-103)
									} else {
										v77 = int32(4395184)
										v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
										v79 = v78
									}
								}
							}
						}
					}
				}
			}
		}
	}
	if v79 < int32(0) {
		return v79
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v79
		return int32(0)
	}
}
func F_pgp_set_symkey(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	v5 = int32(-13)
	if l1 == int32(0) {
		v13 = v5
	} else {
		if l2 <= int32(0) {
			v13 = v5
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = l1
			v13 = int32(0)
		}
	}
	return v13
}
