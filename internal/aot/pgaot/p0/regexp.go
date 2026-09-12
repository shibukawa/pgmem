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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
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
	v19 = v2
	goto L3
L3:
	;
	F_parse_re_flags(m, v7+int32(8), v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return int32(0)
L5:
	;
	v19 = v15
	goto L3
L6:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)))
	if v22 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v25 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)) = uint8(v25)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v28 = F_pg_detoast_datum_packed(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
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
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L22
	}
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v31 = F_pg_detoast_datum_packed(m, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v35 = int32(0)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v38 = int32(1)
	v40 = F_setup_regexp_matches(m, v28, v31, v7+int32(8), v35, v36, v35, v38, v38)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v42 <= v43 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v47 = v2
	goto L16
L14:
	;
	v65 = v2
	goto L15
L15:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v69 = F_makeArrayResult(m, v65, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L21
	}
L16:
	;
	v49 = F_build_regexp_split_result(m, v40)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L18
	}
L17:
	;
	v65 = v55
	goto L15
L18:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v55 = F_accumArrayResult(m, v47, v49, int32(0), int32(25), v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	v59 = v57 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v59 <= v61 {
		v47 = v55
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
	return v69
L22:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(697197)
	F_errmsg(m, int32(251948), v7)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(507189), int32(1826), int32(24396))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
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
