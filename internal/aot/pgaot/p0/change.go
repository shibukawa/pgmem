package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_changeDependencyFor(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	v6 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(112)
	m.G0 = v13
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = l2
	goto L1
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = l4
	v52 = base.B2i32(base.B2i32(l2 == int32(2613))|base.B2i32(base.Ui32(int32(_a_F_changeDependencyFor_0)) < base.Ui32(l4)) == int32(0)) & ((base.B2i32(l2 != int32(2615)) | base.B2i32(l4 != int32(2200))) & base.B2i32(l2 != int32(1262)))
	goto L2
L2:
	;
	if base.B2i32(base.B2i32(l2 == int32(2613))|base.B2i32(base.Ui32(int32(_a_F_changeDependencyFor_0)) < base.Ui32(l3)) == v6)&((base.B2i32(l2 != int32(2615))|base.B2i32(l3 != int32(2200)))&base.B2i32(l2 != int32(1262))) != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v13 + int32(112)
	return v159
L4:
	;
	v53 = int32(1)
	if v52 != 0 {
		v159 = v53
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v70 = F_table_open(m, int32(2608), int32(3))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L8
	} else {
		goto L10
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l0
	F_recordMultipleDependencies(m, v13+int32(16), v13+int32(4), int32(1), int32(110))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	v159 = v53
	goto L3
L10:
	;
	v73 = v13 + int32(16)
	F_ScanKeyInit(m, v73, int32(1), int32(3), int32(184), l0)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	F_ScanKeyInit(m, v13-int32(-64), int32(2), int32(3), int32(184), l1)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v90 = F_systable_beginscan(m, v70, int32(2673), int32(1), int32(0), int32(2), v73)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v92 = F_systable_getnext(m, v90)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	if v92 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v94 = v92
	v102 = v6
	goto L18
L16:
	;
	v144 = v6
	goto L17
L17:
	;
	F_systable_endscan(m, v90)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L8
	} else {
		goto L33
	}
L18:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v94)+16))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+22)))
	v106 = v104 + v105
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	if v107 != l2 {
		v133 = v102
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v144 = v133
	goto L17
L20:
	;
	v134 = F_systable_getnext(m, v90)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L8
	} else {
		goto L31
	}
L21:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
	if v109 != l3 {
		v133 = v102
		goto L20
	} else {
		goto L22
	}
L22:
	;
	if v52 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v133 = v102 + int32(1)
	goto L20
L24:
	;
	F_simple_heap_delete(m, v70, v94+int32(4))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L8
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v115 = F_heap_copytuple(m, v94)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L8
	} else {
		goto L28
	}
L27:
	;
	goto L23
L28:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v115)+16))
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v117+v118)+16)) = l4
	F_CatalogTupleUpdate(m, v70, v115+int32(4), v115)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L8
	} else {
		goto L29
	}
L29:
	;
	F_pfree(m, v115)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L8
	} else {
		goto L30
	}
L30:
	;
	goto L23
L31:
	;
	if v134 != 0 {
		v94 = v134
		v102 = v133
		goto L18
	} else {
		goto L32
	}
L32:
	;
	goto L19
L33:
	;
	F_relation_close(m, v70, int32(3))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L8
	} else {
		goto L34
	}
L34:
	;
	v159 = v144
	goto L3
}
