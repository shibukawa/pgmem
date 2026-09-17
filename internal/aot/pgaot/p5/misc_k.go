package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_KnownAssignedXidsDisplay(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsDisplay[0]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	F_initStringInfo(m, v9+int32(-16))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v16 < v15 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsDisplay[1]))
	v26 = v16
	v27 = v23
	v30 = v2
	goto L6
L4:
	;
	v66 = v2
	goto L5
L5:
	;
	v69 = F_errstart(m, l0, int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L13
	}
L6:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v27))))
	if v33 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v66 = v56
	goto L5
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsDisplay[2]))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37+v26<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v41
	F_appendStringInfo(m, v9+int32(-16), int32(_a_F_KnownAssignedXidsDisplay_0), v9+int32(-32))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v55 = v27
	v56 = v30
	goto L10
L10:
	;
	v58 = v26 + int32(1)
	if v58 != v15 {
		v26 = v58
		v27 = v55
		v30 = v56
		goto L6
	} else {
		goto L12
	}
L11:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsDisplay[1]))
	v55 = v54
	v56 = v30 + int32(1)
	goto L10
L12:
	;
	goto L7
L13:
	;
	if v69 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v71 = *(*int64)(unsafe.Add(mBase, uint32(v14)+16))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v72
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v71
	F_errmsg_internal(m, int32(_a_F_KnownAssignedXidsDisplay_1), v11)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	F_pfree(m, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	F_errfinish(m, int32(_a_F_KnownAssignedXidsDisplay_2), int32(_a_F_KnownAssignedXidsDisplay_3), int32(_a_F_KnownAssignedXidsDisplay_4))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	m.G0 = v11 - int32(-64)
	return
}
func F_KnownAssignedXidsRemoveTree(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_KnownAssignedXidsRemove(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	if int32(0) < l1 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v15 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsRemoveTree[0]))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v46 = v44 - v45
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	if v46 == v47 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l2+v15<<(uint(int32(2))%32))))
	F_KnownAssignedXidsRemove(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v31 = v15 + int32(1)
	if v31 != l1 {
		v15 = v31
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	return
L14:
	;
	v49 = int32(_a_F_KnownAssignedXidsRemoveTree_0)
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsRemoveTree[1]))
	v52 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsRemoveTree[1])) = v51 + v52
	if v51&int32(127)|base.B2i32(v46 < v47<<(uint(v52)%32)) != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v61 = int32(0)
	if v44 <= v45 {
		v156 = v61
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+20)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = int32(0)
	v171 = m.G0
	v172 = int32(16)
	v173 = v171 - v172
	m.G0 = v173
	F_gettimeofday(m, v173)
	mBase = m.M
	v176 = *(*int64)(unsafe.Add(mBase, uint32(v173)))
	v177 = int64(*(*int32)(unsafe.Add(mBase, uint32(v173)+8)))
	m.G0 = v173 + v172
	goto L32
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsRemoveTree[2]))
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsRemoveTree[3]))
	if v45+int32(1) != v44 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v75 = v45
	v76 = v61
	v78 = int32(0)
	goto L21
L19:
	;
	v129 = v45
	v130 = v61
	goto L20
L20:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129+v66))))
	if v139 != int32(1) {
		v156 = v130
		goto L16
	} else {
		goto L31
	}
L21:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75+v66))))
	if v85 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v46&int32(1) == int32(0) {
		v156 = v121
		goto L16
	} else {
		goto L30
	}
L23:
	;
	v88 = int32(2)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v64+v75<<(uint(v88)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v64+v76<<(uint(v88)%32)))) = v94
	v97 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v76+v66))) = uint8(v97)
	v101 = v76 + v97
	goto L25
L24:
	;
	v101 = v76
	goto L25
L25:
	;
	v102 = int32(1)
	v103 = v75 + v102
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66+v103))))
	if v105 == v102 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v108 = int32(2)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v64+v103<<(uint(v108)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v64+v101<<(uint(v108)%32)))) = v114
	v117 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v101+v66))) = uint8(v117)
	v121 = v101 + v117
	goto L28
L27:
	;
	v121 = v101
	goto L28
L28:
	;
	v122 = int32(2)
	v123 = v75 + v122
	v125 = v78 + v122
	if v125 != v46&int32(-2) {
		v75 = v123
		v76 = v121
		v78 = v125
		goto L21
	} else {
		goto L29
	}
L29:
	;
	goto L22
L30:
	;
	v129 = v123
	v130 = v121
	goto L20
L31:
	;
	v142 = int32(2)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v64+v129<<(uint(v142)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v64+v130<<(uint(v142)%32)))) = v148
	v151 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v130+v66))) = uint8(v151)
	v156 = v130 + v151
	goto L16
L32:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_KnownAssignedXidsRemoveTree[4])) = v177 + v176*int64(1000000) - int64(946684800000000)
	goto L13
}
func F_koi8r_to_utf8(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v3 = int32(0)
	v7 = Fn13848(m, l0, int32(22), v3, v3, v3, int32(_a_F_koi8r_to_utf8_0))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_koi8r_to_win1251(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13936(m, l0, int32(_a_F_koi8r_to_win1251_0), int32(23), int32(22))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
