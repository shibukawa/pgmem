package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_WalSndKill(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v3 = int32(_a_F_WalSndKill_0)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKill[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_WalSndKill[0])) = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v4)+76)) = int32(1)
	if v8 != 0 {
		F_s_lock(m, v4+int32(76), int32(_a_F_WalSndKill_1), int32(3023), int32(_a_F_WalSndKill_2))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v18 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+76)) = v18
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = v18
			return
		}
	} else {
		v18 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+76)) = v18
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = v18
		return
	}
}
func F_WalSndLastCycleHandler(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	*(*int32)(unsafe.Add(mBase, _c_F_WalSndLastCycleHandler[0])) = int32(1)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLastCycleHandler[1]))
	F_SetLatch(m, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F_WalSndPrepareWrite(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v37 int64
	_ = v37
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v69 int64
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	v5 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v7))) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v5
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	F_enlargeStringInfo(m, v14, int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		v21 = int32(119)
		*(*uint8)(unsafe.Add(mBase, uint32(v18+v19))) = uint8(v21)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v18 + int32(1)
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
		F_enlargeStringInfo(m, v26, int32(8))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
			v33 = int64(56)
			v35 = int64(65280)
			v37 = int64(40)
			v40 = int64(16711680)
			v42 = int64(24)
			v44 = int64(4278190080)
			v46 = int64(8)
			if l3 != 0 {
				v69 = l1<<(uint(v33)%64) | l1&v35<<(uint(v37)%64) | (l1&v40<<(uint(v42)%64) | l1&v44<<(uint(v46)%64)) | (int64(base.Ui64(l1)>>(uint(v46)%64))&v44 | int64(base.Ui64(l1)>>(uint(v42)%64))&v40 | (int64(base.Ui64(l1)>>(uint(v37)%64))&v35 | int64(base.Ui64(l1)>>(uint(v33)%64))))
			} else {
				v69 = int64(0)
			}
			*(*int64)(unsafe.Add(mBase, uint32(v30+v31))) = v69
			v71 = int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v30 + v71
			v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
			F_enlargeStringInfo(m, v74, v71)
			mBase = m.M
			v77 = m.ExcPending
			if v77 != 0 {
				return
			} else {
				v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
				v79 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
				*(*int64)(unsafe.Add(mBase, uint32(v78+v79))) = v69
				v82 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = v78 + v82
				v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
				F_enlargeStringInfo(m, v85, v82)
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return
				} else {
					v89 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
					v90 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
					*(*int64)(unsafe.Add(mBase, uint32(v89+v90))) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = v89 + int32(8)
					return
				}
			}
		}
	}
}
func F_check_wal_consistency_checking(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v196 int32
	_ = v196
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v298 int32
	_ = v298
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(288)
	m.G0 = v13
	base.MemoryFill(m, v13+int32(16), v4, int32(256))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = F_pstrdup(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(288)
	return v298
L2:
	;
	return int32(0)
L3:
	;
	v28 = F_SplitIdentifierString(m, v21, int32(44), v13+int32(284))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v28 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_check_wal_consistency_checking[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_wal_consistency_checking[1])) = v33
	goto L8
L6:
	;
	goto L7
L7:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v13)+284))
	if v47 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v39 = F_format_elog_string(m, int32(_a_F_check_wal_consistency_checking_0), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_wal_consistency_checking[2])) = v39
	F_pfree(m, v21)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v13)+284))
	F_list_free(m, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v298 = v4
	goto L1
L12:
	;
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_check_wal_consistency_checking[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_wal_consistency_checking[1])) = v275
	goto L70
L13:
	;
	F_pfree(m, v21)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L2
	} else {
		goto L66
	}
L14:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v50 <= int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v60 = v4
	goto L16
L16:
	;
	v63 = int32(0)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+v60<<(uint(int32(2))%32))))
	v72 = v68
	v73 = int32(_a_F_check_wal_consistency_checking_1)
	goto L20
L17:
	;
	goto L13
L18:
	;
	v245 = v60 + int32(1)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v245 < v246 {
		v60 = v245
		goto L16
	} else {
		goto L65
	}
L19:
	;
	if v110 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L20:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v76 == v77 {
		v99 = v76
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v110 = int32(0)
	goto L19
L22:
	;
	v101 = int32(1)
	if v99 != 0 {
		v72 = v72 + v101
		v73 = v73 + v101
		goto L20
	} else {
		goto L31
	}
L23:
	;
	if base.Ui32((v76-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v87 = v76 | int32(32)
	goto L26
L25:
	;
	v87 = v76
	goto L26
L26:
	;
	if base.Ui32((v77-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v96 = v77 | int32(32)
	goto L29
L28:
	;
	v96 = v77
	goto L29
L29:
	;
	if v87 == v96 {
		v99 = v87
		goto L22
	} else {
		goto L30
	}
L30:
	;
	v110 = v87 - v96
	goto L19
L31:
	;
	goto L21
L32:
	;
	v113 = v63
	goto L35
L33:
	;
	goto L34
L34:
	;
	v159 = v63
	goto L44
L35:
	;
	v124 = v113 << (uint(int32(5)) % 32)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+uint32(_c_F_check_wal_consistency_checking[3])))
	if v125 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v139 = v113 | int32(1)
	v141 = v139 << (uint(int32(5)) % 32)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+uint32(_c_F_check_wal_consistency_checking[3])))
	if v142 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v124)+uint32(_c_F_check_wal_consistency_checking[4])))
	if v130 == int32(0) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v136 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(16)+v113))) = uint8(v136)
	goto L37
