package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_list_concat_unique_ptr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	v3 = int32(0)
	if l1 == v3 {
		v69 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v69
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v10 <= int32(0) {
		v69 = l0
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = l0
	v16 = v3
	goto L4
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20+v16<<(uint(int32(2))%32))))
	if v13 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v69 = v58
	goto L1
L6:
	;
	v66 = v16 + int32(1)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v66 < v67 {
		v13 = v58
		v16 = v66
		goto L4
	} else {
		goto L16
	}
L7:
	;
	v54 = F_lappend(m, v13, v24)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L14
	} else {
		goto L15
	}
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v27 <= int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v34 = int32(0)
	goto L10
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v30+v34<<(uint(int32(2))%32))))
	if v42 == v24 {
		v58 = v13
		goto L6
	} else {
		goto L12
	}
L11:
	;
	goto L7
L12:
	;
	v45 = v34 + int32(1)
	if v27 != v45 {
		v34 = v45
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	return int32(0)
L15:
	;
	v58 = v54
	goto L6
L16:
	;
	goto L5
}
func F_list_delete(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	v3 = int32(0)
	if l0 == v3 {
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
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v11 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return l0
L5:
	;
	v18 = v3
	goto L6
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v23 = v20 + v18<<(uint(int32(2))%32)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v25 = F_equal(m, v24, l1)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v29 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	goto L7
L9:
	;
	return int32(0)
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v25 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v31 = v18 + int32(1)
	if v31 < v29 {
		v18 = v31
		goto L6
	} else {
		goto L12
	}
L12:
	;
	goto L4
L13:
	;
	if l0+int32(16) != v33 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v48 = int32(2)
	v52 = (v29 + int32(base.Ui32(v23-v33^int32(-1))>>(uint(v48)%32))) << (uint(v48) % 32)
	if v52 != 0 {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	F_pfree(m, v33)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L9
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	F_pfree(m, l0)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L9
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	return int32(0)
L21:
	;
	base.MemoryCopy(m, v23, v23+int32(4), v52)
	goto L23
L22:
	;
	goto L23
L23:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v56 - int32(1)
	goto L4
}
func F_list_make2_impl(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v5 = F_palloc(m, int32(32))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v5)+4)) = int64(17179869186)
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v5 + int32(16)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v16
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		*(*int32)(unsafe.Add(mBase, uint32(v5)+20)) = v18
		return v5
	}
}
func F_list_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	v2 = int32(0)
	if l0 == v2 {
		v68 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v68
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v9 <= int32(0) {
		v68 = v2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = v2
	v16 = v2
	goto L4
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v21 = v18 + v16<<(uint(int32(2))%32)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v13 == int32(0) {
		v51 = v22
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v68 = v58
	goto L1
L6:
	;
	v64 = v16 + int32(1)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v64 < v65 {
		v13 = v58
		v16 = v64
		goto L4
	} else {
		goto L17
	}
L7:
	;
	v55 = F_lappend(m, v13, v51)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L12
	} else {
		goto L16
	}
L8:
	;
	v25 = int32(0)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v26 <= v25 {
		v51 = v22
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v32 = v25
	goto L10
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v32<<(uint(int32(2))%32))))
	v40 = F_equal(m, v39, v22)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v51 = v48
	goto L7
L12:
	;
	return int32(0)
L13:
	;
	if v40 != 0 {
		v58 = v13
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v45 = v32 + int32(1)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v45 < v46 {
		v32 = v45
		goto L10
	} else {
		goto L15
	}
L15:
	;
	goto L11
L16:
	;
	v58 = v55
	goto L6
L17:
	;
	goto L5
}
func F_writeListPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v177 int64
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	v6 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(_a_F_writeListPage_0)
	m.G0 = v17
	if l1 < v6 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v37 = int32(1)
	v38 = int32(_a_F_writeListPage_1)
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_writeListPage[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_writeListPage[0])) = v40 + v37
	v44 = int32(16)
	if l1 < int32(0) {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_writeListPage[1]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22+(l1^int32(-1))<<(uint(int32(2))%32))))
	v36 = v28
	goto L1
