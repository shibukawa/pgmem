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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
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
	v41 = F_palloc0_mul(m, int32(4), int32(2))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+20)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v33)+4)) = int64(8589934593)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v47
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = v51
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v55 | v57
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	*(*uint8)(unsafe.Add(mBase, uint32(v60))) = uint8(v62)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)) = uint8(v66)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = v33
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v72 = int32(2)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v70+v71<<(uint(v72)%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v77
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v81 = v79 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v81
	if v72 < v81 {
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
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
		v122 = v55
		v123 = v56
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L3
L15:
	;
	v127 = v122 + int32(1)
	if v127 < v123 {
		v55 = v127
		v56 = v123
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
		v122 = v55
		v123 = v56
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	v77 = v56 + v74 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v77
	v80 = F_repalloc_mul(m, v61, int32(4), v77)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v80
	if v56 != v55+int32(1) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v90 = (v56 + (v55 ^ int32(-1))) << (uint(int32(2)) % 32)
	if v90 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v102 = v80
	goto L21
L21:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	v105 = v103 << (uint(int32(2)) % 32)
	if v105 != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v91 = v80 + v63
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	base.MemoryCopy(m, v91+v92<<(uint(int32(2))%32), v91+int32(4), v90)
	goto L24
L23:
	;
	goto L24
L24:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v102 = v100
	goto L21
L25:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	base.MemoryCopy(m, v102+v63, v107, v105)
	goto L27
L26:
	;
	goto L27
L27:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+4)))
	if v111&int32(1) != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	F_pfree(m, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
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
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v122 = v55 + v109 - int32(1)
	v123 = v121
	goto L15
L33:
	;
	goto L14
}
