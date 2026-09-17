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
	var v79 int32
	_ = v79
	var v86 int64
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int64
	_ = v102
	var v112 int32
	_ = v112
	var v113 int64
	_ = v113
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v175 int32
	_ = v175
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
	v175 = m.ExcPending
	if v175 != 0 {
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
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L46
	}
L13:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
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
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexMultipleInternal[0]))
	if v51 == v55 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexMultipleInternal[1]))
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
	v79 = int32(0)
	if base.B2i32(v76&int32(8) == v79)|base.B2i32(v74 == int32(116)) == v79 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v86 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = base.I32_wrap_i64(v86) | int32(4)
	v94 = F_ReindexRelationConcurrently(m, l0, v36, v12+int32(8))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	if v72 == int32(105) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexMultipleInternal[2]))
	goto L29
L29:
	;
	if v97 != int32(0) {
		goto L13
	} else {
		goto L30
	}
L30:
	;
	goto L12
L31:
	;
	v102 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = base.I32_wrap_i64(v102) | int32(6)
	F_reindex_index(m, l0, v36, int32(0), v74, v12+int32(8))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v113 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = base.I32_wrap_i64(v113) | int32(6)
	v122 = F_reindex_relation(m, l0, v36, int32(5), v12+int32(8))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L13
L35:
	;
	if v122 == int32(0) {
		goto L13
	} else {
		goto L36
	}
L36:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v126&int32(1) == int32(0) {
		goto L13
	} else {
		goto L37
	}
L37:
	;
	v133 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v133 == int32(0) {
		goto L13
	} else {
		goto L39
	}
L39:
	;
	v137 = F_get_rel_namespace(m, v36)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v139 = F_get_namespace_name(m, v137)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v141 = F_get_rel_name(m, v36)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v139
	F_errmsg(m, int32(_a_F_ReindexMultipleInternal_0), v12)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_ReindexMultipleInternal_1), int32(3532), int32(_a_F_ReindexMultipleInternal_2))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
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
	v162 = v29 + int32(1)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v162 < v163 {
		v29 = v162
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
		v15 = int32(_a_F_reindex_error_callback_0)
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
			v15 = int32(_a_F_reindex_error_callback_1)
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
