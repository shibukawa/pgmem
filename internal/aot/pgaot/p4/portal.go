package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ChoosePortalStrategy(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = int32(4)
	if l0 == v2 {
		v129 = v13
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L14
	} else {
		goto L53
	}
L2:
	;
	m.G0 = v11 + int32(32)
	return v129
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
		v129 = v13
		goto L2
	} else {
		goto L29
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
	v65 = m.ExcPending
	if v65 != 0 {
		goto L14
	} else {
		goto L26
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
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+26)))
	if v45 != int32(1) {
		goto L4
	} else {
		goto L19
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
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v40 = F_UtilityReturnsTuples(m, v39)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+42)))
	v129 = v32 << (uint(int32(1)) % 32) & int32(2)
	goto L2
L14:
	;
	return int32(0)
L15:
	;
	if v40 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v44 = int32(3)
	goto L18
L17:
	;
	v44 = int32(4)
	goto L18
L18:
	;
	v129 = v44
	goto L2
L19:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	switch v48 - int32(1) {
	case 0:
		goto L21
	default:
		goto L4
	case 5:
		goto L20
	}
L20:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v20)+88))
	v59 = F_UtilityReturnsTuples(m, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L14
	} else {
		goto L22
	}
L21:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+25)))
	v129 = v51 << (uint(int32(1)) % 32) & int32(2)
	goto L2
L22:
	;
	if v59 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v61 = int32(3)
	goto L25
L24:
	;
	v61 = int32(4)
	goto L25
L25:
	;
	v129 = v61
	goto L2
L26:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v66
	F_errmsg_internal(m, int32(_a_F_ChoosePortalStrategy_0), v11)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L14
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_ChoosePortalStrategy_1), int32(270), int32(_a_F_ChoosePortalStrategy_2))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L14
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	v80 = int32(0)
	if v80 < v16 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v83 = v16
	goto L32
L31:
	;
	v83 = v80
	goto L32
L32:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v86 = int32(0)
	v92 = v2
	goto L33
L33:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v84+v86<<(uint(int32(2))%32))))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if v98 != int32(330) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	if v120 != 0 {
		goto L50
	} else {
		goto L51
	}
L35:
	;
	v122 = v86 + int32(1)
	if v122 != v83 {
		v86 = v122
		v92 = v120
		goto L33
	} else {
		goto L49
	}
L36:
	;
	v120 = int32(1)
	goto L35
L37:
	;
	if v98 != int32(67) {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+26)))
	if v110 != int32(1) {
		v120 = v92
		goto L35
	} else {
		goto L45
	}
L40:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+24)))
	if v103 != int32(1) {
		v120 = v92
		goto L35
	} else {
		goto L41
	}
L41:
	;
	if v92 != 0 {
		v129 = v13
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	if v106 == int32(6) {
		v129 = v13
		goto L2
	} else {
		goto L43
	}
L43:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v97)+96))
	if v109 != 0 {
		goto L36
	} else {
		goto L44
	}
L44:
	;
	v129 = v13
	goto L2
L45:
	;
	if v92 != 0 {
		v129 = v13
		goto L2
	} else {
		goto L46
	}
L46:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	if v113 == int32(6) {
		v129 = v13
		goto L2
	} else {
		goto L47
	}
L47:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+24)))
	if v116 == int32(0) {
		v129 = v13
		goto L2
	} else {
		goto L48
	}
L48:
	;
	goto L36
L49:
	;
	goto L34
L50:
	;
	v126 = int32(1)
	goto L52
L51:
	;
	v126 = int32(4)
	goto L52
L52:
	;
	v129 = v126
	goto L2
L53:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v143
	F_errmsg_internal(m, int32(_a_F_ChoosePortalStrategy_0), v11+int32(16))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L14
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_ChoosePortalStrategy_1), int32(310), int32(_a_F_ChoosePortalStrategy_2))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L14
	} else {
		goto L55
	}
