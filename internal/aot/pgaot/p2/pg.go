package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_RemovePgTempFiles(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	v4 = m.G0
	v6 = v4 - int32(1120)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = int32(234164)
	v16 = F_pg_snprintf(m, v6+int32(48), int32(1060), int32(176250), v6+int32(32))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_RemovePgTempFilesInDir(m, v6+int32(48), int32(1), int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_RemovePgTempRelationFiles(m, int32(360346))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v28 = F_AllocateDir(m, int32(485645))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v32 = F_ReadDirExtended(m, v28, int32(485645), int32(15))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v32 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v35 = v32
	goto L10
L8:
	;
	goto L9
L9:
	;
	F_FreeDir(m, v28)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L24
	}
L10:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+19)))
	if v37 != int32(46) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L9
L12:
	;
	v90 = F_ReadDirExtended(m, v28, int32(485645), int32(15))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L22
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = int32(234164)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = int32(550670)
	v54 = v35 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(485645)
	v64 = F_pg_snprintf(m, v6+int32(48), int32(1060), int32(176158), v6+int32(16))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L18
	}
L14:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+20)))
	if v40 == int32(0) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+20)))
	if v43 != int32(46) {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+21)))
	if v46 == int32(0) {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L13
L18:
	;
	F_RemovePgTempFilesInDir(m, v6+int32(48), int32(1), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = int32(550670)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(485645)
	v81 = F_pg_snprintf(m, v6+int32(48), int32(1060), int32(176161), v6)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_RemovePgTempRelationFiles(m, v6+int32(48))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	goto L12
L22:
	;
	if v90 != 0 {
		v35 = v90
		goto L10
	} else {
		goto L23
	}
L23:
	;
	goto L11
L24:
	;
	m.G0 = v6 + int32(1120)
	return
}
func F_pg_any_to_server(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l1 <= int32(0) {
		v92 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return v92
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _consts[358]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v15 != l2 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v17 = l2
	goto L5
L4:
	;
	v17 = int32(0)
	goto L5
L5:
	;
	if v17 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v15*int32(28))+uint32(_consts[1222])))
	v25 = m.T0[v24].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	if v15 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	return int32(0)
L10:
	;
	if l1 == v25 {
		v92 = l0
		goto L1
	} else {
		goto L11
	}
L11:
	;
	F_report_invalid_encoding(m, v15, l0+v25, l1-v25)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	if base.Ui32(int32(34)) < base.Ui32(l2) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	goto L15
L15:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _consts[357]))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v85 == l2 {
		goto L33
	} else {
		goto L34
	}
L16:
	;
	v53 = int32(0)
	goto L23
L17:
	;
	goto L16
L18:
	;
	goto L19
L19:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l2*int32(28))+uint32(_consts[1222])))
	v44 = m.T0[v43].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	if l1 == v44 {
		v92 = l0
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_report_invalid_encoding(m, l2, l0+v44, l1-v44)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	v56 = l0 + v53
	v57 = int32(*(*int8)(unsafe.Add(mBase, uint32(v56))))
	if int32(0) < v57 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L9
	} else {
		goto L29
	}
L25:
	;
	v61 = v53 + int32(1)
	if l1 != v61 {
		v53 = v61
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	goto L24
L28:
	;
	v92 = l0
	goto L1
L29:
	;
	F_errcode(m, int32(17301634))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L9
	} else {
		goto L30
	}
L30:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v70
	v73 = *(*int32)(unsafe.Add(mBase, _consts[359]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v73
	F_errmsg(m, int32(29564), v8)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L9
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(490789), int32(724), int32(212927))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L9
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	v88 = F_perform_default_encoding_conversion(m, l0, l1, int32(1))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L9
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v90 = F_pg_do_encoding_conversion(m, l0, l1, l2, v15)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L9
	} else {
		goto L37
	}
L36:
	;
	v92 = v88
	goto L1
L37:
	;
	v92 = v90
	goto L1
}
func F_pg_ascii_verifystr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	v3 = int32(0)
	v8 = base.B2i32(l1 != v3)
	if l0&int32(3) == v3 {
		v34 = l0
		v36 = l1
		v37 = v8
		goto L4
	} else {
		goto L5
	}
L1:
	;
	if v107 != 0 {
		goto L27
	} else {
		goto L28
	}
L2:
	;
	v107 = int32(0)
	goto L1
L3:
	;
	v85 = v78
	v87 = v80
	goto L21
L4:
	;
	if v37 == int32(0) {
		goto L2
	} else {
		goto L12
	}
L5:
	;
	if l1 == int32(0) {
		v34 = l0
		v36 = l1
		v37 = v8
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v17 = l0
	v19 = l1
	goto L7
L7:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v22 == int32(0) {
		v78 = v17
		v80 = v19
		goto L3
	} else {
		goto L9
	}
L8:
	;
	v34 = v29
	v36 = v25
	v37 = v27
	goto L4
L9:
	;
	v24 = int32(1)
	v25 = v19 - v24
	v26 = int32(0)
	v27 = base.B2i32(v25 != v26)
	v29 = v17 + v24
	if v29&int32(3) == v26 {
		v34 = v29
		v36 = v25
		v37 = v27
		goto L4
	} else {
		goto L10
	}
L10:
	;
	if v25 != 0 {
		v17 = v29
		v19 = v25
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v41 == int32(0) {
		v71 = v34
		v73 = v36
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v73 == int32(0) {
		goto L2
	} else {
		goto L20
	}
L14:
	;
	if base.Ui32(v36) < base.Ui32(int32(4)) {
		v71 = v34
		v73 = v36
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v51 = v34
	v53 = v36
	goto L16
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v58 = v57 ^ int32(0)
	v61 = int32(-2139062144)
	if (int32(16843008)-v58|v58)&v61 != v61 {
		v78 = v51
		v80 = v53
		goto L3
	} else {
		goto L18
	}
L17:
	;
	v71 = v66
	v73 = v68
	goto L13
L18:
	;
	v65 = int32(4)
	v66 = v51 + v65
	v68 = v53 - v65
	if base.Ui32(int32(3)) < base.Ui32(v68) {
		v51 = v66
		v53 = v68
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v78 = v71
	v80 = v73
	goto L3
L21:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	if int32(0) == v90 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L2
L23:
	;
	v107 = v85
	goto L1
L24:
	;
	goto L25
L25:
	;
	v92 = int32(1)
	v95 = v87 - v92
	if v95 != 0 {
		v85 = v85 + v92
		v87 = v95
		goto L21
	} else {
		goto L26
	}
L26:
	;
	goto L22
L27:
	;
	v109 = v107 - l0
	goto L29
L28:
	;
	v109 = l1
	goto L29
L29:
	;
	return v109
}
func F_pg_backend_pid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, _consts[350]))
	return v3
}
func F_pg_base64_enc_len(m *base.Module, l0 int32, l1 int32) int64 {
	var v4 int32
	_ = v4
	var v7 int64
	_ = v7
	var v10 int64
	_ = v10
	v4 = base.I32_div_u_s(l1, int32(57))
	v7 = int64(2)
	v10 = base.I64_div_u_s(base.I64_extend_i32_u(l1)+v7, int64(3))
	return base.I64_extend_i32_u(v4) + v10<<(uint(v7)%64)
}
func F_pg_create_physical_replication_slot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_get_call_result_type(m, l0, int32(0), v8)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 == int32(1) {
			F_CheckSlotPermissions(m)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_CheckSlotRequirements(m)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = int32(0)
					F_ReplicationSlotCreate(m, v12, v24, base.B2i32(v10 != v24)<<(uint(int32(1))%32), v24, v24, v24)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						if v11 != 0 {
							F_ReplicationSlotReserveWal(m)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								F_ReplicationSlotMarkDirty(m)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int32(0)
								} else {
									F_ReplicationSlotSave(m)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return int32(0)
									} else {
										v40 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v8)+6)) = uint8(v40)
										v44 = *(*int32)(unsafe.Add(mBase, _consts[642]))
										*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v44 + int32(24)
										v48 = *(*int64)(unsafe.Add(mBase, uint32(v44)+104))
										v49 = F_Int64GetDatum(m, v48)
										mBase = m.M
										v50 = m.ExcPending
										if v50 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v49
											v60 = v40
											*(*uint8)(unsafe.Add(mBase, uint32(v8)+7)) = uint8(v60)
											v63 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
											v68 = F_heap_form_tuple(m, v63, v8+int32(8), v8+int32(6))
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int32(0)
											} else {
												v70 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
												v71 = F_HeapTupleHeaderGetDatum(m, v70)
												mBase = m.M
												v72 = m.ExcPending
												if v72 != 0 {
													return int32(0)
												} else {
													F_ReplicationSlotRelease(m)
													mBase = m.M
													v74 = m.ExcPending
													if v74 != 0 {
														return int32(0)
													} else {
														m.G0 = v8 + int32(16)
														return v71
													}
												}
											}
										}
									}
								}
							}
						} else {
							v52 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+6)) = uint8(v52)
							v55 = *(*int32)(unsafe.Add(mBase, _consts[642]))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v55 + int32(24)
							v60 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+7)) = uint8(v60)
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
							v68 = F_heap_form_tuple(m, v63, v8+int32(8), v8+int32(6))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
								v71 = F_HeapTupleHeaderGetDatum(m, v70)
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int32(0)
								} else {
									F_ReplicationSlotRelease(m)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										m.G0 = v8 + int32(16)
										return v71
									}
								}
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(365102), int32(0))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(491427), int32(77), int32(84557))
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_pg_cryptohash_final(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	if l0 == int32(0) {
		return int32(-1)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v9 = m.Env.Pgmem_hash_final(m, v8, l1, l2)
		mBase = m.M
		if int32(0) <= v9 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1)
			return int32(-1)
		}
	}
}
func F_pg_cryptohash_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	if l0 != 0 {
		v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		m.Env.Pgmem_hash_free(m, v2)
		mBase = m.M
		v6 = F___memset(m, l0, int32(0), int32(12))
		mBase = m.M
		F_pfree(m, l0)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_pg_cursor(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[1238]))
	F_hash_seq_init(m, v6+int32(-20), v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = F_hash_seq_search(m, v6+int32(-20))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if v24 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v28 = v24
	goto L8
L6:
	;
	goto L7
L7:
	;
	m.G0 = v8 - int32(-64)
	return int32(0)
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+64))
	v34 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v6+int32(-52)))) = uint16(v34)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v34
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+136)))
	if v38 != int32(1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v81 = F_hash_seq_search(m, v6+int32(-20))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L17
	}
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v33)+32))
	if v41 == int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v45 = F_cstring_to_text(m, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v45
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v33)+32))
	v49 = F_cstring_to_text(m, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v49
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v33)+76))
	v53 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v52 & v53
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(base.Ui32(v52)>>(uint(v53)%32)) & v53
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(base.Ui32(v52)>>(uint(int32(5))%32)) & v53
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v33)+128))
	v67 = F_Int64GetDatum(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v67
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	F_tuplestore_putvalues(m, v70, v71, v6+int32(-48), v6+int32(-56))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	goto L10
L17:
	;
	if v81 != 0 {
		v28 = v81
		goto L8
	} else {
		goto L18
	}
L18:
	;
	goto L9
}
func F_pg_database_size_oid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_SearchSysCacheExists(m, int32(21), v10, v2, v2, v2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 != 0 {
			v18 = F_calculate_database_size(m, v10)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v18 == int64(0) {
					v22 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
					v27 = int32(0)
					m.G0 = v7 + int32(16)
					return v27
				} else {
					v25 = F_Int64GetDatum(m, v18)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = v25
						m.G0 = v7 + int32(16)
						return v27
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(67137668))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10
					F_errmsg(m, int32(68981), v7)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(495188), int32(180), int32(431939))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_euccn2wchar_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	v4 = int32(0)
	if l2 <= v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v9
	return v9
L2:
	;
	goto L3
L3:
	;
	v13 = l0
	v14 = l1
	v15 = l2
	v17 = v4
	goto L4
L4:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	switch v19 - int32(142) {
	case 0:
		goto L11
	case 1:
		goto L10
	default:
		goto L9
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76))) = int32(0)
	return v79
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v64
	v69 = v17 + int32(1)
	v71 = v14 + int32(4)
	v72 = v15 + v65
	if int32(0) < v72 {
		v13 = v66
		v14 = v71
		v15 = v72
		v17 = v69
		goto L4
	} else {
		goto L19
	}
L8:
	;
	v64 = v60
	v65 = int32(-3)
	v66 = v13 + int32(3)
	goto L7
L9:
	;
	if v19 == int32(0) {
		v76 = v14
		v79 = v17
		goto L6
	} else {
		goto L14
	}
L10:
	;
	if base.Ui32(v15) < base.Ui32(int32(3)) {
		v76 = v14
		v79 = v17
		goto L6
	} else {
		goto L13
	}
L11:
	;
	if base.Ui32(v15) < base.Ui32(int32(3)) {
		v76 = v14
		v79 = v17
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v28 = v24<<(uint(int32(8))%32) | int32(9306112)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v28
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
	v60 = v28 | v30
	goto L8
L13:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v38 = v34<<(uint(int32(8))%32) | int32(9371648)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v38
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
	v60 = v38 | v40
	goto L8
L14:
	;
	if base.I32_extend8_s(v19) < int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if v15 == int32(1) {
		v76 = v14
		v79 = v17
		goto L6
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v64 = v19
	v65 = int32(-1)
	v66 = v13 + int32(1)
	goto L7
L18:
	;
	v50 = v19 << (uint(int32(8)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v50
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v64 = v50 | v52
	v65 = int32(-2)
	v66 = v13 + int32(2)
	goto L7
L19:
	;
	v76 = v71
	v79 = v69
	goto L6
}
func F_pg_euctw_verifychar(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	v5 = int32(-1)
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v6 - int32(142) {
	case 0:
		if l1 < int32(4) {
			v49 = v5
		} else {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
			if base.Ui32((v11+int32(88))&int32(255)) < base.Ui32(int32(249)) {
				v49 = v5
			} else {
				v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
				if base.Ui32(int32(93)) < base.Ui32((v18+int32(95))&int32(255)) {
					v49 = v5
				} else {
					v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
					if base.Ui32(int32(94)) <= base.Ui32((v25+int32(95))&int32(255)) {
						v49 = v5
					} else {
						v47 = int32(4)
						v49 = v47
					}
				}
			}
		}
	case 1:
		v49 = v5
	default:
		if int32(0) <= base.I32_extend8_s(v6) {
			v47 = int32(1)
			v49 = v47
		} else {
			v37 = int32(2)
			if l1 < v37 {
				v49 = v5
			} else {
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
				if base.Ui32(int32(93)) < base.Ui32((v40+int32(95))&int32(255)) {
					v49 = v5
				} else {
					v47 = v37
					v49 = v47
				}
			}
		}
	}
	return v49
}
func F_pg_gb18030_verifystr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	if l1 <= int32(0) {
		v73 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v73 - l0
L2:
	;
	v9 = l1
	v11 = l0
	goto L3
L3:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v14 = base.I32_extend8_s(v13)
	if int32(0) <= v14 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v73 = v66
	goto L1
L5:
	;
	v66 = v64 + v11
	v67 = v9 - v64
	if int32(0) < v67 {
		v9 = v67
		v11 = v66
		goto L3
	} else {
		goto L24
	}
L6:
	;
	if v14 == int32(0) {
		v73 = v11
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v9) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v64 = int32(1)
	goto L5
L10:
	;
	if v13 == int32(128) {
		v73 = v11
		goto L1
	} else {
		goto L20
	}
L11:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if base.Ui32(int32(9)) < base.Ui32((v22-int32(48))&int32(255)) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if v9 == int32(1) {
		v73 = v11
		goto L1
	} else {
		goto L19
	}
L14:
	;
	if v13 == int32(128) {
		v73 = v11
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v13 == int32(255) {
		v73 = v11
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+2)))
	if base.Ui32((v33+int32(1))&int32(255)) < base.Ui32(int32(130)) {
		v73 = v11
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+3)))
	if base.Ui32(int32(10)) <= base.Ui32((v40-int32(48))&int32(255)) {
		v73 = v11
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v64 = int32(4)
	goto L5
L19:
	;
	goto L10
L20:
	;
	if v13 == int32(255) {
		v73 = v11
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v54 = int32(2)
	v55 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v55 < int32(-1) {
		v64 = v54
		goto L5
	} else {
		goto L22
	}
L22:
	;
	if base.Ui32((v55-int32(127))&int32(255)) < base.Ui32(int32(193)) {
		v73 = v11
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v64 = v54
	goto L5
L24:
	;
	goto L4
}
func F_pg_generic_charinc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	v8 = l0 + l1 - int32(1)
	v10 = *(*int32)(unsafe.Add(mBase, _consts[358]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11*int32(28))+uint32(_consts[1223])))
	goto L1
L1:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v22 != int32(255) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return base.B2i32(v22 != int32(255))
L3:
	;
	v26 = v22 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v26)
	v28 = m.T0[v16].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	goto L2
L6:
	;
	return int32(0)
L7:
	;
	if v28 != l1 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L5
}
func F_pg_get_catalog_foreign_keys(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int64
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v14 == int32(0) {
		v17 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int32(4486928)
			v22 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v24
			v29 = F_get_call_result_type(m, l0, int32(0), v11+int32(16))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				if v29 != int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v133 = m.ExcPending
					if v133 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(365102), int32(0))
						mBase = m.M
						v137 = m.ExcPending
						if v137 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(496562), int32(510), int32(112330))
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
					v34 = F_BlessTupleDesc(m, v33)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v34
						v39 = F_palloc(m, int32(28))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							F_fmgr_info(m, int32(750), v39)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v39
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v22
								v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
								v51 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
								if base.Ui64(v51) <= base.Ui64(int64(218)) {
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
									v55 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(v11)+12)) = uint16(v55)
									*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v55
									v61 = base.I32_wrap_i64(v51) * int32(20)
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+uint32(_consts[1098])))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v64
									v68 = *(*int32)(unsafe.Add(mBase, uint32(v61)+uint32(_consts[1099])))
									v72 = *(*int32)(unsafe.Add(mBase, uint32(v61)+uint32(_consts[1100])))
									v75 = F_FunctionCall3Coll(m, v54, v55, v72, int32(25), int32(-1))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v68
										*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v75
										v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+uint32(_consts[1101]))))
										v85 = *(*int32)(unsafe.Add(mBase, uint32(v61)+uint32(_consts[1102])))
										v88 = F_FunctionCall3Coll(m, v54, int32(0), v85, int32(25), int32(-1))
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v81
											v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+uint32(_consts[1103]))))
											*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v93
											*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v88
											v96 = *(*int32)(unsafe.Add(mBase, uint32(v50)+28))
											v101 = F_heap_form_tuple(m, v96, v11+int32(16), v11+int32(8))
											mBase = m.M
											v102 = m.ExcPending
											if v102 != 0 {
												return int32(0)
											} else {
												v103 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
												*(*int64)(unsafe.Add(mBase, uint32(v50))) = v103 + int64(1)
												v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v107)+20)) = int32(1)
												v110 = *(*int32)(unsafe.Add(mBase, uint32(v101)+16))
												v111 = F_HeapTupleHeaderGetDatum(m, v110)
												mBase = m.M
												v112 = m.ExcPending
												if v112 != 0 {
													return int32(0)
												} else {
													v125 = v111
													m.G0 = v11 + int32(48)
													return v125
												}
											}
										}
									}
								} else {
									F_end_MultiFuncCall(m, l0)
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return int32(0)
									} else {
										v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v115)+20)) = int32(2)
										v118 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v118)
										v125 = int32(0)
										m.G0 = v11 + int32(48)
										return v125
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
		v51 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
		if base.Ui64(v51) <= base.Ui64(int64(218)) {
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
			v55 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v11)+12)) = uint16(v55)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v55
			v61 = base.I32_wrap_i64(v51) * int32(20)
			v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+uint32(_consts[1098])))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v64
			v68 = *(*int32)(unsafe.Add(mBase, uint32(v61)+uint32(_consts[1099])))
			v72 = *(*int32)(unsafe.Add(mBase, uint32(v61)+uint32(_consts[1100])))
			v75 = F_FunctionCall3Coll(m, v54, v55, v72, int32(25), int32(-1))
			mBase = m.M
			v76 = m.ExcPending
			if v76 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v68
				*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v75
				v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+uint32(_consts[1101]))))
				v85 = *(*int32)(unsafe.Add(mBase, uint32(v61)+uint32(_consts[1102])))
				v88 = F_FunctionCall3Coll(m, v54, int32(0), v85, int32(25), int32(-1))
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v81
					v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+uint32(_consts[1103]))))
					*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v93
					*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v88
					v96 = *(*int32)(unsafe.Add(mBase, uint32(v50)+28))
					v101 = F_heap_form_tuple(m, v96, v11+int32(16), v11+int32(8))
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return int32(0)
					} else {
						v103 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
						*(*int64)(unsafe.Add(mBase, uint32(v50))) = v103 + int64(1)
						v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v107)+20)) = int32(1)
						v110 = *(*int32)(unsafe.Add(mBase, uint32(v101)+16))
						v111 = F_HeapTupleHeaderGetDatum(m, v110)
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return int32(0)
						} else {
							v125 = v111
							m.G0 = v11 + int32(48)
							return v125
						}
					}
				}
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v114 = m.ExcPending
			if v114 != 0 {
				return int32(0)
			} else {
				v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v115)+20)) = int32(2)
				v118 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v118)
				v125 = int32(0)
				m.G0 = v11 + int32(48)
				return v125
			}
		}
	}
}
func F_pg_get_indexdef_ext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	v2 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v13 != 0 {
		v14 = int32(7)
	} else {
		v14 = int32(2)
	}
	v16 = F_pg_get_indexdef_worker(m, v3, v4, v2, base.B2i32(v4 != v2), v2, v2, v2, v14, int32(1))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		if v16 == int32(0) {
			v22 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
			return int32(0)
		} else {
			v26 = F_cstring_to_text(m, v16)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v16)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					return v26
				}
			}
		}
	}
}
func F_pg_get_ruledef_ext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v6 != 0 {
		v7 = int32(7)
	} else {
		v7 = int32(2)
	}
	v8 = F_pg_get_ruledef_worker(m, v3, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 == int32(0) {
			v14 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
			return int32(0)
		} else {
			v18 = F_cstring_to_text(m, v8)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v8)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					return v18
				}
			}
		}
	}
}
func F_pg_get_statisticsobj_worker(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v367 int32
	_ = v367
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v434 int32
	_ = v434
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v645 int32
	_ = v645
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v741 int32
	_ = v741
	var v749 int32
	_ = v749
	var v755 int32
	_ = v755
	var v760 int32
	_ = v760
	v4 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(192)
	m.G0 = v23
	v26 = F_SearchSysCache1(m, int32(64), l0)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L3
	} else {
		goto L161
	}
L2:
	;
	m.G0 = v23 + int32(192)
	return v741
L3:
	;
	return int32(0)
L4:
	;
	if v26 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if l2 != 0 {
		v741 = int32(0)
		goto L2
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v48 = F_heap_attisnull(m, v26, int32(9), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L3
	} else {
		goto L12
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = l0
	F_errmsg_internal(m, int32(41722), v23)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	F_errfinish(m, int32(490777), int32(1680), int32(219147))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+22)))
	v52 = v50 + v51
	if v48 != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	F_initStringInfo(m, v23+int32(120))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L3
	} else {
		goto L24
	}
L14:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v76 = v68
	v77 = v62
	v78 = v67
	v79 = v4
	v80 = v74
	goto L13
L15:
	;
	v76 = v70
	v77 = v4
	v78 = v71
	v79 = int32(1)
	v80 = int32(0)
	goto L13
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v52)+96))
	v70 = v55
	v71 = v52 + int32(96)
	goto L15
L17:
	;
	goto L18
L18:
	;
	v58 = F_SysCacheGetAttrNotNull(m, int32(64), v26, int32(9))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	v60 = F_text_to_cstring(m, v58)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v62 = F_stringToNode(m, v60)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	F_pfree(m, v60)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v67 = v52 + int32(96)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v52)+96))
	if v62 != 0 {
		goto L14
	} else {
		goto L23
	}
L23:
	;
	v70 = v68
	v71 = v67
	goto L15
L24:
	;
	if l1 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v52)+72))
	v90 = F_get_namespace_name_or_temp(m, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L3
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v455 = int32(1)
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	if v456 <= int32(0) {
		goto L113
	} else {
		goto L114
	}
L28:
	;
	F_initStringInfo(m, v23+int32(136))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	if v90 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v96 = F_quote_identifier(m, v90)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L3
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v108 = F_quote_identifier(m, v52+int32(8))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L3
	} else {
		goto L35
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+112)) = v96
	F_appendStringInfo(m, v23+int32(136), int32(588192), v23+int32(112))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	F_appendStringInfoString(m, v23+int32(136), v108)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v23)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+96)) = v112
	F_appendStringInfo(m, v23+int32(120), int32(197051), v23+int32(96))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L3
	} else {
		goto L37
	}
L37:
	;
	v123 = F_SysCacheGetAttrNotNull(m, int32(64), v26, int32(8))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L3
	} else {
		goto L42
	}
L38:
	;
	F_appendStringInfoString(m, v23+int32(120), int32(725763))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L3
	} else {
		goto L111
	}
L39:
	;
	F_appendStringInfoString(m, v23+int32(120), int32(668382))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L3
	} else {
		goto L90
	}
L40:
	;
	if v134 < int32(2) {
		goto L38
	} else {
		goto L89
	}
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L3
	} else {
		goto L86
	}
L42:
	;
	v125 = F_pg_detoast_datum(m, v123)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L3
	} else {
		goto L43
	}
L43:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	if v127 != int32(1) {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
	if v130 != 0 {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	if v131 != int32(18) {
		goto L41
	} else {
		goto L46
	}
L46:
	;
	v134 = v80 + v76
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
	if v135 <= int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v325 = int32(0)
	v329 = v4
	v330 = v4
	goto L40
L48:
	;
	goto L49
L49:
	;
	v140 = v125 + int32(24)
	v142 = v135 & int32(3)
	if base.Ui32(v135) < base.Ui32(int32(4)) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if v142 != 0 {
		goto L73
	} else {
		goto L74
	}
L51:
	;
	v146 = int32(0)
	v222 = v146
	v226 = v146
	v230 = v4
	v231 = v4
	goto L50
L52:
	;
	goto L53
L53:
	;
	v150 = int32(0)
	v152 = v150
	v156 = v150
	v160 = v4
	v161 = v4
	v171 = v4
	goto L54
L54:
	;
	v172 = int32(1)
	v174 = v152 + v140
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	switch v175 - int32(100) {
	case 0:
		v182 = v160
		v183 = v161
		v184 = v172
		goto L56
	default:
		v180 = v160
		v181 = v161
		goto L57
	case 2:
		goto L59
	case 9:
		goto L58
	}
L55:
	;
	v222 = v218
	v226 = v214
	v230 = v215
	v231 = v216
	goto L50
L56:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
	switch v185 - int32(100) {
	case 0:
		v192 = v172
		v193 = v182
		v194 = v183
		goto L60
	default:
		v190 = v182
		v191 = v183
		goto L61
	case 2:
		goto L62
	case 9:
		goto L63
	}
L57:
	;
	v182 = v180
	v183 = v181
	v184 = v156
	goto L56
L58:
	;
	v180 = v160
	v181 = int32(1)
	goto L57
L59:
	;
	v180 = int32(1)
	v181 = v161
	goto L57
L60:
	;
	v195 = int32(1)
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+2)))
	switch v197 - int32(100) {
	case 0:
		v204 = v193
		v205 = v194
		v206 = v195
		goto L64
	default:
		v202 = v193
		v203 = v194
		goto L65
	case 2:
		goto L66
	case 9:
		goto L67
	}
L61:
	;
	v192 = v184
	v193 = v190
	v194 = v191
	goto L60
L62:
	;
	v190 = int32(1)
	v191 = v183
	goto L61
L63:
	;
	v190 = v182
	v191 = int32(1)
	goto L61
L64:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+3)))
	switch v207 - int32(100) {
	case 0:
		v214 = v195
		v215 = v204
		v216 = v205
		goto L68
	default:
		v212 = v204
		v213 = v205
		goto L69
	case 2:
		goto L70
	case 9:
		goto L71
	}
L65:
	;
	v204 = v202
	v205 = v203
	v206 = v192
	goto L64
L66:
	;
	v202 = int32(1)
	v203 = v194
	goto L65
L67:
	;
	v202 = v193
	v203 = int32(1)
	goto L65
L68:
	;
	v217 = int32(4)
	v218 = v152 + v217
	v220 = v171 + v217
	if v220 != v135&int32(2147483644) {
		v152 = v218
		v156 = v214
		v160 = v215
		v161 = v216
		v171 = v220
		goto L54
	} else {
		goto L72
	}
L69:
	;
	v214 = v206
	v215 = v212
	v216 = v213
	goto L68
L70:
	;
	v212 = int32(1)
	v213 = v205
	goto L69
L71:
	;
	v212 = v204
	v213 = int32(1)
	goto L69
L72:
	;
	goto L55
L73:
	;
	v242 = v222
	v244 = v226
	v248 = int32(0)
	v250 = v230
	v251 = v231
	goto L76
L74:
	;
	v283 = v226
	v287 = v230
	v288 = v231
	goto L75
L75:
	;
	if v283&int32(1) == int32(0) {
		v325 = v283
		v329 = v287
		v330 = v288
		goto L40
	} else {
		goto L83
	}
L76:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242+v140))))
	switch v264 - int32(100) {
	case 0:
		v271 = int32(1)
		v272 = v250
		v273 = v251
		goto L78
	default:
		v269 = v250
		v270 = v251
		goto L79
	case 2:
		goto L80
	case 9:
		goto L81
	}
L77:
	;
	v283 = v271
	v287 = v272
	v288 = v273
	goto L75
L78:
	;
	v274 = int32(1)
	v277 = v248 + v274
	if v277 != v142 {
		v242 = v242 + v274
		v244 = v271
		v248 = v277
		v250 = v272
		v251 = v273
		goto L76
	} else {
		goto L82
	}
L79:
	;
	v271 = v244
	v272 = v269
	v273 = v270
	goto L78
L80:
	;
	v269 = int32(1)
	v270 = v251
	goto L79
L81:
	;
	v269 = v250
	v270 = int32(1)
	goto L79
L82:
	;
	goto L77
L83:
	;
	if v287 == int32(0) {
		v325 = v283
		v329 = v287
		v330 = v288
		goto L40
	} else {
		goto L84
	}
L84:
	;
	if base.B2i32(v134 < int32(2))|v288 != 0 {
		goto L38
	} else {
		goto L85
	}
L85:
	;
	v347 = v283
	v351 = v287
	v352 = v288
	goto L39
L86:
	;
	F_errmsg_internal(m, int32(25015), int32(0))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L3
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(490777), int32(1729), int32(219147))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L3
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	v347 = v325
	v351 = v329
	v352 = v330
	goto L39
L90:
	;
	if v347&int32(1) != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	F_appendStringInfoString(m, v23+int32(120), int32(108509))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L3
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	if v351 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	goto L93
L95:
	;
	if v352 != 0 {
		goto L103
	} else {
		goto L104
	}
L96:
	;
	v391 = v347
	goto L95
L97:
	;
	goto L98
L98:
	;
	v377 = int32(1)
	if v347&v377 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v382 = int32(727439)
	goto L101
L100:
	;
	v382 = int32(738681)
	goto L101
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v382
	F_appendStringInfo(m, v23+int32(120), int32(167619), v23+int32(80))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L3
	} else {
		goto L102
	}
L102:
	;
	v391 = v377
	goto L95
L103:
	;
	if v391&int32(1) != 0 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	goto L105
L105:
	;
	F_appendStringInfoChar(m, v23+int32(120), int32(41))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L3
	} else {
		goto L110
	}
L106:
	;
	v396 = int32(727439)
	goto L108
L107:
	;
	v396 = int32(738681)
	goto L108
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v396
	F_appendStringInfo(m, v23+int32(120), int32(35661), v23-int32(-64))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L3
	} else {
		goto L109
	}
L109:
	;
	goto L105
L110:
	;
	goto L38
L111:
	;
	goto L27
L112:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v540 = F_get_rel_name(m, v539)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L3
	} else {
		goto L127
	}
L113:
	;
	v519 = int32(0)
	goto L112
L114:
	;
	goto L115
L115:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v463 = int32(*(*int16)(unsafe.Add(mBase, uint32(v52)+104)))
	v465 = F_get_attname(m, v462, v463, int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L3
	} else {
		goto L116
	}
L116:
	;
	v467 = F_quote_identifier(m, v465)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L3
	} else {
		goto L117
	}
L117:
	;
	F_appendStringInfoString(m, v23+int32(120), v467)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L3
	} else {
		goto L118
	}
L118:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v52)+96))
	if v471 < int32(2) {
		v519 = v455
		goto L112
	} else {
		goto L119
	}
L119:
	;
	v476 = v455
	goto L120
L120:
	;
	v499 = int32(*(*int16)(unsafe.Add(mBase, uint32(v52+int32(104)+v476<<(uint(int32(1))%32)))))
	F_appendStringInfoString(m, v23+int32(120), int32(727439))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L3
	} else {
		goto L122
	}
L121:
	;
	v519 = v516
	goto L112