L3:
	;
	goto L4
L4:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_writeListPage[2]))
	v36 = v30 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	if int32(0) < l3 {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	F_PageInit(m, v62, int32(_a_F_writeListPage_2), int32(8))
	mBase = m.M
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+16)))
	v67 = v62 + v66
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v67)+6)) = uint16(v44)
	goto L5
L7:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_writeListPage[1]))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v48+(l1^int32(-1))<<(uint(int32(2))%32))))
	v62 = v54
	goto L6
L8:
	;
	goto L9
L9:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_writeListPage[2]))
	v62 = v56 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L6
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L19
	} else {
		goto L44
	}
L11:
	;
	v83 = v37
	v84 = v17 + int32(16)
	v85 = v6
	v87 = v6
	goto L14
L12:
	;
	v125 = v6
	goto L13
L13:
	;
	v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v36+v127))) = l4
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+16)))
	if l4 == int32(-1) {
		goto L23
	} else {
		goto L24
	}
L14:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l2+v85<<(uint(int32(2))%32))))
	v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v92)+6)))
	v95 = v93 & int32(_a_F_writeListPage_3)
	if v95 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v125 = v109
	goto L13
L16:
	;
	base.MemoryCopy(m, v84, v92, v95)
	goto L18
L17:
	;
	goto L18
L18:
	;
	v100 = F_PageAddItemExtended(m, v36, v92, v95, v83&int32(_a_F_writeListPage_4), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	if v100 == int32(0) {
		goto L10
	} else {
		goto L21
	}
L21:
	;
	v106 = int32(1)
	v109 = v95 + v87
	v111 = v85 + v106
	if v111 != l3 {
		v83 = v83 + v106
		v84 = v95 + v84
		v85 = v111
		v87 = v109
		goto L14
	} else {
		goto L22
	}
L22:
	;
	goto L15
L23:
	;
	v133 = v130 + v36
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133)+6)))
	v136 = v134 | int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v133)+6)) = uint16(v136)
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+16)))
	v141 = int32(1)
	v142 = v139
	goto L25
L24:
	;
	v141 = v6
	v142 = v130
	goto L25
L25:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v142+v36)+4)) = uint16(v141)
	F_MarkBufferDirty(m, l1)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L19
	} else {
		goto L26
	}
L26:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+118)))
	if v148 != int32(112) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+14)))
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+12)))
	v184 = v182 - v183
	v185 = int32(0)
	if v185 < v184 {
		goto L40
	} else {
		goto L41
	}
L28:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_writeListPage[3]))
	if v152 <= int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v155 != 0 {
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = l4
	F_XLogBeginInsert(m)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L19
	} else {
		goto L34
	}
L32:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v156 != 0 {
		goto L27
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v161 = int32(8)
	F_XLogRegisterData(m, v17+v161, v161)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L19
	} else {
		goto L35
	}
L35:
	;
	F_XLogRegisterBuffer(m, int32(0), l1, int32(6))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L19
	} else {
		goto L36
	}
L36:
	;
	F_XLogRegisterBufData(m, int32(0), v17+int32(16), v125)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L19
	} else {
		goto L37
	}
L37:
	;
	v177 = F_XLogInsert(m, int32(13), int32(112))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L19
	} else {
		goto L38
	}
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v36))) = base.I64_rotr(v177, int64(32))
	goto L27
L39:
	;
	F_UnlockReleaseBuffer(m, l1)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L19
	} else {
		goto L43
	}
L40:
	;
	v188 = v184
	goto L42
L41:
	;
	v188 = v185
	goto L42
L42:
	;
	goto L39
L43:
	;
	v191 = int32(_a_F_writeListPage_1)
	v193 = *(*int32)(unsafe.Add(mBase, _c_F_writeListPage[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_writeListPage[0])) = v193 - int32(1)
	m.G0 = v17 + int32(_a_F_writeListPage_0)
	return v188
L44:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v205 + int32(4)
	F_errmsg_internal(m, int32(_a_F_writeListPage_5), v17)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L19
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_writeListPage_6), int32(90), int32(_a_F_writeListPage_7))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L19
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
