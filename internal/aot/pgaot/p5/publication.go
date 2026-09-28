package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_publication_relations(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v14 = F_table_open(m, int32(_a_F_get_publication_relations_0), int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = v8 + int32(-56)
	F_ScanKeyInit(m, v19, int32(2), int32(3), int32(184), base.I64_extend_i32_u(l0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = int32(1)
	v30 = F_systable_beginscan(m, v14, int32(_a_F_get_publication_relations_1), v27, int32(0), v27, v19)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v32 = F_systable_getnext(m, v30)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v32 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v34 = v32
	v37 = v4
	goto L9
L7:
	;
	v55 = v4
	goto L8
L8:
	;
	F_systable_endscan(m, v30)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L17
	}
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+22)))
	v43 = v41 + v42
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+12)))
	if l2 == v44 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v55 = v49
	goto L8
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v47 = F_GetPubPartitionOptionRelations(m, v37, l1, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v49 = v37
	goto L13
L13:
	;
	v50 = F_systable_getnext(m, v30)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	v49 = v47
	goto L13
L15:
	;
	if v50 != 0 {
		v34 = v50
		v37 = v49
		goto L9
	} else {
		goto L16
	}
L16:
	;
	goto L10
L17:
	;
	F_relation_close(m, v14, int32(1))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_list_sort(m, v55, int32(502))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v67 = int32(0)
	if v55 == v67 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	m.G0 = v10 - int32(-64)
	return v55
L21:
	;
	goto L20
L22:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v77 < int32(2) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v80 = int32(1)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	if v77 != int32(2) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v162 + int32(1)
	goto L21
L25:
	;
	v84 = int32(1)
	v85 = v77 - v84
	v90 = int32(0)
	v93 = v90
	v95 = v80
	v96 = v90
	goto L28
L26:
	;
	v138 = v67
	v140 = v80
	goto L27
L27:
	;
	v146 = int32(2)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v81+v140<<(uint(v146)%32))))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v81+v138<<(uint(v146)%32))))
	if v149 == v153 {
		v162 = v138
		goto L24
	} else {
		goto L38
	}
L28:
	;
	v101 = int32(2)
	v103 = v81 + v95<<(uint(v101)%32)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v81+v93<<(uint(v101)%32))))
	if v104 != v108 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v85&v84 == int32(0) {
		v162 = v129
		goto L24
	} else {
		goto L37
	}
L30:
	;
	v111 = v93 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v81+v111<<(uint(int32(2))%32)))) = v104
	v116 = v111
	goto L32
L31:
	;
	v116 = v93
	goto L32
L32:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v81+v116<<(uint(int32(2))%32))))
	if v117 != v121 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v124 = v116 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v81+v124<<(uint(int32(2))%32)))) = v117
	v129 = v124
	goto L35
L34:
	;
	v129 = v116
	goto L35
L35:
	;
	v130 = int32(2)
	v131 = v95 + v130
	v133 = v96 + v130
	if v133 != v85&int32(-2) {
		v93 = v129
		v95 = v131
		v96 = v133
		goto L28
	} else {
		goto L36
	}
L36:
	;
	goto L29
L37:
	;
	v138 = v129
	v140 = v131
	goto L27
L38:
	;
	v156 = v138 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v81+v156<<(uint(int32(2))%32)))) = v149
	v162 = v156
	goto L24
}