L122:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v509 = F_get_attname(m, v507, v499, int32(0))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L3
	} else {
		goto L123
	}
L123:
	;
	v511 = F_quote_identifier(m, v509)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L3
	} else {
		goto L124
	}
L124:
	;
	F_appendStringInfoString(m, v23+int32(120), v511)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L3
	} else {
		goto L125
	}
L125:
	;
	v516 = v476 + int32(1)
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v52)+96))
	if v516 < v517 {
		v476 = v516
		goto L120
	} else {
		goto L126
	}
L126:
	;
	goto L121
L127:
	;
	if v540 == int32(0) {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v546 = F_palloc0(m, int32(80))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L3
	} else {
		goto L129
	}
L129:
	;
	v549 = F_palloc0(m, int32(136))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L3
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v549)+24)) = int32(1)
	v553 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(v549)+21)) = uint8(v553)
	*(*int32)(unsafe.Add(mBase, uint32(v549)+16)) = v544
	v556 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v549)+12)) = v556
	*(*int32)(unsafe.Add(mBase, uint32(v549))) = int32(101)
	v562 = F_makeAlias(m, v540, v556)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L3
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v549)+8)) = v562
	*(*int32)(unsafe.Add(mBase, uint32(v549)+4)) = v562
	v566 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v549)+124)) = uint16(v566)
	v568 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v549)+20)) = uint8(v568)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v549
	*(*int32)(unsafe.Add(mBase, uint32(v23)+136)) = v549
	v575 = F_list_make1_impl(m, int32(1), v23+int32(60))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L3
	} else {
		goto L132
	}
L132:
	;
	v577 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v546)+20)) = v577
	*(*int64)(unsafe.Add(mBase, uint32(v546)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v546))) = v575
	F_set_rtable_names(m, v546, v577, v577)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L3
	} else {
		goto L133
	}
L133:
	;
	F_set_simple_column_names(m, v546)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L3
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+56)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v23)+176)) = v546
	v593 = F_list_make1_impl(m, int32(1), v23+int32(56))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L3
	} else {
		goto L135
	}
L135:
	;
	if v79 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	if l1 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L137:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	if v595 <= int32(0) {
		goto L136
	} else {
		goto L138
	}
L138:
	;
	v598 = v519
	v600 = v556
	goto L139
L139:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v618+v600<<(uint(int32(2))%32))))
	F_initStringInfo(m, v23+int32(176))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L3
	} else {
		goto L141
	}
L140:
	;
	goto L136
L141:
	;
	v627 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+168)) = uint8(v627)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+152)) = v627
	*(*int64)(unsafe.Add(mBase, uint32(v23)+144)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+140)) = v593
	*(*int32)(unsafe.Add(mBase, uint32(v23)+172)) = v627
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+171)) = uint8(v627)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+164)) = v627
	*(*int64)(unsafe.Add(mBase, uint32(v23)+156)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+136)) = v23 + int32(176)
	v645 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+169)) = uint16(v645)
	F_get_rule_expr(m, v622, v23+int32(136), v627)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L3
	} else {
		goto L142
	}
L142:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v23)+176))
	if int32(0) < v598 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	F_appendStringInfoString(m, v23+int32(120), int32(727439))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L3
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	if v622 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	goto L145
L147:
	;
	v678 = int32(1)
	v681 = v600 + v678
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	if v681 < v682 {
		v598 = v598 + v678
		v600 = v681
		goto L139
	} else {
		goto L154
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v652
	F_appendStringInfo(m, v23+int32(120), int32(657170), v23+int32(48))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L3
	} else {
		goto L153
	}
L149:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v622)))
	switch v662 - int32(15) {
	case 0:
		goto L151
	default:
		goto L148
	case 4, 23, 24, 25, 26, 33:
		goto L150
	}
L150:
	;
	F_appendStringInfoString(m, v23+int32(120), v652)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L3
	} else {
		goto L152
	}
L151:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v622)+16))
	switch v665 {
	case 0, 3:
		goto L150
	default:
		goto L148
	}
L152:
	;
	goto L147
L153:
	;
	goto L147
L154:
	;
	goto L140
L155:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v708 = F_generate_relation_name(m, v706, int32(0))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L3
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	F_ReleaseCatCache(m, v26)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L3
	} else {
		goto L160
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v708
	F_appendStringInfo(m, v23+int32(120), int32(197180), v23+int32(32))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L3
	} else {
		goto L159
	}
L159:
	;
	goto L157
L160:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v23)+120))
	v741 = v720
	goto L2
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v539
	F_errmsg_internal(m, int32(46145), v23+int32(16))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L3
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(490777), int32(13138), int32(376537))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L3
	} else {
		goto L163
	}
L163:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_get_statisticsobjdef(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_get_statisticsobj_worker(m, v3, int32(0), int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v12 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v12)
			return int32(0)
		} else {
			v16 = F_cstring_to_text(m, v6)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v6)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					return v16
				}
			}
		}
	}
}
func F_pg_get_statisticsobjdef_columns(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = int32(1)
	v6 = F_pg_get_statisticsobj_worker(m, v3, v4, v4)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v12 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v12)
			return int32(0)
		} else {
			v16 = F_cstring_to_text(m, v6)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v6)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					return v16
				}
			}
		}
	}
}
func F_pg_get_userbyid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = F_palloc(m, int32(64))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v22 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v18))) = v22
		v25 = v18 + int32(56)
		*(*int64)(unsafe.Add(mBase, uint32(v25))) = v22
		v29 = v18 + int32(48)
		*(*int64)(unsafe.Add(mBase, uint32(v29))) = v22
		v33 = v18 + int32(40)
		*(*int64)(unsafe.Add(mBase, uint32(v33))) = v22
		v37 = v18 + int32(32)
		*(*int64)(unsafe.Add(mBase, uint32(v37))) = v22
		v41 = v18 + int32(24)
		*(*int64)(unsafe.Add(mBase, uint32(v41))) = v22
		v45 = v18 + int32(16)
		*(*int64)(unsafe.Add(mBase, uint32(v45))) = v22
		v49 = v18 + int32(8)
		*(*int64)(unsafe.Add(mBase, uint32(v49))) = v22
		v53 = F_SearchSysCache1(m, int32(11), v16)
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return int32(0)
		} else {
			if v53 != 0 {
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
				v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+22)))
				v57 = v55 + v56
				v58 = *(*int64)(unsafe.Add(mBase, uint32(v57)+4))
				*(*int64)(unsafe.Add(mBase, uint32(v18))) = v58
				v60 = *(*int64)(unsafe.Add(mBase, uint32(v57)+60))
				*(*int64)(unsafe.Add(mBase, uint32(v25))) = v60
				v62 = *(*int64)(unsafe.Add(mBase, uint32(v57)+52))
				*(*int64)(unsafe.Add(mBase, uint32(v29))) = v62
				v64 = *(*int64)(unsafe.Add(mBase, uint32(v57)+44))
				*(*int64)(unsafe.Add(mBase, uint32(v33))) = v64
				v66 = *(*int64)(unsafe.Add(mBase, uint32(v57)+36))
				*(*int64)(unsafe.Add(mBase, uint32(v37))) = v66
				v68 = *(*int64)(unsafe.Add(mBase, uint32(v57)+28))
				*(*int64)(unsafe.Add(mBase, uint32(v41))) = v68
				v70 = *(*int64)(unsafe.Add(mBase, uint32(v57)+20))
				*(*int64)(unsafe.Add(mBase, uint32(v45))) = v70
				v72 = *(*int64)(unsafe.Add(mBase, uint32(v57)+12))
				*(*int64)(unsafe.Add(mBase, uint32(v49))) = v72
				F_ReleaseCatCache(m, v53)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int32(0)
				} else {
					m.G0 = v14 + int32(16)
					return v18
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v16
				v78 = F_pg_sprintf(m, v18, int32(652779), v14)
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return int32(0)
				} else {
					m.G0 = v14 + int32(16)
					return v18
				}
			}
		}
	}
}
func F_pg_hba_file_rules(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v24 int32
	_ = v24
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
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int64
	_ = v127
	var v133 int32
	_ = v133
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v218 int32
	_ = v218
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v917 int32
	_ = v917
	var v921 int32
	_ = v921
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v933 int32
	_ = v933
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v1000 int32
	_ = v1000
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	v2 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(736)
	m.G0 = v26
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+28))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	v36 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+348)) = v36
	v39 = *(*int32)(unsafe.Add(mBase, _consts[1087]))
	v43 = F_open_auth_file(m, v39, int32(21), v36, v36)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[1087]))
	F_tokenize_auth_file(m, v46, v43, v26+int32(348), int32(12), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v59 = F_AllocSetContextCreateInternal(m, v54, int32(60354), int32(0), int32(1024), int32(8192))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v61 = int32(4486928)
	v62 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v59
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v26)+348))
	if v65 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_free_auth_file(m, v43)
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L1
	} else {
		goto L241
	}
L7:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v68 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v94 = v2
	v96 = v2
	goto L9
L9:
	;
	v110 = int32(0)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v111+v94<<(uint(int32(2))%32))))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+16))
	if v116 == v110 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L6
L11:
	;
	v120 = F_parse_hba_line(m, v115, int32(12))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v123 = v110
	v124 = v116
	goto L13
L13:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	v127 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v26+int32(640)))) = v127
	*(*int64)(unsafe.Add(mBase, uint32(v26+int32(648)))) = v127
	*(*int64)(unsafe.Add(mBase, uint32(v26+int32(656)))) = v127
	v133 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(664)))) = v133
	*(*int64)(unsafe.Add(mBase, uint32(v26)+624)) = v127
	*(*int64)(unsafe.Add(mBase, uint32(v26)+632)) = v127
	*(*int64)(unsafe.Add(mBase, uint32(v26)+608)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v26)+615)) = v133
	v144 = v96 + int32(1)
	if v124 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v115)+16))
	v123 = v120
	v124 = v122
	goto L13
L15:
	;
	v148 = F_cstring_to_text(m, v126)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L19
	}
L16:
	;
	v145 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+608)) = uint8(v145)
	goto L15
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+624)) = v144
	goto L15
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+632)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v26)+628)) = v148
	if v123 != 0 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	if v124 != 0 {
		goto L231
	} else {
		goto L232
	}
L21:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v123)+20))
	if v262 != 0 {
		goto L44
	} else {
		goto L45
	}
L22:
	;
	v236 = F_strlist_to_textarray(m, v218)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L41
	}
L23:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	if base.Ui32(v152) <= base.Ui32(int32(5)) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	goto L25
L25:
	;
	v209 = int32(16843009)
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(611)))) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(614)))) = v209
	goto L20
L26:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v123)+16))
	if v165 != 0 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v152<<(uint(int32(2))%32))+uint32(_consts[1088])))
	v160 = F_cstring_to_text(m, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v163 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+611)) = uint8(v163)
	goto L26
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+636)) = v160
	goto L26
L31:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	if v166 <= int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	v207 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+612)) = uint8(v207)
	goto L21
L34:
	;
	v218 = int32(0)
	goto L22
L35:
	;
	goto L36
L36:
	;
	v170 = int32(0)
	v173 = v170
	v177 = v170
	goto L37
L37:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v195+v173<<(uint(int32(2))%32))))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	v201 = F_lappend(m, v177, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	v218 = v201
	goto L22
L39:
	;
	v204 = v173 + int32(1)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	if v204 < v205 {
		v173 = v204
		v177 = v201
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+640)) = v236
	goto L21
L42:
	;
	v356 = int32(0)
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v123)+288))
	switch v357 {
	case 0:
		goto L60
	case 1:
		goto L58
	case 2:
		goto L57
	case 3:
		v431 = v356
		v432 = int32(303564)
		goto L56
	default:
		v424 = v356
		goto L59
	}
L43:
	;
	v329 = F_strlist_to_textarray(m, v311)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L54
	}
L44:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	if v263 <= int32(0) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	v304 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+613)) = uint8(v304)
	goto L42
L47:
	;
	v311 = int32(0)
	goto L43
L48:
	;
	goto L49
L49:
	;
	v267 = int32(0)
	v270 = v267
	v274 = v267
	goto L50
L50:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v262)+12))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v292+v270<<(uint(int32(2))%32))))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
	v298 = F_lappend(m, v274, v297)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L52
	}
L51:
	;
	v311 = v298
	goto L43
L52:
	;
	v301 = v270 + int32(1)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	if v301 < v302 {
		v270 = v301
		v274 = v298
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+644)) = v329
	goto L42
L55:
	;
	if v437 != 0 {
		goto L89
	} else {
		goto L90
	}
L56:
	;
	v434 = F_cstring_to_text(m, v432)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L87
	}
L57:
	;
	v431 = v356
	v432 = int32(106897)
	goto L56
L58:
	;
	v431 = v356
	v432 = int32(67808)
	goto L56
L59:
	;
	v427 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+614)) = uint8(v427)
	v437 = v424
	goto L55
L60:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v123)+292))
	if v358 != 0 {
		v431 = v356
		v432 = v358
		goto L56
	} else {
		goto L61
	}
L61:
	;
	v359 = int32(0)
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v123)+152))
	if v359 < v360 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v364 = v123 + int32(24)
	v368 = int32(0)
	v371 = F_pg_getnameinfo_all(m, v364, v360, v26+int32(352), int32(255), v368, v368, int32(1))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	v391 = v359
	goto L64
L64:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v123)+284))
	if int32(0) < v392 {
		goto L74
	} else {
		goto L75
	}
L65:
	;
	if v371 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v375 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v364))))
	if v375 != int32(10) {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	goto L68
L68:
	;
	v389 = F_pstrdup(m, v26+int32(352))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L73
	}
L69:
	;
	goto L68
L70:
	;
	goto L69
L71:
	;
	v381 = F_strchr(m, v26+int32(352), int32(37))
	mBase = m.M
	if v381 == int32(0) {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v384 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v381))) = uint8(v384)
	goto L70
L73:
	;
	v391 = v389
	goto L64
L74:
	;
	v396 = v123 + int32(156)
	v400 = int32(0)
	v403 = F_pg_getnameinfo_all(m, v396, v392, v26+int32(352), int32(255), v400, v400, int32(1))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	v423 = v356
	goto L76
L76:
	;
	if v391 != 0 {
		v431 = v423
		v432 = v391
		goto L56
	} else {
		goto L86
	}
L77:
	;
	if v403 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v407 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v396))))
	if v407 != int32(10) {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	goto L80
L80:
	;
	v421 = F_pstrdup(m, v26+int32(352))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L85
	}
L81:
	;
	goto L80
L82:
	;
	goto L81
L83:
	;
	v413 = F_strchr(m, v26+int32(352), int32(37))
	mBase = m.M
	if v413 == int32(0) {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v416 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v413))) = uint8(v416)
	goto L82
L85:
	;
	v423 = v421
	goto L76
L86:
	;
	v424 = v423
	goto L59
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+648)) = v434
	v437 = v431
	goto L55
L88:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v123)+296))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v445<<(uint(int32(2))%32))+uint32(_consts[1089])))
	goto L93
L89:
	;
	v440 = F_cstring_to_text(m, v437)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v443 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+615)) = uint8(v443)
	goto L88
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+652)) = v440
	goto L88
L93:
	;
	v451 = F_cstring_to_text(m, v450)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+656)) = v451
	v454 = int32(0)
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v123)+296))
	if base.Ui32(int32(1)) < base.Ui32(v455-int32(7)) {
		v486 = v454
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v123)+300))
	if v489 != 0 {
		goto L105
	} else {
		goto L106
	}
L96:
	;
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+368)))
	if v460 != int32(1) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v123)+364))
	if v472 == int32(0) {
		v486 = v470
		goto L95
	} else {
		goto L102
	}
L98:
	;
	v470 = v454
	v471 = v26 + int32(672)
	goto L97
L99:
	;
	goto L100
L100:
	;
	v466 = F_cstring_to_text(m, int32(342086))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+672)) = v466
	v470 = int32(1)
	v471 = v26 + int32(672) | int32(4)
	goto L97
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+336)) = v472
	v479 = F_psprintf(m, int32(175821), v26+int32(336))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v481 = F_cstring_to_text(m, v479)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v471))) = v481
	v486 = v470 + int32(1)
	goto L95
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+320)) = v489
	v499 = F_psprintf(m, int32(175786), v26+int32(320))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L108
	}
L106:
	;
	v506 = v486
	goto L107
L107:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v123)+356))
	if v507 != 0 {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	v501 = F_cstring_to_text(m, v499)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(672)|v486<<(uint(int32(2))%32)))) = v501
	v506 = v486 + int32(1)
	goto L107
L110:
	;
	if v507 == int32(1) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	v529 = v506
	goto L112
L112:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v123)+304))
	if v530 != 0 {
		goto L118
	} else {
		goto L119
	}
L113:
	;
	v512 = int32(503976)
	goto L115
L114:
	;
	v512 = int32(301688)
	goto L115
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+304)) = v512
	v522 = F_psprintf(m, int32(175285), v26+int32(304))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v524 = F_cstring_to_text(m, v522)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(672)+v506<<(uint(int32(2))%32)))) = v524
	v529 = v506 + int32(1)
	goto L112
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+288)) = v530
	v540 = F_psprintf(m, int32(176036), v26+int32(288))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L121
	}
L119:
	;
	v547 = v529
	goto L120
L120:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v123)+296))
	if v548 == int32(11) {
		goto L124
	} else {
		goto L125
	}
L121:
	;
	v542 = F_cstring_to_text(m, v540)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(672)+v529<<(uint(int32(2))%32)))) = v542
	v547 = v529 + int32(1)
	goto L120
L123:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v123)+296))
	if v845 != int32(15) {
		v917 = v843
		goto L207
	} else {
		goto L208
	}
L124:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v123)+316))
	if v551 != 0 {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	v765 = v547
	v767 = v548
	goto L126
L126:
	;
	if v767 != int32(13) {
		v843 = v765
		goto L123
	} else {
		goto L186
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v551
	v561 = F_psprintf(m, int32(175537), v26+int32(272))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L130
	}
L128:
	;
	v568 = v547
	goto L129
L129:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v123)+320))
	if v569 != 0 {
		goto L132
	} else {
		goto L133
	}
L130:
	;
	v563 = F_cstring_to_text(m, v561)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(672)+v547<<(uint(int32(2))%32)))) = v563
	v568 = v547 + int32(1)
	goto L129
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+256)) = v569
	v579 = F_psprintf(m, int32(463269), v26+int32(256))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L1
	} else {
		goto L135
	}
L133:
	;
	v586 = v568
	goto L134
L134:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v123)+312))
	if v587 != 0 {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	v581 = F_cstring_to_text(m, v579)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(672)+v568<<(uint(int32(2))%32)))) = v581
	v586 = v568 + int32(1)
	goto L134
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+240)) = v587
	v597 = F_psprintf(m, int32(175879), v26+int32(240))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L1
	} else {
		goto L140
	}
L138:
	;
	v604 = v586
	goto L139
L139:
	;
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+309)))
	if v605 == int32(1) {
		goto L142
	} else {
		goto L143
	}
L140:
	;
	v599 = F_cstring_to_text(m, v597)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(672)+v586<<(uint(int32(2))%32)))) = v599
	v604 = v586 + int32(1)
	goto L139
L142:
	;
	v614 = F_cstring_to_text(m, int32(342073))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L1
	} else {
		goto L145
	}
L143:
	;
	v619 = v604
	goto L144
L144:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v123)+348))
	if v620 != 0 {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(672)+v604<<(uint(int32(2))%32)))) = v614
	v619 = v604 + int32(1)
	goto L144
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v620
	v630 = F_psprintf(m, int32(175200), v26+int32(224))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L1
	} else {
		goto L149
	}
L147:
	;
	v637 = v619
	goto L148
L148:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v123)+352))
	if v638 != 0 {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	v632 = F_cstring_to_text(m, v630)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(672)+v619<<(uint(int32(2))%32)))) = v632
	v637 = v619 + int32(1)
	goto L148
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+208)) = v638
	v648 = F_psprintf(m, int32(175186), v26+int32(208))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L1
	} else {
		goto L154
	}
L152:
	;
	v655 = v637
	goto L153
L153:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v123)+340))
	if v656 != 0 {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	v650 = F_cstring_to_text(m, v648)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(672)+v637<<(uint(int32(2))%32)))) = v650
	v655 = v637 + int32(1)
	goto L153
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+192)) = v656
	v666 = F_psprintf(m, int32(175793), v26+int32(192))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L1
	} else {
		goto L159
	}
L157:
	;
	v673 = v655
	goto L158
L158:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v123)+324))
	if v674 != 0 {
		goto L161
	} else {
		goto L162
	}
L159:
	;
	v668 = F_cstring_to_text(m, v666)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(672)+v655<<(uint(int32(2))%32)))) = v668
	v673 = v655 + int32(1)
	goto L158
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+176)) = v674
	v684 = F_psprintf(m, int32(175807), v26+int32(176))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L164
	}
L162:
	;
	v691 = v673
	goto L163
L163:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v123)+328))
	if v692 != 0 {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	v686 = F_cstring_to_text(m, v684)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(672)+v673<<(uint(int32(2))%32)))) = v686
	v691 = v673 + int32(1)
	goto L163
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+160)) = v692
	v702 = F_psprintf(m, int32(176050), v26+int32(160))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L1
	} else {
		goto L169
	}
L167:
	;
	v709 = v691
	goto L168
L168:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v123)+332))
	if v710 != 0 {
		goto L171
	} else {
		goto L172
	}
L169:
	;
	v704 = F_cstring_to_text(m, v702)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(672)+v691<<(uint(int32(2))%32)))) = v704
	v709 = v691 + int32(1)
	goto L168
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+144)) = v710
	v720 = F_psprintf(m, int32(175834), v26+int32(144))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L1
	} else {
		goto L174
	}
L172:
	;
	v727 = v709
	goto L173
L173:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v123)+336))
	if v728 != 0 {
		goto L176
	} else {
		goto L177
	}
L174:
	;
	v722 = F_cstring_to_text(m, v720)
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(672)+v709<<(uint(int32(2))%32)))) = v722
	v727 = v709 + int32(1)
	goto L173
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+128)) = v728
	v738 = F_psprintf(m, int32(175561), v26+int32(128))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L1
	} else {
		goto L179
	}
L177:
	;
	v745 = v727
	goto L178
L178:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v123)+344))
	if v746 != 0 {
		goto L181
	} else {
		goto L182
	}
L179:
	;
	v740 = F_cstring_to_text(m, v738)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(672)+v727<<(uint(int32(2))%32)))) = v740
	v745 = v727 + int32(1)
	goto L178
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+112)) = v746
	v756 = F_psprintf(m, int32(463707), v26+int32(112))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L1
	} else {
		goto L184
	}
L182:
	;
	v763 = v745
	goto L183
L183:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v123)+296))
	v765 = v763
	v767 = v764
	goto L126
L184:
	;
	v758 = F_cstring_to_text(m, v756)
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(672)+v745<<(uint(int32(2))%32)))) = v758
	v763 = v745 + int32(1)
	goto L183
L186:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v123)+376))
	if v770 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+96)) = v770
	v780 = F_psprintf(m, int32(175382), v26+int32(96))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L1
	} else {
		goto L190
	}
L188:
	;
	v787 = v765
	goto L189
L189:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v123)+384))
	if v788 != 0 {
		goto L192
	} else {
		goto L193
	}
L190:
	;
	v782 = F_cstring_to_text(m, v780)
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(672)+v765<<(uint(int32(2))%32)))) = v782
	v787 = v765 + int32(1)
	goto L189
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+80)) = v788
	v798 = F_psprintf(m, int32(175365), v26+int32(80))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L1
	} else {
		goto L195
	}
L193:
	;
	v805 = v787
	goto L194
L194:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v123)+392))
	if v806 != 0 {
		goto L197
	} else {
		goto L198
	}
L195:
	;
	v800 = F_cstring_to_text(m, v798)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(672)+v787<<(uint(int32(2))%32)))) = v800
	v805 = v787 + int32(1)
	goto L194
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = v806
	v816 = F_psprintf(m, int32(175399), v26-int32(-64))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L1
	} else {
		goto L200
	}
L198:
	;
	v823 = v805
	goto L199
L199:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v123)+400))
	if v824 == int32(0) {
		v843 = v823
		goto L123
	} else {
		goto L202
	}
L200:
	;
	v818 = F_cstring_to_text(m, v816)
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(672)+v805<<(uint(int32(2))%32)))) = v818
	v823 = v805 + int32(1)
	goto L199
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v824
	v836 = F_psprintf(m, int32(175350), v26+int32(48))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	v838 = F_cstring_to_text(m, v836)
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(672)+v823<<(uint(int32(2))%32)))) = v838
	v843 = v823 + int32(1)
	goto L123
L205:
	;
	v933 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+617)) = uint8(v933)
	goto L20
L206:
	;
	v926 = F_construct_array_builtin(m, v26+int32(672), v921, int32(25))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L1
	} else {
		goto L228
	}
L207:
	;
	if v917 == int32(0) {
		goto L205
	} else {
		goto L227
	}
L208:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v123)+404))
	if v848 != 0 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v848
	v858 = F_psprintf(m, int32(175551), v26+int32(32))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L1
	} else {
		goto L212
	}
L210:
	;
	v865 = v843
	goto L211
L211:
	;
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v123)+408))
	if v866 != 0 {
		goto L214
	} else {
		goto L215
	}
L212:
	;
	v860 = F_cstring_to_text(m, v858)
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(672)+v843<<(uint(int32(2))%32)))) = v860
	v865 = v843 + int32(1)
	goto L211
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v866
	v876 = F_psprintf(m, int32(175870), v26+int32(16))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L1
	} else {
		goto L217
	}
L215:
	;
	v883 = v865
	goto L216
L216:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v123)+412))
	if v884 != 0 {
		goto L219
	} else {
		goto L220
	}
L217:
	;
	v878 = F_cstring_to_text(m, v876)
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(672)+v865<<(uint(int32(2))%32)))) = v878
	v883 = v865 + int32(1)
	goto L216
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v884
	v892 = F_psprintf(m, int32(175524), v26)
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L1
	} else {
		goto L222
	}
L220:
	;
	v899 = v883
	goto L221
L221:
	;
	v900 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+416)))
	if v900 != int32(1) {
		v917 = v899
		goto L207
	} else {
		goto L224
	}
L222:
	;
	v894 = F_cstring_to_text(m, v892)
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L1
	} else {
		goto L223
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(672)+v883<<(uint(int32(2))%32)))) = v894
	v899 = v883 + int32(1)
	goto L221
L224:
	;
	v910 = F_psprintf(m, int32(342105), int32(0))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	v912 = F_cstring_to_text(m, v910)
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(672)+v899<<(uint(int32(2))%32)))) = v912
	v921 = v899 + int32(1)
	goto L206
L227:
	;
	v921 = v917
	goto L206
L228:
	;
	if v926 == int32(0) {
		goto L205
	} else {
		goto L229
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+660)) = v926
	goto L20
L230:
	;
	if v124 != 0 {
		goto L235
	} else {
		goto L236
	}
L231:
	;
	v958 = F_cstring_to_text(m, v124)
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L1
	} else {
		goto L234
	}
L232:
	;
	goto L233
L233:
	;
	v961 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+618)) = uint8(v961)
	goto L230
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+664)) = v958
	goto L230
L235:
	;
	v963 = v96
	goto L237
L236:
	;
	v963 = v144
	goto L237
L237:
	;
	v968 = F_heap_form_tuple(m, v34, v26+int32(624), v26+int32(608))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	F_tuplestore_puttuple(m, v35, v968)
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	v973 = v94 + int32(1)
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v973 < v974 {
		v94 = v973
		v96 = v963
		goto L9
	} else {
		goto L240
	}
L240:
	;
	goto L10
