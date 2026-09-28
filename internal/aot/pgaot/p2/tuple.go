package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateTupleDescTruncatedCopy(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int64
	_ = v22
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	v3 = int32(0)
	v15 = F_palloc(m, l1*int32(108)+int32(28))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l1
	v22 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(2249)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v22
	if l1 <= v19 {
		v75 = l1
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v86
	if int32(0) < v75 {
		goto L13
	} else {
		goto L14
	}
L4:
	;
	v31 = l1 * int32(100)
	if v31 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v32 = int32(3)
	v35 = int32(28)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	base.MemoryCopy(m, v15+l1<<(uint(v32)%32)+v35, l0+v37<<(uint(v32)%32)+v35, v31)
	goto L7
L6:
	;
	goto L7
L7:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v44 <= int32(0) {
		v75 = v44
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v49 = v44
	v51 = int32(0)
	goto L9
L9:
	;
	v63 = v15 + v49<<(uint(int32(3))%32) + v51*int32(100)
	v64 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+118)) = uint8(v64)
	*(*int32)(unsafe.Add(mBase, uint32(v63)+114)) = v64
	F_populate_compact_attribute(m, v15, v51)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v75 = v72
	goto L3
L11:
	;
	v71 = v51 + int32(1)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v71 < v72 {
		v49 = v72
		v51 = v71
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v91 = v15 + int32(28)
	v95 = v75
	v100 = v3
	v104 = v3
	goto L17
L14:
	;
	v157 = v3
	v162 = v75
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v157
	return v15
L16:
	;
	v157 = v150
	v162 = v128
	goto L15
L17:
	;
	v107 = v91 + v75<<(uint(int32(3))%32) + v100*int32(100)
	v110 = v91 + v100<<(uint(int32(3))%32)
	if v95 != v75 {
		v128 = v95
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v150 = v75
	goto L16
L19:
	;
	v129 = int32(*(*int16)(unsafe.Add(mBase, uint32(v110)+2)))
	if v129 <= int32(0) {
		v150 = v100
		goto L16
	} else {
		goto L27
	}
L20:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+7)))
	if v112 != int32(118) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v128 = v100
	goto L19
L22:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+4)))
	if v115 != int32(1) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+6)))
	if v118&int32(6) != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v110)+2)))
	if v121 <= int32(0) {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+90)))
	if v124 != int32(118) {
		v128 = v75
		goto L19
	} else {
		goto L26
	}
L26:
	;
	goto L21
L27:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+90)))
	if v132 == int32(118) {
		v150 = v100
		goto L16
	} else {
		goto L28
	}
L28:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+5)))
	v141 = (v104 + v135 - int32(1)) & (int32(0) - v135)
	if int32(_a_F_CreateTupleDescTruncatedCopy_0) < v141 {
		v150 = v100
		goto L16
	} else {
		goto L29
	}
L29:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v110))) = uint16(v141)
	v147 = v100 + int32(1)
	if v147 != v75 {
		v95 = v128
		v100 = v147
		v104 = v141 + v129
		goto L17
	} else {
		goto L30
	}
L30:
	;
	goto L18
}
func F_CreateTupleQueueDestReceiver(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_palloc0(m, int32(24))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+20)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = int32(11)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = int32(826)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = int32(827)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = int32(828)
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(829)
		return v4
	}
}
func F_ExecInitScanTupleSlot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	v6 = F_MakeTupleTableSlot(m, l2, l3, l4)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
		v9 = F_lappend(m, v8, v6)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v9
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+96)) = uint8(base.B2i32(l2 != int32(0)))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+76)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(l1)+112)) = v6
			v17 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+100)) = uint8(v17)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = l3
			return
		}
	}
}
func F_GetTupleTransactionInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v15 = m.T0[v14].(func(*base.Module, int32, int32, int32) int64)(m, l0, int32(-2), v8+int32(15))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = base.I32_wrap_i64(v15)
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v19
		v21 = int32(0)
		v23 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetTupleTransactionInfo[0])))
		if v23 == v21 {
			v26 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v26)
			*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(0)
			v32 = v21
			m.G0 = v8 + int32(16)
			return v32
		} else {
			v30 = F_TransactionIdGetCommitTsData(m, v19, l3, l2)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				v32 = v30
				m.G0 = v8 + int32(16)
				return v32
			}
		}
	}
}
func F_TupleDescGetAttInMetadata(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	v11 = m.G0
	v12 = int32(16)
	v13 = v11 - v12
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = F_palloc(m, v12)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v21 != int32(2249) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = l0
	v32 = F_palloc0(m, v15*int32(28))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L7
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(0) <= v24 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	F_assign_record_type_typmod(m, l0)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	goto L3
L7:
	;
	v35 = v15 << (uint(int32(2)) % 32)
	v36 = F_palloc0(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v38 = F_palloc0(m, v35)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if int32(0) < v15 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v44 = int32(0)
	goto L13
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v32
	m.G0 = v13 + int32(16)
	return v17
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v59 = l0 + v53<<(uint(int32(3))%32) + v44*int32(100)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+119)))
	if v60 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L12
