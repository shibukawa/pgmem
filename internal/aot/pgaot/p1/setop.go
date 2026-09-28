package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_setop_query(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
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
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_get_setop_query[0]))
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v19 != int32(142) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L4
	} else {
		goto L63
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
	} else {
		goto L60
	}
L9:
	;
	m.G0 = v10 + int32(32)
	return
L10:
	;
	if v19 != int32(63) {
		goto L7
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v60 != int32(142) {
		goto L27
	} else {
		goto L28
	}
L13:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v25+v26<<(uint(int32(2))%32)-int32(4))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+36))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	if v34 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+33)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	F_get_query_def(m, v33, v12, v46, v47, v48, v49, v50, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L23
	}
L15:
	;
	F_appendStringInfoChar(m, v12, int32(40))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L22
	}
L16:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+124))
	if v35 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+140))
	if v36 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+128))
	if v37 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+132))
	if v38 != 0 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)+144))
	if v39 != 0 {
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v45 = int32(0)
	goto L14
L22:
	;
	v45 = int32(1)
	goto L14
L23:
	;
	if v45 == int32(0) {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	F_appendStringInfoChar(m, v12, int32(41))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	goto L9
L26:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v103 = v101 - int32(1)
	if base.Ui32(int32(3)) <= base.Ui32(v103) {
		goto L8
	} else {
		goto L43
	}
L27:
	;
	F_get_setop_query(m, v59, l1, l2)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L37
	}
L28:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v63 == v64 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+8)))
	if v66 == v67 {
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	F_appendStringInfoChar(m, v12, int32(40))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	v74 = int32(0)
	F_appendContextKeyword(m, l2, int32(_a_F_get_setop_query_4), int32(8), v74, v74)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_get_setop_query(m, v78, l1, l2)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v83 = int32(0)
	F_appendContextKeyword(m, l2, int32(_a_F_get_setop_query_9), int32(-8), v83, v83)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	goto L26
L37:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v89&int32(2) != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v93 = int32(0)
	F_appendContextKeyword(m, l2, int32(_a_F_get_setop_query_4), v93, v93, v93)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	F_appendStringInfoChar(m, v12, int32(32))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L42
	}
L41:
	;
	goto L26
L42:
	;
	goto L26
L43:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v103<<(uint(int32(2))%32))+uint32(_c_F_get_setop_query[1])))
	F_appendStringInfoString(m, v12, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v111 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_appendStringInfoString(m, v12, int32(_a_F_get_setop_query_7))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	if v119 == int32(142) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L47
L49:
	;
	F_appendStringInfoChar(m, v12, int32(40))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L4
	} else {
		goto L52
	}
L50:
	;
	v126 = int32(0)
	goto L51
L51:
	;
	v128 = int32(0)
	F_appendContextKeyword(m, l2, int32(_a_F_get_setop_query_4), v126, v128, v128)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L53
	}
L52:
	;
	v126 = int32(8)
	goto L51
L53:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+33)))
	v133 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+33)) = uint8(v133)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_get_setop_query(m, v135, l1, l2)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+33)) = uint8(v132)
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v139&int32(2) != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v142 - v126
	goto L57
L56:
	;
	goto L57
L57:
	;
	if v119 != int32(142) {
		goto L9
	} else {
		goto L58
	}
L58:
	;
	v148 = int32(0)
	F_appendContextKeyword(m, l2, int32(_a_F_get_setop_query_8), v148, v148, v148)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	goto L9
L60:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v164
	F_errmsg_internal(m, int32(_a_F_get_setop_query_5), v10+int32(16))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_get_setop_query_1), int32(_a_F_get_setop_query_6), int32(_a_F_get_setop_query_3))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v180
	F_errmsg_internal(m, int32(_a_F_get_setop_query_0), v10)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_get_setop_query_1), int32(_a_F_get_setop_query_2), int32(_a_F_get_setop_query_3))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_setop_compare_slots(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v103 int32
	_ = v103
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v11 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v11 < v10 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	m.T0[v14].(func(*base.Module, int32, int32))(m, l0, v10)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v21 < v20 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	m.T0[v24].(func(*base.Module, int32, int32))(m, l1, v20)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)+120))
	if int32(0) < v27 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L8
L10:
	;
	return int32(1)
L11:
	;
	return v103
L12:
	;
	v33 = v27
	v35 = int32(0)
	goto L15
L13:
	;
	goto L14
L14:
	;
	v103 = int32(0)
	goto L11
L15:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+124))
	v41 = v38 + v35*int32(36)
	v42 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+10)))
	v43 = int32(1)
	v44 = v42 - v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v45))))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+v44))))
	if v50 == v43 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L14
L17:
	;
	v89 = v35 + int32(1)
	if v89 < v87 {
		v33 = v87
		v35 = v89
		goto L15
	} else {
		goto L37
	}
L18:
	;
	if v47&int32(1) != 0 {
		v87 = v33
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v47&int32(1) != 0 {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+9)))
	if v57 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v58 = int32(-1)
	goto L24
L23:
	;
	v58 = int32(1)
	goto L24
L24:
	;
	return v58
L25:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+9)))
	if v64 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v68 = v44 << (uint(int32(3)) % 32)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v71 = *(*int64)(unsafe.Add(mBase, uint32(v68+v69)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v74 = *(*int64)(unsafe.Add(mBase, uint32(v72+v68)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v76 = m.T0[v75].(func(*base.Module, int64, int64, int32) int32)(m, v71, v74, v41)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L31
	}
L28:
	;
	v65 = int32(1)
	goto L30
L29:
	;
	v65 = int32(-1)
	goto L30
L30:
	;
	return v65
L31:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+8)))
	if v78 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	if v76 < int32(0) {
		goto L10
	} else {
		goto L35
	}
L33:
	;
	v85 = v76
	goto L34
L34:
	;
	if v85 != 0 {
		v103 = v85
		goto L11
	} else {
		goto L36
	}
L35:
	;
	v85 = int32(0) - v76
	goto L34
L36:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l2)+120))
	v87 = v86
	goto L17
L37:
	;
	goto L16
}