L55:
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	v3 = m.G0
	v5 = v3 - int32(80)
	m.G0 = v5
	goto L1
L1:
	;
	v9 = int32(_a_F_CreateNewPortal_0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_CreateNewPortal[0]))
	v13 = v11 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CreateNewPortal[0])) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v13
	v17 = v5 + int32(16)
	v19 = F_pg_sprintf(m, v17, int32(_a_F_CreateNewPortal_1), v5)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v34 = int32(0)
	v36 = F_CreatePortal(m, v5+int32(16), v34, v34)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
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
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_CreateNewPortal[1]))
	v25 = int32(0)
	v27 = F_hash_search(m, v24, v17, v25, v25)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v27 == int32(0) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+64))
	if v31 != 0 {
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
	return v36
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
	var v40 int32
	_ = v40
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
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int64
	_ = v156
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
		goto L60
	} else {
		goto L61
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
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_PortalRunMulti[0]))
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
	v40 = v7
	goto L10
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v47 = v44 + v40<<(uint(int32(2))%32)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_PortalRunMulti[1]))
	if v50 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v108&int32(1) == int32(0) {
		goto L1
	} else {
		goto L58
	}
L12:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+88))
	if v53 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	return
L16:
	;
	goto L14
L17:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_MemoryContextDeleteChildren(m, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L15
	} else {
		goto L49
	}
L18:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PortalRunMulti[2])))
	if v57 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+26)))
	if v98 == int32(1) {
		goto L44
	} else {
		goto L45
	}
L21:
	;
	F_getrusage(m, int32(_a_F_PortalRunMulti_0))
	mBase = m.M
	F_gettimeofday(m, int32(_a_F_PortalRunMulti_1))
	mBase = m.M
	goto L24
L22:
	;
	goto L23
L23:
	;
	if v34&int32(1) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L23
L25:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+26)))
	if v82 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L26:
	;
	v68 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L15
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	F_UpdateActiveSnapshotCommandId(m)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L15
	} else {
		goto L35
	}
L29:
	;
	if l2 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v70 = F_RegisterSnapshot(m, v68)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L15
	} else {
		goto L33
	}
L31:
	;
	v73 = v68
	goto L32
L32:
	;
	F_PushCopiedSnapshot(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L15
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v70
	v73 = v70
	goto L32
L34:
	;
	goto L25
L35:
	;
	goto L25
L36:
	;
	v90 = int32(1)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PortalRunMulti[2])))
	if v92 != v90 {
		v108 = v90
		goto L17
	} else {
		goto L42
	}
L37:
	;
	F_ProcessQuery(m, v48, v81, v80, v79, v29, l5)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L15
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	F_ProcessQuery(m, v48, v81, v80, v79, v25, int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L15
	} else {
		goto L41
	}
L40:
	;
	goto L36
L41:
	;
	goto L36
L42:
	;
	F_ShowUsage(m, int32(_a_F_PortalRunMulti_2))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L15
	} else {
		goto L43
	}
L43:
	;
	v108 = v90
	goto L17
L44:
	;
	F_PortalRunUtility(m, l0, v48, l1, int32(0), v29, l5)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L15
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v104 = int32(0)
	F_PortalRunUtility(m, l0, v48, l1, v104, v25, v104)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L15
	} else {
		goto L48
	}
L47:
	;
	v108 = v34
	goto L17
L48:
	;
	v108 = v34
	goto L17
L49:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v114 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if base.Ui32(v47+int32(4)) < base.Ui32(v117+v118<<(uint(int32(2))%32)) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	goto L11
L53:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L15
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v126 = v40 + int32(1)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v126 < v127 {
		v34 = v108
		v40 = v126
		goto L10
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	goto L52
L58:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L15
	} else {
		goto L59
	}
L59:
	;
	goto L1
L60:
	;
	return
L61:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v151 != 0 {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v152 == int32(0) {
		goto L60
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v152
	v156 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = v156
	goto L60
}
