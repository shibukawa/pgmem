package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_wal_replay_resume(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_wal_replay_resume[0])))
	if v8 == int32(1) {
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wal_replay_resume[1]))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+316))
		v16 = base.B2i32(v14 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_pg_wal_replay_resume[0])) = uint8(v16)
		v18 = v16
	} else {
		v18 = int32(0)
	}
	if v18 != 0 {
		v19 = F_PromoteIsTriggered(m)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			if v19 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_pg_wal_replay_resume_0), int32(0))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(_a_F_pg_wal_replay_resume_1)
							F_errhint(m, int32(_a_F_pg_wal_replay_resume_2), v4)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_pg_wal_replay_resume_3), int32(561), int32(_a_F_pg_wal_replay_resume_4))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
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
				F_SetRecoveryPause(m, int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					m.G0 = v4 + int32(16)
					return int32(0)
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
			F_errcode(m, int32(325))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_pg_wal_replay_resume_5), int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					F_errhint(m, int32(_a_F_pg_wal_replay_resume_6), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_wal_replay_resume_3), int32(554), int32(_a_F_pg_wal_replay_resume_4))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
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
func F_pg_wal_summary_contents(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int64
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v128 int64
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	v8 = m.G0
	v10 = v8 - int32(1152)
	m.G0 = v10
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+1116)) = uint16(v18)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+1112)) = v18
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
	if base.Ui64(int64(-2147483648)) < base.Ui64(v23-int64(2147483648)) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v10)+1104)) = uint32(v23)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+1088)) = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v33 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+1080)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+1096)) = v33
	v39 = F_OpenWalSummaryFile(m, v10+int32(1088))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L45
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+1072)) = v39
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wal_summary_contents[0]))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v39*int32(48))+32))
	goto L7
L7:
	;
	v50 = F_CreateBlockRefTableReader(m, v10+int32(1072), v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v58 = F_BlockRefTableReaderNextRelation(m, v50, v10+int32(1060), v10+int32(1056), v10+int32(1052))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v58 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	goto L13
L11:
	;
	goto L12
L12:
	;
	F_DestroyBlockRefTableReader(m, v50)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L43
	}
L13:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wal_summary_contents[1]))
	if v68 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L12
L15:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1068))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+1120)) = v71
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v10)+1060))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+1124)) = v73
	v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10)+1056)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+1132)) = v75
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1052))
	if v77 != int32(-1) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	v81 = F_Int64GetDatum(m, base.I64_extend_i32_u(v77))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	goto L25
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+1140)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+1136)) = v81
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v91 = F_heap_form_tuple(m, v86, v10+int32(1120), v10+int32(1112))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	F_tuplestore_puttuple(m, v93, v91)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wal_summary_contents[1]))
	if v105 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v151 = F_BlockRefTableReaderNextRelation(m, v50, v10+int32(1060), v10+int32(1056), v10+int32(1052))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L41
	}
L27:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v111 = F_BlockRefTableReaderGetBlocks(m, v50, v10+int32(16), int32(256))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	if v111 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v113 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+1140)) = v113
	v116 = v113
	goto L35
L33:
	;
	goto L34
L34:
	;
	goto L26
L35:
	;
	v128 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v10+int32(16)+v116<<(uint(int32(2))%32)))))
	v129 = F_Int64GetDatum(m, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L37
	}
L36:
	;
	goto L25
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+1136)) = v129
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v137 = F_heap_form_tuple(m, v132, v10+int32(1120), v10+int32(1112))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	F_tuplestore_puttuple(m, v139, v137)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v143 = v116 + int32(1)
	if v143 != v111 {
		v116 = v143
		goto L35
	} else {
		goto L40
	}
L40:
	;
	goto L36
L41:
	;
	if v151 != 0 {
		goto L13
	} else {
		goto L42
	}
L42:
	;
	goto L14
L43:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1072))
	F_FileClose(m, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	m.G0 = v10 + int32(1152)
	return int32(0)
L45:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v23
	F_errmsg(m, int32(_a_F_pg_wal_summary_contents_0), v10)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_pg_wal_summary_contents_1), int32(95), int32(_a_F_pg_wal_summary_contents_2))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
