package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_regexp_count_no_start(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_regexp_count(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_regexp_split_to_array(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v11 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if int32(3) <= v11 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v20 = int32(0)
	goto L3
L3:
	;
	F_parse_re_flags(m, v7+int32(8), v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return int32(0)
L5:
	;
	v20 = v15
	goto L3
L6:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)))
	if v23 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v26 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)) = uint8(v26)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v29 = F_pg_detoast_datum_packed(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
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
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L22
	}
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v32 = F_pg_detoast_datum_packed(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v36 = int32(0)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v39 = int32(1)
	v41 = F_setup_regexp_matches(m, v29, v32, v7+int32(8), v36, v37, v36, v39, v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v43 <= v44 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v48 = v2
	goto L16
L14:
	;
	v66 = v2
	goto L15
L15:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_regexp_split_to_array[0]))
	v70 = F_makeArrayResult(m, v66, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L21
	}
L16:
	;
	v50 = F_build_regexp_split_result(m, v41)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L18
	}
L17:
	;
	v66 = v56
	goto L15
L18:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_regexp_split_to_array[0]))
	v56 = F_accumArrayResult(m, v48, v50, int32(0), int32(25), v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v60 = v58 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v60
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v60 <= v62 {
		v48 = v56
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	m.G0 = v7 + int32(16)
	return v70
L22:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_regexp_split_to_array_0)
	F_errmsg(m, int32(_a_F_regexp_split_to_array_1), v7)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_regexp_split_to_array_2), int32(1826), int32(_a_F_regexp_split_to_array_3))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
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