L15:
	;
	v64 = v59 + int32(28)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+68))
	v69 = v44 << (uint(int32(2)) % 32)
	F_getTypeInputInfo(m, v65, v13+int32(12), v36+v69)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v85 = v44 + int32(1)
	if v85 != v15 {
		v44 = v85
		goto L13
	} else {
		goto L20
	}
L18:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_fmgr_info(m, v73, v32+v44*int32(28))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v64)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v38+v69))) = v80
	goto L17
L20:
	;
	goto L14
}
func F_TupleDescGetDefault(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v6 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+12)))
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v16 = int32(0)
	goto L8
L5:
	;
	v42 = int32(0)
	goto L6
L6:
	;
	return v42
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v32 = F_stringToNode(m, v31)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	v23 = v13 + v16<<(uint(int32(3))%32)
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23))))
	if v24 == l1&int32(_a_F_TupleDescGetDefault_0) {
		goto L7
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v27 = v16 + int32(1)
	if v27 != v12 {
		v16 = v27
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	return int32(0)
L13:
	;
	v42 = v32
	goto L6
}
func F_TupleDescInitEntryCollation(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0+v4<<(uint(int32(3))%32)+l1*int32(100))+24)) = l2
	return
}
func F_tuple_data_split(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int64
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int64
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v434 int64
	_ = v434
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	v2 = int32(0)
	v20 = m.G0
	v22 = v20 + int32(-64)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v25 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v29 = F_pg_detoast_datum(m, v28)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v33 = v2
	goto L3
L3:
	;
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+72)))
	v35 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)))
	if v36 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int64(0)
L5:
	;
	v33 = v29
	goto L3
L6:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v40 = F_pg_detoast_datum_packed(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	v44 = v2
	goto L8
L8:
	;
	v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if int32(6) <= v45 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v42 = F_text_to_cstring(m, v40)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v44 = v42
	goto L8
L11:
	;
	v48 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	v51 = base.B2i32(v48 != int64(0))
	goto L13
L12:
	;
	v51 = v2
	goto L13
L13:
	;
	v52 = F_superuser(m)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L21
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L4
	} else {
		goto L161
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L4
	} else {
		goto L157
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L4
	} else {
		goto L153
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L4
	} else {
		goto L149
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L4
	} else {
		goto L144
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L4
	} else {
		goto L140
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L4
	} else {
		goto L136
	}
L21:
	;
	if v52 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if v33 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	goto L24
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L4
	} else {
		goto L132
	}
L25:
	;
	m.G0 = v22 - int32(-64)
	return v434
L26:
	;
	v56 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v56)
	v434 = int64(0)
	goto L25
L27:
	;
	goto L28
L28:
	;
	v62 = base.B2i32(v35&int64(1) == int64(0))
	if v62 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v146 = F_relation_open(m, v24, int32(1))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L48
	}
L30:
	;
	if v44 == int32(0) {
		goto L20
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if v44 != 0 {
		goto L17
	} else {
		goto L47
	}
L33:
	;
	v67 = F_strlen(m, v44)
	mBase = m.M
	v73 = (v34&int32(2047) + int32(7)) & int32(4088)
	if v67 != v73 {
		goto L19
	} else {
		goto L34
	}
L34:
	;
	v77 = F_palloc(m, v73|int32(1))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	if v73 == int32(0) {
		v135 = v77
		goto L29
	} else {
		goto L36
	}
L36:
	;
	v82 = int32(0)
	v84 = v2
	goto L37
L37:
	;
	v101 = v82 + v44
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	if v102&int32(254) != int32(48) {
		goto L18
	} else {
		goto L39
	}
L38:
	;
	v135 = v77
	goto L29
L39:
	;
	v110 = v82 & int32(7)
	if v110 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v114 = base.I32_extend8_s(v84)
	goto L42
L41:
	;
	v114 = int32(0)
	goto L42
L42:
	;
	v115 = (v102-int32(48))<<(uint(v110)%32) | v114
	if v110 == int32(7) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v77+int32(base.Ui32(v82)>>(uint(int32(3))%32))))) = uint8(v115)
	goto L45
L44:
	;
	goto L45
