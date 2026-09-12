package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_check_temp_buffers(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	v4 = int32(1)
	if l2 == int32(12) {
		v26 = v4
		return v26
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _consts[640]))
		if v8 == int32(0) {
			v26 = v4
			return v26
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v8 == v11 {
				v26 = v4
				return v26
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, _consts[86]))
				*(*int32)(unsafe.Add(mBase, _consts[87])) = v15
				v21 = F_format_elog_string(m, int32(618061), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[88])) = v21
					v26 = int32(0)
					return v26
				}
			}
		}
	}
}
func F_check_temp_tablespaces(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v212 int32
	_ = v212
	v4 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v20 = F_pstrdup(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v17 + int32(16)
	return v212
L2:
	;
	F_pfree(m, v20)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L3
	} else {
		goto L57
	}
L3:
	;
	return int32(0)
L4:
	;
	v27 = F_SplitIdentifierString(m, v20, int32(44), v17+int32(12))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if v27 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	*(*int32)(unsafe.Add(mBase, _consts[87])) = v32
	goto L9
L7:
	;
	goto L8
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
	goto L11
L9:
	;
	v38 = F_format_elog_string(m, int32(646598), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, _consts[88])) = v38
	goto L2
L11:
	;
	if base.B2i32(v43 == int32(2)) == int32(0) {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _consts[107]))
	if v49 == int32(0) {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v53 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v58 = v54 << (uint(int32(2)) % 32)
	goto L16
L15:
	;
	v58 = int32(0)
	goto L16
L16:
	;
	v59 = F_palloc(m, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v61 == int32(0) {
		v159 = v4
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v169 = v159 << (uint(int32(2)) % 32)
	v172 = F_guc_malloc(m, v169+int32(4))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L3
	} else {
		goto L48
	}
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v64 <= int32(0) {
		v159 = v4
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v73 = int32(0)
	v78 = v4
	goto L21
L21:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87+v73<<(uint(int32(2))%32))))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	if v92 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v159 = v147
	goto L18
L23:
	;
	v151 = v73 + int32(1)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v151 < v152 {
		v73 = v151
		v78 = v147
		goto L21
	} else {
		goto L47
	}
L24:
	;
	v147 = v78 + int32(1)
	goto L23
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59+v78<<(uint(int32(2))%32)))) = int32(0)
	goto L24
L26:
	;
	goto L27
L27:
	;
	v100 = F_get_tablespace_oid(m, v91, base.B2i32(base.Ui32(l2) < base.Ui32(int32(13))))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	if v100 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if l2 != int32(12) {
		v147 = v78
		goto L23
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _consts[108]))
	if v123 == v100 {
		goto L38
	} else {
		goto L39
	}
L32:
	;
	v106 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	if v106 == int32(0) {
		v147 = v78
		goto L23
	} else {
		goto L34
	}
L34:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v91
	F_errmsg(m, int32(72660), v17)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(499979), int32(1258), int32(171792))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L3
	} else {
		goto L37
	}
L37:
	;
	v147 = v78
	goto L23
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59+v78<<(uint(int32(2))%32)))) = int32(0)
	goto L24
L39:
	;
	goto L40
L40:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _consts[239]))
	v134 = F_object_aclcheck(m, int32(1213), v100, v132, int64(512))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	if v134 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	if base.Ui32(l2) < base.Ui32(int32(11)) {
		v147 = v78
		goto L23
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59+v78<<(uint(int32(2))%32)))) = v100
	goto L24
L45:
	;
	F_aclcheck_error(m, v134, int32(42), v91)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L3
	} else {
		goto L46
	}
L46:
	;
	v147 = v78
	goto L23
L47:
	;
	goto L22
L48:
	;
	if v172 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v212 = int32(0)
	goto L1
L50:
	;
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172))) = v159
	if v169 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v172
	F_pfree(m, v59)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L3
	} else {
		goto L56
	}
L53:
	;
	v180 = F__emscripten_memcpy_bulkmem(m, v172+int32(4), v59, v169)
	mBase = m.M
	goto L55
L54:
	;
	goto L55
L55:
	;
	goto L52
L56:
	;
	goto L2
L57:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	F_list_free(m, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L3
	} else {
		goto L58
	}
L58:
	;
	v212 = v27
	goto L1
}
