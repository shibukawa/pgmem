package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dataFindChildPtr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+v8)+4)))
	v12 = l3 - int32(1)
	if base.Ui32(v12&int32(_a_F_dataFindChildPtr_0)) < base.Ui32(v10) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v93 & int32(_a_F_dataFindChildPtr_0)
L2:
	;
	v18 = l3*int32(10) + l1
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+22)))
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+24)))
	if v19<<(uint(int32(16))%32)|v22 == l2 {
		v93 = l3
		goto L1
	} else {
		goto L5
	}
L3:
	;
	v54 = v10
	goto L4
L4:
	;
	v57 = int32(0)
	if v54&int32(_a_F_dataFindChildPtr_0) == v57 {
		v93 = v57
		goto L1
	} else {
		goto L12
	}
L5:
	;
	v30 = l3
	goto L6
L6:
	;
	v35 = v30 + int32(1)
	v36 = int32(_a_F_dataFindChildPtr_0)
	v37 = v35 & v36
	if base.Ui32(v37) <= base.Ui32(v10&v36) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v54 = v12
	goto L4
L8:
	;
	v43 = l1 + int32(22) + v37*int32(10)
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43))))
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+2)))
	if v44<<(uint(int32(16))%32)|v47 != l2 {
		v30 = v35
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	goto L7
L11:
	;
	v93 = v35
	goto L1
L12:
	;
	v65 = int32(1)
	goto L13
L13:
	;
	v76 = l1 + int32(22) + v65&int32(_a_F_dataFindChildPtr_0)*int32(10)
	v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76))))
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76)+2)))
	if l2 == v77<<(uint(int32(16))%32)|v80 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v93 = v57
	goto L1
L15:
	;
	v93 = v65
	goto L1
L16:
	;
	goto L17
L17:
	;
	v84 = v65 + int32(1)
	v85 = int32(_a_F_dataFindChildPtr_0)
	if base.Ui32(v84&v85) <= base.Ui32(v54&v85) {
		v65 = v84
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L14
}
func F_dataGetLeftMostPage(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+34)))
	v4 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)))
	return v3 | v4<<(uint(int32(16))%32)
}
func F_dataIsMoveRight(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v29 int64
	_ = v29
	v4 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
	v5 = l1 + v4
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	if v6 == int32(-1) {
		return int32(0)
	} else {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+6)))
		if v11&int32(4) != 0 {
			return int32(1)
		} else {
			v16 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+66)))
			v17 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)))
			v18 = int64(32)
			v20 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+62)))
			v21 = int64(48)
			v25 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l1)+28)))
			v26 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)))
			v29 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l1)+24)))
			return base.B2i32(base.Ui64(v25|(v26<<(uint(v18)%64)|v29<<(uint(v21)%64))) < base.Ui64(v16|(v17<<(uint(v18)%64)|v20<<(uint(v21)%64))))
		}
	}
}
func F_dataLocateItem(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int64
	_ = v80
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v88 int64
	_ = v88
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v95 int64
	_ = v95
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v129 int32
	_ = v129
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v12 < int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
	if v31 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_dataLocateItem[0]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16+(v12^int32(-1))<<(uint(int32(2))%32))))
	v30 = v22
	goto L1
L3:
	;
	goto L4
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_dataLocateItem[1]))
	v30 = v24 + v12<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	v34 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v34)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+16)))
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+v37)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v36 * v39
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v43 = m.T0[v42].(func(*base.Module, int32, int32) int32)(m, l0, v30)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+16)))
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+v48)+4)))
	v52 = v50 + int32(1)
	if base.Ui32(int32(2)) <= base.Ui32(v52&int32(_a_F_dataLocateItem_0)) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	return int32(0)
L9:
	;
	return v43
L10:
	;
	v62 = int32(1)
	v65 = v52
	goto L13
L11:
	;
	v129 = v52
	goto L12
L12:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v129)
	v140 = v30 + v129&int32(_a_F_dataLocateItem_0)*int32(10)
	v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140)+22)))
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140)+24)))
	return v141<<(uint(int32(16))%32) | v144
L13:
	;
	v76 = int32(base.Ui32((v65-v62)&int32(_a_F_dataLocateItem_1))>>(uint(int32(1))%32)) + v62
	v78 = v76 & int32(_a_F_dataLocateItem_0)
	if v50 == v78 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v129 = v118
	goto L12
L15:
	;
	v119 = int32(_a_F_dataLocateItem_0)
	if base.Ui32(v114&v119) < base.Ui32(v118&v119) {
		v62 = v114
		v65 = v118
		goto L13
	} else {
		goto L28
	}
L16:
	;
	v114 = v62
	v118 = v76
	goto L15
L17:
	;
	goto L18
L18:
	;
	v80 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+66)))
	v81 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)))
	v82 = int64(32)
	v84 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+62)))
	v85 = int64(48)
	v88 = v80 | (v81<<(uint(v82)%64) | v84<<(uint(v85)%64))
	v91 = v30 + int32(22) + v78*int32(10)
	v92 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v91)+6)))
	v95 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v91)+4)))
	v99 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v91)+8)))
	v100 = v92<<(uint(v82)%64) | v95<<(uint(v85)%64) | v99
	if v88 == v100 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v76)
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+2)))
	v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91))))
	return v103 | v104<<(uint(int32(16))%32)
L20:
	;
	goto L21
L21:
	;
	v111 = base.B2i32(base.Ui64(v100) < base.Ui64(v88))
	if base.Ui64(v100) < base.Ui64(v88) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v112 = v76 + int32(1)
	goto L24
L23:
	;
	v112 = v62
	goto L24
L24:
	;
	if base.Ui64(v100) < base.Ui64(v88) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v113 = v65
	goto L27
L26:
	;
	v113 = v76
	goto L27
L27:
	;
	v114 = v112
	v118 = v113
	goto L15
L28:
	;
	goto L14
}