L45:
	;
	v123 = v82 + int32(1)
	if v123 != v73 {
		v82 = v123
		v84 = v115
		goto L37
	} else {
		goto L46
	}
L46:
	;
	goto L38
L47:
	;
	v135 = v2
	goto L29
L48:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v146)+52))
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_tuple_data_split[0]))
	v153 = F_initArrayResult(m, int32(17), v151, int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v146)+48))
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+119)))
	if v157 != int32(83) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v156)+84))
	if v160 != int32(2) {
		goto L16
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v164 = v34 & int32(2047)
	if v164 <= v155 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	goto L52
L54:
	;
	if v387 != v393 {
		goto L14
	} else {
		goto L127
	}
L55:
	;
	v190 = v33 + int32(4)
	v192 = v169 & int32(_a_F_tuple_data_split_0)
	v197 = int32(0)
	v199 = v197
	v201 = v197
	v206 = v153
	goto L64
L56:
	;
	v169 = int32(base.Ui32(v144)>>(uint(int32(2))%32)) - int32(4)
	if v155 != 0 {
		goto L55
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L4
	} else {
		goto L60
	}
L59:
	;
	v387 = int32(0)
	v392 = v153
	v393 = v169 & int32(_a_F_tuple_data_split_0)
	goto L54
L60:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	F_errmsg(m, int32(_a_F_tuple_data_split_1), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_tuple_data_split_2), int32(336), int32(_a_F_tuple_data_split_3))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	if base.Ui32(v164) <= base.Ui32(v199) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v387 = v369
	v392 = v378
	v393 = v192
	goto L54
L66:
	;
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_tuple_data_split[0]))
	v378 = F_accumArrayResult(m, v206, base.I64_extend_i32_u(v370), v368, int32(17), v377)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L4
	} else {
		goto L121
	}
L67:
	;
	v368 = int32(1)
	v369 = v201
	v370 = int32(0)
	goto L66
L68:
	;
	goto L69
L69:
	;
	if v35&int64(1) == int64(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v235 = v148 + int32(28) + v199<<(uint(int32(3))%32)
	v236 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v235)+2)))
	v238 = base.B2i32(v236 != int32(_a_F_tuple_data_split_0))
	if v238 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	v221 = int32(1)
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v199)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v225)>>(uint(v199&int32(7))%32))&v221 != 0 {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v368 = v221
	v369 = v201
	v370 = int32(0)
	goto L66
L73:
	;
	if v192 < v307+v304 {
		goto L15
	} else {
		goto L96
	}
L74:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201+v190))))
	if v242 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+5)))
	v304 = (v201 + v295 - int32(1)) & (int32(0) - v295)
	v307 = base.I32_extend16_s(v236)
	goto L73
L77:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+5)))
	v251 = (v201 + v245 - int32(1)) & (int32(0) - v245)
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190+v251))))
	v254 = v253
	v255 = v251
	goto L79
L78:
	;
	v254 = v242
	v255 = v201
	goto L79
L79:
	;
	v256 = v255 + v190
	v258 = v254 & int32(255)
	if v258 == int32(1) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+1)))
	v263 = v261 - int32(1)
	if v263 != 0 {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	goto L82
L82:
	;
	v288 = int32(1)
	if v254&v288 != 0 {
		v304 = v255
		v307 = int32(base.Ui32(v258) >> (uint(v288) % 32))
		goto L73
	} else {
		goto L95
	}
L83:
	;
	if base.Ui32(v261) < base.Ui32(int32(4)) {
		goto L92
	} else {
		goto L93
	}
L84:
	;
	if v263 == int32(17) {
		v283 = v261
		goto L83
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v283 = int32(2)
	goto L83
L87:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v199
	F_errmsg(m, int32(_a_F_tuple_data_split_4), v22)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_tuple_data_split_2), int32(376), int32(_a_F_tuple_data_split_3))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	v287 = int32(6)
	goto L94
L93:
	;
	v287 = v283
	goto L94
L94:
	;
	v304 = v255
	v307 = v287
	goto L73
L95:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v304 = v255
	v307 = int32(base.Ui32(v292) >> (uint(int32(2)) % 32))
	goto L73
L96:
	;
	if v238|(v51^int32(1)) == int32(0) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v331 = int32(*(*int16)(unsafe.Add(mBase, uint32(v235)+2)))
	if int32(0) < v331 {
		v364 = v331
		goto L104
	} else {
		goto L105
	}
L98:
	;
	v314 = F_pg_detoast_datum_copy(m, v304+v190)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L4
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v317 = v307 + int32(4)
	v318 = F_palloc(m, v317)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L4
	} else {
		goto L102
	}