L241:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v62
	F_MemoryContextDelete(m, v59)
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	v1005 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v1005)
	m.G0 = v26 + int32(736)
	return int32(0)
}
func F_pg_hmac_error(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	if l0 == int32(0) {
		return int32(13869)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v7 != 0 {
			v19 = v7
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v11 == int32(2) {
				v14 = int32(211461)
			} else {
				v14 = int32(129434)
			}
			if v11 == int32(1) {
				v17 = int32(13869)
			} else {
				v17 = v14
			}
			v19 = v17
		}
		return v19
	}
}
func F_pg_hmac_init(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v225 int32
	_ = v225
	v4 = int32(0)
	v13 = int32(-1)
	if l0 == v4 {
		v225 = v13
		return v225
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v22 = F__emscripten_memset_bulkmem(m, l0+int32(152), base.I32_extend8_s(int32(92)), v20)
		mBase = m.M
		v27 = F__emscripten_memset_bulkmem(m, l0+int32(24), base.I32_extend8_s(int32(54)), v20)
		mBase = m.M
		if base.Ui32(l2) <= base.Ui32(v20) {
			v88 = l2
			v90 = l1
			v92 = v4
			if v88 == int32(0) {
			} else {
				v95 = int32(1)
				v97 = int32(0)
				if v88 != v95 {
					v104 = int32(0)
					v105 = v97
					for {
						v115 = v105 + v27
						v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
						v117 = v105 + v90
						v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
						v119 = v116 ^ v118
						*(*uint8)(unsafe.Add(mBase, uint32(v115))) = uint8(v119)
						v121 = v105 + v22
						v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
						v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
						v124 = v122 ^ v123
						*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v124)
						v127 = v105 | int32(1)
						v128 = v27 + v127
						v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
						v130 = v127 + v90
						v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
						v132 = v129 ^ v131
						*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v132)
						v134 = v127 + v22
						v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
						v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
						v137 = v135 ^ v136
						*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v137)
						v139 = int32(2)
						v140 = v105 + v139
						v142 = v104 + v139
						if v142 != v88&int32(-2) {
							v104 = v142
							v105 = v140
							continue
						} else {
							break
						}
						break
					}
					v146 = v140
				} else {
					v146 = v97
				}
				if v88&v95 == int32(0) {
				} else {
					v158 = v146 + v27
					v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
					v160 = v146 + v90
					v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
					v162 = v159 ^ v161
					*(*uint8)(unsafe.Add(mBase, uint32(v158))) = uint8(v162)
					v164 = v146 + v22
					v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
					v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
					v167 = v165 ^ v166
					*(*uint8)(unsafe.Add(mBase, uint32(v164))) = uint8(v167)
				}
			}
			v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v182 = F_pg_cryptohash_init(m, v181)
			mBase = m.M
			if int32(0) <= v182 {
				v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v187 = F_pg_cryptohash_update(m, v185, v27, v186)
				mBase = m.M
				if int32(0) <= v187 {
					v209 = int32(0)
					if v92 == v209 {
						v225 = v209
						return v225
					} else {
						v212 = v209
						F_pfree(m, v92)
						mBase = m.M
						v214 = m.ExcPending
						if v214 != 0 {
							return int32(0)
						} else {
							v225 = v212
							return v225
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
					v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					if v192 == int32(0) {
						v207 = int32(13869)
					} else {
						v199 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
						if v199 == int32(1) {
							v202 = int32(302852)
						} else {
							v202 = int32(129434)
						}
						if v199 == int32(2) {
							v205 = int32(13869)
						} else {
							v205 = v202
						}
						v207 = v205
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v207
					if v92 != 0 {
						v212 = v13
						F_pfree(m, v92)
						mBase = m.M
						v214 = m.ExcPending
						if v214 != 0 {
							return int32(0)
						} else {
							v225 = v212
							return v225
						}
					} else {
						v225 = v13
						return v225
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
				v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v192 == int32(0) {
					v207 = int32(13869)
				} else {
					v199 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
					if v199 == int32(1) {
						v202 = int32(302852)
					} else {
						v202 = int32(129434)
					}
					if v199 == int32(2) {
						v205 = int32(13869)
					} else {
						v205 = v202
					}
					v207 = v205
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v207
				if v92 != 0 {
					v212 = v13
					F_pfree(m, v92)
					mBase = m.M
					v214 = m.ExcPending
					if v214 != 0 {
						return int32(0)
					} else {
						v225 = v212
						return v225
					}
				} else {
					v225 = v13
					return v225
				}
			}
		} else {
			v29 = F_palloc(m, v16)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				if v29 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(1)
					return int32(-1)
				} else {
					v41 = F__emscripten_memset_bulkmem(m, v29, base.I32_extend8_s(int32(0)), v16)
					mBase = m.M
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v43 = F_pg_cryptohash_create(m, v42)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						if v43 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(1)
							F_pfree(m, v41)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								return int32(-1)
							}
						} else {
							v53 = F_pg_cryptohash_init(m, v43)
							mBase = m.M
							if v53 < int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
								if v43 == int32(0) {
									v78 = int32(13869)
								} else {
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
									if v70 == int32(1) {
										v73 = int32(302852)
									} else {
										v73 = int32(129434)
									}
									if v70 == int32(2) {
										v76 = int32(13869)
									} else {
										v76 = v73
									}
									v78 = v76
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v78
								F_pg_cryptohash_free(m, v43)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									F_pfree(m, v41)
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return int32(0)
									} else {
										return int32(-1)
									}
								}
							} else {
								v56 = F_pg_cryptohash_update(m, v43, l1, l2)
								mBase = m.M
								if v56 < int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
									if v43 == int32(0) {
										v78 = int32(13869)
									} else {
										v70 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
										if v70 == int32(1) {
											v73 = int32(302852)
										} else {
											v73 = int32(129434)
										}
										if v70 == int32(2) {
											v76 = int32(13869)
										} else {
											v76 = v73
										}
										v78 = v76
									}
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v78
									F_pg_cryptohash_free(m, v43)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v41)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return int32(0)
										} else {
											return int32(-1)
										}
									}
								} else {
									v59 = F_pg_cryptohash_final(m, v43, v41, v16)
									mBase = m.M
									if int32(0) <= v59 {
										F_pg_cryptohash_free(m, v43)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int32(0)
										} else {
											v88 = v16
											v90 = v29
											v92 = v41
											if v88 == int32(0) {
											} else {
												v95 = int32(1)
												v97 = int32(0)
												if v88 != v95 {
													v104 = int32(0)
													v105 = v97
													for {
														v115 = v105 + v27
														v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
														v117 = v105 + v90
														v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
														v119 = v116 ^ v118
														*(*uint8)(unsafe.Add(mBase, uint32(v115))) = uint8(v119)
														v121 = v105 + v22
														v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
														v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
														v124 = v122 ^ v123
														*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v124)
														v127 = v105 | int32(1)
														v128 = v27 + v127
														v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
														v130 = v127 + v90
														v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
														v132 = v129 ^ v131
														*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v132)
														v134 = v127 + v22
														v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
														v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
														v137 = v135 ^ v136
														*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v137)
														v139 = int32(2)
														v140 = v105 + v139
														v142 = v104 + v139
														if v142 != v88&int32(-2) {
															v104 = v142
															v105 = v140
															continue
														} else {
															break
														}
														break
													}
													v146 = v140
												} else {
													v146 = v97
												}
												if v88&v95 == int32(0) {
												} else {
													v158 = v146 + v27
													v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
													v160 = v146 + v90
													v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
													v162 = v159 ^ v161
													*(*uint8)(unsafe.Add(mBase, uint32(v158))) = uint8(v162)
													v164 = v146 + v22
													v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
													v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
													v167 = v165 ^ v166
													*(*uint8)(unsafe.Add(mBase, uint32(v164))) = uint8(v167)
												}
											}
											v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v182 = F_pg_cryptohash_init(m, v181)
											mBase = m.M
											if int32(0) <= v182 {
												v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												v187 = F_pg_cryptohash_update(m, v185, v27, v186)
												mBase = m.M
												if int32(0) <= v187 {
													v209 = int32(0)
													if v92 == v209 {
														v225 = v209
														return v225
													} else {
														v212 = v209
														F_pfree(m, v92)
														mBase = m.M
														v214 = m.ExcPending
														if v214 != 0 {
															return int32(0)
														} else {
															v225 = v212
															return v225
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
													v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													if v192 == int32(0) {
														v207 = int32(13869)
													} else {
														v199 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
														if v199 == int32(1) {
															v202 = int32(302852)
														} else {
															v202 = int32(129434)
														}
														if v199 == int32(2) {
															v205 = int32(13869)
														} else {
															v205 = v202
														}
														v207 = v205
													}
													*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v207
													if v92 != 0 {
														v212 = v13
														F_pfree(m, v92)
														mBase = m.M
														v214 = m.ExcPending
														if v214 != 0 {
															return int32(0)
														} else {
															v225 = v212
															return v225
														}
													} else {
														v225 = v13
														return v225
													}
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
												v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												if v192 == int32(0) {
													v207 = int32(13869)
												} else {
													v199 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
													if v199 == int32(1) {
														v202 = int32(302852)
													} else {
														v202 = int32(129434)
													}
													if v199 == int32(2) {
														v205 = int32(13869)
													} else {
														v205 = v202
													}
													v207 = v205
												}
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v207
												if v92 != 0 {
													v212 = v13
													F_pfree(m, v92)
													mBase = m.M
													v214 = m.ExcPending
													if v214 != 0 {
														return int32(0)
													} else {
														v225 = v212
														return v225
													}
												} else {
													v225 = v13
													return v225
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
										if v43 == int32(0) {
											v78 = int32(13869)
										} else {
											v70 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
											if v70 == int32(1) {
												v73 = int32(302852)
											} else {
												v73 = int32(129434)
											}
											if v70 == int32(2) {
												v76 = int32(13869)
											} else {
												v76 = v73
											}
											v78 = v76
										}
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v78
										F_pg_cryptohash_free(m, v43)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											F_pfree(m, v41)
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return int32(0)
											} else {
												return int32(-1)
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_hypot(m *base.Module, l0 float64, l1 float64) float64 {
	var v5 float64
	_ = v5
	var v6 float64
	_ = v6
	var v9 float64
	_ = v9
	var v12 float64
	_ = v12
	var v19 int32
	_ = v19
	var v20 float64
	_ = v20
	var v21 float64
	_ = v21
	var v24 float64
	_ = v24
	var v29 float64
	_ = v29
	var v37 float64
	_ = v37
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	v5 = math.Float64frombits(uint64(0x7ff0000000000000))
	v6 = base.F64_abs(l0)
	if base.F64_eq(v6, v5) != 0 {
		v37 = v5
		return v37
	} else {
		v9 = base.F64_abs(l1)
		if base.F64_eq(v9, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v37 = v5
			return v37
		} else {
			v12 = math.Float64frombits(uint64(0x7ff8000000000000))
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v6)) {
				v37 = v12
				return v37
			} else {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v9)) {
					v37 = v12
					return v37
				} else {
					v19 = base.F64_lt(v6, v9)
					if v19 != 0 {
						v20 = v9
					} else {
						v20 = v6
					}
					if v19 != 0 {
						v21 = v6
					} else {
						v21 = v9
					}
					if base.F64_eq(v21, float64(0)) != 0 {
						v37 = v20
						return v37
					} else {
						v24 = base.F64_div(v21, v20)
						v29 = base.F64_mul(v20, base.F64_sqrt(base.F64_add(base.F64_mul(v24, v24), float64(1))))
						if base.F64_eq(base.F64_abs(v29), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							F_float_overflow_error(m)
							v43 = m.ExcPending
							if v43 != 0 {
								return float64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							if base.F64_eq(v29, float64(0)) != 0 {
								F_float_underflow_error(m)
								v45 = m.ExcPending
								if v45 != 0 {
									return float64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v37 = v29
								return v37
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_input_error_info(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v15 = F_pg_detoast_datum_packed(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _consts[1105]))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+56)) = v18
			v21 = *(*int64)(unsafe.Add(mBase, _consts[1106]))
			*(*int64)(unsafe.Add(mBase, uint32(v7)+48)) = v21
			v26 = F_get_call_result_type(m, l0, int32(0), v5+int32(-20))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v26 == int32(1) {
					v30 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+53)) = uint8(v30)
					v34 = F_pg_input_is_valid_common(m, l0, v10, v15, v5+int32(-16))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						if v34 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(16843009)
							v110 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
							v115 = F_heap_form_tuple(m, v110, v5+int32(-48), v5+int32(-52))
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return int32(0)
							} else {
								v117 = *(*int32)(unsafe.Add(mBase, uint32(v115)+16))
								v118 = F_HeapTupleHeaderGetDatum(m, v117)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int32(0)
								} else {
									m.G0 = v7 - int32(-64)
									return v118
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)+56))
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
							v42 = F_cstring_to_text(m, v41)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v42
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v7)+56))
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+36))
								if v46 != 0 {
									v47 = F_cstring_to_text(m, v46)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v47
										v50 = *(*int32)(unsafe.Add(mBase, uint32(v7)+56))
										v53 = v50
										v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+44))
										if v54 != 0 {
											v55 = F_cstring_to_text(m, v54)
											mBase = m.M
											v56 = m.ExcPending
											if v56 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v55
												v58 = *(*int32)(unsafe.Add(mBase, uint32(v7)+56))
												v61 = v58
												v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+28))
												v63 = int32(4480664)
												v64 = int32(63)
												v66 = int32(48)
												v67 = v62&v64 + v66
												*(*uint8)(unsafe.Add(mBase, _consts[1107])) = uint8(v67)
												v75 = int32(base.Ui32(v62)>>(uint(int32(24))%32))&v64 + v66
												*(*uint8)(unsafe.Add(mBase, _consts[1108])) = uint8(v75)
												v83 = int32(base.Ui32(v62)>>(uint(int32(18))%32))&v64 + v66
												*(*uint8)(unsafe.Add(mBase, _consts[1109])) = uint8(v83)
												v91 = int32(base.Ui32(v62)>>(uint(int32(12))%32))&v64 + v66
												*(*uint8)(unsafe.Add(mBase, _consts[1110])) = uint8(v91)
												v99 = int32(base.Ui32(v62)>>(uint(int32(6))%32))&v64 + v66
												*(*uint8)(unsafe.Add(mBase, _consts[1111])) = uint8(v99)
												v102 = int32(0)
												*(*uint8)(unsafe.Add(mBase, _consts[1112])) = uint8(v102)
												v105 = F_cstring_to_text(m, v63)
												mBase = m.M
												v106 = m.ExcPending
												if v106 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v105
													v110 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
													v115 = F_heap_form_tuple(m, v110, v5+int32(-48), v5+int32(-52))
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return int32(0)
													} else {
														v117 = *(*int32)(unsafe.Add(mBase, uint32(v115)+16))
														v118 = F_HeapTupleHeaderGetDatum(m, v117)
														mBase = m.M
														v119 = m.ExcPending
														if v119 != 0 {
															return int32(0)
														} else {
															m.G0 = v7 - int32(-64)
															return v118
														}
													}
												}
											}
										} else {
											v59 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v59)
											v61 = v53
											v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+28))
											v63 = int32(4480664)
											v64 = int32(63)
											v66 = int32(48)
											v67 = v62&v64 + v66
											*(*uint8)(unsafe.Add(mBase, _consts[1107])) = uint8(v67)
											v75 = int32(base.Ui32(v62)>>(uint(int32(24))%32))&v64 + v66
											*(*uint8)(unsafe.Add(mBase, _consts[1108])) = uint8(v75)
											v83 = int32(base.Ui32(v62)>>(uint(int32(18))%32))&v64 + v66
											*(*uint8)(unsafe.Add(mBase, _consts[1109])) = uint8(v83)
											v91 = int32(base.Ui32(v62)>>(uint(int32(12))%32))&v64 + v66
											*(*uint8)(unsafe.Add(mBase, _consts[1110])) = uint8(v91)
											v99 = int32(base.Ui32(v62)>>(uint(int32(6))%32))&v64 + v66
											*(*uint8)(unsafe.Add(mBase, _consts[1111])) = uint8(v99)
											v102 = int32(0)
											*(*uint8)(unsafe.Add(mBase, _consts[1112])) = uint8(v102)
											v105 = F_cstring_to_text(m, v63)
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v105
												v110 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
												v115 = F_heap_form_tuple(m, v110, v5+int32(-48), v5+int32(-52))
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return int32(0)
												} else {
													v117 = *(*int32)(unsafe.Add(mBase, uint32(v115)+16))
													v118 = F_HeapTupleHeaderGetDatum(m, v117)
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return int32(0)
													} else {
														m.G0 = v7 - int32(-64)
														return v118
													}
												}
											}
										}
									}
								} else {
									v51 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v7)+13)) = uint8(v51)
									v53 = v45
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+44))
									if v54 != 0 {
										v55 = F_cstring_to_text(m, v54)
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v55
											v58 = *(*int32)(unsafe.Add(mBase, uint32(v7)+56))
											v61 = v58
											v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+28))
											v63 = int32(4480664)
											v64 = int32(63)
											v66 = int32(48)
											v67 = v62&v64 + v66
											*(*uint8)(unsafe.Add(mBase, _consts[1107])) = uint8(v67)
											v75 = int32(base.Ui32(v62)>>(uint(int32(24))%32))&v64 + v66
											*(*uint8)(unsafe.Add(mBase, _consts[1108])) = uint8(v75)
											v83 = int32(base.Ui32(v62)>>(uint(int32(18))%32))&v64 + v66
											*(*uint8)(unsafe.Add(mBase, _consts[1109])) = uint8(v83)
											v91 = int32(base.Ui32(v62)>>(uint(int32(12))%32))&v64 + v66
											*(*uint8)(unsafe.Add(mBase, _consts[1110])) = uint8(v91)
											v99 = int32(base.Ui32(v62)>>(uint(int32(6))%32))&v64 + v66
											*(*uint8)(unsafe.Add(mBase, _consts[1111])) = uint8(v99)
											v102 = int32(0)
											*(*uint8)(unsafe.Add(mBase, _consts[1112])) = uint8(v102)
											v105 = F_cstring_to_text(m, v63)
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v105
												v110 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
												v115 = F_heap_form_tuple(m, v110, v5+int32(-48), v5+int32(-52))
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return int32(0)
												} else {
													v117 = *(*int32)(unsafe.Add(mBase, uint32(v115)+16))
													v118 = F_HeapTupleHeaderGetDatum(m, v117)
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return int32(0)
													} else {
														m.G0 = v7 - int32(-64)
														return v118
													}
												}
											}
										}
									} else {
										v59 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v59)
										v61 = v53
										v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+28))
										v63 = int32(4480664)
										v64 = int32(63)
										v66 = int32(48)
										v67 = v62&v64 + v66
										*(*uint8)(unsafe.Add(mBase, _consts[1107])) = uint8(v67)
										v75 = int32(base.Ui32(v62)>>(uint(int32(24))%32))&v64 + v66
										*(*uint8)(unsafe.Add(mBase, _consts[1108])) = uint8(v75)
										v83 = int32(base.Ui32(v62)>>(uint(int32(18))%32))&v64 + v66
										*(*uint8)(unsafe.Add(mBase, _consts[1109])) = uint8(v83)
										v91 = int32(base.Ui32(v62)>>(uint(int32(12))%32))&v64 + v66
										*(*uint8)(unsafe.Add(mBase, _consts[1110])) = uint8(v91)
										v99 = int32(base.Ui32(v62)>>(uint(int32(6))%32))&v64 + v66
										*(*uint8)(unsafe.Add(mBase, _consts[1111])) = uint8(v99)
										v102 = int32(0)
										*(*uint8)(unsafe.Add(mBase, _consts[1112])) = uint8(v102)
										v105 = F_cstring_to_text(m, v63)
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v105
											v110 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
											v115 = F_heap_form_tuple(m, v110, v5+int32(-48), v5+int32(-52))
											mBase = m.M
											v116 = m.ExcPending
											if v116 != 0 {
												return int32(0)
											} else {
												v117 = *(*int32)(unsafe.Add(mBase, uint32(v115)+16))
												v118 = F_HeapTupleHeaderGetDatum(m, v117)
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return int32(0)
												} else {
													m.G0 = v7 - int32(-64)
													return v118
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v127 = m.ExcPending
					if v127 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(365102), int32(0))
						mBase = m.M
						v131 = m.ExcPending
						if v131 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(496562), int32(726), int32(240408))
							mBase = m.M
							v136 = m.ExcPending
							if v136 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_input_is_valid_common(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = F_text_to_cstring(m, l1)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
		if v17 == int32(0) {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
			v22 = F_MemoryContextAlloc(m, v20, int32(48))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v22
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
				v28 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v27))) = v28
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v30 == v28 {
					v77 = v28
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
					if v36 == int32(0) {
						v77 = v28
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
						v41 = v39 - int32(11)
						if base.Ui32(int32(9)) < base.Ui32(v41) {
							v77 = v28
						} else {
							if int32(base.Ui32(int32(977))>>(uint(v41)%32))&int32(1) == int32(0) {
								v77 = v28
							} else {
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v41<<(uint(int32(2))%32))+uint32(_consts[1104])))
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v36+v56)))
								if v58 == int32(0) {
									v77 = v28
								} else {
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
									if v61 <= int32(1) {
										v77 = v28
									} else {
										v63 = int32(1)
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
										v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+int32(4))))
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
										switch v69 - int32(7) {
										case 0:
											v77 = v63
										case 1:
											v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
											if v72 == int32(0) {
												v77 = v63
											} else {
												v77 = int32(0)
											}
										default:
											v77 = int32(0)
										}
									}
								}
							}
						}
					}
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v27)+8)) = uint8(v77)
				v80 = v27
				v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
				if v81 != 0 {
					v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+8)))
					if v82 != 0 {
						v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
						v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
						v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(12))
						mBase = m.M
						v118 = m.ExcPending
						if v118 != 0 {
							return int32(0)
						} else {
							m.G0 = v10 + int32(16)
							return v117
						}
					} else {
						v83 = F_text_to_cstring(m, l2)
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
							return int32(0)
						} else {
							v90 = F_parseTypeString(m, v83, v10+int32(8), v80+int32(4), int32(0))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								v92 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
								v93 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
								if v92 == v93 {
									v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
									v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
									v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(12))
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return int32(0)
									} else {
										m.G0 = v10 + int32(16)
										return v117
									}
								} else {
									F_getTypeInputInfo(m, v92, v80+int32(12), v80+int32(16))
									mBase = m.M
									v100 = m.ExcPending
									if v100 != 0 {
										return int32(0)
									} else {
										v101 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
										v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+20))
										F_fmgr_info_cxt(m, v101, v80+int32(20), v105)
										mBase = m.M
										v107 = m.ExcPending
										if v107 != 0 {
											return int32(0)
										} else {
											v108 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v80))) = v108
											v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
											v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
											v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(12))
											mBase = m.M
											v118 = m.ExcPending
											if v118 != 0 {
												return int32(0)
											} else {
												m.G0 = v10 + int32(16)
												return v117
											}
										}
									}
								}
							}
						}
					}
				} else {
					v83 = F_text_to_cstring(m, l2)
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return int32(0)
					} else {
						v90 = F_parseTypeString(m, v83, v10+int32(8), v80+int32(4), int32(0))
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
							v93 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
							if v92 == v93 {
								v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
								v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
								v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(12))
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return int32(0)
								} else {
									m.G0 = v10 + int32(16)
									return v117
								}
							} else {
								F_getTypeInputInfo(m, v92, v80+int32(12), v80+int32(16))
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return int32(0)
								} else {
									v101 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
									v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+20))
									F_fmgr_info_cxt(m, v101, v80+int32(20), v105)
									mBase = m.M
									v107 = m.ExcPending
									if v107 != 0 {
										return int32(0)
									} else {
										v108 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v80))) = v108
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
										v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
										v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(12))
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int32(0)
										} else {
											m.G0 = v10 + int32(16)
											return v117
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v80 = v17
			v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
			if v81 != 0 {
				v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+8)))
				if v82 != 0 {
					v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
					v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
					v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(12))
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return int32(0)
					} else {
						m.G0 = v10 + int32(16)
						return v117
					}
				} else {
					v83 = F_text_to_cstring(m, l2)
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return int32(0)
					} else {
						v90 = F_parseTypeString(m, v83, v10+int32(8), v80+int32(4), int32(0))
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
							v93 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
							if v92 == v93 {
								v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
								v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
								v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(12))
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return int32(0)
								} else {
									m.G0 = v10 + int32(16)
									return v117
								}
							} else {
								F_getTypeInputInfo(m, v92, v80+int32(12), v80+int32(16))
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return int32(0)
								} else {
									v101 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
									v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+20))
									F_fmgr_info_cxt(m, v101, v80+int32(20), v105)
									mBase = m.M
									v107 = m.ExcPending
									if v107 != 0 {
										return int32(0)
									} else {
										v108 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v80))) = v108
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
										v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
										v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(12))
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int32(0)
										} else {
											m.G0 = v10 + int32(16)
											return v117
										}
									}
								}
							}
						}
					}
				}
			} else {
				v83 = F_text_to_cstring(m, l2)
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int32(0)
				} else {
					v90 = F_parseTypeString(m, v83, v10+int32(8), v80+int32(4), int32(0))
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						v92 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
						v93 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
						if v92 == v93 {
							v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
							v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
							v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(12))
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return int32(0)
							} else {
								m.G0 = v10 + int32(16)
								return v117
							}
						} else {
							F_getTypeInputInfo(m, v92, v80+int32(12), v80+int32(16))
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int32(0)
							} else {
								v101 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
								v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+20))
								F_fmgr_info_cxt(m, v101, v80+int32(20), v105)
								mBase = m.M
								v107 = m.ExcPending
								if v107 != 0 {
									return int32(0)
								} else {
									v108 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v80))) = v108
									v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
									v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
									v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(12))
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return int32(0)
									} else {
										m.G0 = v10 + int32(16)
										return v117
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_is_ascii(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v3 = l0
	for {
		v5 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3))))
		if int32(0) < v5 {
			v3 = v3 + int32(1)
			continue
		} else {
			break
		}
		break
	}
	return base.B2i32(v5 == int32(0))
}
func F_pg_isblank(m *base.Module, l0 int32) int32 {
	return base.I32_wrap_i64(int64(base.Ui64(int64(4294976000))>>(uint(base.I64_extend_i32_s(l0))%64))) & base.B2i32(base.Ui32(l0) < base.Ui32(int32(33)))
}
func F_pg_listening_channels(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
	if v6 == int32(0) {
		v9 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v16 = *(*int32)(unsafe.Add(mBase, _consts[351]))
			if v16 == int32(0) {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = int32(2)
					v43 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v43)
					return int32(0)
				}
			} else {
				v19 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
				v20 = int64(*(*int32)(unsafe.Add(mBase, uint32(v16)+4)))
				if base.Ui64(v20) <= base.Ui64(v19) {
					F_end_MultiFuncCall(m, l0)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = int32(2)
						v43 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v43)
						return int32(0)
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v22+base.I32_wrap_i64(v19)<<(uint(int32(2))%32))))
					*(*int64)(unsafe.Add(mBase, uint32(v14))) = v19 + int64(1)
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = int32(1)
					v34 = F_cstring_to_text(m, v27)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						return v34
					}
				}
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
		v16 = *(*int32)(unsafe.Add(mBase, _consts[351]))
		if v16 == int32(0) {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = int32(2)
				v43 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v43)
				return int32(0)
			}
		} else {
			v19 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
			v20 = int64(*(*int32)(unsafe.Add(mBase, uint32(v16)+4)))
			if base.Ui64(v20) <= base.Ui64(v19) {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = int32(2)
					v43 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v43)
					return int32(0)
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v22+base.I32_wrap_i64(v19)<<(uint(int32(2))%32))))
				*(*int64)(unsafe.Add(mBase, uint32(v14))) = v19 + int64(1)
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = int32(1)
				v34 = F_cstring_to_text(m, v27)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					return v34
				}
			}
		}
	}
}
func F_pg_lock_status(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v228 int64
	_ = v228
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v255 int64
	_ = v255
	var v256 int32
	_ = v256
	var v258 int64
	_ = v258
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int64
	_ = v316
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v340 int32
	_ = v340
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int64
	_ = v673
	var v675 int64
	_ = v675
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v697 int64
	_ = v697
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1063 int32
	_ = v1063
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int64
	_ = v1068
	var v1070 int64
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1105 int32
	_ = v1105
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1133 int32
	_ = v1133
	var v1135 int32
	_ = v1135
	var v1139 int32
	_ = v1139
	var v1141 int32
	_ = v1141
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1151 int32
	_ = v1151
	var v1153 int32
	_ = v1153
	var v1157 int32
	_ = v1157
	var v1159 int32
	_ = v1159
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1193 int32
	_ = v1193
	var v1195 int32
	_ = v1195
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1205 int32
	_ = v1205
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1267 int64
	_ = v1267
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1342 int32
	_ = v1342
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1358 int32
	_ = v1358
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1383 int32
	_ = v1383
	var v1387 int32
	_ = v1387
	var v1389 int32
	_ = v1389
	var v1391 int32
	_ = v1391
	var v1395 int32
	_ = v1395
	var v1397 int32
	_ = v1397
	var v1399 int32
	_ = v1399
	var v1401 int32
	_ = v1401
	var v1403 int32
	_ = v1403
	var v1407 int32
	_ = v1407
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1412 int32
	_ = v1412
	var v1419 int64
	_ = v1419
	var v1423 int32
	_ = v1423
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1434 int32
	_ = v1434
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1450 int32
	_ = v1450
	var v1452 int32
	_ = v1452
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1462 int32
	_ = v1462
	var v1464 int32
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1468 int32
	_ = v1468
	var v1472 int32
	_ = v1472
	var v1475 int64
	_ = v1475
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1490 int32
	_ = v1490
	var v1492 int32
	_ = v1492
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1509 int32
	_ = v1509
	var v1511 int64
	_ = v1511
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1518 int32
	_ = v1518
	var v1521 int32
	_ = v1521
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1531 int64
	_ = v1531
	var v1535 int32
	_ = v1535
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1565 int64
	_ = v1565
	var v1591 int32
	_ = v1591
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1598 int32
	_ = v1598
	var v1601 int32
	_ = v1601
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1610 int32
	_ = v1610
	var v1612 int32
	_ = v1612
	var v1616 int32
	_ = v1616
	var v1618 int32
	_ = v1618
	var v1624 int32
	_ = v1624
	var v1626 int32
	_ = v1626
	var v1628 int32
	_ = v1628
	var v1632 int64
	_ = v1632
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1642 int32
	_ = v1642
	var v1643 int32
	_ = v1643
	var v1645 int32
	_ = v1645
	var v1647 int32
	_ = v1647
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1653 int32
	_ = v1653
	var v1657 int32
	_ = v1657
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1667 int64
	_ = v1667
	var v1671 int32
	_ = v1671
	var v1675 int32
	_ = v1675
	var v1676 int32
	_ = v1676
	var v1679 int32
	_ = v1679
	var v1685 int32
	_ = v1685
	v2 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(208)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	if v24 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v27 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v1232)+16))
	goto L171
L4:
	;
	return int32(0)
L5:
	;
	v31 = int32(4486928)
	v32 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v34
	v37 = F_CreateTemplateTupleDesc(m, int32(16))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	F_TupleDescInitEntry(m, v37, int32(1), int32(363663), int32(25), int32(-1), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	F_TupleDescInitEntry(m, v37, int32(2), int32(360248), int32(26), int32(-1), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	F_TupleDescInitEntry(m, v37, int32(3), int32(262542), int32(26), int32(-1), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	F_TupleDescInitEntry(m, v37, int32(4), int32(405958), int32(23), int32(-1), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	F_TupleDescInitEntry(m, v37, int32(5), int32(382399), int32(21), int32(-1), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	F_TupleDescInitEntry(m, v37, int32(6), int32(430103), int32(25), int32(-1), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	F_TupleDescInitEntry(m, v37, int32(7), int32(432577), int32(28), int32(-1), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	F_TupleDescInitEntry(m, v37, int32(8), int32(430611), int32(26), int32(-1), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	F_TupleDescInitEntry(m, v37, int32(9), int32(433403), int32(26), int32(-1), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	F_TupleDescInitEntry(m, v37, int32(10), int32(433485), int32(21), int32(-1), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	F_TupleDescInitEntry(m, v37, int32(11), int32(254105), int32(25), int32(-1), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	F_TupleDescInitEntry(m, v37, int32(12), int32(430790), int32(23), int32(-1), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	F_TupleDescInitEntry(m, v37, int32(13), int32(411053), int32(25), int32(-1), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	F_TupleDescInitEntry(m, v37, int32(14), int32(443930), int32(16), int32(-1), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	F_TupleDescInitEntry(m, v37, int32(15), int32(319366), int32(16), int32(-1), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	F_TupleDescInitEntry(m, v37, int32(16), int32(81620), int32(1184), int32(-1), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v151 = F_BlessTupleDesc(m, v37)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = v151
	v155 = F_palloc(m, int32(16))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v155
	v158 = m.G0
	v160 = v158 - int32(32)
	m.G0 = v160
	v163 = F_palloc(m, int32(8))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v166 = *(*int32)(unsafe.Add(mBase, _consts[786]))
	v169 = F_palloc(m, v166*int32(56))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+4)) = v169
	v173 = *(*int32)(unsafe.Add(mBase, _consts[153]))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+16))
	if v174 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v177 = v173
	v182 = v2
	v184 = v166
	v187 = v2
	goto L30
L28:
	;
	v436 = v2
	v438 = v166
	goto L29
L29:
	;
	v448 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v452 = F_LWLockAcquire(m, v448+int32(23296), int32(1))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L4
	} else {
		goto L64
	}
L30:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	v196 = v193 + v187*int32(640)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)+44))
	if v197 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v436 = v414
	v438 = v416
	goto L29
L32:
	;
	v199 = v196 + int32(584)
	v201 = F_LWLockAcquire(m, v199, int32(1))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L35
	}
L33:
	;
	v409 = v177
	v414 = v182
	v416 = v184
	goto L34
L34:
	;
	v426 = v187 + int32(1)
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v409)+16))
	if base.Ui32(v426) < base.Ui32(v427) {
		v177 = v409
		v182 = v414
		v184 = v416
		v187 = v426
		goto L30
	} else {
		goto L63
	}
L35:
	;
	v204 = *(*int32)(unsafe.Add(mBase, _consts[801]))
	if v204 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v208 = v204
	v213 = v182
	v215 = v184
	v216 = int32(0)
	goto L39
L37:
	;
	v349 = v182
	v351 = v184
	goto L38
L38:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+608)))
	if v360 != 0 {
		goto L55
	} else {
		goto L56
	}
L39:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v196)+600))
	v228 = *(*int64)(unsafe.Add(mBase, uint32(v224+v216<<(uint(int32(3))%32))))
	if v228 != int64(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v349 = v328
	v351 = v330
	goto L38
L41:
	;
	v245 = v213
	v247 = v215
	v255 = int64(0)
	goto L44
L42:
	;
	v323 = v208
	v328 = v213
	v330 = v215
	goto L43
L43:
	;
	v340 = v216 + int32(1)
	if base.Ui32(v340) < base.Ui32(v323) {
		v208 = v323
		v213 = v328
		v215 = v330
		v216 = v340
		goto L39
	} else {
		goto L54
	}
L44:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v196)+600))
	v258 = *(*int64)(unsafe.Add(mBase, uint32(v256+v216&int32(268435455)<<(uint(int32(3))%32))))
	v264 = base.I32_wrap_i64(int64(base.Ui64(v258)>>(uint(v255*int64(3))%64))) & int32(7)
	if v264 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v320 = *(*int32)(unsafe.Add(mBase, _consts[801]))
	v323 = v320
	v328 = v312
	v330 = v313
	goto L43
L46:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	if v247 <= v245 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v312 = v245
	v313 = v247
	goto L48
L48:
	;
	v316 = v255 + int64(1)
	if v316 != int64(16) {
		v245 = v312
		v247 = v313
		v255 = v316
		goto L44
	} else {
		goto L53
	}
L49:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _consts[786]))
	v269 = v268 + v247
	v272 = F_repalloc(m, v265, v269*int32(56))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L4
	} else {
		goto L52
	}
L50:
	;
	v275 = v265
	v276 = v247
	goto L51
L51:
	;
	v279 = v275 + v245*int32(56)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v196)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v279))) = v280
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v196)+604))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v282+base.I32_wrap_i64(v255)<<(uint(int32(2))%32)+v216<<(uint(int32(6))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v279)+20)) = int32(0)
	v291 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v279)+16)) = v264 << (uint(v291) % 32)
	*(*int64)(unsafe.Add(mBase, uint32(v279)+8)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v279)+4)) = v288
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v196)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v279)+24)) = v297
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v196)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v279)+28)) = v299
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v196)+44))
	*(*uint8)(unsafe.Add(mBase, uint32(v279)+48)) = uint8(v291)
	*(*int32)(unsafe.Add(mBase, uint32(v279)+44)) = v301
	*(*int32)(unsafe.Add(mBase, uint32(v279)+40)) = v301
	*(*int64)(unsafe.Add(mBase, uint32(v279)+32)) = int64(0)
	v312 = v245 + v291
	v313 = v276
	goto L48
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+4)) = v272
	v275 = v272
	v276 = v269
	goto L51
L53:
	;
	goto L45
L54:
	;
	goto L40
L55:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	if v351 <= v349 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	v400 = v349
	v401 = v351
	goto L57
L57:
	;
	F_LWLockRelease(m, v199)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L4
	} else {
		goto L62
	}
L58:
	;
	v364 = *(*int32)(unsafe.Add(mBase, _consts[786]))
	v365 = v364 + v351
	v368 = F_repalloc(m, v361, v365*int32(56))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L4
	} else {
		goto L61
	}
L59:
	;
	v371 = v361
	v372 = v351
	goto L60
L60:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v196)+52))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v196)+612))
	v377 = v371 + v349*int32(56)
	*(*int64)(unsafe.Add(mBase, uint32(v377)+16)) = int64(128)
	*(*int64)(unsafe.Add(mBase, uint32(v377)+8)) = int64(73746443898191872)
	*(*int32)(unsafe.Add(mBase, uint32(v377)+4)) = v374
	*(*int32)(unsafe.Add(mBase, uint32(v377))) = v373
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v196)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v377)+24)) = v384
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v196)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v377)+28)) = v386
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v196)+44))
	v389 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v377)+48)) = uint8(v389)
	*(*int32)(unsafe.Add(mBase, uint32(v377)+44)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v377)+40)) = v388
	*(*int64)(unsafe.Add(mBase, uint32(v377)+32)) = int64(0)
	v400 = v349 + v389
	v401 = v372
	goto L57
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+4)) = v368
	v371 = v368
	v372 = v365
	goto L60
