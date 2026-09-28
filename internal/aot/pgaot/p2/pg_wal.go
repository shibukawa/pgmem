package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_wal_replay_resume(m *base.Module, l0 int32) int64 {
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
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+308))
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
			return int64(0)
		} else {
			if v19 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_pg_wal_replay_resume_0), int32(0))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(_a_F_pg_wal_replay_resume_1)
							F_errhint(m, int32(_a_F_pg_wal_replay_resume_2), v4)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_pg_wal_replay_resume_3), int32(592), int32(_a_F_pg_wal_replay_resume_4))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int64(0)
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
					return int64(0)
				} else {
					m.G0 = v4 + int32(16)
					return int64(0)
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_pg_wal_replay_resume_5), int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int64(0)
				} else {
					F_errhint(m, int32(_a_F_pg_wal_replay_resume_6), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_wal_replay_resume_3), int32(585), int32(_a_F_pg_wal_replay_resume_4))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int64(0)
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
func F_pg_wal_summary_contents(m *base.Module, l0 int32) int64 {
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
	var v22 int64
	_ = v22
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v125 int64
	_ = v125
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	v8 = m.G0
	v10 = v8 - int32(1168)
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
	return int64(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+1116)) = uint16(v18)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+1112)) = v18
	v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(int64(-2147483648)) < base.Ui64(v22-int64(2147483648)) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v10)+1104)) = uint32(v22)
	v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+1088)) = v28
	v30 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+1080)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+1096)) = v30
	v36 = F_OpenWalSummaryFile(m, v10+int32(1088))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
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
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L43
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+1072)) = v36
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wal_summary_contents[0]))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v36*int32(48))+32))
	goto L7
L7:
	;
	v47 = F_CreateBlockRefTableReader(m, v10+int32(1072), v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v55 = F_BlockRefTableReaderNextRelation(m, v47, v10+int32(1060), v10+int32(1056), v10+int32(1052))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v55 != 0 {
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
	F_DestroyBlockRefTableReader(m, v47)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L41
	}
L13:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wal_summary_contents[1]))
	if v65 != 0 {
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
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v68 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v10)+1068)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+1120)) = v68
	v70 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v10)+1060)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+1128)) = v70
	v72 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v10)+1064)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+1136)) = v72
	v74 = int64(*(*int16)(unsafe.Add(mBase, uint32(v10)+1056)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+1144)) = v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1052))
	if v76 != int32(-1) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+1160)) = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+1152)) = base.I64_extend_i32_u(v76)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v88 = F_heap_form_tuple(m, v83, v10+int32(1120), v10+int32(1112))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	goto L24
L22:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	F_tuplestore_puttuple(m, v90, v88)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wal_summary_contents[1]))
	if v102 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v146 = F_BlockRefTableReaderNextRelation(m, v47, v10+int32(1060), v10+int32(1056), v10+int32(1052))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L39
	}
L26:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v108 = F_BlockRefTableReaderGetBlocks(m, v47, v10+int32(16), int32(256))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	if v108 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+1160)) = int64(0)
	v113 = int32(0)
	goto L34
L32:
	;
	goto L33
L33:
	;
	goto L25
L34:
	;
	v125 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v10+int32(16)+v113<<(uint(int32(2))%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+1152)) = v125
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v132 = F_heap_form_tuple(m, v127, v10+int32(1120), v10+int32(1112))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L36
	}
L35:
	;
	goto L24
L36:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	F_tuplestore_puttuple(m, v134, v132)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v138 = v113 + int32(1)
	if v138 != v108 {
		v113 = v138
		goto L34
	} else {
		goto L38
	}
L38:
	;
	goto L35
L39:
	;
	if v146 != 0 {
		goto L13
	} else {
		goto L40
	}
L40:
	;
	goto L14
L41:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1072))
	F_FileClose(m, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	m.G0 = v10 + int32(1168)
	return int64(0)
L43:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v22
	F_errmsg(m, int32(_a_F_pg_wal_summary_contents_0), v10)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_pg_wal_summary_contents_1), int32(97), int32(_a_F_pg_wal_summary_contents_2))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