L101:
	;
	v329 = v314
	goto L97
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v318))) = v317 << (uint(int32(2)) % 32)
	if v307 == int32(0) {
		v329 = v318
		goto L97
	} else {
		goto L103
	}
L103:
	;
	base.MemoryCopy(m, v318+int32(4), v304+v190, v307)
	v329 = v318
	goto L97
L104:
	;
	v368 = int32(0)
	v369 = v364 + v304
	v370 = v329
	goto L66
L105:
	;
	v334 = v304 + v190
	if v331 == int32(-1) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334))))
	if v337 == int32(1) {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	goto L108
L108:
	;
	v361 = F_strlen(m, v334)
	mBase = m.M
	v364 = v361 + int32(1)
	goto L104
L109:
	;
	v341 = int32(18)
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334)+1)))
	if v343 == v341 {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	goto L111
L111:
	;
	if v337&int32(1) != 0 {
		goto L118
	} else {
		goto L119
	}
L112:
	;
	v346 = v341
	goto L114
L113:
	;
	v346 = int32(2)
	goto L114
L114:
	;
	if base.Ui32((v343-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v353 = int32(6)
	goto L117
L116:
	;
	v353 = v346
	goto L117
L117:
	;
	v364 = v353
	goto L104
L118:
	;
	v364 = int32(base.Ui32(v337) >> (uint(int32(1)) % 32))
	goto L104
L119:
	;
	goto L120
L120:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v334)))
	v364 = int32(base.Ui32(v358) >> (uint(int32(2)) % 32))
	goto L104
L121:
	;
	if v370 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	F_pfree(m, v370)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L4
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v383 = v199 + int32(1)
	if v383 != v155 {
		v199 = v383
		v201 = v369
		v206 = v378
		goto L64
	} else {
		goto L126
	}
L125:
	;
	goto L124
L126:
	;
	goto L65
L127:
	;
	F_relation_close(m, v146, int32(1))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L4
	} else {
		goto L128
	}
L128:
	;
	v409 = *(*int32)(unsafe.Add(mBase, _c_F_tuple_data_split[0]))
	v410 = F_makeArrayResult(m, v392, v409)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L4
	} else {
		goto L129
	}
L129:
	;
	if v135 == int32(0) {
		v434 = v410
		goto L25
	} else {
		goto L130
	}
L130:
	;
	F_pfree(m, v135)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L4
	} else {
		goto L131
	}
L131:
	;
	v434 = v410
	goto L25
L132:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	F_errmsg(m, int32(_a_F_tuple_data_split_5), int32(0))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_tuple_data_split_2), int32(453), int32(_a_F_tuple_data_split_6))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	F_errmsg(m, int32(_a_F_tuple_data_split_7), int32(0))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L4
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(_a_F_tuple_data_split_2), int32(471), int32(_a_F_tuple_data_split_6))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L4
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L4
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v67
	F_errmsg(m, int32(_a_F_tuple_data_split_8), v20+int32(-16))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L4
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F_tuple_data_split_2), int32(478), int32(_a_F_tuple_data_split_6))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L4
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L144:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	v497 = F_pg_mblen_cstr(m, v101)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L4
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v497
	F_errmsg(m, int32(_a_F_tuple_data_split_9), v20+int32(-32))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L4
	} else {
		goto L147
	}
L147:
	;
	F_errfinish(m, int32(_a_F_tuple_data_split_2), int32(105), int32(_a_F_tuple_data_split_10))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L4
	} else {
		goto L148
	}
L148:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L149:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	v518 = F_strlen(m, v44)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v518
	F_errmsg(m, int32(_a_F_tuple_data_split_11), v20+int32(-48))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L4
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_tuple_data_split_2), int32(489), int32(_a_F_tuple_data_split_6))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L4
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	F_errmsg(m, int32(_a_F_tuple_data_split_12), int32(0))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L4
	} else {
		goto L155
	}
L155:
	;
	F_errfinish(m, int32(_a_F_tuple_data_split_2), int32(331), int32(_a_F_tuple_data_split_3))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L4
	} else {
		goto L156
	}
L156:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L157:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L4
	} else {
		goto L158
	}
L158:
	;
	F_errmsg(m, int32(_a_F_tuple_data_split_13), int32(0))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L4
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(_a_F_tuple_data_split_2), int32(389), int32(_a_F_tuple_data_split_3))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L4
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L161:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L4
	} else {
		goto L162
	}
L162:
	;
	F_errmsg(m, int32(_a_F_tuple_data_split_14), int32(0))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L4
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(_a_F_tuple_data_split_2), int32(413), int32(_a_F_tuple_data_split_3))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L4
	} else {
		goto L164
	}
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