L62:
	;
	v406 = *(*int32)(unsafe.Add(mBase, _consts[153]))
	v409 = v406
	v414 = v400
	v416 = v401
	goto L34
L63:
	;
	goto L31
L64:
	;
	v455 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v459 = F_LWLockAcquire(m, v455+int32(23424), int32(1))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	v462 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v466 = F_LWLockAcquire(m, v462+int32(23552), int32(1))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	v469 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v473 = F_LWLockAcquire(m, v469+int32(23680), int32(1))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	v476 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v480 = F_LWLockAcquire(m, v476+int32(23808), int32(1))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	v483 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v487 = F_LWLockAcquire(m, v483+int32(23936), int32(1))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	v490 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v494 = F_LWLockAcquire(m, v490+int32(24064), int32(1))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	v497 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v501 = F_LWLockAcquire(m, v497+int32(24192), int32(1))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	v504 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v508 = F_LWLockAcquire(m, v504+int32(24320), int32(1))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	v511 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v515 = F_LWLockAcquire(m, v511+int32(24448), int32(1))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	v518 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v522 = F_LWLockAcquire(m, v518+int32(24576), int32(1))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	v525 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v529 = F_LWLockAcquire(m, v525+int32(24704), int32(1))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L4
	} else {
		goto L75
	}
L75:
	;
	v532 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v536 = F_LWLockAcquire(m, v532+int32(24832), int32(1))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L4
	} else {
		goto L76
	}
L76:
	;
	v539 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v543 = F_LWLockAcquire(m, v539+int32(24960), int32(1))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	v546 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v550 = F_LWLockAcquire(m, v546+int32(25088), int32(1))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	v553 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v557 = F_LWLockAcquire(m, v553+int32(25216), int32(1))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	v560 = *(*int32)(unsafe.Add(mBase, _consts[803]))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v560)))
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v562)+4))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v562)+412))
	if v564 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v630 = v629 + v436
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v630
	if v438 < v630 {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v562)+376))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v562)+364))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v562)+352))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v562)+340))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v562)+328))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v562)+316))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v562)+304))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v562)+292))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v562)+280))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v562)+268))
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v562)+256))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v562)+244))
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v562)+232))
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v562)+220))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v562)+208))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v562)+196))
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v562)+184))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v562)+172))
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v562)+160))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v562)+148))
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v562)+136))
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v562)+124))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v562)+112))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v562)+100))
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v562)+88))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v562)+76))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v562-int32(-64))))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v562)+52))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v562)+40))
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v562)+28))
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v562)+16))
	v629 = v565 + (v566 + (v567 + (v568 + (v569 + (v570 + (v571 + (v572 + (v573 + (v574 + (v575 + (v576 + (v577 + (v578 + (v579 + (v580 + (v581 + (v582 + (v583 + (v584 + (v585 + (v586 + (v587 + (v588 + (v589 + (v590 + (v593 + (v594 + (v595 + (v596 + (v597 + v563))))))))))))))))))))))))))))))
	goto L83
L82:
	;
	v629 = v563
	goto L83
L83:
	;
	goto L80
L84:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	v636 = F_repalloc(m, v633, v630*int32(56))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L4
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v642 = *(*int32)(unsafe.Add(mBase, _consts[803]))
	F_hash_seq_init(m, v160+int32(12), v642)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L4
	} else {
		goto L88
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+4)) = v636
	goto L86
L88:
	;
	v647 = F_hash_seq_search(m, v160+int32(12))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	if v647 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v655 = v647
	v656 = v436
	goto L93
L91:
	;
	goto L92
L92:
	;
	v725 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v725+int32(25216))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L4
	} else {
		goto L100
	}
L93:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v655)+4))
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	v671 = v668 + v656*int32(56)
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v655)))
	v673 = *(*int64)(unsafe.Add(mBase, uint32(v672)))
	*(*int64)(unsafe.Add(mBase, uint32(v671))) = v673
	v675 = *(*int64)(unsafe.Add(mBase, uint32(v672)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v671)+8)) = v675
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v655)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v671)+16)) = v677
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v667)+92))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v655)))
	if v679 == v680 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	goto L92
L95:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v667)+100))
	v684 = v682
	goto L97
L96:
	;
	v684 = int32(0)
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v671)+20)) = v684
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v667)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v671)+24)) = v686
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v667)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v671)+28)) = v688
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v667)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v671)+40)) = v690
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v655)+8))
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v692)+44))
	v694 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v671)+48)) = uint8(v694)
	*(*int32)(unsafe.Add(mBase, uint32(v671)+44)) = v693
	v697 = *(*int64)(unsafe.Add(mBase, uint32(v667)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v667)+112)) = v697
	*(*int64)(unsafe.Add(mBase, uint32(v671)+32)) = v697
	v704 = F_hash_seq_search(m, v160+int32(12))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	if v704 != 0 {
		v655 = v704
		v656 = v656 + int32(1)
		goto L93
	} else {
		goto L99
	}
L99:
	;
	goto L94
L100:
	;
	v731 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v731+int32(25088))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L4
	} else {
		goto L101
	}
L101:
	;
	v737 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v737+int32(24960))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	v743 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v743+int32(24832))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	v749 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v749+int32(24704))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	v755 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v755+int32(24576))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	v761 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v761+int32(24448))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L4
	} else {
		goto L106
	}
L106:
	;
	v767 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v767+int32(24320))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L4
	} else {
		goto L107
	}
L107:
	;
	v773 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v773+int32(24192))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	v779 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v779+int32(24064))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	v785 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v785+int32(23936))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	v791 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v791+int32(23808))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L4
	} else {
		goto L111
	}
L111:
	;
	v797 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v797+int32(23680))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L4
	} else {
		goto L112
	}
L112:
	;
	v803 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v803+int32(23552))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L4
	} else {
		goto L113
	}
L113:
	;
	v809 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v809+int32(23424))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L4
	} else {
		goto L114
	}
L114:
	;
	v815 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v815+int32(23296))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L4
	} else {
		goto L115
	}
L115:
	;
	v820 = int32(32)
	m.G0 = v160 + v820
	v823 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v155)+4)) = v823
	*(*int32)(unsafe.Add(mBase, uint32(v155))) = v163
	v827 = m.G0
	v829 = v827 - v820
	m.G0 = v829
	v832 = F_palloc(m, int32(12))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L4
	} else {
		goto L116
	}
L116:
	;
	v835 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v839 = F_LWLockAcquire(m, v835+int32(25344), int32(1))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L4
	} else {
		goto L117
	}
L117:
	;
	v842 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v846 = F_LWLockAcquire(m, v842+int32(25472), int32(1))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L4
	} else {
		goto L118
	}
L118:
	;
	v849 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v853 = F_LWLockAcquire(m, v849+int32(25600), int32(1))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L4
	} else {
		goto L119
	}
L119:
	;
	v856 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v860 = F_LWLockAcquire(m, v856+int32(25728), int32(1))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L4
	} else {
		goto L120
	}
L120:
	;
	v863 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v867 = F_LWLockAcquire(m, v863+int32(25856), int32(1))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L4
	} else {
		goto L121
	}
L121:
	;
	v870 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v874 = F_LWLockAcquire(m, v870+int32(25984), int32(1))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L4
	} else {
		goto L122
	}
L122:
	;
	v877 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v881 = F_LWLockAcquire(m, v877+int32(26112), int32(1))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L4
	} else {
		goto L123
	}
L123:
	;
	v884 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v888 = F_LWLockAcquire(m, v884+int32(26240), int32(1))
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L4
	} else {
		goto L124
	}
L124:
	;
	v891 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v895 = F_LWLockAcquire(m, v891+int32(26368), int32(1))
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L4
	} else {
		goto L125
	}
L125:
	;
	v898 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v902 = F_LWLockAcquire(m, v898+int32(26496), int32(1))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L4
	} else {
		goto L126
	}
L126:
	;
	v905 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v909 = F_LWLockAcquire(m, v905+int32(26624), int32(1))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	v912 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v916 = F_LWLockAcquire(m, v912+int32(26752), int32(1))
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L4
	} else {
		goto L128
	}
L128:
	;
	v919 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v923 = F_LWLockAcquire(m, v919+int32(26880), int32(1))
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L4
	} else {
		goto L129
	}
L129:
	;
	v926 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v930 = F_LWLockAcquire(m, v926+int32(27008), int32(1))
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	v933 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v937 = F_LWLockAcquire(m, v933+int32(27136), int32(1))
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L4
	} else {
		goto L131
	}
L131:
	;
	v940 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v944 = F_LWLockAcquire(m, v940+int32(27264), int32(1))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	v947 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v951 = F_LWLockAcquire(m, v947+int32(3584), int32(1))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	v954 = *(*int32)(unsafe.Add(mBase, _consts[818]))
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v954)))
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v956)+4))
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v956)+412))
	if v958 != 0 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v832))) = v1023
	v1027 = F_palloc(m, v1023<<(uint(int32(4))%32))
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L4
	} else {
		goto L138
	}
L135:
	;
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v956)+376))
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v956)+364))
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v956)+352))
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v956)+340))
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v956)+328))
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v956)+316))
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v956)+304))
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v956)+292))
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v956)+280))
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v956)+268))
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v956)+256))
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v956)+244))
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v956)+232))
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v956)+220))
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v956)+208))
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v956)+196))
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v956)+184))
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v956)+172))
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v956)+160))
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v956)+148))
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v956)+136))
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v956)+124))
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v956)+112))
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v956)+100))
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v956)+88))
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v956)+76))
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v956-int32(-64))))
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v956)+52))
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v956)+40))
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v956)+28))
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v956)+16))
	v1023 = v959 + (v960 + (v961 + (v962 + (v963 + (v964 + (v965 + (v966 + (v967 + (v968 + (v969 + (v970 + (v971 + (v972 + (v973 + (v974 + (v975 + (v976 + (v977 + (v978 + (v979 + (v980 + (v981 + (v982 + (v983 + (v984 + (v987 + (v988 + (v989 + (v990 + (v991 + v957))))))))))))))))))))))))))))))
	goto L137
L136:
	;
	v1023 = v957
	goto L137
L137:
	;
	goto L134
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v832)+4)) = v1027
	v1032 = F_palloc(m, v1023*int32(120))
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L4
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v832)+8)) = v1032
	v1038 = *(*int32)(unsafe.Add(mBase, _consts[818]))
	F_hash_seq_init(m, v829+int32(12), v1038)
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L4
	} else {
		goto L140
	}
L140:
	;
	v1043 = F_hash_seq_search(m, v829+int32(12))
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L4
	} else {
		goto L141
	}
L141:
	;
	if v1043 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v1049 = v1043
	v1051 = v823
	goto L145
L143:
	;
	goto L144
L144:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v1105+int32(3584))
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L4
	} else {
		goto L153
	}
L145:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v832)+4))
	v1066 = v1063 + v1051<<(uint(int32(4))%32)
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v1049)))
	v1068 = *(*int64)(unsafe.Add(mBase, uint32(v1067)))
	*(*int64)(unsafe.Add(mBase, uint32(v1066))) = v1068
	v1070 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1066)+8)) = v1070
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v832)+8))
	v1073 = int32(120)
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v1049)+4))
	goto L148
L146:
	;
	goto L144
L147:
	;
	v1084 = F_hash_seq_search(m, v829+int32(12))
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L4
	} else {
		goto L151
	}
L148:
	;
	v1078 = F__emscripten_memcpy_bulkmem(m, v1072+v1051*v1073, v1076, v1073)
	mBase = m.M
	goto L150
L150:
	;
	goto L147
L151:
	;
	if v1084 != 0 {
		v1049 = v1084
		v1051 = v1051 + int32(1)
		goto L145
	} else {
		goto L152
	}
L152:
	;
	goto L146
L153:
	;
	v1111 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v1111+int32(27264))
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v1117+int32(27136))
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L4
	} else {
		goto L155
	}
L155:
	;
	v1123 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v1123+int32(27008))
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L4
	} else {
		goto L156
	}
L156:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v1129+int32(26880))
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L4
	} else {
		goto L157
	}
L157:
	;
	v1135 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v1135+int32(26752))
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L4
	} else {
		goto L158
	}
L158:
	;
	v1141 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v1141+int32(26624))
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		goto L4
	} else {
		goto L159
	}
L159:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v1147+int32(26496))
	mBase = m.M
	v1151 = m.ExcPending
	if v1151 != 0 {
		goto L4
	} else {
		goto L160
	}
L160:
	;
	v1153 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v1153+int32(26368))
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	v1159 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v1159+int32(26240))
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L4
	} else {
		goto L162
	}
L162:
	;
	v1165 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v1165+int32(26112))
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L4
	} else {
		goto L163
	}
L163:
	;
	v1171 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v1171+int32(25984))
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L4
	} else {
		goto L164
	}
L164:
	;
	v1177 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v1177+int32(25856))
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L4
	} else {
		goto L165
	}
L165:
	;
	v1183 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v1183+int32(25728))
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		goto L4
	} else {
		goto L166
	}
L166:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v1189+int32(25600))
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L4
	} else {
		goto L167
	}
L167:
	;
	v1195 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v1195+int32(25472))
	mBase = m.M
	v1199 = m.ExcPending
	if v1199 != 0 {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	v1201 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v1201+int32(25344))
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L4
	} else {
		goto L169
	}
L169:
	;
	m.G0 = v829 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v155)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v155)+8)) = v832
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v32
	goto L3
L170:
	;
	m.G0 = v21 + int32(208)
	return v1685
L171:
	;
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v1233)+16))
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v1234)+4))
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(v1234)))
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v1236)))
	if v1235 < v1237 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	goto L175
L173:
	;
	goto L174
L174:
	;
	v1559 = *(*int32)(unsafe.Add(mBase, uint32(v1234)+12))
	v1560 = *(*int32)(unsafe.Add(mBase, uint32(v1234)+8))
	v1561 = *(*int32)(unsafe.Add(mBase, uint32(v1560)))
	if v1559 < v1561 {
		goto L245
	} else {
		goto L246
	}
L175:
	;
	v1267 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+136)) = v1267
	*(*int64)(unsafe.Add(mBase, uint32(v21+int32(128)))) = v1267
	*(*int64)(unsafe.Add(mBase, uint32(v21+int32(120)))) = v1267
	*(*int64)(unsafe.Add(mBase, uint32(v21+int32(112)))) = v1267
	*(*int64)(unsafe.Add(mBase, uint32(v21+int32(104)))) = v1267
	*(*int64)(unsafe.Add(mBase, uint32(v21+int32(96)))) = v1267
	*(*int64)(unsafe.Add(mBase, uint32(v21)+88)) = v1267
	*(*int64)(unsafe.Add(mBase, uint32(v21)+80)) = v1267
	*(*int64)(unsafe.Add(mBase, uint32(v21)+72)) = v1267
	*(*int64)(unsafe.Add(mBase, uint32(v21)+64)) = v1267
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v1236)+4))
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1234)+4))
	v1291 = v1287 + v1288*int32(56)
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v1291)+16))
	if v1292 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L176:
	;
	goto L174
L177:
	;
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(v1234)+4))
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(v1236)))
	if v1538 < v1539 {
		goto L175
	} else {
		goto L244
	}
L178:
	;
	v1351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1291)+14)))
	if base.Ui32(v1351) <= base.Ui32(int32(11)) {
		goto L212
	} else {
		goto L213
	}
L179:
	;
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v1291)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1234)+4)) = v1288 + int32(1)
	if v1342 == int32(0) {
		goto L177
	} else {
		goto L210
	}
L180:
	;
	if v1292&int32(1) != 0 {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1291)+16)) = v1338 & v1292
	v1349 = v1337
	v1350 = int32(1)
	goto L178
L182:
	;
	v1337 = int32(0)
	v1338 = int32(-2)
	goto L181
L183:
	;
	goto L184
L184:
	;
	if v1292&int32(2) != 0 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v1337 = int32(1)
	v1338 = int32(-3)
	goto L181
L186:
	;
	goto L187
L187:
	;
	if v1292&int32(4) != 0 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v1337 = int32(2)
	v1338 = int32(-5)
	goto L181
L189:
	;
	goto L190
L190:
	;
	if v1292&int32(8) != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v1337 = int32(3)
	v1338 = int32(-9)
	goto L181
L192:
	;
	goto L193
L193:
	;
	if v1292&int32(16) != 0 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v1337 = int32(4)
	v1338 = int32(-17)
	goto L181
L195:
	;
	goto L196
L196:
	;
	if v1292&int32(32) != 0 {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v1337 = int32(5)
	v1338 = int32(-33)
	goto L181
L198:
	;
	goto L199
L199:
	;
	if v1292&int32(64) != 0 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v1337 = int32(6)
	v1338 = int32(-65)
	goto L181
L201:
	;
	goto L202
L202:
	;
	if v1292&int32(128) != 0 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v1337 = int32(7)
	v1338 = int32(-129)
	goto L181
L204:
	;
	goto L205
L205:
	;
	if v1292&int32(256) != 0 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v1337 = int32(8)
	v1338 = int32(-257)
	goto L181
L207:
	;
	goto L208
L208:
	;
	if v1292&int32(512) == int32(0) {
		goto L179
	} else {
		goto L209
	}
L209:
	;
	v1337 = int32(9)
	v1338 = int32(-513)
	goto L181
L210:
	;
	v1349 = v1342
	v1350 = int32(0)
	goto L178
L211:
	;
	v1371 = F_cstring_to_text(m, v1370)
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		goto L4
	} else {
		goto L216
	}
L212:
	;
	v1358 = *(*int32)(unsafe.Add(mBase, uint32(v1351<<(uint(int32(2))%32))+uint32(_consts[1093])))
	v1370 = v1358
	goto L211
L213:
	;
	goto L214
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v1351
	v1366 = F_pg_snprintf(m, v21+int32(144), int32(32), int32(469647), v21+int32(48))
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L4
	} else {
		goto L215
	}
L215:
	;
	v1370 = v21 + int32(144)
	goto L211
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = v1371
	v1374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1291)+14)))
	switch v1374 {
	case 0, 1:
		goto L226
	case 2:
		goto L225
	case 3:
		goto L224
	case 4:
		goto L223
	case 5:
		goto L222
	case 6:
		goto L221
	case 7:
		goto L220
	default:
		goto L218
	case 11:
		goto L219
	}
L217:
	;
	v1475 = *(*int64)(unsafe.Add(mBase, uint32(v1291)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+16)) = v1475
	v1483 = F_pg_snprintf(m, v21+int32(176), int32(32), int32(39061), v21+int32(16))
	mBase = m.M
	v1484 = m.ExcPending
	if v1484 != 0 {
		goto L4
	} else {
		goto L229
	}
L218:
	;
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = v1462
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(v1291)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+108)) = v1464
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(v1291)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+112)) = v1466
	v1468 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1291)+12)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+66)) = int32(16843009)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+116)) = v1468
	v1472 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+70)) = uint8(v1472)
	goto L217
L219:
	;
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = v1450
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(v1291)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+112)) = v1452
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v1291)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+104)) = v1454
	v1456 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1291)+12)))
	v1457 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+71)) = uint8(v1457)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+116)) = v1456
	*(*int32)(unsafe.Add(mBase, uint32(v21)+66)) = int32(16843009)
	goto L217
L220:
	;
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+104)) = v1438
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v1291)+4))
	v1441 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+73)) = uint8(v1441)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+71)) = uint8(v1441)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+112)) = v1440
	*(*int32)(unsafe.Add(mBase, uint32(v21)+65)) = int32(16843009)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+69)) = uint8(v1441)
	goto L217
L221:
	;
	v1419 = *(*int64)(unsafe.Add(mBase, uint32(v1291)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+32)) = v1419
	v1423 = int32(32)
	v1427 = F_pg_snprintf(m, v21+int32(176), v1423, int32(39061), v21+v1423)
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L4
	} else {
		goto L227
	}
L222:
	;
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	v1410 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+73)) = uint8(v1410)
	v1412 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+71)) = uint16(v1412)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+104)) = v1409
	*(*int32)(unsafe.Add(mBase, uint32(v21)+65)) = int32(16843009)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+69)) = uint8(v1410)
	goto L217
L223:
	;
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = v1397
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(v1291)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+88)) = v1399
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v1291)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+92)) = v1401
	v1403 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1291)+12)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+69)) = int32(16843009)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+96)) = v1403
	v1407 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+73)) = uint8(v1407)
	goto L217
L224:
	;
	v1387 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = v1387
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(v1291)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+88)) = v1389
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v1291)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = int32(16843009)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+92)) = v1391
	v1395 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+72)) = uint16(v1395)
	goto L217
L225:
	;
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+66)) = int64(72340172838076673)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = v1383
	goto L217
L226:
	;
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = v1375
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(v1291)+4))
	v1378 = int32(16843009)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+67)) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v21)+88)) = v1377
	*(*int32)(unsafe.Add(mBase, uint32(v21)+70)) = v1378
	goto L217
L227:
	;
	v1431 = F_cstring_to_text(m, v21+int32(176))
	mBase = m.M
	v1432 = m.ExcPending
	if v1432 != 0 {
		goto L4
	} else {
		goto L228
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+100)) = v1431
	v1434 = int32(16843009)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+70)) = v1434
	*(*int32)(unsafe.Add(mBase, uint32(v21)+65)) = v1434
	goto L217
L229:
	;
	v1487 = F_cstring_to_text(m, v21+int32(176))
	mBase = m.M
	v1488 = m.ExcPending
	if v1488 != 0 {
		goto L4
	} else {
		goto L230
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+120)) = v1487
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(v1291)+40))
	if v1490 != 0 {
		goto L232
	} else {
		goto L233
	}
L231:
	;
	v1494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1291)+15)))
	v1495 = int32(2)
	v1499 = *(*int32)(unsafe.Add(mBase, uint32(v1494<<(uint(v1495)%32))+uint32(_consts[97])))
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(v1499)+8))
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(v1500+v1349<<(uint(v1495)%32))))
	goto L235
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+124)) = v1490
	goto L231
L233:
	;
	goto L234
L234:
	;
	v1492 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+75)) = uint8(v1492)
	goto L231
L235:
	;
	v1505 = F_cstring_to_text(m, v1504)
	mBase = m.M
	v1506 = m.ExcPending
	if v1506 != 0 {
		goto L4
	} else {
		goto L236
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+132)) = v1350
	*(*int32)(unsafe.Add(mBase, uint32(v21)+128)) = v1505
	v1509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1291)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+136)) = v1509
	if v1350 != 0 {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(v1233)+28))
	v1526 = F_heap_form_tuple(m, v1521, v21+int32(80), v21-int32(-64))
	mBase = m.M
	v1527 = m.ExcPending
	if v1527 != 0 {
		goto L4
	} else {
		goto L242
	}
L238:
	;
	v1518 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+79)) = uint8(v1518)
	goto L237
L239:
	;
	v1511 = *(*int64)(unsafe.Add(mBase, uint32(v1291)+32))
	if v1511 == int64(0) {
		goto L238
	} else {
		goto L240
	}
L240:
	;
	v1514 = F_Int64GetDatum(m, v1511)
	mBase = m.M
	v1515 = m.ExcPending
	if v1515 != 0 {
		goto L4
	} else {
		goto L241
	}
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+140)) = v1514
	goto L237
L242:
	;
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(v1526)+16))
	v1529 = F_HeapTupleHeaderGetDatum(m, v1528)
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L4
	} else {
		goto L243
	}
L243:
	;
	v1531 = *(*int64)(unsafe.Add(mBase, uint32(v1233)))
	*(*int64)(unsafe.Add(mBase, uint32(v1233))) = v1531 + int64(1)
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1535)+20)) = int32(1)
	v1685 = v1529
	goto L170
L244:
	;
	goto L176
L245:
	;
	v1563 = *(*int32)(unsafe.Add(mBase, uint32(v1560)+4))
	v1564 = *(*int32)(unsafe.Add(mBase, uint32(v1560)+8))
	v1565 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+136)) = v1565
	*(*int64)(unsafe.Add(mBase, uint32(v21)+128)) = v1565
	*(*int64)(unsafe.Add(mBase, uint32(v21)+120)) = v1565
	*(*int64)(unsafe.Add(mBase, uint32(v21)+112)) = v1565
	*(*int64)(unsafe.Add(mBase, uint32(v21)+104)) = v1565
	*(*int64)(unsafe.Add(mBase, uint32(v21)+96)) = v1565
	*(*int64)(unsafe.Add(mBase, uint32(v21)+88)) = v1565
	*(*int64)(unsafe.Add(mBase, uint32(v21)+80)) = v1565
	*(*int64)(unsafe.Add(mBase, uint32(v21)+152)) = v1565
	*(*int64)(unsafe.Add(mBase, uint32(v21)+144)) = v1565
	*(*int32)(unsafe.Add(mBase, uint32(v1234)+12)) = v1559 + int32(1)
	v1591 = v1559*int32(120) + v1564
	v1594 = v1563 + v1559<<(uint(int32(4))%32)
	v1595 = *(*int32)(unsafe.Add(mBase, uint32(v1594)+12))
	if v1595 == int32(0) {
		goto L248
	} else {
		goto L249
	}
L246:
	;
	goto L247
L247:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v1675 = m.ExcPending
	if v1675 != 0 {
		goto L4
	} else {
		goto L268
	}
L248:
	;
	v1598 = *(*int32)(unsafe.Add(mBase, uint32(v1594)+8))
	v1601 = base.B2i32(v1598 != int32(-1))
	goto L250
L249:
	;
	v1601 = int32(2)
	goto L250
L250:
	;
	v1606 = *(*int32)(unsafe.Add(mBase, uint32(v1601<<(uint(int32(2))%32))+uint32(_consts[1094])))
	v1607 = F_cstring_to_text(m, v1606)
	mBase = m.M
	v1608 = m.ExcPending
	if v1608 != 0 {
		goto L4
	} else {
		goto L251
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = v1607
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(v1594)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = v1610
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(v1594)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+88)) = v1612
	if v1601 == int32(2) {
		goto L255
	} else {
		goto L256
	}
L252:
	;
	v1628 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+153)) = uint8(v1628)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+149)) = int32(16843009)
	v1632 = *(*int64)(unsafe.Add(mBase, uint32(v1591)))
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = v1632
	v1638 = F_pg_snprintf(m, v21+int32(176), int32(32), int32(39061), v21)
	mBase = m.M
	v1639 = m.ExcPending
	if v1639 != 0 {
		goto L4
	} else {
		goto L259
	}
L253:
	;
	v1626 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+147)) = uint8(v1626)
	goto L252
L254:
	;
	v1624 = *(*int32)(unsafe.Add(mBase, uint32(v1594)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+92)) = v1624
	goto L252
L255:
	;
	v1616 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1594)+12)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+96)) = v1616
	goto L254
L256:
	;
	goto L257
L257:
	;
	v1618 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+148)) = uint8(v1618)
	if base.Ui32(v1618) < base.Ui32(v1601-v1618) {
		goto L253
	} else {
		goto L258
	}
L258:
	;
	goto L254
L259:
	;
	v1642 = F_cstring_to_text(m, v21+int32(176))
	mBase = m.M
	v1643 = m.ExcPending
	if v1643 != 0 {
		goto L4
	} else {
		goto L260
	}
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+120)) = v1642
	v1645 = *(*int32)(unsafe.Add(mBase, uint32(v1591)+112))
	if v1645 != 0 {
		goto L262
	} else {
		goto L263
	}
L261:
	;
	v1650 = F_cstring_to_text(m, int32(315885))
	mBase = m.M
	v1651 = m.ExcPending
	if v1651 != 0 {
		goto L4
	} else {
		goto L265
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+124)) = v1645
	goto L261
L263:
	;
	goto L264
L264:
	;
	v1647 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+155)) = uint8(v1647)
	goto L261
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+128)) = v1650
	v1653 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+159)) = uint8(v1653)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+132)) = int64(1)
	v1657 = *(*int32)(unsafe.Add(mBase, uint32(v1233)+28))
	v1662 = F_heap_form_tuple(m, v1657, v21+int32(80), v21+int32(144))
	mBase = m.M
	v1663 = m.ExcPending
	if v1663 != 0 {
		goto L4
	} else {
		goto L266
	}
L266:
	;
	v1664 = *(*int32)(unsafe.Add(mBase, uint32(v1662)+16))
	v1665 = F_HeapTupleHeaderGetDatum(m, v1664)
	mBase = m.M
	v1666 = m.ExcPending
	if v1666 != 0 {
		goto L4
	} else {
		goto L267
	}
L267:
	;
	v1667 = *(*int64)(unsafe.Add(mBase, uint32(v1233)))
	*(*int64)(unsafe.Add(mBase, uint32(v1233))) = v1667 + int64(1)
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1671)+20)) = int32(1)
	v1685 = v1665
	goto L170
L268:
	;
	v1676 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1676)+20)) = int32(2)
	v1679 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v1679)
	v1685 = int32(0)
	goto L170
}
func F_pg_log_backend_memory_contexts(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_BackendPidGetProc(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v7 + int32(16)
	return v91
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
	F_errmsg(m, v81, v7)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L24
	}
L3:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _consts[153]))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v66 = base.I32_div_s(v58-v63, int32(640))
	v67 = F_SendProcSignal(m, v9, int32(5), v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L20
	}
L4:
	;
	return int32(0)
L5:
	;
	if v10 != 0 {
		v58 = v10
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v14 = int32(0)
	if v9 == v14 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v48 != 0 {
		v58 = v48
		goto L3
	} else {
		goto L17
	}
L8:
	;
	v48 = int32(0)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _consts[1097]))
	v23 = v14
	goto L12
L11:
	;
	v48 = v43
	goto L7
L12:
	;
	v28 = v21 + v23*int32(640)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+44))
	if v29 == v9 {
		v43 = v28
		goto L11
	} else {
		goto L14
	}
L13:
	;
	v48 = int32(0)
	goto L7
