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
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _consts[186])))
	if v8 == int32(1) {
		v13 = *(*int32)(unsafe.Add(mBase, _consts[179]))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+316))
		v16 = base.B2i32(v14 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v16)
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
						F_errmsg(m, int32(334693), int32(0))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(684513)
							F_errhint(m, int32(649225), v4)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(495145), int32(561), int32(374721))
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
				F_errmsg(m, int32(128025), int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					F_errhint(m, int32(573204), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(495145), int32(554), int32(374721))
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
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int64
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v130 int64
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
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
	v20 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v10+int32(1116)))) = uint16(v20)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+1112)) = v20
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
	if base.Ui64(int64(-2147483648)) < base.Ui64(v25-int64(2147483648)) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v10)+1104)) = uint32(v25)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v32 = *(*int64)(unsafe.Add(mBase, uint32(v31)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+1088)) = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v34)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+1080)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+1096)) = v35
	v41 = F_OpenWalSummaryFile(m, v10+int32(1088))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
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
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L45
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+1072)) = v41
	v47 = *(*int32)(unsafe.Add(mBase, _consts[266]))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47+v41*int32(48))+32))
	goto L7
L7:
	;
	v52 = F_CreateBlockRefTableReader(m, v10+int32(1072), v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v60 = F_BlockRefTableReaderNextRelation(m, v52, v10+int32(1060), v10+int32(1056), v10+int32(1052))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v60 != 0 {
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
	F_DestroyBlockRefTableReader(m, v52)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L43
	}
L13:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v70 != 0 {
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
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1068))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+1120)) = v73
	v75 = *(*int64)(unsafe.Add(mBase, uint32(v10)+1060))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+1124)) = v75
	v77 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10)+1056)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+1132)) = v77
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1052))
	if v79 != int32(-1) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	v83 = F_Int64GetDatum(m, base.I64_extend_i32_u(v79))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v10)+1136)) = v83
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v93 = F_heap_form_tuple(m, v88, v10+int32(1120), v10+int32(1112))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	F_tuplestore_puttuple(m, v95, v93)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v107 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v153 = F_BlockRefTableReaderNextRelation(m, v52, v10+int32(1060), v10+int32(1056), v10+int32(1052))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L41
	}
L27:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v113 = F_BlockRefTableReaderGetBlocks(m, v52, v10+int32(16), int32(256))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	if v113 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v115 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+1140)) = v115
	v118 = v115
	goto L35
L33:
	;
	goto L34
L34:
	;
	goto L26
L35:
	;
	v130 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v10+int32(16)+v118<<(uint(int32(2))%32)))))
	v131 = F_Int64GetDatum(m, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L37
	}
L36:
	;
	goto L25
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+1136)) = v131
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v139 = F_heap_form_tuple(m, v134, v10+int32(1120), v10+int32(1112))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	F_tuplestore_puttuple(m, v141, v139)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v145 = v118 + int32(1)
	if v145 != v113 {
		v118 = v145
		goto L35
	} else {
		goto L40
	}
L40:
	;
	goto L36
L41:
	;
	if v153 != 0 {
		goto L13
	} else {
		goto L42
	}
L42:
	;
	goto L14
L43:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1072))
	F_FileClose(m, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
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
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v25
	F_errmsg(m, int32(430391), v10)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(494788), int32(95), int32(120540))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
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
