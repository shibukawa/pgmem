package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_QTNBinary(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	F_check_stack_depth(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v7 != int32(2) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v10 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v14 = int32(0)
	goto L6
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16+v14<<(uint(int32(2))%32))))
	F_QTNBinary(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	if v25 < int32(3) {
		goto L3
	} else {
		goto L10
	}
L8:
	;
	v24 = v14 + int32(1)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v24 < v25 {
		v14 = v24
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	goto L11
L11:
	;
	v33 = F_palloc0(m, int32(24))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L3
L13:
	;
	v36 = F_palloc0(m, int32(12))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v36
	v40 = F_palloc0(m, int32(8))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+20)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v33)+4)) = int64(8589934593)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v50
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v54 | v56
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	*(*uint8)(unsafe.Add(mBase, uint32(v59))) = uint8(v61)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)) = uint8(v65)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v33
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v71 = int32(2)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v69+v70<<(uint(v71)%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v69)+4)) = v76
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v80 = v78 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v80
	if v71 < v80 {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	goto L12
}
func F_QTNTernary(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	F_check_stack_depth(m)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v11 != int32(2) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(0) < v14 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v19 = int32(0)
	goto L8
L6:
	;
	v39 = v14
	v44 = v10
	goto L7
L7:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	if base.B2i32(v45&int32(254) != int32(2))|base.B2i32(v39 <= int32(0)) != 0 {
		goto L3
	} else {
		goto L12
	}
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25+v19<<(uint(int32(2))%32))))
	F_QTNTernary(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = v34
	v44 = v36
	goto L7
L10:
	;
	v33 = v19 + int32(1)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v33 < v34 {
		v19 = v33
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v55 = int32(0)
	v56 = v39
	goto L13
L13:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v62 = int32(2)
	v63 = v55 << (uint(v62) % 32)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61+v63)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v67 != v62 {
		v123 = v55
		v124 = v56
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L3
L15:
	;
	v128 = v123 + int32(1)
	if v128 < v124 {
		v55 = v128
		v56 = v124
		goto L13
	} else {
		goto L33
	}
L16:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
	if v71 != v72 {
		v123 = v55
		v124 = v56
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	v77 = v56 + v74 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v77
	v81 = F_repalloc(m, v61, v77<<(uint(int32(2))%32))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v81
	if v56 != v55+int32(1) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v91 = (v56 + (v55 ^ int32(-1))) << (uint(int32(2)) % 32)
	if v91 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v103 = v81
	goto L21
L21:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	v106 = v104 << (uint(int32(2)) % 32)
	if v106 != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v92 = v81 + v63
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	base.MemoryCopy(m, v92+v93<<(uint(int32(2))%32), v92+int32(4), v91)
	goto L24
L23:
	;
	goto L24
L24:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v103 = v101
	goto L21
L25:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	base.MemoryCopy(m, v103+v63, v108, v106)
	goto L27
L26:
	;
	goto L27
L27:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+4)))
	if v112&int32(1) != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	F_pfree(m, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	F_pfree(m, v65)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v123 = v55 + v110 - int32(1)
	v124 = v122
	goto L15
L33:
	;
	goto L14
}