L14:
	;
	v35 = v21 + (v23|int32(1))*int32(640)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+44))
	if v36 == v9 {
		v43 = v35
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v39 = v23 + int32(2)
	if v39 != int32(38) {
		v23 = v39
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	v49 = int32(0)
	v52 = F_errstart(m, int32(19), v49)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	if v52 == int32(0) {
		v91 = v49
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v81 = int32(128225)
	v82 = int32(293)
	goto L2
L20:
	;
	if int32(0) <= v67 {
		v91 = int32(1)
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v71 = int32(0)
	v74 = F_errstart(m, int32(19), v71)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	if v74 == int32(0) {
		v91 = v71
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v81 = int32(294051)
	v82 = int32(302)
	goto L2
L24:
	;
	F_errfinish(m, int32(491404), v82, int32(114597))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v91 = int32(0)
	goto L1
}
func F_pg_md5_binary(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	v5 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v5
	v9 = F_pg_cryptohash_create(m, v5)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(13869)
			return int32(0)
		} else {
			v34 = F_pg_cryptohash_init(m, v9)
			mBase = m.M
			if v34 < int32(0) {
				if v9 == int32(0) {
					v58 = int32(13869)
				} else {
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
					if v50 == int32(1) {
						v53 = int32(302852)
					} else {
						v53 = int32(129434)
					}
					if v50 == int32(2) {
						v56 = int32(13869)
					} else {
						v56 = v53
					}
					v58 = v56
				}
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v58
				F_pg_cryptohash_free(m, v9)
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			} else {
				v37 = F_pg_cryptohash_update(m, v9, l0, l1)
				mBase = m.M
				if v37 < int32(0) {
					if v9 == int32(0) {
						v58 = int32(13869)
					} else {
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
						if v50 == int32(1) {
							v53 = int32(302852)
						} else {
							v53 = int32(129434)
						}
						if v50 == int32(2) {
							v56 = int32(13869)
						} else {
							v56 = v53
						}
						v58 = v56
					}
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v58
					F_pg_cryptohash_free(m, v9)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				} else {
					v41 = F_pg_cryptohash_final(m, v9, l2, int32(16))
					mBase = m.M
					if int32(0) <= v41 {
						F_pg_cryptohash_free(m, v9)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							return int32(1)
						}
					} else {
						if v9 == int32(0) {
							v58 = int32(13869)
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
							if v50 == int32(1) {
								v53 = int32(302852)
							} else {
								v53 = int32(129434)
							}
							if v50 == int32(2) {
								v56 = int32(13869)
							} else {
								v56 = v53
							}
							v58 = v56
						}
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v58
						F_pg_cryptohash_free(m, v9)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							return int32(0)
						}
					}
				}
			}
		}
	}
}
func F_pg_ndistinct_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(108474)
			F_errmsg(m, int32(191561), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(490244), int32(343), int32(277670))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_pg_node_tree_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(407669)
			F_errmsg(m, int32(191561), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(490975), int32(335), int32(36537))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_pg_parse_json(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	if l0 == int32(4488392) {
		return int32(16)
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
		if v8 != 0 {
			return int32(2)
		} else {
			v11 = F_json_lex(m, l0)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				if v11 != 0 {
					v35 = v11
					return v35
				} else {
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					switch v15 - int32(3) {
					case 0:
						v18 = F_parse_object(m, l0, l1)
						mBase = m.M
						v19 = m.ExcPending
						if v19 != 0 {
							return int32(0)
						} else {
							v24 = v18
							if v24 != 0 {
								v35 = v24
								return v35
							} else {
								v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v25 == int32(12) {
									v28 = F_json_lex(m, l0)
									mBase = m.M
									v29 = m.ExcPending
									if v29 != 0 {
										return int32(0)
									} else {
										return v28
									}
								} else {
									v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v33 != 0 {
										v34 = int32(9)
									} else {
										v34 = int32(11)
									}
									v35 = v34
									return v35
								}
							}
						}
					default:
						v22 = F_parse_scalar(m, l0, l1)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							v24 = v22
							if v24 != 0 {
								v35 = v24
								return v35
							} else {
								v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v25 == int32(12) {
									v28 = F_json_lex(m, l0)
									mBase = m.M
									v29 = m.ExcPending
									if v29 != 0 {
										return int32(0)
									} else {
										return v28
									}
								} else {
									v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v33 != 0 {
										v34 = int32(9)
									} else {
										v34 = int32(11)
									}
									v35 = v34
									return v35
								}
							}
						}
					case 2:
						v20 = F_parse_array(m, l0, l1)
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return int32(0)
						} else {
							v24 = v20
							if v24 != 0 {
								v35 = v24
								return v35
							} else {
								v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v25 == int32(12) {
									v28 = F_json_lex(m, l0)
									mBase = m.M
									v29 = m.ExcPending
									if v29 != 0 {
										return int32(0)
									} else {
										return v28
									}
								} else {
									v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v33 != 0 {
										v34 = int32(9)
									} else {
										v34 = int32(11)
									}
									v35 = v34
									return v35
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_plan_query(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v6 == int32(6) {
		v82 = int32(0)
		return v82
	} else {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, _consts[849])))
		if v10 == int32(1) {
			v13 = int32(4410568)
			v14 = int32(0)
			v18 = m.G0
			v20 = v18 - int32(16)
			m.G0 = v20
			v23 = int32(4410584)
			v28 = F___memset(m, int32(4410592), v14, int32(144))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, _consts[845])) = int32(4)
			*(*int64)(unsafe.Add(mBase, _consts[846])) = int64(3)
			*(*int32)(unsafe.Add(mBase, _consts[831])) = int32(2)
			*(*int64)(unsafe.Add(mBase, _consts[834])) = int64(1)
			v41 = F___memcpy(m, v20, v23, int32(16))
			mBase = m.M
			v42 = int64(*(*int32)(unsafe.Add(mBase, uint32(v20))))
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
			v44 = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[847])) = v44
			*(*int32)(unsafe.Add(mBase, _consts[830])) = v43
			*(*int64)(unsafe.Add(mBase, _consts[833])) = v42
			v48 = int64(*(*int32)(unsafe.Add(mBase, uint32(v20)+8)))
			v49 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
			*(*int32)(unsafe.Add(mBase, _consts[848])) = v44
			*(*int32)(unsafe.Add(mBase, _consts[831])) = v49
			*(*int64)(unsafe.Add(mBase, _consts[834])) = v48
			v56 = F___syscall_ret(m, v14)
			mBase = m.M
			m.G0 = v20 + int32(16)
			F___gettimeofday(m, int32(4410720))
			mBase = m.M
		} else {
		}
		v62 = F_planner(m, l0, l1, l2, l3)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			v67 = int32(*(*uint8)(unsafe.Add(mBase, _consts[849])))
			if v67 == int32(1) {
				F_ShowUsage(m, int32(520657))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					v74 = int32(*(*uint8)(unsafe.Add(mBase, _consts[850])))
					if v74 != int32(1) {
						v82 = v62
						return v82
					} else {
						v79 = int32(*(*uint8)(unsafe.Add(mBase, _consts[851])))
						F_elog_node_display(m, int32(281762), v62, v79)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							v82 = v62
							return v82
						}
					}
				}
			} else {
				v74 = int32(*(*uint8)(unsafe.Add(mBase, _consts[850])))
				if v74 != int32(1) {
					v82 = v62
					return v82
				} else {
					v79 = int32(*(*uint8)(unsafe.Add(mBase, _consts[851])))
					F_elog_node_display(m, int32(281762), v62, v79)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						v82 = v62
						return v82
					}
				}
			}
		}
	}
}
func F_pg_prepared_xact(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int64
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int64
	_ = v291
	var v295 int32
	_ = v295
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v18 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+16))
	goto L45
L4:
	;
	return int32(0)
L5:
	;
	v25 = int32(4486928)
	v26 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v28
	v31 = F_CreateTemplateTupleDesc(m, int32(5))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	F_TupleDescInitEntry(m, v31, int32(1), int32(256538), int32(28), int32(-1), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	F_TupleDescInitEntry(m, v31, int32(2), int32(433419), int32(25), int32(-1), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	F_TupleDescInitEntry(m, v31, int32(3), int32(448657), int32(1184), int32(-1), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	F_TupleDescInitEntry(m, v31, int32(4), int32(430646), int32(26), int32(-1), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	F_TupleDescInitEntry(m, v31, int32(5), int32(433520), int32(26), int32(-1), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v68 = F_BlessTupleDesc(m, v31)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v68
	v72 = F_palloc(m, int32(12))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v72
	v76 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v80 = F_LWLockAcquire(m, v76+int32(2304), int32(1))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _consts[152]))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	if v84 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v84
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v26
	goto L3
L16:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v88+int32(2304))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v97 = F_palloc(m, v84*int32(248))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = int32(0)
	goto L15
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v97
	if v84 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v189+int32(2304))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L42
	}
L22:
	;
	v102 = int32(1)
	v104 = int32(0)
	v106 = *(*int32)(unsafe.Add(mBase, _consts[152]))
	v108 = v106 + int32(8)
	if v84 != v102 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v114 = v104
	v121 = v2
	goto L26
L24:
	;
	v153 = v104
	goto L25
L25:
	;
	if v84&v102 == int32(0) {
		goto L21
	} else {
		goto L37
	}
L26:
	;
	v125 = int32(248)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v108+v114<<(uint(int32(2))%32))))
	goto L29
L27:
	;
	v153 = v148
	goto L25
L28:
	;
	v136 = v114 | int32(1)
	v137 = int32(248)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v108+v136<<(uint(int32(2))%32))))
	goto L33
L29:
	;
	v133 = F__emscripten_memcpy_bulkmem(m, v97+v114*v125, v131, v125)
	mBase = m.M
	goto L31
L31:
	;
	goto L28
L32:
	;
	v147 = int32(2)
	v148 = v114 + v147
	v150 = v121 + v147
	if v150 != v84&int32(2147483646) {
		v114 = v148
		v121 = v150
		goto L26
	} else {
		goto L36
	}
L33:
	;
	v145 = F__emscripten_memcpy_bulkmem(m, v97+v136*v137, v143, v137)
	mBase = m.M
	goto L35
L35:
	;
	goto L32
L36:
	;
	goto L27
L37:
	;
	v166 = int32(248)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v108+v153<<(uint(int32(2))%32))))
	goto L39
L38:
	;
	goto L21
L39:
	;
	v174 = F__emscripten_memcpy_bulkmem(m, v97+v153*v166, v172, v166)
	mBase = m.M
	goto L41
L41:
	;
	goto L38
L42:
	;
	goto L15
L43:
	;
	m.G0 = v15 + int32(48)
	return v320
L44:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L4
	} else {
		goto L58
	}
L45:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+16))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	if v226 == int32(0) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v225)+8))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v225)+4))
	if v230 <= v229 {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v233 = *(*int32)(unsafe.Add(mBase, _consts[153]))
	v237 = v229
	goto L48
L48:
	;
	v248 = int32(1)
	v249 = v237 + v248
	*(*int32)(unsafe.Add(mBase, uint32(v225)+8)) = v249
	v253 = v226 + v237*int32(248)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)+4))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	v256 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15+int32(12)))) = uint8(v256)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v256
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253)+44)))
	if v260 == v248 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L44
L50:
	;
	v265 = v255 + v254*int32(640)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v266
	v270 = F_cstring_to_text(m, v253+int32(47))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L4
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	if v249 < v230 {
		v237 = v249
		goto L48
	} else {
		goto L57
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v270
	v273 = *(*int64)(unsafe.Add(mBase, uint32(v253)+8))
	v274 = F_Int64GetDatum(m, v273)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v274
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v253)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v277
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v265)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v279
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v224)+28))
	v286 = F_heap_form_tuple(m, v281, v15+int32(16), v15+int32(8))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v286)+16))
	v289 = F_HeapTupleHeaderGetDatum(m, v288)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	v291 = *(*int64)(unsafe.Add(mBase, uint32(v224)))
	*(*int64)(unsafe.Add(mBase, uint32(v224))) = v291 + int64(1)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v295)+20)) = int32(1)
	v320 = v289
	goto L43
L57:
	;
	goto L49
L58:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v313)+20)) = int32(2)
	v316 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v316)
	v320 = int32(0)
	goto L43
}
func F_pg_promote(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
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
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, _consts[182])))
	if v14 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	v220 = F_unlink(m, int32(347095))
	mBase = m.M
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L13
	} else {
		goto L70
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L13
	} else {
		goto L66
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L13
	} else {
		goto L62
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L13
	} else {
		goto L58
	}
L5:
	;
	if v24 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+316))
	v22 = base.B2i32(v20 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[182])) = uint8(v22)
	v24 = v22
	goto L8
L7:
	;
	v24 = int32(0)
	goto L8
L8:
	;
	goto L5
L9:
	;
	if v10 <= int32(0) {
		goto L4
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L13
	} else {
		goto L53
	}
L12:
	;
	v29 = F_AllocateFile(m, int32(347095), int32(32454))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	if v29 == int32(0) {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v35 = F_FreeFile(m, v29)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	if v35 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _consts[232]))
	v40 = F_kill(m, v38, int32(10))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	if v40 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v11 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	m.G0 = v8 + int32(48)
	return v144
L21:
	;
	v144 = int32(1)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v45 = int32(0)
	v47 = v10 * int32(10)
	if v47 <= v45 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v121 = int32(0)
	v124 = F_errstart(m, int32(19), v121)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L13
	} else {
		goto L49
	}
L25:
	;
	v51 = v45
	goto L26
L26:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _consts[96]))
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = int32(0)
	goto L28
L27:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L13
	} else {
		goto L43
	}
L28:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, _consts[182])))
	if v62 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v72 == int32(0) {
		v144 = int32(1)
		goto L20
	} else {
		goto L33
	}
L30:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+316))
	v70 = base.B2i32(v68 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[182])) = uint8(v70)
	v72 = v70
	goto L32
L31:
	;
	v72 = int32(0)
	goto L32
L32:
	;
	goto L29
L33:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v76 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L13
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _consts[96]))
	v84 = F_WaitLatch(m, v80, int32(25), int32(100), int32(134217771))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L13
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	if v84&int32(16) == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v91 = v51 + int32(1)
	if v91 == v47 {
		goto L24
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	goto L27
L42:
	;
	v51 = v91
	goto L26
L43:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L13
	} else {
		goto L44
	}
L44:
	;
	F_errmsg(m, int32(99021), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L13
	} else {
		goto L46
	}
L46:
	;
	F_errcontext_msg(m, int32(246680), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(491704), int32(741), int32(347092))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L13
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	if v124 == int32(0) {
		v144 = v121
		goto L20
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v10
	F_errmsg_plural(m, int32(422655), int32(171909), v10, v8+int32(16))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L13
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(491704), int32(748), int32(347092))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L13
	} else {
		goto L52
	}
L52:
	;
	v144 = v121
	goto L20
L53:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L13
	} else {
		goto L54
	}
L54:
	;
	F_errmsg(m, int32(127191), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L13
	} else {
		goto L55
	}
L55:
	;
	F_errhint(m, int32(554560), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L13
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(491704), int32(681), int32(347092))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L13
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L13
	} else {
		goto L59
	}
L59:
	;
	F_errmsg(m, int32(238363), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L13
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(491704), int32(686), int32(347092))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L13
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L13
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(347095)
	F_errmsg(m, int32(297547), v8)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L13
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(491704), int32(694), int32(347092))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L13
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L13
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(347095)
	F_errmsg(m, int32(297453), v8+int32(32))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L13
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(491704), int32(700), int32(347092))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L13
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	F_errcode(m, int32(517))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L13
	} else {
		goto L71
	}
L71:
	;
	F_errmsg(m, int32(291930), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L13
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(491704), int32(708), int32(347092))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L13
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_read_binary_file_all_missing(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = F_convert_and_check_filename(m, v5)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v16 = F_read_binary_file(m, v10, int64(0), int64(-1), base.B2i32(v9 != int32(0)))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				if v16 == int32(0) {
					v20 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v20)
					return int32(0)
				} else {
					return v16
				}
			}
		}
	}
}
func F_pg_read_binary_file_off_len_missing(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		if int64(0) <= v12 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v17 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
			v18 = F_convert_and_check_filename(m, v7)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v22 = F_read_binary_file(m, v18, v17, v12, base.B2i32(v15 != int32(0)))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					if v22 == int32(0) {
						v26 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
						return int32(0)
					} else {
						return v22
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(341102), int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(495772), int32(269), int32(244623))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_reg_getcolor(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v254 int32
	_ = v254
	var v262 int32
	_ = v262
	var v275 int32
	_ = v275
	var v286 int32
	_ = v286
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v443 int32
	_ = v443
	var v450 int32
	_ = v450
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v550 int32
	_ = v550
	var v564 int32
	_ = v564
	var v570 int32
	_ = v570
	var v578 int32
	_ = v578
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v614 int32
	_ = v614
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v713 int32
	_ = v713
	var v723 int32
	_ = v723
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v761 int32
	_ = v761
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v878 int32
	_ = v878
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v894 int32
	_ = v894
	var v900 int32
	_ = v900
	var v911 int32
	_ = v911
	var v926 int32
	_ = v926
	var v932 int32
	_ = v932
	var v938 int32
	_ = v938
	var v949 int32
	_ = v949
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v967 int32
	_ = v967
	var v974 int32
	_ = v974
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v1002 int32
	_ = v1002
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1015 int32
	_ = v1015
	var v1025 int32
	_ = v1025
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1065 int32
	_ = v1065
	var v1072 int32
	_ = v1072
	var v1087 int32
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1097 int32
	_ = v1097
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1123 int32
	_ = v1123
	var v1129 int32
	_ = v1129
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1147 int32
	_ = v1147
	var v1151 int32
	_ = v1151
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1200 int32
	_ = v1200
	var v1206 int32
	_ = v1206
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1224 int32
	_ = v1224
	var v1228 int32
	_ = v1228
	var v1241 int32
	_ = v1241
	var v1244 int32
	_ = v1244
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1251 int32
	_ = v1251
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1277 int32
	_ = v1277
	var v1283 int32
	_ = v1283
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1291 int32
	_ = v1291
	var v1294 int32
	_ = v1294
	var v1297 int32
	_ = v1297
	var v1303 int32
	_ = v1303
	var v1308 int32
	_ = v1308
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1325 int32
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1330 int32
	_ = v1330
	var v1336 int32
	_ = v1336
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1346 int32
	_ = v1346
	var v1360 int32
	_ = v1360
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1369 int32
	_ = v1369
	var v1381 int32
	_ = v1381
	var v1384 int32
	_ = v1384
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1390 int32
	_ = v1390
	var v1396 int32
	_ = v1396
	var v1398 int32
	_ = v1398
	var v1402 int32
	_ = v1402
	v3 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v3 < v8 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if int32(2) <= v52 {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v51 = v43
	goto L1
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v14 = v3
	v15 = v8
	goto L6
L4:
	;
	goto L5
L5:
	;
	v51 = int32(0)
	goto L1
L6:
	;
	v21 = base.I32_div_s(v15-v14, int32(2))
	v22 = v21 + v14
	v25 = v11 + v22*int32(12)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if base.Ui32(l1) < base.Ui32(v26) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L5
L8:
	;
	if v32 < v33 {
		v14 = v32
		v15 = v33
		goto L6
	} else {
		goto L13
	}
L9:
	;
	v32 = v14
	v33 = v22
	goto L8
L10:
	;
	goto L11
L11:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if base.Ui32(l1) <= base.Ui32(v28) {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v32 = v22 + int32(1)
	v33 = v15
	goto L8
L13:
	;
	goto L7
L14:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v55 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L16
L16:
	;
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v1402 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1398+v51<<(uint(int32(1))%32)))))
	return v1402
L17:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v314 == int32(0) {
		v487 = v313
		goto L92
	} else {
		goto L93
	}
L18:
	;
	v313 = int32(0)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v59 = int32(0)
	v61 = *(*int32)(unsafe.Add(mBase, _consts[586]))
	switch v61 - int32(1) {
	case 0:
		goto L24
	case 1:
		goto L23
	case 2:
		goto L22
	default:
		goto L25
	}
L21:
	;
	if v309 == int32(0) {
		v313 = v59
		goto L17
	} else {
		goto L91
	}
L22:
	;
	if base.Ui32(int32(255)) < base.Ui32(l1) {
		v313 = v59
		goto L17
	} else {
		goto L89
	}
L23:
	;
	if base.Ui32(l1) <= base.Ui32(int32(254)) {
		goto L82
	} else {
		goto L83
	}
L24:
	;
	if base.Ui32(int32(127)) < base.Ui32(l1) {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	v309 = base.B2i32(base.Ui32(l1-int32(32)) < base.Ui32(int32(95)))
	goto L21
L26:
	;
	v309 = v262
	goto L21
L27:
	;
	if base.Ui32(int32(128)) <= base.Ui32(l1) {
		goto L47
	} else {
		goto L48
	}
L28:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	if v112 != int32(15) {
		goto L27
	} else {
		goto L41
	}
L29:
	;
	v111 = v84 + int32(1922968)
	goto L28
L30:
	;
	v76 = int32(3367)
	v77 = int32(0)
	goto L33
L31:
	;
	goto L32
L32:
	;
	v111 = l1<<(uint(int32(1))%32) + int32(1898352)
	goto L28
L33:
	;
	v82 = base.I32_div_s(v76+v77, int32(2))
	v84 = v82 * int32(12)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v84)+uint32(_consts[588])))
	if base.Ui32(v87) < base.Ui32(l1) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L27
L35:
	;
	if v98 <= v97 {
		v76 = v97
		v77 = v98
		goto L33
	} else {
		goto L40
	}
L36:
	;
	v97 = v76
	v98 = v82 + int32(1)
	goto L35
L37:
	;
	goto L38
L38:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v84)+uint32(_consts[589])))
	if base.Ui32(v93) <= base.Ui32(l1) {
		goto L29
	} else {
		goto L39
	}
L39:
	;
	v97 = v82 - int32(1)
	v98 = v77
	goto L35
L40:
	;
	goto L34
L41:
	;
	v262 = int32(0)
	goto L26
L42:
	;
	v262 = v254
	goto L26
L43:
	;
	v254 = base.B2i32(v244&int32(255) == int32(12))
	goto L42
L44:
	;
	if int32(1)<<(uint(v176)%32)&int32(294913) != 0 {
		goto L62
	} else {
		goto L63
	}
L45:
	;
	if l1 == int32(9) {
		v254 = v151
		goto L42
	} else {
		goto L60
	}
L46:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+uint32(_consts[590]))))
	v176 = v165
	goto L44
L47:
	;
	v126 = int32(3367)
	v127 = int32(0)
	goto L50
L48:
	;
	goto L49
L49:
	;
	v151 = int32(1)
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1<<(uint(v151)%32))+uint32(_consts[591]))))
	if v151<<(uint(v157)%32)&int32(294913) == int32(0) {
		goto L45
	} else {
		goto L58
	}
L50:
	;
	v132 = base.I32_div_s(v126+v127, int32(2))
	v134 = v132 * int32(12)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v134)+uint32(_consts[588])))
	if base.Ui32(v137) < base.Ui32(l1) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v176 = int32(0)
	goto L44
L52:
	;
	if v148 <= v147 {
		v126 = v147
		v127 = v148
		goto L50
	} else {
		goto L57
	}
L53:
	;
	v147 = v126
	v148 = v132 + int32(1)
	goto L52
L54:
	;
	goto L55
L55:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v134)+uint32(_consts[589])))
	if base.Ui32(v143) <= base.Ui32(l1) {
		goto L46
	} else {
		goto L56
	}
L56:
	;
	v147 = v132 - int32(1)
	v148 = v127
	goto L52
L57:
	;
	goto L51
L58:
	;
	if l1 != int32(9) {
		v244 = v157
		goto L43
	} else {
		goto L59
	}
L59:
	;
	v254 = v151
	goto L42
L60:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1<<(uint(int32(1))%32))+uint32(_consts[592]))))
	if v172&int32(32) != 0 {
		v244 = v157
		goto L43
	} else {
		goto L61
	}
L61:
	;
	v254 = v151
	goto L42
L62:
	;
	v217 = int32(3367)
	v218 = int32(0)
	goto L73
L63:
	;
	v184 = int32(10)
	v185 = int32(0)
	goto L64
L64:
	;
	v190 = base.I32_div_s(v184+v185, int32(2))
	v192 = v190 << (uint(int32(3)) % 32)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v192)+uint32(_consts[593])))
	if base.Ui32(v195) < base.Ui32(l1) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v262 = int32(1)
	goto L26
L66:
	;
	if v206 <= v205 {
		v184 = v205
		v185 = v206
		goto L64
	} else {
		goto L71
	}
L67:
	;
	v205 = v184
	v206 = v190 + int32(1)
	goto L66
L68:
	;
	goto L69
L69:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v192)+uint32(_consts[594])))
	if base.Ui32(v201) <= base.Ui32(l1) {
		goto L62
	} else {
		goto L70
	}
L70:
	;
	v205 = v190 - int32(1)
	v206 = v185
	goto L66
L71:
	;
	goto L65
L72:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+uint32(_consts[590]))))
	v244 = v242
	goto L43
L73:
	;
	v223 = base.I32_div_s(v217+v218, int32(2))
	v225 = v223 * int32(12)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v225)+uint32(_consts[588])))
	if base.Ui32(v228) < base.Ui32(l1) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v244 = int32(0)
	goto L43
L75:
	;
	if v239 <= v238 {
		v217 = v238
		v218 = v239
		goto L73
	} else {
		goto L80
	}
L76:
	;
	v238 = v217
	v239 = v223 + int32(1)
	goto L75
L77:
	;
	goto L78
L78:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v225)+uint32(_consts[589])))
	if base.Ui32(v234) <= base.Ui32(l1) {
		goto L72
	} else {
		goto L79
	}
L79:
	;
	v238 = v223 - int32(1)
	v239 = v218
	goto L75
L80:
	;
	goto L74
L81:
	;
	v309 = v297
	goto L21
L82:
	;
	v297 = base.B2i32(base.Ui32(int32(32)) < base.Ui32((l1+int32(1))&int32(127)))
	goto L81
L83:
	;
	goto L84
L84:
	;
	v275 = int32(1)
	if base.Ui32(l1-int32(57344)) < base.Ui32(int32(8185)) {
		v295 = v275
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v297 = v295
	goto L81
L86:
	;
	if base.Ui32(l1) < base.Ui32(int32(8232)) {
		v295 = v275
		goto L85
	} else {
		goto L87
	}
L87:
	;
	if base.Ui32(l1-int32(8234)) < base.Ui32(int32(47062)) {
		v295 = v275
		goto L85
	} else {
		goto L88
	}
L88:
	;
	v286 = int32(65534)
	v295 = base.B2i32(l1&v286 != v286) & base.B2i32(base.Ui32(l1-int32(65532)) < base.Ui32(int32(1048580)))
	goto L85
L89:
	;
	goto L90
L90:
	;
	v309 = base.B2i32(base.B2i32(base.Ui32(l1-int32(32)) < base.Ui32(int32(95))) != int32(0))
	goto L21
L91:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v313 = v312
	goto L17
L92:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v488 == int32(0) {
		v597 = v487
		goto L137
	} else {
		goto L138
	}
L93:
	;
	v318 = *(*int32)(unsafe.Add(mBase, _consts[586]))
	switch v318 - int32(1) {
	case 0:
		goto L97
	case 1:
		goto L96
	case 2:
		goto L95
	default:
		goto L98
	}
L94:
	;
	if v482 == int32(0) {
		v487 = v313
		goto L92
	} else {
		goto L136
	}
L95:
	;
	if base.Ui32(int32(255)) < base.Ui32(l1) {
		v487 = v313
		goto L92
	} else {
		goto L134
	}
L96:
	;
	if base.Ui32(int32(10)) <= base.Ui32(l1-int32(48)) {
		goto L131
	} else {
		goto L132
	}
L97:
	;
	v331 = *(*int32)(unsafe.Add(mBase, _consts[587]))
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331)+16)))
	v336 = (v332 ^ int32(-1)) & int32(1)
	if base.Ui32(int32(128)) <= base.Ui32(l1) {
		goto L105
	} else {
		goto L106
	}
L98:
	;
	if base.Ui32(int32(127)) < base.Ui32(l1) {
		v487 = v313
		goto L92
	} else {
		goto L99
	}
L99:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[595]))))
	v482 = base.B2i32(v325&int32(3) != int32(0))
	goto L94
L100:
	;
	v482 = v450
	goto L94
L101:
	;
	v450 = v443
	goto L100
L102:
	;
	v443 = base.B2i32(v432&int32(255) == int32(9))
	goto L101
L103:
	;
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1<<(uint(int32(1))%32))+uint32(_consts[591]))))
	v432 = v425
	goto L102
L104:
	;
	v450 = base.B2i32(base.Ui32(l1-int32(48)) < base.Ui32(int32(10)))
	goto L100
L105:
	;
	v346 = int32(1178)
	v347 = int32(0)
	goto L108
L106:
	;
	goto L107
L107:
	;
	v401 = int32(1)
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1<<(uint(v401)%32))+uint32(_consts[592]))))
	if v406&v401 != 0 {
		v443 = v401
		goto L101
	} else {
		goto L128
	}
L108:
	;
	v352 = base.I32_div_s(v346+v347, int32(2))
	v354 = v352 << (uint(int32(3)) % 32)
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v354)+uint32(_consts[596])))
	if base.Ui32(v357) < base.Ui32(l1) {
		goto L111
	} else {
		goto L112
	}
L109:
	;
	if v336 != 0 {
		goto L104
	} else {
		goto L118
	}
L110:
	;
	if v369 <= v368 {
		v346 = v368
		v347 = v369
		goto L108
	} else {
		goto L117
	}
L111:
	;
	v368 = v346
	v369 = v352 + int32(1)
	goto L110
L112:
	;
	goto L113
L113:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v354)+uint32(_consts[597])))
	if base.Ui32(v363) <= base.Ui32(l1) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v450 = int32(1)
	goto L100
L115:
	;
	goto L116
L116:
	;
	v368 = v352 - int32(1)
	v369 = v347
	goto L110
L117:
	;
	goto L109
L118:
	;
	v375 = int32(3367)
	v376 = int32(0)
	goto L120
L119:
	;
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383)+uint32(_consts[590]))))
	v432 = v400
	goto L102
L120:
	;
	v381 = base.I32_div_s(v375+v376, int32(2))
	v383 = v381 * int32(12)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v383)+uint32(_consts[588])))
	if base.Ui32(v386) < base.Ui32(l1) {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	v432 = int32(0)
	goto L102
L122:
	;
	if v397 <= v396 {
		v375 = v396
		v376 = v397
		goto L120
	} else {
		goto L127
	}
L123:
	;
	v396 = v375
	v397 = v381 + int32(1)
	goto L122
L124:
	;
	goto L125
L125:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v383)+uint32(_consts[589])))
	if base.Ui32(v392) <= base.Ui32(l1) {
		goto L119
	} else {
		goto L126
	}
L126:
	;
	v396 = v381 - int32(1)
	v397 = v376
	goto L122
L127:
	;
	goto L121
L128:
	;
	if v336 == int32(0) {
		goto L103
	} else {
		goto L129
	}
L129:
	;
	goto L104
L130:
	;
	v482 = v463
	goto L94
L131:
	;
	v460 = F_iswalpha(m, l1)
	mBase = m.M
	v463 = base.B2i32(v460 != int32(0))
	goto L133
L132:
	;
	v463 = int32(1)
	goto L133
L133:
	;
	goto L130
L134:
	;
	goto L135
L135:
	;
	v482 = base.B2i32(base.B2i32(base.Ui32(l1-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(l1|int32(32)-int32(97)) < base.Ui32(int32(26))) != int32(0))
	goto L94
L136:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v487 = v485 | v313
	goto L92
L137:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v598 == int32(0) {
		v646 = v597
		goto L166
	} else {
		goto L167
	}
L138:
	;
	v492 = *(*int32)(unsafe.Add(mBase, _consts[586]))
	switch v492 - int32(1) {
	case 0:
		goto L142
	case 1:
		goto L141
	case 2:
		goto L140
	default:
		goto L143
	}
L139:
	;
	if v592 == int32(0) {
		v597 = v487
		goto L137
	} else {
		goto L165
	}
L140:
	;
	if base.Ui32(int32(255)) < base.Ui32(l1) {
		v597 = v487
		goto L137
	} else {
		goto L163
	}
L141:
	;
	if base.Ui32(l1) <= base.Ui32(int32(131071)) {
		goto L160
	} else {
		goto L161
	}
L142:
	;
	if base.Ui32(int32(127)) < base.Ui32(l1) {
		goto L146
	} else {
		goto L147
	}
L143:
	;
	if base.Ui32(int32(127)) < base.Ui32(l1) {
		v597 = v487
		goto L137
	} else {
		goto L144
	}
L144:
	;
	v499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[595]))))
	v500 = int32(1)
	v592 = int32(base.Ui32(v499)>>(uint(v500)%32)) & v500
	goto L139
L145:
	;
	v592 = v550
	goto L139
L146:
	;
	v512 = int32(0)
	v513 = int32(1178)
	goto L149
L147:
	;
	goto L148
L148:
	;
	v538 = int32(1)
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1<<(uint(v538)%32))+uint32(_consts[592]))))
	v550 = v542 & v538
	goto L145
L149:
	;
	v518 = base.I32_div_s(v512+v513, int32(2))
	v520 = v518 << (uint(int32(3)) % 32)
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v520)+uint32(_consts[596])))
	if base.Ui32(v523) < base.Ui32(l1) {
		goto L152
	} else {
		goto L153
	}
L150:
	;
	v550 = int32(0)
	goto L145
L151:
	;
	if v534 <= v535 {
		v512 = v534
		v513 = v535
		goto L149
	} else {
		goto L158
	}
L152:
	;
	v534 = v518 + int32(1)
	v535 = v513
	goto L151
L153:
	;
	goto L154
L154:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v520)+uint32(_consts[597])))
	if base.Ui32(v529) <= base.Ui32(l1) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v550 = int32(1)
	goto L145
L156:
	;
	goto L157
L157:
	;
	v534 = v512
	v535 = v518 - int32(1)
	goto L151
L158:
	;
	goto L150
L159:
	;
	v592 = v578
	goto L139
L160:
	;
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l1)>>(uint(int32(8))%32)))+uint32(_consts[598]))))
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l1)>>(uint(int32(3))%32))&int32(31)|v564<<(uint(int32(5))%32))+uint32(_consts[598]))))
	v578 = int32(base.Ui32(v570)>>(uint(l1&int32(7))%32)) & int32(1)
	goto L159
L161:
	;
	goto L162
L162:
	;
	v578 = base.B2i32(base.Ui32(l1) < base.Ui32(int32(196606)))
	goto L159
L163:
	;
	goto L164
L164:
	;
	v592 = base.B2i32(base.B2i32(base.Ui32(l1|int32(32)-int32(97)) < base.Ui32(int32(26))) != int32(0))
	goto L139
L165:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v597 = v595 | v487
	goto L137
L166:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v647 == int32(0) {
		v747 = v646
		goto L178
	} else {
		goto L179
	}
L167:
	;
	if l1 == int32(95) {
		v639 = int32(1)
		goto L169
	} else {
		goto L170
	}
L168:
	;
	if v641 == int32(0) {
		v646 = v597
		goto L166
	} else {
		goto L177
	}
L169:
	;
	v641 = v639
	goto L168
L170:
	;
	v605 = int32(0)
	v607 = *(*int32)(unsafe.Add(mBase, _consts[586]))
	switch v607 - int32(1) {
	case 0:
		goto L173
	case 1:
		goto L172
	case 2:
		goto L171
	default:
		goto L174
	}
L171:
	;
	if base.Ui32(int32(255)) < base.Ui32(l1) {
		v639 = v605
		goto L169
	} else {
		goto L176
	}
L172:
	;
	v630 = F_iswalnum(m, l1)
	mBase = m.M
	v641 = v630
	goto L168
L173:
	;
	v620 = *(*int32)(unsafe.Add(mBase, _consts[587]))
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v620)+16)))
	v626 = F_pg_u_isalnum(m, l1, (v621^int32(-1))&int32(1))
	mBase = m.M
	v641 = v626
	goto L168
L174:
	;
	if base.Ui32(int32(127)) < base.Ui32(l1) {
		v639 = v605
		goto L169
	} else {
		goto L175
	}
L175:
	;
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[595]))))
	v641 = base.B2i32(v614&int32(3) != int32(0))
	goto L168
L176:
	;
	v636 = F_isalnum(m, l1)
	mBase = m.M
	v639 = base.B2i32(v636 != int32(0))
	goto L169
L177:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v646 = v644 | v597
	goto L166
L178:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(-64))))
	if v750 == int32(0) {
		v962 = v747
		goto L206
	} else {
		goto L207
	}