L40:
	;
	v156 = v113 + int32(2)
	if v156 != int32(256) {
		v113 = v156
		goto L35
	} else {
		goto L43
	}
L41:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v141)+uint32(_c_F_check_wal_consistency_checking[4])))
	if v147 == int32(0) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v153 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(16)+v139))) = uint8(v153)
	goto L40
L43:
	;
	goto L18
L44:
	;
	v170 = v159 << (uint(int32(5)) % 32)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+uint32(_c_F_check_wal_consistency_checking[3])))
	if v171 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_wal_consistency_checking[5])))
	if v230 != 0 {
		goto L12
	} else {
		goto L64
	}
L46:
	;
	v226 = v159 + int32(1)
	if v226 != int32(256) {
		v159 = v226
		goto L44
	} else {
		goto L63
	}
L47:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v170)+uint32(_c_F_check_wal_consistency_checking[4])))
	if v176 == int32(0) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v181 = v68
	v182 = v171
	goto L50
L49:
	;
	if v219 != 0 {
		goto L46
	} else {
		goto L62
	}
L50:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181))))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	if v185 == v186 {
		v208 = v185
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v219 = int32(0)
	goto L49
L52:
	;
	v210 = int32(1)
	if v208 != 0 {
		v181 = v181 + v210
		v182 = v182 + v210
		goto L50
	} else {
		goto L61
	}
L53:
	;
	if base.Ui32((v185-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v196 = v185 | int32(32)
	goto L56
L55:
	;
	v196 = v185
	goto L56
L56:
	;
	if base.Ui32((v186-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v205 = v186 | int32(32)
	goto L59
L58:
	;
	v205 = v186
	goto L59
L59:
	;
	if v196 == v205 {
		v208 = v196
		goto L52
	} else {
		goto L60
	}
L60:
	;
	v219 = v196 - v205
	goto L49
L61:
	;
	goto L51
L62:
	;
	v223 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(16)+v159))) = uint8(v223)
	goto L18
L63:
	;
	goto L45
L64:
	;
	v232 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_check_wal_consistency_checking[6])) = uint8(v232)
	goto L18
L65:
	;
	goto L17
L66:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v13)+284))
	F_list_free(m, v260)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	v264 = F_guc_malloc(m, int32(256))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v264
	if v264 == int32(0) {
		v298 = v4
		goto L1
	} else {
		goto L69
	}
L69:
	;
	base.MemoryCopy(m, v264, v13+int32(16), int32(256))
	v298 = int32(1)
	goto L1
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v68
	v281 = F_format_elog_string(m, int32(_a_F_check_wal_consistency_checking_2), v13)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_wal_consistency_checking[2])) = v281
	F_pfree(m, v21)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v13)+284))
	F_list_free(m, v286)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	v298 = v4
	goto L1
}
func F_check_wal_segment_size(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = int32(0)
	v14 = base.B2i32(base.Ui32(v4+int32(-1048576)) < base.Ui32(int32(1072693249))) & base.B2i32(v4&(v4-int32(1)) == v12)
	if v14 == v12 {
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_check_wal_segment_size[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_check_wal_segment_size[1])) = v18
		v24 = F_format_elog_string(m, int32(_a_F_check_wal_segment_size_0), int32(0))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_check_wal_segment_size[2])) = v24
			return v14
		}
	} else {
		return v14
	}
}
func F_wal_segment_open(m *base.Module, l0 int32, l1 int64, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v19 int64
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v7 = m.G0
	v9 = v7 - int32(1072)
	m.G0 = v9
	v11 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v12
	v15 = base.I64_div_u_s(int64(4294967296), v11)
	v16 = base.I64_div_u_s(l1, v15)
	*(*uint32)(unsafe.Add(mBase, uint32(v9)+36)) = uint32(v16)
	v19 = l1 - v15*v16
	*(*uint32)(unsafe.Add(mBase, uint32(v9)+40)) = uint32(v19)
	v22 = v9 + int32(48)
	v27 = F_pg_snprintf(m, v22, int32(1024), int32(_a_F_wal_segment_open_0), v9+int32(32))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		return
	} else {
		v30 = F_BasicOpenFile(m, v22, int32(0))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+1168)) = v30
			if v30 < int32(0) {
				v36 = *(*int32)(unsafe.Add(mBase, _c_F_wal_segment_open[0]))
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					F_errcode_for_file_access(m)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						if v36 == int32(44) {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(48)
							F_errmsg(m, int32(_a_F_wal_segment_open_1), v9)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_wal_segment_open_2), int32(821), int32(_a_F_wal_segment_open_3))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v22
							F_errmsg(m, int32(_a_F_wal_segment_open_4), v9+int32(16))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_wal_segment_open_2), int32(826), int32(_a_F_wal_segment_open_3))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			} else {
				m.G0 = v9 + int32(1072)
				return
			}
		}
	}
}
