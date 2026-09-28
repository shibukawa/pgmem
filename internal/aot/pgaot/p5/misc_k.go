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
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_KnownAssignedXidsRemove(m, l0)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
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
	v12 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsRemoveTree[0]))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	v37 = v35 - v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	if v37 == v38 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2+v12<<(uint(int32(2))%32))))
	F_KnownAssignedXidsRemove(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v25 = v12 + int32(1)
	if v25 != l1 {
		v12 = v25
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
	v40 = int32(_a_F_KnownAssignedXidsRemoveTree_0)
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsRemoveTree[1]))
	v43 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsRemoveTree[1])) = v42 + v43
	if v42&int32(127)|base.B2i32(v37 < v38<<(uint(v43)%32)) != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v52 = int32(0)
	if v36 < v35 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsRemoveTree[2]))
	v56 = v36
	v57 = v52
	v58 = v55
	goto L19
L17:
	;
	v91 = v52
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = int32(0)
	v103 = m.G0
	v104 = int32(16)
	v105 = v103 - v104
	m.G0 = v105
	F_gettimeofday(m, v105)
	mBase = m.M
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v105)))
	v109 = int64(*(*int32)(unsafe.Add(mBase, uint32(v105)+8)))
	m.G0 = v105 + v104
	goto L25
L19:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v58))))
	if v63 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v91 = v85
	goto L18
L21:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsRemoveTree[3]))
	v68 = int32(2)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v67+v56<<(uint(v68)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v67+v57<<(uint(v68)%32)))) = v74
	v76 = int32(_a_F_KnownAssignedXidsRemoveTree_1)
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsRemoveTree[2]))
	v79 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v77+v57))) = uint8(v79)
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsRemoveTree[2]))
	v85 = v57 + v79
	v86 = v82
	goto L23
L22:
	;
	v85 = v57
	v86 = v58
	goto L23
L23:
	;
	v88 = v56 + int32(1)
	if v88 != v35 {
		v56 = v88
		v57 = v85
		v58 = v86
		goto L19
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_KnownAssignedXidsRemoveTree[4])) = v109 + v108*int64(1000000) - int64(946684800000000)
	goto L13
}
func F_koi8r_to_utf8(m *base.Module, l0 int32) int64 {
	var v3 int32
	_ = v3
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v3 = int32(0)
	v7 = Fn14231(m, l0, int32(22), v3, v3, v3, int32(_a_F_koi8r_to_utf8_0))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_koi8r_to_win1251(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14323(m, l0, int32(_a_F_koi8r_to_win1251_0), int32(23), int32(22))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