L179:
	;
	v651 = *(*int32)(unsafe.Add(mBase, _consts[586]))
	switch v651 - int32(1) {
	case 0:
		goto L183
	case 1:
		goto L182
	case 2:
		goto L181
	default:
		goto L184
	}
L180:
	;
	if v742 == int32(0) {
		v747 = v646
		goto L178
	} else {
		goto L205
	}
L181:
	;
	if base.Ui32(int32(255)) < base.Ui32(l1) {
		v747 = v646
		goto L178
	} else {
		goto L203
	}
L182:
	;
	goto L202
L183:
	;
	v659 = *(*int32)(unsafe.Add(mBase, _consts[587]))
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v659)+16)))
	if (v660^int32(-1))&int32(1) != 0 {
		goto L186
	} else {
		goto L187
	}
L184:
	;
	v742 = base.B2i32(base.Ui32(l1-int32(48)) < base.Ui32(int32(10)))
	goto L180
L185:
	;
	v742 = v723
	goto L180
L186:
	;
	v723 = base.B2i32(base.Ui32(l1-int32(48)) < base.Ui32(int32(10)))
	goto L185
L187:
	;
	goto L188
L188:
	;
	if base.Ui32(int32(127)) < base.Ui32(l1) {
		goto L191
	} else {
		goto L192
	}
L189:
	;
	v723 = base.B2i32(v713&int32(255) == int32(9))
	goto L185
L190:
	;
	v707 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v685)+uint32(_consts[590]))))
	v713 = v707
	goto L189
L191:
	;
	v677 = int32(0)
	v678 = int32(3367)
	goto L194
L192:
	;
	goto L193
L193:
	;
	v706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1<<(uint(int32(1))%32))+uint32(_consts[591]))))
	v713 = v706
	goto L189
L194:
	;
	v683 = base.I32_div_s(v677+v678, int32(2))
	v685 = v683 * int32(12)
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_consts[588])))
	if base.Ui32(v688) < base.Ui32(l1) {
		goto L197
	} else {
		goto L198
	}
L195:
	;
	v713 = int32(0)
	goto L189
L196:
	;
	if v698 <= v699 {
		v677 = v698
		v678 = v699
		goto L194
	} else {
		goto L201
	}
L197:
	;
	v698 = v683 + int32(1)
	v699 = v678
	goto L196
L198:
	;
	goto L199
L199:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_consts[589])))
	if base.Ui32(v694) <= base.Ui32(l1) {
		goto L190
	} else {
		goto L200
	}
L200:
	;
	v698 = v677
	v699 = v683 - int32(1)
	goto L196
L201:
	;
	goto L195
L202:
	;
	v742 = base.B2i32(base.Ui32(l1-int32(48)) < base.Ui32(int32(10)))
	goto L180
L203:
	;
	goto L204
L204:
	;
	v742 = base.B2i32(base.B2i32(base.Ui32(l1-int32(48)) < base.Ui32(int32(10))) != int32(0))
	goto L180
L205:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v747 = v745 | v646
	goto L178
L206:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v963 == int32(0) {
		v1092 = v962
		goto L265
	} else {
		goto L266
	}
L207:
	;
	v754 = *(*int32)(unsafe.Add(mBase, _consts[586]))
	switch v754 - int32(1) {
	case 0:
		goto L211
	case 1:
		goto L210
	case 2:
		goto L209
	default:
		goto L212
	}
L208:
	;
	if v955 == int32(0) {
		v962 = v747
		goto L206
	} else {
		goto L264
	}
L209:
	;
	if base.Ui32(int32(255)) < base.Ui32(l1) {
		v962 = v747
		goto L206
	} else {
		goto L259
	}
L210:
	;
	if base.Ui32(l1) <= base.Ui32(int32(131071)) {
		goto L256
	} else {
		goto L257
	}
L211:
	;
	v767 = *(*int32)(unsafe.Add(mBase, _consts[587]))
	v768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v767)+16)))
	v771 = int32(1)
	if (v768^int32(-1))&v771 != 0 {
		goto L221
	} else {
		goto L222
	}
L212:
	;
	if base.Ui32(int32(127)) < base.Ui32(l1) {
		v962 = v747
		goto L206
	} else {
		goto L213
	}
L213:
	;
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[595]))))
	v955 = int32(base.Ui32(v761)>>(uint(int32(6))%32)) & int32(1)
	goto L208
L214:
	;
	v955 = v911
	goto L208
L215:
	;
	v911 = base.B2i32(v771<<(uint(v900)%32)&int32(1073217536) != int32(0))
	goto L214
L216:
	;
	v884 = int32(1)
	v885 = l1 << (uint(v884) % 32)
	v888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v885)+uint32(_consts[592]))))
	if v888&v884 != 0 {
		goto L252
	} else {
		goto L253
	}
L217:
	;
	v911 = base.B2i32(v771<<(uint(v878)%32)&int32(821559296) != int32(0))
	goto L214
L218:
	;
	v872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v849)+uint32(_consts[590]))))
	v878 = v872
	goto L217
L219:
	;
	v871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1<<(uint(int32(1))%32))+uint32(_consts[591]))))
	v878 = v871
	goto L217
L220:
	;
	v866 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v819)+uint32(_consts[590]))))
	v900 = v866
	goto L215
L221:
	;
	if base.Ui32(l1) < base.Ui32(int32(128)) {
		goto L216
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	if base.Ui32(l1) <= base.Ui32(int32(127)) {
		goto L219
	} else {
		goto L243
	}
L224:
	;
	v783 = int32(0)
	v784 = int32(1178)
	goto L225
L225:
	;
	v789 = base.I32_div_s(v783+v784, int32(2))
	v791 = v789 << (uint(int32(3)) % 32)
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v791)+uint32(_consts[596])))
	if base.Ui32(v794) < base.Ui32(l1) {
		goto L228
	} else {
		goto L229
	}
L226:
	;
	v811 = int32(0)
	v812 = int32(3367)
	goto L235
L227:
	;
	if v805 <= v806 {
		v783 = v805
		v784 = v806
		goto L225
	} else {
		goto L234
	}
L228:
	;
	v805 = v789 + int32(1)
	v806 = v784
	goto L227
L229:
	;
	goto L230
L230:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v791)+uint32(_consts[597])))
	if base.Ui32(v800) <= base.Ui32(l1) {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v911 = int32(0)
	goto L214
L232:
	;
	goto L233
L233:
	;
	v805 = v783
	v806 = v789 - int32(1)
	goto L227
L234:
	;
	goto L226
L235:
	;
	v817 = base.I32_div_s(v811+v812, int32(2))
	v819 = v817 * int32(12)
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v819)+uint32(_consts[588])))
	if base.Ui32(v822) < base.Ui32(l1) {
		goto L238
	} else {
		goto L239
	}
L236:
	;
	v900 = int32(0)
	goto L215
L237:
	;
	if v832 <= v833 {
		v811 = v832
		v812 = v833
		goto L235
	} else {
		goto L242
	}
L238:
	;
	v832 = v817 + int32(1)
	v833 = v812
	goto L237
L239:
	;
	goto L240
L240:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v819)+uint32(_consts[589])))
	if base.Ui32(v828) <= base.Ui32(l1) {
		goto L220
	} else {
		goto L241
	}
L241:
	;
	v832 = v811
	v833 = v817 - int32(1)
	goto L237
L242:
	;
	goto L236
L243:
	;
	v841 = int32(0)
	v842 = int32(3367)
	goto L244
L244:
	;
	v847 = base.I32_div_s(v841+v842, int32(2))
	v849 = v847 * int32(12)
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v849)+uint32(_consts[588])))
	if base.Ui32(v852) < base.Ui32(l1) {
		goto L247
	} else {
		goto L248
	}
L245:
	;
	v878 = int32(0)
	goto L217
L246:
	;
	if v862 <= v863 {
		v841 = v862
		v842 = v863
		goto L244
	} else {
		goto L251
	}
L247:
	;
	v862 = v847 + int32(1)
	v863 = v842
	goto L246
L248:
	;
	goto L249
L249:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v849)+uint32(_consts[589])))
	if base.Ui32(v858) <= base.Ui32(l1) {
		goto L218
	} else {
		goto L250
	}
L250:
	;
	v862 = v841
	v863 = v847 - int32(1)
	goto L246
L251:
	;
	goto L245
L252:
	;
	v911 = int32(0)
	goto L214
L253:
	;
	goto L254
L254:
	;
	v894 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v885)+uint32(_consts[591]))))
	v900 = v894
	goto L215
L255:
	;
	v955 = v938
	goto L208
L256:
	;
	v926 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l1)>>(uint(int32(8))%32)))+uint32(_consts[599]))))
	v932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l1)>>(uint(int32(3))%32))&int32(31)|v926<<(uint(int32(5))%32))+uint32(_consts[599]))))
	v938 = int32(base.Ui32(v932)>>(uint(l1&int32(7))%32)) & int32(1)
	goto L258
L257:
	;
	v938 = int32(0)
	goto L258
L258:
	;
	goto L255
L259:
	;
	if base.Ui32(l1-int32(33)) <= base.Ui32(int32(93)) {
		goto L261
	} else {
		goto L262
	}
L260:
	;
	v955 = base.B2i32(v952 != int32(0))
	goto L208
L261:
	;
	v949 = F_isalnum(m, l1)
	mBase = m.M
	v952 = base.B2i32(v949 == int32(0))
	goto L263
L262:
	;
	v952 = int32(0)
	goto L263
L263:
	;
	goto L260
L264:
	;
	v960 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(-64))))
	v962 = v960 | v747
	goto L206
L265:
	;
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v1093 == int32(0) {
		v1169 = v1092
		goto L310
	} else {
		goto L311
	}
L266:
	;
	v967 = *(*int32)(unsafe.Add(mBase, _consts[586]))
	switch v967 - int32(1) {
	case 0:
		goto L270
	case 1:
		goto L269
	case 2:
		goto L268
	default:
		goto L271
	}
L267:
	;
	if v1087 == int32(0) {
		v1092 = v962
		goto L265
	} else {
		goto L309
	}
L268:
	;
	if base.Ui32(int32(255)) < base.Ui32(l1) {
		v1092 = v962
		goto L265
	} else {
		goto L307
	}
L269:
	;
	if l1 == int32(0) {
		goto L288
	} else {
		goto L289
	}
L270:
	;
	if base.Ui32(int32(127)) < base.Ui32(l1) {
		goto L274
	} else {
		goto L275
	}
L271:
	;
	if base.Ui32(int32(127)) < base.Ui32(l1) {
		v1092 = v962
		goto L265
	} else {
		goto L272
	}
L272:
	;
	v974 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[595]))))
	v1087 = int32(base.Ui32(v974) >> (uint(int32(7)) % 32))
	goto L267
L273:
	;
	v1087 = v1025
	goto L267
L274:
	;
	v985 = int32(0)
	v986 = int32(10)
	goto L277
L275:
	;
	goto L276
L276:
	;
	v1015 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1<<(uint(int32(1))%32))+uint32(_consts[592]))))
	v1025 = int32(base.Ui32(v1015&int32(32)) >> (uint(int32(5)) % 32))
	goto L273
L277:
	;
	v991 = base.I32_div_s(v985+v986, int32(2))
	v993 = v991 << (uint(int32(3)) % 32)
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v993)+uint32(_consts[593])))
	if base.Ui32(v996) < base.Ui32(l1) {
		goto L280
	} else {
		goto L281
	}
L278:
	;
	v1025 = int32(0)
	goto L273
L279:
	;
	if v1007 <= v1008 {
		v985 = v1007
		v986 = v1008
		goto L277
	} else {
		goto L286
	}
L280:
	;
	v1007 = v991 + int32(1)
	v1008 = v986
	goto L279
L281:
	;
	goto L282
L282:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v993)+uint32(_consts[594])))
	if base.Ui32(v1002) <= base.Ui32(l1) {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v1025 = int32(1)
	goto L273
L284:
	;
	goto L285
L285:
	;
	v1007 = v985
	v1008 = v991 - int32(1)
	goto L279
L286:
	;
	goto L278
L287:
	;
	v1087 = v1072
	goto L267
L288:
	;
	v1072 = int32(0)
	goto L287
L289:
	;
	goto L290
L290:
	;
	if l1 != 0 {
		goto L292
	} else {
		goto L293
	}
L291:
	;
	v1072 = base.B2i32(v1065 != int32(0))
	goto L287
L292:
	;
	v1038 = int32(4072432)
	goto L295
L293:
	;
	goto L294
L294:
	;
	v1048 = int32(4072432)
	goto L304
L295:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v1038)))
	if v1040 != 0 {
		goto L297
	} else {
		goto L298
	}
L296:
	;
	if v1040 != 0 {
		goto L301
	} else {
		goto L302
	}
L297:
	;
	if l1 != v1040 {
		v1038 = v1038 + int32(4)
		goto L295
	} else {
		goto L300
	}
L298:
	;
	goto L299
L299:
	;
	goto L296
L300:
	;
	goto L299
L301:
	;
	v1046 = v1038
	goto L303
L302:
	;
	v1046 = int32(0)
	goto L303
L303:
	;
	v1065 = v1046
	goto L291
L304:
	;
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v1048)))
	if v1054 != 0 {
		v1048 = v1048 + int32(4)
		goto L304
	} else {
		goto L306
	}
L305:
	;
	v1055 = int32(4072432)
	v1065 = (v1048-v1055)&int32(-4) + v1055
	goto L291
L306:
	;
	goto L305
L307:
	;
	goto L308
L308:
	;
	v1087 = base.B2i32(base.B2i32(l1 == int32(32))|base.B2i32(base.Ui32(l1-int32(9)) < base.Ui32(int32(5))) != int32(0))
	goto L267
L309:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v1092 = v1090 | v962
	goto L265
L310:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v1170 == int32(0) {
		v1246 = v1169
		goto L335
	} else {
		goto L336
	}
L311:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, _consts[586]))
	switch v1097 - int32(1) {
	case 0:
		goto L315
	case 1:
		goto L314
	case 2:
		goto L313
	default:
		goto L316
	}
L312:
	;
	if v1164 == int32(0) {
		v1169 = v1092
		goto L310
	} else {
		goto L334
	}
L313:
	;
	if base.Ui32(int32(255)) < base.Ui32(l1) {
		v1169 = v1092
		goto L310
	} else {
		goto L332
	}
L314:
	;
	v1151 = F_towupper(m, l1)
	mBase = m.M
	goto L331
L315:
	;
	if base.Ui32(int32(127)) < base.Ui32(l1) {
		goto L318
	} else {
		goto L319
	}
L316:
	;
	v1164 = base.B2i32(base.Ui32(l1-int32(97)) < base.Ui32(int32(26)))
	goto L312
L317:
	;
	v1164 = v1147
	goto L312
L318:
	;
	v1112 = int32(0)
	v1113 = int32(689)
	goto L321
L319:
	;
	goto L320
L320:
	;
	v1147 = base.B2i32(base.Ui32(l1-int32(97)) < base.Ui32(int32(26)))
	goto L317
L321:
	;
	v1118 = base.I32_div_s(v1112+v1113, int32(2))
	v1120 = v1118 << (uint(int32(3)) % 32)
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v1120)+uint32(_consts[600])))
	if base.Ui32(v1123) < base.Ui32(l1) {
		goto L324
	} else {
		goto L325
	}
L322:
	;
	v1147 = int32(0)
	goto L317
L323:
	;
	if v1134 <= v1135 {
		v1112 = v1134
		v1113 = v1135
		goto L321
	} else {
		goto L330
	}
L324:
	;
	v1134 = v1118 + int32(1)
	v1135 = v1113
	goto L323
L325:
	;
	goto L326
L326:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1120)+uint32(_consts[601])))
	if base.Ui32(v1129) <= base.Ui32(l1) {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	v1147 = int32(1)
	goto L317
L328:
	;
	goto L329
L329:
	;
	v1134 = v1112
	v1135 = v1118 - int32(1)
	goto L323
L330:
	;
	goto L322
L331:
	;
	v1164 = base.B2i32(v1151 != l1)
	goto L312
L332:
	;
	goto L333
L333:
	;
	v1164 = base.B2i32(base.B2i32(base.Ui32(l1-int32(97)) < base.Ui32(int32(26))) != int32(0))
	goto L312
L334:
	;
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v1169 = v1167 | v1092
	goto L310
L335:
	;
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v1247 == int32(0) {
		v1386 = v1246
		goto L360
	} else {
		goto L361
	}
L336:
	;
	v1174 = *(*int32)(unsafe.Add(mBase, _consts[586]))
	switch v1174 - int32(1) {
	case 0:
		goto L340
	case 1:
		goto L339
	case 2:
		goto L338
	default:
		goto L341
	}
L337:
	;
	if v1241 == int32(0) {
		v1246 = v1169
		goto L335
	} else {
		goto L359
	}
L338:
	;
	if base.Ui32(int32(255)) < base.Ui32(l1) {
		v1246 = v1169
		goto L335
	} else {
		goto L357
	}
L339:
	;
	v1228 = F_towlower(m, l1)
	mBase = m.M
	goto L356
L340:
	;
	if base.Ui32(int32(127)) < base.Ui32(l1) {
		goto L343
	} else {
		goto L344
	}
L341:
	;
	v1241 = base.B2i32(base.Ui32(l1-int32(65)) < base.Ui32(int32(26)))
	goto L337
L342:
	;
	v1241 = v1224
	goto L337
L343:
	;
	v1189 = int32(0)
	v1190 = int32(655)
	goto L346
L344:
	;
	goto L345
L345:
	;
	v1224 = base.B2i32(base.Ui32(l1-int32(65)) < base.Ui32(int32(26)))
	goto L342
L346:
	;
	v1195 = base.I32_div_s(v1189+v1190, int32(2))
	v1197 = v1195 << (uint(int32(3)) % 32)
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v1197)+uint32(_consts[602])))
	if base.Ui32(v1200) < base.Ui32(l1) {
		goto L349
	} else {
		goto L350
	}
L347:
	;
	v1224 = int32(0)
	goto L342
L348:
	;
	if v1211 <= v1212 {
		v1189 = v1211
		v1190 = v1212
		goto L346
	} else {
		goto L355
	}
L349:
	;
	v1211 = v1195 + int32(1)
	v1212 = v1190
	goto L348
L350:
	;
	goto L351
L351:
	;
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v1197)+uint32(_consts[603])))
	if base.Ui32(v1206) <= base.Ui32(l1) {
		goto L352
	} else {
		goto L353
	}
L352:
	;
	v1224 = int32(1)
	goto L342
L353:
	;
	goto L354
L354:
	;
	v1211 = v1189
	v1212 = v1195 - int32(1)
	goto L348
L355:
	;
	goto L347
L356:
	;
	v1241 = base.B2i32(v1228 != l1)
	goto L337
L357:
	;
	goto L358
L358:
	;
	v1241 = base.B2i32(base.B2i32(base.Ui32(l1-int32(65)) < base.Ui32(int32(26))) != int32(0))
	goto L337
L359:
	;
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v1246 = v1244 | v1169
	goto L335
L360:
	;
	v1387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v1388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v1390 = int32(1)
	v1396 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1387+v1388*v51<<(uint(v1390)%32)+v1386<<(uint(v1390)%32)))))
	return v1396
L361:
	;
	v1251 = *(*int32)(unsafe.Add(mBase, _consts[586]))
	switch v1251 - int32(1) {
	case 0:
		goto L365
	case 1:
		goto L364
	case 2:
		goto L363
	default:
		goto L366
	}
L362:
	;
	if v1381 == int32(0) {
		v1386 = v1246
		goto L360
	} else {
		goto L400
	}
L363:
	;
	if base.Ui32(int32(255)) < base.Ui32(l1) {
		v1386 = v1246
		goto L360
	} else {
		goto L398
	}
L364:
	;
	v1365 = F_iswspace(m, l1)
	mBase = m.M
	if v1365 != 0 {
		goto L395
	} else {
		goto L396
	}
L365:
	;
	if base.Ui32(int32(128)) <= base.Ui32(l1) {
		goto L372
	} else {
		goto L373
	}
L366:
	;
	v1381 = base.B2i32(base.Ui32(l1-int32(33)) < base.Ui32(int32(94)))
	goto L362
L367:
	;
	v1381 = v1360
	goto L362
L368:
	;
	v1360 = int32(0)
	goto L367
L369:
	;
	v1360 = v1346
	goto L367
L370:
	;
	v1312 = int32(0)
	if int32(1)<<(uint(v1311)%32)&int32(294913) != 0 {
		v1346 = v1312
		goto L369
	} else {
		goto L385
	}
L371:
	;
	v1308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1274)+uint32(_consts[590]))))
	v1311 = v1308
	goto L370
L372:
	;
	v1266 = int32(3367)
	v1267 = int32(0)
	goto L375
L373:
	;
	goto L374
L374:
	;
	v1291 = int32(1)
	v1294 = l1 << (uint(v1291) % 32)
	v1297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1294)+uint32(_consts[591]))))
	if v1291<<(uint(v1297)%32)&int32(294913) != 0 {
		goto L368
	} else {
		goto L383
	}
L375:
	;
	v1272 = base.I32_div_s(v1266+v1267, int32(2))
	v1274 = v1272 * int32(12)
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v1274)+uint32(_consts[588])))
	if base.Ui32(v1277) < base.Ui32(l1) {
		goto L378
	} else {
		goto L379
	}
L376:
	;
	v1311 = int32(0)
	goto L370
L377:
	;
	if v1288 <= v1287 {
		v1266 = v1287
		v1267 = v1288
		goto L375
	} else {
		goto L382
	}
L378:
	;
	v1287 = v1266
	v1288 = v1272 + int32(1)
	goto L377
L379:
	;
	goto L380
L380:
	;
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v1274)+uint32(_consts[589])))
	if base.Ui32(v1283) <= base.Ui32(l1) {
		goto L371
	} else {
		goto L381
	}
L381:
	;
	v1287 = v1272 - int32(1)
	v1288 = v1267
	goto L377
L382:
	;
	goto L376
L383:
	;
	v1303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1294)+uint32(_consts[592]))))
	if v1303&int32(32) == int32(0) {
		v1346 = v1291
		goto L369
	} else {
		goto L384
	}
L384:
	;
	goto L368
L385:
	;
	v1319 = int32(10)
	v1320 = v1312
	goto L386
L386:
	;
	v1325 = base.I32_div_s(v1319+v1320, int32(2))
	v1327 = v1325 << (uint(int32(3)) % 32)
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v1327)+uint32(_consts[593])))
	if base.Ui32(v1330) < base.Ui32(l1) {
		goto L389
	} else {
		goto L390
	}
L387:
	;
	v1346 = int32(1)
	goto L369
L388:
	;
	if v1341 <= v1340 {
		v1319 = v1340
		v1320 = v1341
		goto L386
	} else {
		goto L393
	}
L389:
	;
	v1340 = v1319
	v1341 = v1325 + int32(1)
	goto L388
L390:
	;
	goto L391
L391:
	;
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v1327)+uint32(_consts[594])))
	if base.Ui32(v1336) <= base.Ui32(l1) {
		goto L368
	} else {
		goto L392
	}
L392:
	;
	v1340 = v1325 - int32(1)
	v1341 = v1320
	goto L388
L393:
	;
	goto L387
L394:
	;
	v1381 = v1369
	goto L362
L395:
	;
	v1369 = int32(0)
	goto L397
L396:
	;
	v1366 = F_iswprint(m, l1)
	mBase = m.M
	v1369 = base.B2i32(v1366 != int32(0))
	goto L397
L397:
	;
	goto L394
L398:
	;
	goto L399
L399:
	;
	v1381 = base.B2i32(base.B2i32(base.Ui32(l1-int32(33)) < base.Ui32(int32(94))) != int32(0))
	goto L362
L400:
	;
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1386 = v1384 | v1246
	goto L360
}
func F_pg_regerror(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
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
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v672 int32
	_ = v672
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	v8 = m.G0
	v10 = v8 - int32(144)
	m.G0 = v10
	v12 = int32(1601920)
	switch l0 - int32(101) {
	case 0:
		goto L5
	case 1:
		goto L4
	default:
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(144)
	return v720
L2:
	;
	v718 = F_strlen(m, v717)
	mBase = m.M
	v720 = v718 + int32(1)
	if base.Ui32(v720) < base.Ui32(int32(100)) {
		goto L227
	} else {
		goto L228
	}
L3:
	;
	v686 = v12
	goto L217
L4:
	;
	v611 = l1
	goto L192
L5:
	;
	v15 = int32(0)
	v16 = int32(506535)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, _consts[605])))
	if v20 == v15 {
		v39 = v19
		v40 = v20
		goto L8
	} else {
		goto L9
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v595
	v602 = F_pg_sprintf(m, v10+int32(48), int32(485141), v10+int32(16))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L189
	} else {
		goto L190
	}
L7:
	;
	if v40-v39 == int32(0) {
		v595 = v15
		goto L6
	} else {
		goto L15
	}
L8:
	;
	goto L7
L9:
	;
	if v19 != v20 {
		v39 = v19
		v40 = v20
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v24 = v16
	v25 = l1
	goto L11
L11:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if v29 == int32(0) {
		v39 = v28
		v40 = v29
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v39 = v28
	v40 = v29
	goto L8
L13:
	;
	v32 = int32(1)
	if v28 == v29 {
		v24 = v24 + v32
		v25 = v25 + v32
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v45 = int32(532660)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, _consts[606])))
	if v49 == int32(0) {
		v68 = v48
		v69 = v49
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v69-v68 == int32(0) {
		v595 = int32(1)
		goto L6
	} else {
		goto L24
	}
L17:
	;
	goto L16
L18:
	;
	if v48 != v49 {
		v68 = v48
		v69 = v49
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v53 = v45
	v54 = l1
	goto L20
L20:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+1)))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1)))
	if v58 == int32(0) {
		v68 = v57
		v69 = v58
		goto L17
	} else {
		goto L22
	}
L21:
	;
	v68 = v57
	v69 = v58
	goto L17
L22:
	;
	v61 = int32(1)
	if v57 == v58 {
		v53 = v53 + v61
		v54 = v54 + v61
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v74 = int32(518907)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, _consts[607])))
	if v78 == int32(0) {
		v97 = v77
		v98 = v78
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v98-v97 == int32(0) {
		v595 = int32(2)
		goto L6
	} else {
		goto L33
	}
L26:
	;
	goto L25
L27:
	;
	if v77 != v78 {
		v97 = v77
		v98 = v78
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v82 = v74
	v83 = l1
	goto L29
L29:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+1)))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+1)))
	if v87 == int32(0) {
		v97 = v86
		v98 = v87
		goto L26
	} else {
		goto L31
	}
L30:
	;
	v97 = v86
	v98 = v87
	goto L26
L31:
	;
	v90 = int32(1)
	if v86 == v87 {
		v82 = v82 + v90
		v83 = v83 + v90
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v103 = int32(535123)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, _consts[608])))
	if v107 == int32(0) {
		v126 = v106
		v127 = v107
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v127-v126 == int32(0) {
		v595 = int32(3)
		goto L6
	} else {
		goto L42
	}
L35:
	;
	goto L34
L36:
	;
	if v106 != v107 {
		v126 = v106
		v127 = v107
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v111 = v103
	v112 = l1
	goto L38
L38:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+1)))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+1)))
	if v116 == int32(0) {
		v126 = v115
		v127 = v116
		goto L35
	} else {
		goto L40
	}
L39:
	;
	v126 = v115
	v127 = v116
	goto L35
L40:
	;
	v119 = int32(1)
	if v115 == v116 {
		v111 = v111 + v119
		v112 = v112 + v119
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v132 = int32(536269)
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v136 = int32(*(*uint8)(unsafe.Add(mBase, _consts[609])))
	if v136 == int32(0) {
		v155 = v135
		v156 = v136
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v156-v155 == int32(0) {
		v595 = int32(4)
		goto L6
	} else {
		goto L51
	}
L44:
	;
	goto L43
L45:
	;
	if v135 != v136 {
		v155 = v135
		v156 = v136
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v140 = v132
	v141 = l1
	goto L47
L47:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+1)))
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+1)))
	if v145 == int32(0) {
		v155 = v144
		v156 = v145
		goto L44
	} else {
		goto L49
	}
L48:
	;
	v155 = v144
	v156 = v145
	goto L44
L49:
	;
	v148 = int32(1)
	if v144 == v145 {
		v140 = v140 + v148
		v141 = v141 + v148
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v161 = int32(536365)
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, _consts[610])))
	if v165 == int32(0) {
		v184 = v164
		v185 = v165
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if v185-v184 == int32(0) {
		v595 = int32(5)
		goto L6
	} else {
		goto L60
	}
L53:
	;
	goto L52
L54:
	;
	if v164 != v165 {
		v184 = v164
		v185 = v165
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v169 = v161
	v170 = l1
	goto L56
L56:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+1)))
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+1)))
	if v174 == int32(0) {
		v184 = v173
		v185 = v174
		goto L53
	} else {
		goto L58
	}
L57:
	;
	v184 = v173
	v185 = v174
	goto L53
L58:
	;
	v177 = int32(1)
	if v173 == v174 {
		v169 = v169 + v177
		v170 = v170 + v177
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v190 = int32(533426)
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v194 = int32(*(*uint8)(unsafe.Add(mBase, _consts[611])))
	if v194 == int32(0) {
		v213 = v193
		v214 = v194
		goto L62
	} else {
		goto L63
	}
L61:
	;
	if v214-v213 == int32(0) {
		v595 = int32(6)
		goto L6
	} else {
		goto L69
	}
L62:
	;
	goto L61
L63:
	;
	if v193 != v194 {
		v213 = v193
		v214 = v194
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v198 = v190
	v199 = l1
	goto L65
L65:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+1)))
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+1)))
	if v203 == int32(0) {
		v213 = v202
		v214 = v203
		goto L62
	} else {
		goto L67
	}
L66:
	;
	v213 = v202
	v214 = v203
	goto L62
L67:
	;
	v206 = int32(1)
	if v202 == v203 {
		v198 = v198 + v206
		v199 = v199 + v206
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v219 = int32(530754)
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v223 = int32(*(*uint8)(unsafe.Add(mBase, _consts[612])))
	if v223 == int32(0) {
		v242 = v222
		v243 = v223
		goto L71
	} else {
		goto L72
	}
L70:
	;
	if v243-v242 == int32(0) {
		v595 = int32(7)
		goto L6
	} else {
		goto L78
	}
L71:
	;
	goto L70
L72:
	;
	if v222 != v223 {
		v242 = v222
		v243 = v223
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v227 = v219
	v228 = l1
	goto L74
L74:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+1)))
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+1)))
	if v232 == int32(0) {
		v242 = v231
		v243 = v232
		goto L71
	} else {
		goto L76
	}
L75:
	;
	v242 = v231
	v243 = v232
	goto L71
L76:
	;
	v235 = int32(1)
	if v231 == v232 {
		v227 = v227 + v235
		v228 = v228 + v235
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v248 = int32(527378)
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v252 = int32(*(*uint8)(unsafe.Add(mBase, _consts[613])))
	if v252 == int32(0) {
		v271 = v251
		v272 = v252
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v272-v271 == int32(0) {
		v595 = int32(8)
		goto L6
	} else {
		goto L87
	}
L80:
	;
	goto L79
L81:
	;
	if v251 != v252 {
		v271 = v251
		v272 = v252
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v256 = v248
	v257 = l1
	goto L83
L83:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257)+1)))
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+1)))
	if v261 == int32(0) {
		v271 = v260
		v272 = v261
		goto L80
	} else {
		goto L85
	}
L84:
	;
	v271 = v260
	v272 = v261
	goto L80
L85:
	;
	v264 = int32(1)
	if v260 == v261 {
		v256 = v256 + v264
		v257 = v257 + v264
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v277 = int32(538902)
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v281 = int32(*(*uint8)(unsafe.Add(mBase, _consts[614])))
	if v281 == int32(0) {
		v300 = v280
		v301 = v281
		goto L89
	} else {
		goto L90
	}
L88:
	;
	if v301-v300 == int32(0) {
		v595 = int32(9)
		goto L6
	} else {
		goto L96
	}
L89:
	;
	goto L88
L90:
	;
	if v280 != v281 {
		v300 = v280
		v301 = v281
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v285 = v277
	v286 = l1
	goto L92
L92:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+1)))
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285)+1)))
	if v290 == int32(0) {
		v300 = v289
		v301 = v290
		goto L89
	} else {
		goto L94
	}
L93:
	;
	v300 = v289
	v301 = v290
	goto L89
L94:
	;
	v293 = int32(1)
	if v289 == v290 {
		v285 = v285 + v293
		v286 = v286 + v293
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	v306 = int32(522686)
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v310 = int32(*(*uint8)(unsafe.Add(mBase, _consts[615])))
	if v310 == int32(0) {
		v329 = v309
		v330 = v310
		goto L98
	} else {
		goto L99
	}
L97:
	;
	if v330-v329 == int32(0) {
		v595 = int32(10)
		goto L6
	} else {
		goto L105
	}
L98:
	;
	goto L97
L99:
	;
	if v309 != v310 {
		v329 = v309
		v330 = v310
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v314 = v306
	v315 = l1
	goto L101
L101:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+1)))
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+1)))
	if v319 == int32(0) {
		v329 = v318
		v330 = v319
		goto L98
	} else {
		goto L103
	}
