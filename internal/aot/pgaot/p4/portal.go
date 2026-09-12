package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ChoosePortalStrategy(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = int32(4)
	if l0 == int32(0) {
		v138 = v13
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L17
	} else {
		goto L58
	}
L2:
	;
	m.G0 = v11 + int32(32)
	return v138
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v16 != int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if v16 <= int32(0) {
		v138 = v13
		goto L2
	} else {
		goto L35
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v21 != int32(330) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L17
	} else {
		goto L32
	}
L7:
	;
	if v21 != int32(67) {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+26)))
	if v44 != int32(1) {
		goto L4
	} else {
		goto L22
	}
L10:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+24)))
	if v26 != int32(1) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	switch v29 - int32(1) {
	case 0:
		goto L13
	default:
		goto L4
	case 5:
		goto L12
	}
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v39 = F_UtilityReturnsTuples(m, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+42)))
	if v34 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v35 = int32(2)
	goto L16
L15:
	;
	v35 = int32(0)
	goto L16
L16:
	;
	v138 = v35
	goto L2
L17:
	;
	return int32(0)
L18:
	;
	if v39 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v43 = int32(3)
	goto L21
L20:
	;
	v43 = int32(4)
	goto L21
L21:
	;
	v138 = v43
	goto L2
L22:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	switch v47 - int32(1) {
	case 0:
		goto L24
	default:
		goto L4
	case 5:
		goto L23
	}
L23:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v20)+88))
	v57 = F_UtilityReturnsTuples(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L17
	} else {
		goto L28
	}
L24:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+25)))
	if v52 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v53 = int32(2)
	goto L27
L26:
	;
	v53 = int32(0)
	goto L27
L27:
	;
	v138 = v53
	goto L2
L28:
	;
	if v57 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v59 = int32(3)
	goto L31
L30:
	;
	v59 = int32(4)
	goto L31
L31:
	;
	v138 = v59
	goto L2
L32:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v64
	F_errmsg_internal(m, int32(509011), v11)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L17
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(515894), int32(270), int32(21044))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L17
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	v78 = int32(0)
	if v78 < v16 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v82 = v16
	goto L38
L37:
	;
	v82 = v78
	goto L38
L38:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v86 = v78
	v87 = int32(0)
	goto L39
L39:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v83+v86<<(uint(int32(2))%32))))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	if v97 != int32(330) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v130 = int32(1)
	if v126 == v130 {
		goto L55
	} else {
		goto L56
	}
L41:
	;
	v128 = v86 + int32(1)
	if v128 != v82 {
		v86 = v128
		v87 = v126
		goto L39
	} else {
		goto L54
	}
L42:
	;
	if v97 != int32(67) {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+26)))
	if v113 != int32(1) {
		v126 = v87
		goto L41
	} else {
		goto L50
	}
L45:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+24)))
	if v102 != int32(1) {
		v126 = v87
		goto L41
	} else {
		goto L46
	}
L46:
	;
	v105 = int32(1)
	v106 = v87 + v105
	if v105 < v106 {
		v138 = v13
		goto L2
	} else {
		goto L47
	}
L47:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v109 == int32(6) {
		v138 = v13
		goto L2
	} else {
		goto L48
	}
L48:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v96)+96))
	if v112 != 0 {
		v126 = v106
		goto L41
	} else {
		goto L49
	}
L49:
	;
	v138 = v13
	goto L2
L50:
	;
	v116 = int32(1)
	v117 = v87 + v116
	if v116 < v117 {
		v138 = v13
		goto L2
	} else {
		goto L51
	}
L51:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v120 == int32(6) {
		v138 = v13
		goto L2
	} else {
		goto L52
	}
L52:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+24)))
	if v123 == int32(0) {
		v138 = v13
		goto L2
	} else {
		goto L53
	}
L53:
	;
	v126 = v117
	goto L41
L54:
	;
	goto L40
L55:
	;
	v134 = v130
	goto L57
L56:
	;
	v134 = int32(4)
	goto L57
L57:
	;
	v138 = v134
	goto L2
L58:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v151
	F_errmsg_internal(m, int32(509011), v11+int32(16))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L17
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(515894), int32(310), int32(21044))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L17
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CreateNewPortal(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	v3 = m.G0
	v5 = v3 - int32(80)
	m.G0 = v5
	goto L1
L1:
	;
	v9 = int32(4562128)
	v11 = *(*int32)(unsafe.Add(mBase, _consts[1451]))
	v13 = v11 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[1451])) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v13
	v19 = F_pg_sprintf(m, v5+int32(16), int32(572886), v5)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v36 = int32(0)
	v38 = F_CreatePortal(m, v5+int32(16), v36, v36)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L9
	}
