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
	var v45 int64
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v87 int64
	_ = v87
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int64
	_ = v103
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
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
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v176 int32
	_ = v176
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
	v176 = m.ExcPending
	if v176 != 0 {
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
	v45 = int64(0)
	v48 = F_SearchSysCacheExists(m, int32(57), base.I64_extend_i32_u(v36), v45, v45, v45)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L46
	}
L13:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L45
	}
L14:
	;
	if v48 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v52 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v73 = F_get_rel_relkind(m, v36)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L23
	}
L17:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexMultipleInternal[0]))
	if v52 == v56 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexMultipleInternal[1]))
	v62 = F_object_aclcheck(m, int32(1213), v52, v60, int64(512))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v62 == int32(0) {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v68 = F_get_tablespace_name(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_aclcheck_error(m, v62, int32(43), v68)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	goto L16
L23:
	;
	v75 = F_get_rel_persistence(m, v36)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v80 = int32(0)
	if base.B2i32(v77&int32(8) == v80)|base.B2i32(v75 == int32(116)) == v80 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v87 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = base.I32_wrap_i64(v87) | int32(4)
	v95 = F_ReindexRelationConcurrently(m, l0, v36, v12+int32(8))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	if v73 == int32(105) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexMultipleInternal[2]))
	goto L29
L29:
	;
	if v98 != int32(0) {
		goto L13
	} else {
		goto L30
	}
L30:
	;
	goto L12
L31:
	;
	v103 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = base.I32_wrap_i64(v103) | int32(6)
	F_reindex_index(m, l0, v36, int32(0), v75, v12+int32(8))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v114 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = base.I32_wrap_i64(v114) | int32(6)
	v123 = F_reindex_relation(m, l0, v36, int32(5), v12+int32(8))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L13
L35:
	;
	if v123 == int32(0) {
		goto L13
	} else {
		goto L36
	}
L36:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v127&int32(1) == int32(0) {
		goto L13
	} else {
		goto L37
	}
L37:
	;
	v134 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v134 == int32(0) {
		goto L13
	} else {
		goto L39
	}
L39:
	;
	v138 = F_get_rel_namespace(m, v36)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v140 = F_get_namespace_name(m, v138)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v142 = F_get_rel_name(m, v36)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v140
	F_errmsg(m, int32(_a_F_ReindexMultipleInternal_0), v12)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_ReindexMultipleInternal_1), int32(3576), int32(_a_F_ReindexMultipleInternal_2))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
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
	v163 = v29 + int32(1)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v163 < v164 {
		v29 = v163
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
