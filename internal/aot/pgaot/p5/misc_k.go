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
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
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
	v14 = *(*int32)(unsafe.Add(mBase, _consts[515]))
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
	v23 = *(*int32)(unsafe.Add(mBase, _consts[675]))
	v26 = v16
	v27 = v2
	v28 = v23
	goto L6
L4:
	;
	v63 = v2
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
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v28))))
	if v33 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v63 = v55
	goto L5
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[676]))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37+v26<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v41
	F_appendStringInfo(m, v9+int32(-16), int32(731558), v9+int32(-32))
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
	v56 = v28
	goto L10
L10:
	;
	v58 = v26 + int32(1)
	if v58 != v15 {
		v26 = v58
		v27 = v55
		v28 = v56
		goto L6
	} else {
		goto L12
	}
L11:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[675]))
	v55 = v27 + int32(1)
	v56 = v52
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
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v72
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v71
	F_errmsg_internal(m, int32(206142), v11)
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
	F_errfinish(m, int32(492693), int32(5246), int32(26488))
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
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int64
	_ = v175
	var v176 int64
	_ = v176
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
	v43 = *(*int32)(unsafe.Add(mBase, _consts[515]))
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
	v49 = int32(4432368)
	v51 = *(*int32)(unsafe.Add(mBase, _consts[677]))
	v52 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[677])) = v51 + v52
	if v46 < v47<<(uint(v52)%32) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	if v51&int32(127) != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v60 = int32(0)
	if v44 <= v45 {
		v155 = v60
		goto L17
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+20)) = v155
	*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = int32(0)
	v170 = m.G0
	v171 = int32(16)
	v172 = v170 - v171
	m.G0 = v172
	F___gettimeofday(m, v172)
	mBase = m.M
	v175 = *(*int64)(unsafe.Add(mBase, uint32(v172)))
	v176 = int64(*(*int32)(unsafe.Add(mBase, uint32(v172)+8)))
	m.G0 = v172 + v171
	goto L33
L18:
	;
	v62 = int32(1)
	v65 = *(*int32)(unsafe.Add(mBase, _consts[676]))
	v67 = *(*int32)(unsafe.Add(mBase, _consts[675]))
	if v45+v62 != v44 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v74 = v45
	v75 = v60
	v77 = int32(0)
	goto L22
L20:
	;
	v126 = v45
	v127 = v60
	goto L21
L21:
	;
	if v46&v62 == int32(0) {
		v155 = v127
		goto L17
	} else {
		goto L31
	}
L22:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74+v67))))
	if v84 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v126 = v122
	v127 = v120
	goto L21
L24:
	;
	v87 = int32(2)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v65+v74<<(uint(v87)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v65+v75<<(uint(v87)%32)))) = v93
	v96 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v75+v67))) = uint8(v96)
	v100 = v75 + v96
	goto L26
L25:
	;
	v100 = v75
	goto L26
L26:
	;
	v101 = int32(1)
	v102 = v74 + v101
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+v102))))
	if v104 == v101 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v107 = int32(2)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v65+v102<<(uint(v107)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v65+v100<<(uint(v107)%32)))) = v113
	v116 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v100+v67))) = uint8(v116)
	v120 = v100 + v116
	goto L29
L28:
	;
	v120 = v100
	goto L29
L29:
	;
	v121 = int32(2)
	v122 = v74 + v121
	v124 = v77 + v121
	if v124 != v46&int32(-2) {
		v74 = v122
		v75 = v120
		v77 = v124
		goto L22
	} else {
		goto L30
	}
L30:
	;
	goto L23
L31:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126+v67))))
	if v138 != int32(1) {
		v155 = v127
		goto L17
	} else {
		goto L32
	}
L32:
	;
	v141 = int32(2)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v65+v126<<(uint(v141)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v65+v127<<(uint(v141)%32)))) = v147
	v150 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v127+v67))) = uint8(v150)
	v155 = v127 + v150
	goto L17
L33:
	;
	*(*int64)(unsafe.Add(mBase, _consts[678])) = v176 + v175*int64(1000000) - int64(946684800000000)
	goto L13
}
func F_koi8r_to_utf8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(22), int32(6))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		v24 = F_LocalToUtf(m, v6, v10, v5, int32(4389092), v18, v18, v18, int32(22), base.B2i32(v7 != v18))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			return v24
		}
	}
}
func F_koi8r_to_win1251(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(22), int32(23))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v22 = F_local2local(m, v6, v5, v10, int32(22), int32(23), int32(2229600), base.B2i32(v7 != int32(0)))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			return v22
		}
	}
}
