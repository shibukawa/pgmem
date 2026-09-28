package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_regexp_count_no_start(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_regexp_count(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_regexp_split_to_array(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if int32(3) <= v12 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v21 = int32(0)
	goto L3
L3:
	;
	F_parse_re_flags(m, v8+int32(8), v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return int64(0)
L5:
	;
	v21 = v16
	goto L3
L6:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+12)))
	if v24 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v27 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+12)) = uint8(v27)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v30 = F_pg_detoast_datum_packed(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L22
	}
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v33 = F_pg_detoast_datum_packed(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v37 = int32(0)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v40 = int32(1)
	v42 = F_setup_regexp_matches(m, v30, v33, v8+int32(8), v37, v38, v37, v40, v40)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v44 <= v45 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v49 = v2
	goto L16
L14:
	;
	v68 = v2
	goto L15
L15:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_regexp_split_to_array[0]))
	v73 = F_makeArrayResult(m, v68, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L21
	}
L16:
	;
	v52 = F_build_regexp_split_result(m, v42)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L18
	}
L17:
	;
	v68 = v58
	goto L15
L18:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_regexp_split_to_array[0]))
	v58 = F_accumArrayResult(m, v49, v52, int32(0), int32(25), v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v62 = v60 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v62 <= v64 {
		v49 = v58
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	m.G0 = v8 + int32(16)
	return v73
L22:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_regexp_split_to_array_0)
	F_errmsg(m, int32(_a_F_regexp_split_to_array_1), v8)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_regexp_split_to_array_2), int32(1827), int32(_a_F_regexp_split_to_array_3))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