L102:
	;
	v329 = v318
	v330 = v319
	goto L98
L103:
	;
	v322 = int32(1)
	if v318 == v319 {
		v314 = v314 + v322
		v315 = v315 + v322
		goto L101
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	v335 = int32(538015)
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v339 = int32(*(*uint8)(unsafe.Add(mBase, _consts[616])))
	if v339 == int32(0) {
		v358 = v338
		v359 = v339
		goto L107
	} else {
		goto L108
	}
L106:
	;
	if v359-v358 == int32(0) {
		v595 = int32(11)
		goto L6
	} else {
		goto L114
	}
L107:
	;
	goto L106
L108:
	;
	if v338 != v339 {
		v358 = v338
		v359 = v339
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v343 = v335
	v344 = l1
	goto L110
L110:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344)+1)))
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+1)))
	if v348 == int32(0) {
		v358 = v347
		v359 = v348
		goto L107
	} else {
		goto L112
	}
L111:
	;
	v358 = v347
	v359 = v348
	goto L107
L112:
	;
	v351 = int32(1)
	if v347 == v348 {
		v343 = v343 + v351
		v344 = v344 + v351
		goto L110
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	v364 = int32(538913)
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v368 = int32(*(*uint8)(unsafe.Add(mBase, _consts[617])))
	if v368 == int32(0) {
		v387 = v367
		v388 = v368
		goto L116
	} else {
		goto L117
	}
L115:
	;
	if v388-v387 == int32(0) {
		v595 = int32(12)
		goto L6
	} else {
		goto L123
	}
L116:
	;
	goto L115
L117:
	;
	if v367 != v368 {
		v387 = v367
		v388 = v368
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v372 = v364
	v373 = l1
	goto L119
L119:
	;
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+1)))
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372)+1)))
	if v377 == int32(0) {
		v387 = v376
		v388 = v377
		goto L116
	} else {
		goto L121
	}
L120:
	;
	v387 = v376
	v388 = v377
	goto L116
L121:
	;
	v380 = int32(1)
	if v376 == v377 {
		v372 = v372 + v380
		v373 = v373 + v380
		goto L119
	} else {
		goto L122
	}
L122:
	;
	goto L120
L123:
	;
	v393 = int32(514319)
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v397 = int32(*(*uint8)(unsafe.Add(mBase, _consts[618])))
	if v397 == int32(0) {
		v416 = v396
		v417 = v397
		goto L125
	} else {
		goto L126
	}
L124:
	;
	if v417-v416 == int32(0) {
		v595 = int32(13)
		goto L6
	} else {
		goto L132
	}
L125:
	;
	goto L124
L126:
	;
	if v396 != v397 {
		v416 = v396
		v417 = v397
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v401 = v393
	v402 = l1
	goto L128
L128:
	;
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402)+1)))
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401)+1)))
	if v406 == int32(0) {
		v416 = v405
		v417 = v406
		goto L125
	} else {
		goto L130
	}
L129:
	;
	v416 = v405
	v417 = v406
	goto L125
L130:
	;
	v409 = int32(1)
	if v405 == v406 {
		v401 = v401 + v409
		v402 = v402 + v409
		goto L128
	} else {
		goto L131
	}
L131:
	;
	goto L129
L132:
	;
	v422 = int32(514008)
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v426 = int32(*(*uint8)(unsafe.Add(mBase, _consts[619])))
	if v426 == int32(0) {
		v445 = v425
		v446 = v426
		goto L134
	} else {
		goto L135
	}
L133:
	;
	if v446-v445 == int32(0) {
		v595 = int32(15)
		goto L6
	} else {
		goto L141
	}
L134:
	;
	goto L133
L135:
	;
	if v425 != v426 {
		v445 = v425
		v446 = v426
		goto L134
	} else {
		goto L136
	}
L136:
	;
	v430 = v422
	v431 = l1
	goto L137
L137:
	;
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v431)+1)))
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430)+1)))
	if v435 == int32(0) {
		v445 = v434
		v446 = v435
		goto L134
	} else {
		goto L139
	}
L138:
	;
	v445 = v434
	v446 = v435
	goto L134
L139:
	;
	v438 = int32(1)
	if v434 == v435 {
		v430 = v430 + v438
		v431 = v431 + v438
		goto L137
	} else {
		goto L140
	}
L140:
	;
	goto L138
L141:
	;
	v451 = int32(532721)
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v455 = int32(*(*uint8)(unsafe.Add(mBase, _consts[620])))
	if v455 == int32(0) {
		v474 = v454
		v475 = v455
		goto L143
	} else {
		goto L144
	}
L142:
	;
	if v475-v474 == int32(0) {
		v595 = int32(16)
		goto L6
	} else {
		goto L150
	}
L143:
	;
	goto L142
L144:
	;
	if v454 != v455 {
		v474 = v454
		v475 = v455
		goto L143
	} else {
		goto L145
	}
L145:
	;
	v459 = v451
	v460 = l1
	goto L146
L146:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460)+1)))
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459)+1)))
	if v464 == int32(0) {
		v474 = v463
		v475 = v464
		goto L143
	} else {
		goto L148
	}
L147:
	;
	v474 = v463
	v475 = v464
	goto L143
L148:
	;
	v467 = int32(1)
	if v463 == v464 {
		v459 = v459 + v467
		v460 = v460 + v467
		goto L146
	} else {
		goto L149
	}
L149:
	;
	goto L147
L150:
	;
	v480 = int32(540196)
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v484 = int32(*(*uint8)(unsafe.Add(mBase, _consts[621])))
	if v484 == int32(0) {
		v503 = v483
		v504 = v484
		goto L152
	} else {
		goto L153
	}
L151:
	;
	if v504-v503 == int32(0) {
		v595 = int32(17)
		goto L6
	} else {
		goto L159
	}
L152:
	;
	goto L151
L153:
	;
	if v483 != v484 {
		v503 = v483
		v504 = v484
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v488 = v480
	v489 = l1
	goto L155
L155:
	;
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+1)))
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+1)))
	if v493 == int32(0) {
		v503 = v492
		v504 = v493
		goto L152
	} else {
		goto L157
	}
L156:
	;
	v503 = v492
	v504 = v493
	goto L152
L157:
	;
	v496 = int32(1)
	if v492 == v493 {
		v488 = v488 + v496
		v489 = v489 + v496
		goto L155
	} else {
		goto L158
	}
L158:
	;
	goto L156
L159:
	;
	v509 = int32(514330)
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v513 = int32(*(*uint8)(unsafe.Add(mBase, _consts[622])))
	if v513 == int32(0) {
		v532 = v512
		v533 = v513
		goto L161
	} else {
		goto L162
	}
L160:
	;
	if v533-v532 == int32(0) {
		v595 = int32(18)
		goto L6
	} else {
		goto L168
	}
L161:
	;
	goto L160
L162:
	;
	if v512 != v513 {
		v532 = v512
		v533 = v513
		goto L161
	} else {
		goto L163
	}
L163:
	;
	v517 = v509
	v518 = l1
	goto L164
L164:
	;
	v521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+1)))
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517)+1)))
	if v522 == int32(0) {
		v532 = v521
		v533 = v522
		goto L161
	} else {
		goto L166
	}
L165:
	;
	v532 = v521
	v533 = v522
	goto L161
L166:
	;
	v525 = int32(1)
	if v521 == v522 {
		v517 = v517 + v525
		v518 = v518 + v525
		goto L164
	} else {
		goto L167
	}
L167:
	;
	goto L165
L168:
	;
	v538 = int32(533336)
	v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v542 = int32(*(*uint8)(unsafe.Add(mBase, _consts[623])))
	if v542 == int32(0) {
		v561 = v541
		v562 = v542
		goto L170
	} else {
		goto L171
	}
L169:
	;
	if v562-v561 == int32(0) {
		v595 = int32(19)
		goto L6
	} else {
		goto L177
	}
L170:
	;
	goto L169
L171:
	;
	if v541 != v542 {
		v561 = v541
		v562 = v542
		goto L170
	} else {
		goto L172
	}
L172:
	;
	v546 = v538
	v547 = l1
	goto L173
L173:
	;
	v550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547)+1)))
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546)+1)))
	if v551 == int32(0) {
		v561 = v550
		v562 = v551
		goto L170
	} else {
		goto L175
	}
L174:
	;
	v561 = v550
	v562 = v551
	goto L170
L175:
	;
	v554 = int32(1)
	if v550 == v551 {
		v546 = v546 + v554
		v547 = v547 + v554
		goto L173
	} else {
		goto L176
	}
L176:
	;
	goto L174
L177:
	;
	v568 = int32(519501)
	v571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v572 = int32(*(*uint8)(unsafe.Add(mBase, _consts[624])))
	if v572 == int32(0) {
		v591 = v571
		v592 = v572
		goto L179
	} else {
		goto L180
	}
L178:
	;
	if v592-v591 != 0 {
		goto L186
	} else {
		goto L187
	}
L179:
	;
	goto L178
L180:
	;
	if v571 != v572 {
		v591 = v571
		v592 = v572
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v576 = v568
	v577 = l1
	goto L182
L182:
	;
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577)+1)))
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+1)))
	if v581 == int32(0) {
		v591 = v580
		v592 = v581
		goto L179
	} else {
		goto L184
	}
L183:
	;
	v591 = v580
	v592 = v581
	goto L179
L184:
	;
	v584 = int32(1)
	if v580 == v581 {
		v576 = v576 + v584
		v577 = v577 + v584
		goto L182
	} else {
		goto L185
	}
L185:
	;
	goto L183
L186:
	;
	v594 = int32(-1)
	goto L188
L187:
	;
	v594 = int32(20)
	goto L188
L188:
	;
	v595 = v594
	goto L6
L189:
	;
	return int32(0)
L190:
	;
	v717 = v10 + int32(48)
	goto L2
L191:
	;
	v659 = v12
	goto L207
L192:
	;
	v616 = v611 + int32(1)
	v617 = int32(*(*int8)(unsafe.Add(mBase, uint32(v611))))
	v618 = F___isspace(m, v617)
	mBase = m.M
	if v618 != 0 {
		v611 = v616
		goto L192
	} else {
		goto L194
	}
L193:
	;
	v619 = int32(1)
	switch v617&int32(255) - int32(43) {
	case 0:
		v625 = v619
		goto L196
	default:
		v627 = v617
		v628 = v611
		v629 = v619
		goto L195
	case 2:
		goto L197
	}
L194:
	;
	goto L193
L195:
	;
	v630 = int32(0)
	v632 = v627 - int32(48)
	if base.Ui32(v632) <= base.Ui32(int32(9)) {
		goto L198
	} else {
		goto L199
	}
L196:
	;
	v626 = int32(*(*int8)(unsafe.Add(mBase, uint32(v616))))
	v627 = v626
	v628 = v616
	v629 = v625
	goto L195
L197:
	;
	v625 = int32(0)
	goto L196
L198:
	;
	v635 = v630
	v636 = v632
	v637 = v628
	goto L201
L199:
	;
	v649 = v630
	goto L200
L200:
	;
	if v629 != 0 {
		goto L204
	} else {
		goto L205
	}
L201:
	;
	v639 = int32(10)
	v641 = v635*v639 - v636
	v642 = int32(*(*int8)(unsafe.Add(mBase, uint32(v637)+1)))
	v646 = v642 - int32(48)
	if base.Ui32(v646) < base.Ui32(v639) {
		v635 = v641
		v636 = v646
		v637 = v637 + int32(1)
		goto L201
	} else {
		goto L203
	}
L202:
	;
	v649 = v641
	goto L200
L203:
	;
	goto L202
L204:
	;
	v655 = int32(0) - v649
	goto L206
L205:
	;
	v655 = v649
	goto L206
L206:
	;
	goto L191
L207:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v659)))
	if int32(0) <= v663 {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	if int32(0) <= v663 {
		goto L213
	} else {
		goto L214
	}
L209:
	;
	if v655 != v663 {
		v659 = v659 + int32(12)
		goto L207
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	goto L208
L212:
	;
	goto L211
L213:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v659)+4))
	v717 = v672
	goto L2
L214:
	;
	goto L215
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v655
	v679 = F_pg_sprintf(m, v10+int32(48), int32(38437), v10+int32(32))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L189
	} else {
		goto L216
	}
L216:
	;
	v717 = v10 + int32(48)
	goto L2
L217:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v686)))
	v691 = int32(0)
	v692 = base.B2i32(v690 < v691)
	if v692 == v691 {
		goto L219
	} else {
		goto L220
	}
L218:
	;
	if v692 == int32(0) {
		goto L223
	} else {
		goto L224
	}
L219:
	;
	if l0 != v690 {
		v686 = v686 + int32(12)
		goto L217
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	goto L218
L222:
	;
	goto L221
L223:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v686)+8))
	v717 = v701
	goto L2
L224:
	;
	goto L225
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
	v706 = F_pg_sprintf(m, v10+int32(48), int32(1602176), v10)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L189
	} else {
		goto L226
	}
L226:
	;
	v717 = v10 + int32(48)
	goto L2
L227:
	;
	if (v717^l1)&int32(3) != 0 {
		goto L233
	} else {
		goto L234
	}
L228:
	;
	goto L229
L229:
	;
	goto L252
L230:
	;
	goto L1
L231:
	;
	goto L230
L232:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v777))) = uint8(v776)
	if v776&int32(255) == int32(0) {
		goto L231
	} else {
		goto L247
	}
L233:
	;
	v728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v717))))
	v775 = v717
	v776 = v728
	v777 = l1
	goto L232
L234:
	;
	goto L235
L235:
	;
	if v717&int32(3) != 0 {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v732 = v717
	v734 = l1
	goto L239
L237:
	;
	v746 = v717
	v748 = l1
	goto L238
L238:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v746)))
	v753 = int32(-2139062144)
	if (int32(16843008)-v750|v750)&v753 != v753 {
		v775 = v746
		v776 = v750
		v777 = v748
		goto L232
	} else {
		goto L243
	}
L239:
	;
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v732))))
	*(*uint8)(unsafe.Add(mBase, uint32(v734))) = uint8(v735)
	if v735 == int32(0) {
		goto L231
	} else {
		goto L241
	}
L240:
	;
	v746 = v742
	v748 = v740
	goto L238
L241:
	;
	v739 = int32(1)
	v740 = v734 + v739
	v742 = v732 + v739
	if v742&int32(3) != 0 {
		v732 = v742
		v734 = v740
		goto L239
	} else {
		goto L242
	}
L242:
	;
	goto L240
L243:
	;
	v758 = v746
	v759 = v750
	v760 = v748
	goto L244
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v760))) = v759
	v762 = int32(4)
	v763 = v760 + v762
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v758)+4))
	v766 = v758 + v762
	v770 = int32(-2139062144)
	if (v764|(int32(16843008)-v764))&v770 == v770 {
		v758 = v766
		v759 = v764
		v760 = v763
		goto L244
	} else {
		goto L246
	}
L245:
	;
	v775 = v766
	v776 = v764
	v777 = v763
	goto L232
L246:
	;
	goto L245
L247:
	;
	v784 = v775
	v786 = v777
	goto L248
L248:
	;
	v787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v784)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v786)+1)) = uint8(v787)
	v789 = int32(1)
	if v787 != 0 {
		v784 = v784 + v789
		v786 = v786 + v789
		goto L248
	} else {
		goto L250
	}
L249:
	;
	goto L231
L250:
	;
	goto L249
L251:
	;
	v800 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v798)+99)) = uint8(v800)
	goto L1
L252:
	;
	v798 = F__emscripten_memcpy_bulkmem(m, l1, v717, int32(99))
	mBase = m.M
	goto L254
L254:
	;
	goto L251
}
func F_pg_regfree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	if l0 != 0 {
		v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
		m.T0[v3].(func(*base.Module, int32))(m, l0)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_pg_sequence_parameters(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int64
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int64
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, _consts[276]))
	v13 = F_pg_class_aclcheck(m, v9, v11, int64(262))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 == int32(0) {
			v22 = F_get_call_result_type(m, l0, int32(0), v5+int32(-4))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				if v22 != int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(365102), int32(0))
						mBase = m.M
						v105 = m.ExcPending
						if v105 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(496302), int32(1757), int32(131537))
							mBase = m.M
							v110 = m.ExcPending
							if v110 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v26 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v7)+27)) = v26
					*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v26
					v31 = F_SearchSysCache1(m, int32(61), v9)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						if v31 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v114 = m.ExcPending
							if v114 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
								F_errmsg_internal(m, int32(52624), v7)
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(496302), int32(1763), int32(131537))
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
							v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
							v37 = v35 + v36
							v38 = *(*int64)(unsafe.Add(mBase, uint32(v37)+8))
							v39 = F_Int64GetDatum(m, v38)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v39
								v42 = *(*int64)(unsafe.Add(mBase, uint32(v37)+32))
								v43 = F_Int64GetDatum(m, v42)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v43
									v46 = *(*int64)(unsafe.Add(mBase, uint32(v37)+24))
									v47 = F_Int64GetDatum(m, v46)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v47
										v50 = *(*int64)(unsafe.Add(mBase, uint32(v37)+16))
										v51 = F_Int64GetDatum(m, v50)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7)+44)) = v51
											v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+48)))
											*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v54
											v56 = *(*int64)(unsafe.Add(mBase, uint32(v37)+40))
											v57 = F_Int64GetDatum(m, v56)
											mBase = m.M
											v58 = m.ExcPending
											if v58 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = v57
												v60 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v7)+56)) = v60
												F_ReleaseCatCache(m, v31)
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
													return int32(0)
												} else {
													v64 = *(*int32)(unsafe.Add(mBase, uint32(v7)+60))
													v69 = F_heap_form_tuple(m, v64, v5+int32(-32), v5+int32(-40))
													mBase = m.M
													v70 = m.ExcPending
													if v70 != 0 {
														return int32(0)
													} else {
														v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
														v72 = F_HeapTupleHeaderGetDatum(m, v71)
														mBase = m.M
														v73 = m.ExcPending
														if v73 != 0 {
															return int32(0)
														} else {
															m.G0 = v7 - int32(-64)
															return v72
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int32(0)
				} else {
					v85 = F_get_rel_name(m, v9)
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v85
						F_errmsg(m, int32(194843), v5+int32(-48))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(496302), int32(1754), int32(131537))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_server_to_any(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	if l1 <= int32(0) {
		v39 = l0
		return v39
	} else {
		if l2 == int32(0) {
			v39 = l0
			return v39
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, _consts[358]))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
			if l2 == v11 {
				v39 = l0
				return v39
			} else {
				if v11 == int32(0) {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l2*int32(28))+uint32(_consts[1222])))
					v20 = m.T0[v19].(func(*base.Module, int32, int32) int32)(m, l0, l1)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						if l1 == v20 {
							v39 = l0
							return v39
						} else {
							F_report_invalid_encoding(m, l2, l0+v20, l1-v20)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, _consts[357]))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
					if v31 == l2 {
						v34 = F_perform_default_encoding_conversion(m, l0, l1, int32(0))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							return v34
						}
					} else {
						v37 = F_pg_do_encoding_conversion(m, l0, l1, v11, l2)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = v37
							return v39
						}
					}
				}
			}
		}
	}
}
func F_pg_size_pretty_numeric(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var __phi23 int32
	_ = __phi23
	var v24 int32
	_ = v24
	var __phi24 int32
	_ = __phi24
	var v26 int32
	_ = v26
	var __phi26 int32
	_ = __phi26
	var v29 int32
	_ = v29
	var __phi29 int32
	_ = __phi29
	var v30 int32
	_ = v30
	var __phi30 int32
	_ = __phi30
	var v32 int64
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = F_pg_detoast_datum(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	__phi23 = v19
	__phi24 = int32(1644144)
	__phi26 = int32(541476)
	__phi29 = int32(159002)
	__phi30 = int32(1644156)
	v23 = __phi23
	v24 = __phi24
	v26 = __phi26
	v29 = __phi29
	v30 = __phi30
	goto L3
L3:
	;
	v32 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+4)))
	v34 = int32(0)
	v37 = F_DirectFunctionCall1Coll(m, int32(1292), v34, v23)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+8)))
	if v74 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L5:
	;
	goto L4
L6:
	;
	v39 = F_pg_detoast_datum(m, v37)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v41 = F_int64_to_numeric(m, v32)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v43 = F_DirectFunctionCall2Coll(m, int32(1291), v34, v39, v41)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v43 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v65 = v23
	v67 = v24
	v69 = v29
	goto L5
L11:
	;
	goto L12
L12:
	;
	v46 = v30 + int32(12)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v51 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)))
	v52 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v24)+21)))
	v53 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v24)+9)))
	v54 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v24)+20)))
	v59 = F_int64_to_numeric(m, int64(1)<<(uint(v51+(v52-(v53+v54)))%64))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v61 = F_DirectFunctionCall2Coll(m, int32(1293), int32(0), v23, v59)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v63 = F_pg_detoast_datum(m, v61)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v47 != 0 {
		__phi23 = v63
		__phi24 = v30
		__phi26 = v47
		__phi29 = v26
		__phi30 = v46
		v23 = __phi23
		v24 = __phi24
		v26 = __phi26
		v29 = __phi29
		v30 = __phi30
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v65 = v63
	v67 = v30
	v69 = v26
	goto L5
L17:
	;
	v78 = F_int64_to_numeric(m, int64(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	v105 = v65
	goto L19
L19:
	;
	v106 = F_DirectFunctionCall1Coll(m, int32(618), int32(0), v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L30
	}
L20:
	;
	v81 = F_int64_to_numeric(m, int64(1))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v84 = F_int64_to_numeric(m, int64(2))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v87 = int32(0)
	v92 = F_DirectFunctionCall2Coll(m, int32(1295), v87, v65, v78)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if v92 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v94 = int32(1294)
	goto L26
L25:
	;
	v94 = int32(18)
	goto L26
L26:
	;
	v96 = F_DirectFunctionCall2Coll(m, v94, int32(0), v65, v81)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v98 = F_DirectFunctionCall2Coll(m, int32(1293), v87, v96, v84)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v100 = F_pg_detoast_datum(m, v98)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v105 = v100
	goto L19
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v106
	v111 = F_psprintf(m, int32(179794), v12)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v113 = F_cstring_to_text(m, v111)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	m.G0 = v12 + int32(16)
	return v113
}
func F_pg_strftime(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[155]))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v5
	v16 = l0 + l1
	v19 = F__fmt(m, l2, l3, l0, v16, v10+int32(12))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		if v19 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _consts[155])) = int32(61)
			v37 = v5
		} else {
			if v19 == v16 {
				*(*int32)(unsafe.Add(mBase, _consts[155])) = int32(68)
				v37 = v5
			} else {
				v32 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, _consts[155])) = v13
				v37 = v19 - l0
			}
		}
		m.G0 = v10 + int32(16)
		return v37
	}
}
func F_pg_strtoint32(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_pg_strtoint32_safe(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_pg_strtoint32_safe(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v76 int64
	_ = v76
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v241 int32
	_ = v241
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v271 int32
	_ = v271
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v323 int64
	_ = v323
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v18 = base.B2i32(v16 == int32(45))
	v19 = l0 + v18
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v24 = (v20 - int32(48)) & int32(255)
	if base.Ui32(int32(10)) <= base.Ui32(v24) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v13 + int32(32)
	return v402
L2:
	;
	F_errsave_finish(m, l1, int32(490743), v396, int32(407181))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L81
	} else {
		goto L90
	}
L3:
	;
	v369 = int32(0)
	v370 = F_errsave_start(m, l1)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L81
	} else {
		goto L86
	}
L4:
	;
	v342 = int32(0)
	v343 = F_errsave_start(m, l1)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L81
	} else {
		goto L82
	}
L5:
	;
	v96 = l0
	v99 = v16
	goto L22
L6:
	;
	v28 = v19 + int32(1)
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	v31 = v29 - int32(48)
	if base.Ui32(v31&int32(255)) <= base.Ui32(int32(9)) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v38 = v24
	v39 = v31
	v40 = v28
	goto L10
L8:
	;
	v64 = v24
	v67 = v29
	goto L9
L9:
	;
	if v67 != 0 {
		goto L5
	} else {
		goto L14
	}
L10:
	;
	if base.Ui32(int32(214748364)) < base.Ui32(v38) {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	v64 = v52
	v67 = v55
	goto L9
L12:
	;
	v48 = int32(10)
	v50 = int32(255)
	v52 = v38*v48 + v39&v50
	v54 = v40 + int32(1)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	v57 = v55 - int32(48)
	if base.Ui32(v57&v50) < base.Ui32(v48) {
		v38 = v52
		v39 = v57
		v40 = v54
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	if v16 == int32(45) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v76 = int64(0) - base.I64_extend_i32_u(v64)
	if v76 != base.I64_extend32_s(v76) {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if v64 < int32(0) {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	v402 = base.I32_wrap_i64(v76)
	goto L1
L19:
	;
	v402 = v64
	goto L1
L20:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v118 != int32(48) {
		goto L28
	} else {
		goto L29
	}
L21:
	;
	v116 = v96 + int32(1)
	v117 = v18
	goto L20
L22:
	;
	if base.Ui32(v99-int32(9)) < base.Ui32(int32(5)) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v111 = int32(1)
	v116 = v96 + v111
	v117 = v111
	goto L20
L24:
	;
	goto L23
L25:
	;
	v109 = v96 + int32(1)
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	v96 = v109
	v99 = v110
	goto L22
L26:
	;
	switch v99 - int32(32) {
	case 0:
		goto L25
	default:
		v116 = v96
		v117 = v18
		goto L20
	case 11:
		goto L21
	case 13:
		goto L24
	}
L27:
	;
	if v297 == v298 {
		goto L3
	} else {
		goto L69
	}
L28:
	;
	v260 = int32(0)
	v261 = v116
	v263 = v118
	goto L60
L29:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
	switch v121 - int32(66) {
	case 0, 32:
		goto L30
	default:
		goto L28
	case 13, 45:
		goto L31
	case 22, 54:
		goto L32
	}
L30:
	;
	v219 = v116 + int32(2)
	v222 = int32(0)
	v223 = v219
	goto L52
L31:
	;
	v179 = v116 + int32(2)
	v182 = int32(0)
	v183 = v179
	goto L44
L32:
	;
	v126 = v116 + int32(2)
	v129 = int32(0)
	v130 = v126
	goto L33
L33:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	goto L35
L34:
	;
	goto L3
L35:
	;
	if base.B2i32(base.Ui32(v137-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v137|int32(32)-int32(97)) < base.Ui32(int32(6))) != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if base.Ui32(int32(134217728)) < base.Ui32(v129) {
		goto L4
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	if v137 != int32(95) {
		v296 = v129
		v297 = v130
		v298 = v126
		v299 = v137
		goto L27
	} else {
		goto L40
	}
L39:
	;
	v153 = int32(*(*int8)(unsafe.Add(mBase, uint32(v137)+uint32(_consts[1090]))))
	v129 = v153 + v129<<(uint(int32(4))%32)
	v130 = v130 + int32(1)
	goto L33
L40:
	;
	v162 = v130 + int32(1)
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	if v163 == int32(0) {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	goto L42
L42:
	;
	if base.B2i32(base.Ui32(v163-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v163|int32(32)-int32(97)) < base.Ui32(int32(6))) != 0 {
		v130 = v162
		goto L33
	} else {
		goto L43
	}
L43:
	;
	goto L34
L44:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	if v190&int32(248) == int32(48) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L3
L46:
	;
	if base.Ui32(int32(268435456)) < base.Ui32(v182) {
		goto L4
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if v190 != int32(95) {
		v296 = v182
		v297 = v183
		v298 = v179
		v299 = v190
		goto L27
	} else {
		goto L50
	}
L49:
	;
	v182 = (v190-int32(48))&int32(255) | v182<<(uint(int32(3))%32)
	v183 = v183 + int32(1)
	goto L44
L50:
	;
	v209 = v183 + int32(1)
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
	if base.Ui32(int32(248)) <= base.Ui32((v210-int32(56))&int32(255)) {
		v183 = v209
		goto L44
	} else {
		goto L51
	}
L51:
	;
	goto L45
L52:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223))))
	if v230&int32(254) == int32(48) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L3
L54:
	;
	if base.Ui32(int32(1073741824)) < base.Ui32(v222) {
		goto L4
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v230 != int32(95) {
		v296 = v222
		v297 = v223
		v298 = v219
		v299 = v230
		goto L27
	} else {
		goto L58
	}
L57:
	;
	v241 = int32(1)
	v222 = (v230-int32(48))&int32(255) | v222<<(uint(v241)%32)
	v223 = v223 + v241
	goto L52
L58:
	;
	v249 = v223 + int32(1)
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	if base.Ui32(int32(254)) <= base.Ui32((v250-int32(50))&int32(255)) {
		v223 = v249
		goto L52
	} else {
		goto L59
	}
L59:
	;
	goto L53
L60:
	;
	v271 = (v263 - int32(48)) & int32(255)
	if base.Ui32(v271) <= base.Ui32(int32(9)) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	goto L3
L62:
	;
	if base.Ui32(int32(214748364)) < base.Ui32(v260) {
		goto L4
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	if v263 != int32(95) {
		v296 = v260
		v297 = v261
		v298 = v116
		v299 = v263
		goto L27
	} else {
		goto L66
	}
L65:
	;
	v280 = v261 + int32(1)
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280))))
	v260 = v260*int32(10) + v271
	v261 = v280
	v263 = v281
	goto L60
L66:
	;
	if v261 == v116 {
		goto L3
	} else {
		goto L67
	}
L67:
	;
	v286 = v261 + int32(1)
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286))))
	if base.Ui32((v287-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v261 = v286
		v263 = v287
		goto L60
	} else {
		goto L68
	}
L68:
	;
	goto L61
L69:
	;
	v308 = v297
	v310 = v299
	goto L70
L70:
	;
	if base.Ui32(v310-int32(9)) < base.Ui32(int32(5)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v330 = v308 + int32(1)
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330))))
	v308 = v330
	v310 = v331
	goto L70
L73:
	;
	if v310 == int32(32) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	if v310 != 0 {
		goto L3
	} else {
		goto L75
	}
L75:
	;
	if v117 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v323 = int64(0) - base.I64_extend_i32_u(v296)
	if v323 != base.I64_extend32_s(v323) {
		goto L4
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	if int32(0) <= v296 {
		v402 = v296
		goto L1
	} else {
		goto L80
	}
L79:
	;
	v402 = base.I32_wrap_i64(v323)
	goto L1
L80:
	;
	goto L4
L81:
	;
	return int32(0)
L82:
	;
	if v343 == int32(0) {
		v402 = v342
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L81
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(224106)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
	F_errmsg(m, int32(189002), v13)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L81
	} else {
		goto L85
	}
L85:
	;
	v388 = v342
	v396 = int32(612)
	goto L2
L86:
	;
	if v370 == int32(0) {
		v402 = v369
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L81
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(224106)
	F_errmsg(m, int32(706620), v13+int32(16))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L81
	} else {
		goto L89
	}
L89:
	;
	v388 = v369
	v396 = int32(618)
	goto L2
L90:
	;
	v402 = v388
	goto L1
}
func F_pg_strxfrm(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v8 = m.T0[v7].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, l0, l2, l1, int32(-1), l3)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_pg_table_size(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_try_relation_open(m, v4, int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v12 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v12)
			return int32(0)
		} else {
			v16 = F_calculate_table_size(m, v6)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				F_relation_close(m, v6, int32(1))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v21 = F_Int64GetDatum(m, v16)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						return v21
					}
				}
			}
		}
	}
}
func F_pg_timezone_abbrevs_zone(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int64
	_ = v97
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int64
	_ = v196
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 int64
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int64
	_ = v361
	var v366 int64
	_ = v366
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int64
	_ = v376
	var v377 int64
	_ = v377
	var v380 int64
	_ = v380
	var v386 int32
	_ = v386
	var v388 int64
	_ = v388
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int64
	_ = v408
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+50)) = uint8(v2)
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+48)) = uint16(v2)
	v15 = *(*int64)(unsafe.Add(mBase, _consts[429]))
	v17 = base.I64_div_s(v15, int64(1000000))
	goto L1
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = v17 + int64(946684800)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	if v22 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L6
	} else {
		goto L104
	}
L3:
	;
	v25 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	goto L12
L6:
	;
	return int32(0)
L7:
	;
	v29 = int32(4486928)
	v30 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v32
	v35 = F_palloc(m, int32(4))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v37 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v35
	v43 = F_get_call_result_type(m, l0, v37, v6+int32(-56))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	if v43 != int32(1) {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v47
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v30
	goto L5
L11:
	;
	m.G0 = v8 - int32(-64)
	return v458
L12:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	v58 = *(*int32)(unsafe.Add(mBase, _consts[1060]))
	v59 = int32(0)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v62 < v59 {
		v82 = v59
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v82 != 0 {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	goto L13
L15:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v58)+268))
	if v65 <= v62 {
		v82 = v59
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v68 = v58 + int32(22376)
	v71 = v62
	goto L17
L17:
	;
	v76 = v71 + int32(1)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+v68))))
	if v77 != 0 {
		v71 = v76
		goto L17
	} else {
		goto L19
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v76
	v82 = v62 + v68
	goto L14
L19:
	;
	goto L18
L20:
	;
	v87 = v82
	goto L23
L21:
	;
	goto L22
L22:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L6
	} else {
		goto L103
	}
