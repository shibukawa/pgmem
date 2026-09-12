package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetDatabaseEncodingName(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _consts[356]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	return v3
}
func F_LockDatabaseObject(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v4 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(264)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+14)) = uint16(v9)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
	v16 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v16
	v20 = F_LockAcquire(m, v7, l2, v4, v4)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return
	} else {
		F_ReceiveSharedInvalidMessages(m)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			m.G0 = v7 + int32(16)
			return
		}
	}
}
func F_database_is_invalid_form(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	return base.B2i32(v2 == int32(-2))
}
func F_get_database_list(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
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
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	v1 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	F_StartTransactionCommand(m)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = F_table_open(m, int32(1262), int32(1))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = int32(0)
	v25 = F_table_beginscan_catalog(m, v21, v23, v23)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = F_heap_getnext(m, v25)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v27 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v29 = v27
	v31 = v1
	goto L9
L7:
	;
	v92 = v1
	goto L8
L8:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+188))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	m.T0[v100].(func(*base.Module, int32))(m, v25)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L25
	}
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+22)))
	v39 = v37 + v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+80))
	goto L12
L10:
	;
	v92 = v86
	goto L8
L11:
	;
	v88 = F_heap_getnext(m, v25)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L23
	}
L12:
	;
	if v40 == int32(-2) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v45 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v60 = int32(4455216)
	v61 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v14
	v65 = F_palloc(m, int32(20))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L20
	}
L16:
	;
	if v45 == int32(0) {
		v86 = v31
		goto L11
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v39 + int32(4)
	F_errmsg_internal(m, int32(679169), v11)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(480199), int32(1842), int32(71841))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v86 = v31
	goto L11
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v67
	v71 = F_pstrdup(m, v39+int32(4))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = v71
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v39)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v39)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v65)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v65)+12)) = v76
	v80 = F_lappend(m, v31, v65)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v61
	v86 = v80
	goto L11
L23:
	;
	if v88 != 0 {
		v29 = v88
		v31 = v86
		goto L9
	} else {
		goto L24
	}
L24:
	;
	goto L10
L25:
	;
	F_sequence_close(m, v21, int32(1))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v14
	m.G0 = v11 + int32(16)
	return v92
}
func F_get_database_oid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v14 = F_table_open(m, int32(1262), int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		F_ScanKeyInit(m, v8+int32(-48), int32(2), int32(3), int32(62), l0)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v26 = int32(1)
			v31 = F_systable_beginscan(m, v14, int32(2671), v26, int32(0), v26, v8+int32(-48))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				v33 = F_systable_getnext(m, v31)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					if v33 != 0 {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v35+v36)))
						v39 = v38
					} else {
						v39 = int32(0)
					}
					F_systable_endscan(m, v31)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						F_sequence_close(m, v14, int32(1))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							if l1 != 0 {
								m.G0 = v10 - int32(-64)
								return v39
							} else {
								if v39 != 0 {
									m.G0 = v10 - int32(-64)
									return v39
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(1283))
										mBase = m.M
										v51 = m.ExcPending
										if v51 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
											F_errmsg(m, int32(70056), v10)
											mBase = m.M
											v55 = m.ExcPending
											if v55 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(477477), int32(3202), int32(419487))
												mBase = m.M
												v60 = m.ExcPending
												if v60 != 0 {
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
				}
			}
		}
	}
}
