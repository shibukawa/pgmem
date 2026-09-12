package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReindexMultipleInternal(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v83 int64
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int64
	_ = v99
	var v109 int32
	_ = v109
	var v110 int64
	_ = v110
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v174 int32
	_ = v174
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	F_PopActiveSnapshot(m)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l1 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L48
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v20 <= int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v29 = int32(0)
	goto L7
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v29<<(uint(int32(2))%32))))
	F_StartTransactionCommand(m)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L4
L9:
	;
	v39 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	F_PushActiveSnapshot(m, v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v44 = int32(0)
	v47 = F_SearchSysCacheExists(m, int32(57), v36, v44, v44, v44)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L46
	}
L13:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L45
	}
L14:
	;
	if v47 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v51 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v72 = F_get_rel_relkind(m, v36)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L23
	}
L17:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _consts[271]))
	if v51 == v55 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v61 = F_object_aclcheck(m, int32(1213), v51, v59, int64(512))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v61 == int32(0) {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v67 = F_get_tablespace_name(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_aclcheck_error(m, v61, int32(42), v67)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	goto L16
L23:
	;
	v74 = F_get_rel_persistence(m, v36)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v76&int32(8) == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if v72 == int32(105) {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	if v74 == int32(116) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v83 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = base.I32_wrap_i64(v83) | int32(4)
	v91 = F_ReindexRelationConcurrently(m, l0, v36, v12+int32(8))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	goto L29
L29:
	;
	if v94 != int32(0) {
		goto L13
	} else {
		goto L30
	}
L30:
	;
	goto L12
L31:
	;
	v99 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = base.I32_wrap_i64(v99) | int32(6)
	F_reindex_index(m, l0, v36, int32(0), v74, v12+int32(8))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v110 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = base.I32_wrap_i64(v110) | int32(6)
	v119 = F_reindex_relation(m, l0, v36, int32(5), v12+int32(8))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L13
L35:
	;
	if v119 == int32(0) {
		goto L13
	} else {
		goto L36
	}
L36:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v123&int32(1) == int32(0) {
		goto L13
	} else {
		goto L37
	}
L37:
	;
	v130 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v130 == int32(0) {
		goto L13
	} else {
		goto L39
	}
L39:
	;
	v134 = F_get_rel_namespace(m, v36)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v136 = F_get_namespace_name(m, v134)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v138 = F_get_rel_name(m, v36)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v136
	F_errmsg(m, int32(438498), v12)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(494197), int32(3532), int32(312772))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	goto L13
L45:
	;
	goto L12
L46:
	;
	v161 = v29 + int32(1)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v161 < v162 {
		v29 = v161
		goto L7
	} else {
		goto L47
	}
L47:
	;
	goto L8
L48:
	;
	m.G0 = v12 + int32(16)
	return
}
func F_reindex_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v24 int32
	_ = v24
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v8 == int32(112) {
		v15 = int32(691873)
		F_set_errcontext_domain(m, int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			*(*int64)(unsafe.Add(mBase, uint32(v6))) = base.I64_rotl(v19, int64(32))
			F_errcontext_msg(m, v15, v6)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				m.G0 = v6 + int32(16)
				return
			}
		}
	} else {
		if v8 != int32(73) {
			m.G0 = v6 + int32(16)
			return
		} else {
			v15 = int32(690850)
			F_set_errcontext_domain(m, int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*int64)(unsafe.Add(mBase, uint32(v6))) = base.I64_rotl(v19, int64(32))
				F_errcontext_msg(m, v15, v6)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					m.G0 = v6 + int32(16)
					return
				}
			}
		}
	}
}