L23:
	;
	v90 = int32(504697)
	v94 = m.G0
	v96 = v94 - int32(32)
	v97 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v96)+24)) = v97
	*(*int64)(unsafe.Add(mBase, uint32(v96)+16)) = v97
	*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = v97
	*(*int64)(unsafe.Add(mBase, uint32(v96))) = v97
	v105 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1066])))
	if v105 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	goto L22
L25:
	;
	v416 = *(*int32)(unsafe.Add(mBase, _consts[1060]))
	v417 = int32(0)
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v420 < v417 {
		v440 = v417
		goto L96
	} else {
		goto L97
	}
L26:
	;
	v174 = F_strlen(m, v87)
	mBase = m.M
	if v173 != v174 {
		goto L25
	} else {
		goto L47
	}
L27:
	;
	v173 = int32(0)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1067])))
	if v109 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v113 = v87
	goto L33
L31:
	;
	goto L32
L32:
	;
	v123 = v90
	v124 = v105
	goto L36
L33:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	if v119 == v105 {
		v113 = v113 + int32(1)
		goto L33
	} else {
		goto L35
	}
L34:
	;
	v173 = v113 - v87
	goto L26
L35:
	;
	goto L34
L36:
	;
	v131 = v96 + int32(base.Ui32(v124)>>(uint(int32(3))%32))&int32(28)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v133 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = v132 | v133<<(uint(v124)%32)
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+1)))
	if v137 != 0 {
		v123 = v123 + v133
		v124 = v137
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v140 == int32(0) {
		v165 = v87
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L37
L39:
	;
	v173 = v165 - v87
	goto L26
L40:
	;
	v144 = v87
	v145 = v140
	goto L41
L41:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v96+int32(base.Ui32(v145)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v153)>>(uint(v145)%32))&int32(1) == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v165 = v161
	goto L39
L43:
	;
	v165 = v144
	goto L39
L44:
	;
	goto L45
L45:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+1)))
	v161 = v144 + int32(1)
	if v159 != 0 {
		v144 = v161
		v145 = v159
		goto L41
	} else {
		goto L46
	}
L46:
	;
	goto L42
L47:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _consts[1060]))
	v184 = int32(0)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v183)+268))
	if v191 <= v184 {
		v354 = v184
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v354 == int32(0) {
		goto L25
	} else {
		goto L86
	}
L49:
	;
	goto L48
L50:
	;
	v195 = v183 + int32(22376)
	v196 = *(*int64)(unsafe.Add(mBase, uint32(v6+int32(-24))))
	v204 = v184
	goto L51
L51:
	;
	v210 = F_strcmp(m, v87, v195+v204)
	mBase = m.M
	if v210 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v183)+260))
	if v228 <= int32(0) {
		goto L61
	} else {
		goto L62
	}
L53:
	;
	v212 = v204
	goto L56
L54:
	;
	goto L55
L55:
	;
	goto L52
L56:
	;
	v225 = v212 + int32(1)
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212+v195))))
	if v226 != 0 {
		v212 = v225
		goto L56
	} else {
		goto L58
	}
L57:
	;
	if v225 < v191 {
		v204 = v225
		goto L51
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
	;
	v354 = v184
	goto L49
L60:
	;
	v273 = v183 + int32(16280)
	v275 = v183 + int32(18280)
	v282 = v261
	goto L74
L61:
	;
	v261 = int32(0)
	goto L60
L62:
	;
	goto L63
L63:
	;
	v236 = int32(0)
	v241 = v228
	goto L64
L64:
	;
	v248 = int32(1)
	v249 = (v236 + v241) >> (uint(v248) % 32)
	v255 = *(*int64)(unsafe.Add(mBase, uint32(v183+int32(280)+v249<<(uint(int32(3))%32))))
	v256 = base.B2i32(v196 < v255)
	if v196 < v255 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v261 = v257
	goto L60
L66:
	;
	v257 = v236
	goto L68
L67:
	;
	v257 = v249 + v248
	goto L68
L68:
	;
	if v196 < v255 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v258 = v249
	goto L71
L70:
	;
	v258 = v241
	goto L71
L71:
	;
	if v257 < v258 {
		v236 = v257
		v241 = v258
		goto L64
	} else {
		goto L72
	}
L72:
	;
	goto L65
L73:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v333)))
	*(*int32)(unsafe.Add(mBase, uint32(v6+int32(-28)))) = v340
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v6+int32(-32)))) = v342
	v354 = int32(1)
	goto L49
L74:
	;
	if int32(0) < v282 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v183)+uint32(_consts[1068])))
	v302 = v275 + v299<<(uint(int32(4))%32)
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v302)+8))
	if v303 == v204 {
		v333 = v302
		goto L73
	} else {
		goto L80
	}
L76:
	;
	v291 = v282 - int32(1)
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273+v291))))
	v296 = v275 + v293<<(uint(int32(4))%32)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v296)+8))
	if v297 != v204 {
		v282 = v291
		goto L74
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	goto L75
L79:
	;
	v333 = v296
	goto L73
L80:
	;
	if v228 <= v261 {
		v354 = v184
		goto L49
	} else {
		goto L81
	}
L81:
	;
	v307 = v261
	goto L82
L82:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307+v273))))
	v322 = v275 + v319<<(uint(int32(4))%32)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)+8))
	if v323 == v204 {
		v333 = v322
		goto L73
	} else {
		goto L84
	}
L83:
	;
	v354 = v184
	goto L49
L84:
	;
	v326 = v307 + int32(1)
	if v228 != v326 {
		v307 = v326
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	v359 = F_cstring_to_text(m, v87)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
	;
	v361 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v361
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v361
	*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = v359
	v366 = int64(*(*int32)(unsafe.Add(mBase, uint32(v8)+36)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v366 * int64(1000000)
	v371 = v6 + int32(-56)
	v373 = F_palloc(m, int32(16))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L6
	} else {
		goto L88
	}
L88:
	;
	v376 = int64(*(*int32)(unsafe.Add(mBase, uint32(v371)+12)))
	v377 = int64(*(*int32)(unsafe.Add(mBase, uint32(v371)+16)))
	v380 = v376 + v377*int64(12)
	if base.Ui64(int64(-4294967296)) <= base.Ui64(v380-int64(2147483648)) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+56)) = v373
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+60)) = base.B2i32(v394 != int32(0))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v55)+28))
	v403 = F_heap_form_tuple(m, v398, v6+int32(-12), v6+int32(-16))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L6
	} else {
		goto L93
	}
L90:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v373)+12)) = uint32(v380)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v371)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v373)+8)) = v386
	v388 = *(*int64)(unsafe.Add(mBase, uint32(v371)))
	*(*int64)(unsafe.Add(mBase, uint32(v373))) = v388
	goto L92
L91:
	;
	goto L92
L92:
	;
	goto L89
L93:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v403)+16))
	v406 = F_HeapTupleHeaderGetDatum(m, v405)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L6
	} else {
		goto L94
	}
L94:
	;
	v408 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
	*(*int64)(unsafe.Add(mBase, uint32(v55))) = v408 + int64(1)
	v412 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v412)+20)) = int32(1)
	v458 = v406
	goto L11
L95:
	;
	if v440 != 0 {
		v87 = v440
		goto L23
	} else {
		goto L102
	}
L96:
	;
	goto L95
L97:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v416)+268))
	if v423 <= v420 {
		v440 = v417
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v426 = v416 + int32(22376)
	v429 = v420
	goto L99
L99:
	;
	v434 = v429 + int32(1)
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429+v426))))
	if v435 != 0 {
		v429 = v434
		goto L99
	} else {
		goto L101
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v434
	v440 = v420 + v426
	goto L96
L101:
	;
	goto L100
L102:
	;
	goto L24
L103:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v450)+20)) = int32(2)
	v453 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v453)
	v458 = int32(0)
	goto L11
L104:
	;
	F_errmsg_internal(m, int32(365102), int32(0))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L6
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(495687), int32(5160), int32(370131))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L6
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_trigger_depth(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, _consts[416]))
	return v3
}
func F_pg_ts_dict_is_visible(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v2)
	v14 = F_TSDictionaryIsVisibleExt(m, v9, v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v23 = v2
		} else {
			v23 = v14
		}
		m.G0 = v7 + int32(16)
		return v23
	}
}
func F_pg_tzenumerate_next(m *base.Module, l0 int32) int32 {
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v276 int32
	_ = v276
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(2080)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 < v2 {
		v276 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L7
	} else {
		goto L66
	}
L2:
	;
	m.G0 = v11 + int32(2080)
	return v276
L3:
	;
	v19 = l0 + int32(88)
	v21 = l0 + int32(48)
	v23 = l0 + int32(8)
	v25 = v13
	goto L4
L4:
	;
	v33 = v25 << (uint(int32(2)) % 32)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v23+v33)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v21)))
	v38 = F_ReadDir(m, v35, v37)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v276 = int32(0)
	goto L2
L6:
	;
	if int32(0) <= v268 {
		v25 = v268
		goto L4
	} else {
		goto L65
	}
L7:
	;
	return int32(0)
L8:
	;
	if v38 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v23+v44<<(uint(int32(2))%32))))
	F_FreeDir(m, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+19)))
	if v62 == int32(46) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v21+v51<<(uint(int32(2))%32))))
	F_pfree(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v60 = v58 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v60
	v268 = v60
	goto L6
L14:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v268 = v267
	goto L6
L15:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v21+v65<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v38 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v69
	v80 = F_pg_snprintf(m, v11+int32(32), int32(2048), int32(176209), v11+int32(16))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v86 = F_get_dirent_type(m, v11+int32(32), v38, int32(1), int32(21))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	if v86 == int32(3) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(9) <= v90 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v141 = F_tzload(m, v136+(v11+int32(32)), int32(0), l0+int32(344))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L7
	} else {
		goto L29
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90 + int32(1)
	v98 = F_pstrdup(m, v11+int32(32))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v21+v100<<(uint(int32(2))%32)))) = v98
	v107 = F_AllocateDir(m, v11+int32(32))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v110 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v23+v109<<(uint(v110)%32)))) = v107
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v23+v114<<(uint(v110)%32))))
	if v118 != 0 {
		v268 = v114
		goto L6
	} else {
		goto L24
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v11 + int32(32)
	F_errmsg(m, int32(294698), v11)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(489161), int32(463), int32(63306))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L7
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
	if v141 != 0 {
		goto L14
	} else {
		goto L30
	}
L30:
	;
	v143 = F_pg_tz_acceptable(m, v19)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	if v143 == int32(0) {
		goto L14
	} else {
		goto L32
	}
L32:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v150 = v147 + (v11 + int32(32))
	goto L36
L33:
	;
	v276 = v19
	goto L2
L34:
	;
	v263 = F_strlen(m, v252)
	mBase = m.M
	goto L33
L36:
	;
	goto L37
L37:
	;
	v157 = int32(255)
	if (v19^v150)&int32(3) != 0 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v256 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v253))) = uint8(v256)
	goto L34
L39:
	;
	v237 = v232
	v238 = v233
	v239 = v234
	goto L61
L40:
	;
	if v227 == int32(0) {
		v252 = v225
		v253 = v226
		goto L38
	} else {
		goto L60
	}
L41:
	;
	v225 = v150
	v226 = v19
	v227 = v157
	goto L40
L42:
	;
	goto L43
L43:
	;
	if v150&int32(3) == int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v194 == int32(0) {
		v252 = v191
		v253 = v192
		goto L38
	} else {
		goto L53
	}
L45:
	;
	v191 = v150
	v192 = v19
	v193 = v157
	v194 = int32(1)
	goto L44
L46:
	;
	goto L47
L47:
	;
	v170 = v150
	v171 = v19
	v172 = v157
	goto L48
L48:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	*(*uint8)(unsafe.Add(mBase, uint32(v171))) = uint8(v174)
	if v174 == int32(0) {
		v232 = v170
		v233 = v171
		v234 = v172
		goto L39
	} else {
		goto L50
	}
L49:
	;
	v191 = v185
	v192 = v179
	v193 = v181
	v194 = v183
	goto L44
L50:
	;
	v178 = int32(1)
	v179 = v171 + v178
	v181 = v172 - v178
	v182 = int32(0)
	v183 = base.B2i32(v181 != v182)
	v185 = v170 + v178
	if v185&int32(3) == v182 {
		v191 = v185
		v192 = v179
		v193 = v181
		v194 = v183
		goto L44
	} else {
		goto L51
	}
L51:
	;
	if v181 != 0 {
		v170 = v185
		v171 = v179
		v172 = v181
		goto L48
	} else {
		goto L52
	}
L52:
	;
	goto L49
L53:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	if v197 == int32(0) {
		v225 = v191
		v226 = v192
		v227 = v193
		goto L40
	} else {
		goto L54
	}
L54:
	;
	if base.Ui32(v193) < base.Ui32(int32(4)) {
		v225 = v191
		v226 = v192
		v227 = v193
		goto L40
	} else {
		goto L55
	}
L55:
	;
	v203 = v191
	v204 = v192
	v205 = v193
	goto L56
L56:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v211 = int32(-2139062144)
	if (int32(16843008)-v208|v208)&v211 != v211 {
		v232 = v203
		v233 = v204
		v234 = v205
		goto L39
	} else {
		goto L58
	}
L57:
	;
	v225 = v219
	v226 = v217
	v227 = v221
	goto L40
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v204))) = v208
	v216 = int32(4)
	v217 = v204 + v216
	v219 = v203 + v216
	v221 = v205 - v216
	if base.Ui32(int32(3)) < base.Ui32(v221) {
		v203 = v219
		v204 = v217
		v205 = v221
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v232 = v225
	v233 = v226
	v234 = v227
	goto L39
L61:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237))))
	*(*uint8)(unsafe.Add(mBase, uint32(v238))) = uint8(v241)
	if v241 == int32(0) {
		v252 = v237
		v253 = v238
		goto L38
	} else {
		goto L63
	}
L62:
	;
	v252 = v248
	v253 = v246
	goto L38
L63:
	;
	v245 = int32(1)
	v246 = v238 + v245
	v248 = v237 + v245
	v250 = v239 - v245
	if v250 != 0 {
		v237 = v248
		v238 = v246
		v239 = v250
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	goto L5
L66:
	;
	F_errmsg_internal(m, int32(31288), int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L7
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(489161), int32(455), int32(63306))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L7
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_tzset(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(23952)
	m.G0 = v8
	v10 = F_strlen(m, l0)
	mBase = m.M
	if base.Ui32(int32(255)) < base.Ui32(v10) {
		v288 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(23952)
	return v288
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _consts[1259]))
	if v14 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+528)) = int64(102873056674048)
	v25 = F_hash_create(m, int32(162067), int32(4), v8+int32(512), int32(24))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v32 = v14
	goto L5
L5:
	;
	v34 = v8 + int32(256)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v35 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	return int32(0)
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1259])) = v25
	if v25 == int32(0) {
		v288 = v2
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v32 = v25
	goto L5
L9:
	;
	v36 = l0
	v38 = v34
	v40 = v35
	goto L12
L10:
	;
	v62 = v34
	v63 = v32
	goto L11
L11:
	;
	v65 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v65)
	v71 = F_hash_search(m, v63, v8+int32(256), v65, v65)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L6
	} else {
		goto L19
	}
L12:
	;
	if base.Ui32((v40-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _consts[1259]))
	v62 = v54
	v63 = v59
	goto L11
L14:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v51)
	v53 = int32(1)
	v54 = v38 + v53
	v56 = v36 + v53
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v57 != 0 {
		v36 = v56
		v38 = v54
		v40 = v57
		goto L12
	} else {
		goto L18
	}
L15:
	;
	v49 = v40 - int32(32)
	goto L17
L16:
	;
	v49 = v40
	goto L17
L17:
	;
	v51 = v49 & int32(255)
	goto L14
L18:
	;
	goto L13
L19:
	;
	if v71 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v288 = v71 + int32(256)
	goto L1
L21:
	;
	goto L22
L22:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v8)+256))
	if v75 == int32(5524807) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v196 = *(*int32)(unsafe.Add(mBase, _consts[1259]))
	v201 = F_hash_search(m, v196, v8+int32(256), int32(1), int32(0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L6
	} else {
		goto L57
	}
L24:
	;
	v119 = v8 + int32(256)
	if (v119^v8)&int32(3) != 0 {
		goto L39
	} else {
		goto L40
	}
L25:
	;
	v83 = F_tzparse(m, v8+int32(256), v8+int32(512), int32(1))
	mBase = m.M
	if v83 != 0 {
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v101 = F_tzload(m, v8+int32(256), v8, v8+int32(512))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L6
	} else {
		goto L32
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	F_errmsg_internal(m, int32(370308), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(489161), int32(278), int32(104908))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	if v101 == int32(0) {
		goto L23
	} else {
		goto L33
	}
L33:
	;
	v105 = int32(0)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+256)))
	if v106 == int32(58) {
		v288 = v105
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v113 = int32(0)
	v114 = F_tzparse(m, v8+int32(256), v8+int32(512), v113)
	mBase = m.M
	if v114 == v113 {
		v288 = v105
		goto L1
	} else {
		goto L35
	}
L35:
	;
	goto L24
L36:
	;
	goto L23
L37:
	;
	goto L36
L38:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v174))) = uint8(v173)
	if v173&int32(255) == int32(0) {
		goto L37
	} else {
		goto L53
	}
L39:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	v172 = v119
	v173 = v125
	v174 = v8
	goto L38
L40:
	;
	goto L41
L41:
	;
	if v119&int32(3) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v129 = v119
	v131 = v8
	goto L45
L43:
	;
	v143 = v119
	v145 = v8
	goto L44
L44:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	v150 = int32(-2139062144)
	if (int32(16843008)-v147|v147)&v150 != v150 {
		v172 = v143
		v173 = v147
		v174 = v145
		goto L38
	} else {
		goto L49
	}
L45:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
	*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v132)
	if v132 == int32(0) {
		goto L37
	} else {
		goto L47
	}
L46:
	;
	v143 = v139
	v145 = v137
	goto L44
L47:
	;
	v136 = int32(1)
	v137 = v131 + v136
	v139 = v129 + v136
	if v139&int32(3) != 0 {
		v129 = v139
		v131 = v137
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v155 = v143
	v156 = v147
	v157 = v145
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v157))) = v156
	v159 = int32(4)
	v160 = v157 + v159
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	v163 = v155 + v159
	v167 = int32(-2139062144)
	if (v161|(int32(16843008)-v161))&v167 == v167 {
		v155 = v163
		v156 = v161
		v157 = v160
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v172 = v163
	v173 = v161
	v174 = v160
	goto L38
L52:
	;
	goto L51
L53:
	;
	v181 = v172
	v183 = v174
	goto L54
L54:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v183)+1)) = uint8(v184)
	v186 = int32(1)
	if v184 != 0 {
		v181 = v181 + v186
		v183 = v183 + v186
		goto L54
	} else {
		goto L56
	}
L55:
	;
	goto L37
L56:
	;
	goto L55
L57:
	;
	v204 = v201 + int32(256)
	if (v8^v204)&int32(3) != 0 {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	v279 = int32(512)
	goto L80
L59:
	;
	goto L58
L60:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v259))) = uint8(v258)
	if v258&int32(255) == int32(0) {
		goto L59
	} else {
		goto L75
	}
L61:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	v257 = v8
	v258 = v210
	v259 = v204
	goto L60
L62:
	;
	goto L63
L63:
	;
	if v8&int32(3) != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v214 = v8
	v216 = v204
	goto L67
L65:
	;
	v228 = v8
	v230 = v204
	goto L66
L66:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	v235 = int32(-2139062144)
	if (int32(16843008)-v232|v232)&v235 != v235 {
		v257 = v228
		v258 = v232
		v259 = v230
		goto L60
	} else {
		goto L71
	}
L67:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214))))
	*(*uint8)(unsafe.Add(mBase, uint32(v216))) = uint8(v217)
	if v217 == int32(0) {
		goto L59
	} else {
		goto L69
	}
L68:
	;
	v228 = v224
	v230 = v222
	goto L66
L69:
	;
	v221 = int32(1)
	v222 = v216 + v221
	v224 = v214 + v221
	if v224&int32(3) != 0 {
		v214 = v224
		v216 = v222
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	v240 = v228
	v241 = v232
	v242 = v230
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v242))) = v241
	v244 = int32(4)
	v245 = v242 + v244
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	v248 = v240 + v244
	v252 = int32(-2139062144)
	if (v246|(int32(16843008)-v246))&v252 == v252 {
		v240 = v248
		v241 = v246
		v242 = v245
		goto L72
	} else {
		goto L74
	}
L73:
	;
	v257 = v248
	v258 = v246
	v259 = v245
	goto L60
L74:
	;
	goto L73
L75:
	;
	v266 = v257
	v268 = v259
	goto L76
L76:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v268)+1)) = uint8(v269)
	v271 = int32(1)
	if v269 != 0 {
		v266 = v266 + v271
		v268 = v268 + v271
		goto L76
	} else {
		goto L78
	}
L77:
	;
	goto L59
L78:
	;
	goto L77
L79:
	;
	v288 = v204
	goto L1
L80:
	;
	v284 = F__emscripten_memcpy_bulkmem(m, v201+v279, v8+v279, int32(23440))
	mBase = m.M
	goto L82
L82:
	;
	goto L79
}
func F_pg_ulltoa_n(m *base.Module, l0 int64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v30 int64
	_ = v30
	var v32 int32
	_ = v32
	var v36 int64
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int64
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v114 int64
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	v3 = int32(0)
	if l0 == int64(0) {
		v12 = int32(48)
		*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v12)
		return int32(1)
	} else {
		v20 = int32(1233)
		v25 = int32(base.Ui32((base.I32_wrap_i64(base.I64_clz(l0))^int32(63))*v20+v20) >> (uint(int32(12)) % 32))
		v30 = *(*int64)(unsafe.Add(mBase, uint32(v25<<(uint(int32(3))%32))+uint32(_consts[1125])))
		v32 = v25 + base.B2i32(base.Ui64(v30) <= base.Ui64(l0))
		if base.Ui64(l0) < base.Ui64(int64(100000000)) {
			v110 = v3
			v114 = l0
		} else {
			v36 = l0
			v40 = v3
			for {
				v45 = l1 + v32 - v40
				v46 = int32(8)
				v49 = base.I64_div_u_s(v36, int64(100000000))
				v53 = base.I32_wrap_i64(v49*int64(4194967296) + v36)
				v55 = base.I32_div_u_s(v53, int32(1000000))
				v56 = int32(1)
				v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55<<(uint(v56)%32))+uint32(_consts[1062]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v45-v46))) = uint16(v60)
				v64 = int32(10000)
				v65 = base.I32_div_u_s(v53, v64)
				v66 = int32(100)
				v67 = base.I32_rem_u_s(v65, v66)
				v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67<<(uint(v56)%32))+uint32(_consts[1062]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v45-int32(6)))) = uint16(v72)
				v78 = v53 - v65*v64
				v79 = int32(65535)
				v82 = base.I32_div_u_s(v78&v79, v66)
				v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82<<(uint(v56)%32))+uint32(_consts[1062]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v45-int32(4)))) = uint16(v87)
				v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v78-v82*v66)&v79<<(uint(v56)%32))+uint32(_consts[1062]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v45-int32(2)))) = uint16(v100)
				v103 = v40 + v46
				if base.Ui64(int64(9999999999999999)) < base.Ui64(v36) {
					v36 = v49
					v40 = v103
					continue
				} else {
					break
				}
				break
			}
			v110 = v103
			v114 = v49
		}
		v115 = base.I32_wrap_i64(v114)
		if base.Ui64(v114) < base.Ui64(int64(10000)) {
			v149 = v115
			v150 = v110
		} else {
			v119 = l1 + v32 - v110
			v120 = int32(4)
			v123 = base.I32_div_u_s(v115, int32(10000))
			v126 = v123*int32(-10000) + v115
			v127 = int32(100)
			v128 = base.I32_div_u_s(v126, v127)
			v129 = int32(1)
			v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v128<<(uint(v129)%32))+uint32(_consts[1062]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v119-v120))) = uint16(v133)
			v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v126-v128*v127)<<(uint(v129)%32))+uint32(_consts[1062]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v119-int32(2)))) = uint16(v144)
			v149 = v123
			v150 = v110 | v120
		}
		if base.Ui32(v149) < base.Ui32(int32(100)) {
			v172 = v149
			v173 = v150
		} else {
			v157 = int32(2)
			v159 = int32(100)
			v160 = base.I32_div_u_s(v149, v159)
			v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v149-v160*v159)<<(uint(int32(1))%32))+uint32(_consts[1062]))))
			*(*uint16)(unsafe.Add(mBase, uint32(l1+v32-v150-v157))) = uint16(v168)
			v172 = v160
			v173 = v150 + v157
		}
		if base.Ui32(int32(10)) <= base.Ui32(v172) {
			v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v172<<(uint(int32(1))%32))+uint32(_consts[1062]))))
			*(*uint16)(unsafe.Add(mBase, uint32(l1+v32-v173-int32(2)))) = uint16(v184)
			return v32
		} else {
			v188 = v172 | int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v188)
			return v32
		}
	}
}
func F_pg_utf_dsplen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v29 int32
	_ = v29
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	v6 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	v8 = v6 & int32(255)
	if v6 < int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v8&int32(224) == int32(192) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v65 = v8
	goto L3
L3:
	;
	if v65 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v58))))
	v65 = v59 | v61&int32(63)
	goto L3
L5:
	;
	v58 = int32(1)
	v59 = v8 << (uint(int32(6)) % 32) & int32(1984)
	goto L4
L6:
	;
	goto L7
L7:
	;
	if v8&int32(240) == int32(224) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v58 = int32(2)
	v59 = v8<<(uint(int32(12))%32)&int32(61440) | v29&int32(63)<<(uint(int32(6))%32)
	goto L4
L9:
	;
	goto L10
L10:
	;
	if v8&int32(248) != int32(240) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(-1)
L12:
	;
	goto L13
L13:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v47 = int32(63)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	v58 = int32(3)
	v59 = v8<<(uint(int32(18))%32)&int32(1835008) | v46&v47<<(uint(int32(12))%32) | v52&v47<<(uint(int32(6))%32)
	goto L4
L14:
	;
	return int32(0)
L15:
	;
	goto L16
L16:
	;
	v71 = int32(-1)
	if base.Ui32(v65) < base.Ui32(int32(32)) {
		v151 = v71
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return v151
L18:
	;
	if base.Ui32(int32(1114111)) < base.Ui32(v65) {
		v151 = v71
		goto L17
	} else {
		goto L19
	}
L19:
	;
	if base.Ui32(v65-int32(127)) < base.Ui32(int32(33)) {
		v151 = v71
		goto L17
	} else {
		goto L20
	}
L20:
	;
	if base.Ui32(v65-int32(918000)) < base.Ui32(int32(-917827)) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(1)
L22:
	;
	goto L23
L23:
	;
	v88 = int32(0)
	v90 = int32(333)
	goto L24
L24:
	;
	v95 = base.I32_div_s(v88+v90, int32(2))
	v97 = v95 << (uint(int32(3)) % 32)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+uint32(_consts[1287])))
	if base.Ui32(v100) < base.Ui32(v65) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	if base.Ui32(v65-int32(262142)) < base.Ui32(int32(-257790)) {
		goto L34
	} else {
		goto L35
	}
L26:
	;
	if v112 <= v113 {
		v88 = v112
		v90 = v113
		goto L24
	} else {
		goto L33
	}
L27:
	;
	v112 = v95 + int32(1)
	v113 = v90
	goto L26
L28:
	;
	goto L29
L29:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v97)+uint32(_consts[1288])))
	if base.Ui32(v106) <= base.Ui32(v65) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	return int32(0)
L31:
	;
	goto L32
L32:
	;
	v112 = v88
	v113 = v95 - int32(1)
	goto L26
L33:
	;
	goto L25
L34:
	;
	return int32(1)
L35:
	;
	goto L36
L36:
	;
	v125 = int32(0)
	v126 = int32(121)
	goto L37
L37:
	;
	v130 = base.I32_div_s(v125+v126, int32(2))
	v132 = v130 << (uint(int32(3)) % 32)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v132)+uint32(_consts[1289])))
	if base.Ui32(v135) < base.Ui32(v65) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v151 = int32(1)
	goto L17
L39:
	;
	if v147 <= v148 {
		v125 = v147
		v126 = v148
		goto L37
	} else {
		goto L46
	}
L40:
	;
	v147 = v130 + int32(1)
	v148 = v126
	goto L39
L41:
	;
	goto L42
L42:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v132)+uint32(_consts[1290])))
	if base.Ui32(v141) <= base.Ui32(v65) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	return int32(2)
L44:
	;
	goto L45
L45:
	;
	v147 = v125
	v148 = v130 - int32(1)
	goto L39
L46:
	;
	goto L38
}
func F_pg_verify_mbstr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0*int32(28))+uint32(_consts[1222])))
	v11 = m.T0[v10].(func(*base.Module, int32, int32) int32)(m, l1, l2)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if l3 != 0 {
			return base.B2i32(l2 == v11)
		} else {
			if l2 == v11 {
				return base.B2i32(l2 == v11)
			} else {
				F_report_invalid_encoding(m, l0, l1+v11, l2-v11)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_pg_walfile_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v62 int64
	_ = v62
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	v6 = m.G0
	v8 = v6 - int32(96)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, _consts[182])))
	if v14 == int32(1) {
		v19 = *(*int32)(unsafe.Add(mBase, _consts[175]))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+316))
		v22 = base.B2i32(v20 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _consts[182])) = uint8(v22)
		v24 = v22
	} else {
		v24 = int32(0)
	}
	if v24 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(127299), int32(0))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(665833)
					F_errhint(m, int32(554742), v8)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(491704), int32(449), int32(377625))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	} else {
		v49 = int64(*(*int32)(unsafe.Add(mBase, _consts[176])))
		v51 = *(*int32)(unsafe.Add(mBase, _consts[175]))
		v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+308))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v52
		v54 = base.I64_div_u_s(v11, v49)
		v57 = int64(*(*int32)(unsafe.Add(mBase, _consts[176])))
		v58 = base.I64_div_u_s(int64(4294967296), v57)
		v59 = base.I64_div_u_s(v54, v58)
		*(*uint32)(unsafe.Add(mBase, uint32(v8)+20)) = uint32(v59)
		v62 = v54 - v58*v59
		*(*uint32)(unsafe.Add(mBase, uint32(v8)+24)) = uint32(v62)
		v70 = F_pg_snprintf(m, v8+int32(32), int32(64), int32(506773), v8+int32(16))
		mBase = m.M
		v71 = m.ExcPending
		if v71 != 0 {
			return int32(0)
		} else {
			v74 = F_cstring_to_text(m, v8+int32(32))
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return int32(0)
			} else {
				m.G0 = v8 + int32(96)
				return v74
			}
		}
	}
}
func F_pg_walfile_name_offset(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int64
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int64
	_ = v74
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v82 int64
	_ = v82
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	v7 = m.G0
	v9 = v7 - int32(112)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, _consts[182])))
	if v15 == int32(1) {
		v20 = *(*int32)(unsafe.Add(mBase, _consts[175]))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+316))
		v23 = base.B2i32(v21 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _consts[182])) = uint8(v23)
		v25 = v23
	} else {
		v25 = int32(0)
	}
	if v25 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(127299), int32(0))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(664768)
					F_errhint(m, int32(554742), v9)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(491704), int32(391), int32(105079))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	} else {
		v50 = F_CreateTemplateTupleDesc(m, int32(2))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			F_TupleDescInitEntry(m, v50, int32(1), int32(377631), int32(25), int32(-1), int32(0))
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int32(0)
			} else {
				F_TupleDescInitEntry(m, v50, int32(2), int32(105102), int32(23), int32(-1), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return int32(0)
				} else {
					v66 = F_BlessTupleDesc(m, v50)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						v69 = int64(*(*int32)(unsafe.Add(mBase, _consts[176])))
						v71 = *(*int32)(unsafe.Add(mBase, _consts[175]))
						v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+308))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v72
						v74 = base.I64_div_u_s(v12, v69)
						v77 = int64(*(*int32)(unsafe.Add(mBase, _consts[176])))
						v78 = base.I64_div_u_s(int64(4294967296), v77)
						v79 = base.I64_div_u_s(v74, v78)
						*(*uint32)(unsafe.Add(mBase, uint32(v9)+20)) = uint32(v79)
						v82 = v74 - v78*v79
						*(*uint32)(unsafe.Add(mBase, uint32(v9)+24)) = uint32(v82)
						v90 = F_pg_snprintf(m, v9+int32(48), int32(64), int32(506773), v9+int32(16))
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							v94 = F_cstring_to_text(m, v9+int32(48))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v94
								v97 = int32(0)
								*(*uint16)(unsafe.Add(mBase, uint32(v9)+38)) = uint16(v97)
								v101 = *(*int32)(unsafe.Add(mBase, _consts[176]))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = base.I32_wrap_i64(v12) & (v101 - int32(1))
								v110 = F_heap_form_tuple(m, v66, v9+int32(40), v9+int32(38))
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int32(0)
								} else {
									v112 = *(*int32)(unsafe.Add(mBase, uint32(v110)+16))
									v113 = F_HeapTupleHeaderGetDatum(m, v112)
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return int32(0)
									} else {
										m.G0 = v9 + int32(112)
										return v113
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