L3:
	;
	goto L2
L4:
	;
	return int32(0)
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _consts[228]))
	v27 = int32(0)
	v29 = F_hash_search(m, v24, v5+int32(16), v27, v27)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v29 == int32(0) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)+64))
	if v33 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L3
L9:
	;
	m.G0 = v5 + int32(80)
	return v38
}
func F_MarkPortalFailed(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = int32(5)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v5 != 0 {
		m.T0[v5].(func(*base.Module, int32))(m, l0)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
			return
		}
	} else {
		return
	}
}
func F_PortalRunMulti(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v142 int32
	_ = v142
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int64
	_ = v163
	v7 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v14 == v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l5 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v17 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _consts[430]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	if v22 == int32(3) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v25 = v21
	goto L6
L5:
	;
	v25 = l4
	goto L6
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v26 == int32(3) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v29 = v21
	goto L9
L8:
	;
	v29 = l3
	goto L9
L9:
	;
	v34 = int32(0)
	v37 = v7
	goto L12
L10:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L17
	} else {
		goto L60
	}
L11:
	;
	if v108&int32(1) == int32(0) {
		goto L1
	} else {
		goto L59
	}
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v47 = v44 + v34<<(uint(int32(2))%32)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v50 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v50 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v108&int32(1) != 0 {
		goto L10
	} else {
		goto L58
	}
L14:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+88))
	if v53 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	return
L18:
	;
	goto L16
L19:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_MemoryContextDeleteChildren(m, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L17
	} else {
		goto L51
	}
L20:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1223])))
	if v57 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+26)))
	if v98 == int32(1) {
		goto L46
	} else {
		goto L47
	}
L23:
	;
	F_getrusage(m, int32(4485720))
	mBase = m.M
	F___gettimeofday(m, int32(4485872))
	mBase = m.M
	goto L26
L24:
	;
	goto L25
L25:
	;
	if v37&int32(1) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L25
L27:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+26)))
	if v82 == int32(1) {
		goto L39
	} else {
		goto L40
	}
L28:
	;
	v68 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L17
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	F_UpdateActiveSnapshotCommandId(m)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L17
	} else {
		goto L37
	}
L31:
	;
	if l2 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v70 = F_RegisterSnapshot(m, v68)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L17
	} else {
		goto L35
	}
L33:
	;
	v73 = v68
	goto L34
L34:
	;
	F_PushCopiedSnapshot(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L17
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v70
	v73 = v70
	goto L34
L36:
	;
	goto L27
L37:
	;
	goto L27
L38:
	;
	v90 = int32(1)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1223])))
	if v92 != v90 {
		v108 = v90
		goto L19
	} else {
		goto L44
	}
L39:
	;
	F_ProcessQuery(m, v48, v81, v80, v79, v29, l5)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L17
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	F_ProcessQuery(m, v48, v81, v80, v79, v25, int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L17
	} else {
		goto L43
	}
L42:
	;
	goto L38
L43:
	;
	goto L38
L44:
	;
	F_ShowUsage(m, int32(549983))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L17
	} else {
		goto L45
	}
L45:
	;
	v108 = v90
	goto L19
L46:
	;
	F_PortalRunUtility(m, l0, v48, l1, int32(0), v29, l5)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L17
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v104 = int32(0)
	F_PortalRunUtility(m, l0, v48, l1, v104, v25, v104)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L17
	} else {
		goto L50
	}
L49:
	;
	v108 = v37
	goto L19
L50:
	;
	v108 = v37
	goto L19
L51:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v114 == int32(0) {
		goto L11
	} else {
		goto L52
	}
L52:
	;
	v118 = v47 + int32(4)
	if v118 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v130 = v34 + int32(1)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v130 < v131 {
		v34 = v130
		v37 = v108
		goto L12
	} else {
		goto L57
	}
L54:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if base.Ui32(v121+v122<<(uint(int32(2))%32)) <= base.Ui32(v118) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L17
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	goto L13
L58:
	;
	goto L1
L59:
	;
	goto L10
L60:
	;
	goto L1
L61:
	;
	return
L62:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v158 != 0 {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v159 == int32(0) {
		goto L61
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v159
	v163 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = v163
	goto L61
}
