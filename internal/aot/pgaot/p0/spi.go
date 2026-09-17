package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SPI_copytuple(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	if l0 == int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_SPI_copytuple[0])) = int32(-6)
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_copytuple[1]))
		if v12 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _c_F_SPI_copytuple[0])) = int32(-4)
			return int32(0)
		} else {
			v20 = int32(_a_F_SPI_copytuple_0)
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_copytuple[2]))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
			*(*int32)(unsafe.Add(mBase, _c_F_SPI_copytuple[2])) = v23
			v25 = F_heap_copytuple(m, l0)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_SPI_copytuple[2])) = v21
				return v25
			}
		}
	}
}
func F_SPI_cursor_fetch(m *base.Module, l0 int32, l1 int32) {
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v5 = F_CreateDestReceiver(m, int32(5))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		F__SPI_cursor_operation(m, l0, int32(0), l1, v5)
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	}
}
func F_SPI_execute_extended(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	if base.B2i32(l0 == v3)|base.B2i32(l1 == v3) != 0 {
		v71 = int32(-6)
		m.G0 = v8 + int32(48)
		return v71
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_extended[0]))
		if v17 == int32(0) {
			v71 = int32(-4)
			m.G0 = v8 + int32(48)
			return v71
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_extended[1]))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
			v25 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_extended[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v23
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
			*(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_extended[2])) = v28
			v30 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v8)+28)) = v30
			*(*int64)(unsafe.Add(mBase, uint32(v8)+12)) = v30
			*(*int64)(unsafe.Add(mBase, uint32(v8)+20)) = v30
			*(*int64)(unsafe.Add(mBase, uint32(v8)+36)) = v30
			*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(569278163)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = int32(2048)
			v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			if v44 != 0 {
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v45
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v47
			} else {
			}
			v50 = v8 + int32(8)
			F__SPI_prepare_oneshot_plan(m, l0, v50)
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				v55 = int32(0)
				v58 = F__SPI_execute_plan(m, v50, l1, v55, v55, int32(1))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					v62 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_extended[0]))
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+20))
					*(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_extended[2])) = v63
					*(*int32)(unsafe.Add(mBase, uint32(v62)+12)) = int32(0)
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v62)+24))
					F_MemoryContextReset(m, v67)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						v71 = v58
						m.G0 = v8 + int32(48)
						return v71
					}
				}
			}
		}
	}
}
func F_SPI_execute_plan_extended(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	v3 = int32(0)
	v5 = int32(-6)
	if base.B2i32(l0 == v3)|base.B2i32(l1 == v3) != 0 {
		v47 = v5
		return v47
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v11 != int32(569278163) {
			v47 = v5
			return v47
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_plan_extended[0]))
			if v15 == int32(0) {
				return int32(-4)
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_plan_extended[1]))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_plan_extended[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v22
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
				*(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_plan_extended[2])) = v27
				v29 = int32(0)
				v32 = F__SPI_execute_plan(m, l0, l1, v29, v29, int32(1))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_plan_extended[0]))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
					*(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_plan_extended[2])) = v39
					*(*int32)(unsafe.Add(mBase, uint32(v38)+12)) = int32(0)
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
					F_MemoryContextReset(m, v43)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						v47 = v32
						return v47
					}
				}
			}
		}
	}
}
func F_SPI_gettypeid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_SPI_gettypeid[0])) = v3
	if base.B2i32(l1 == v3)|base.B2i32(l1 < int32(-6)) == v3 {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if l1 <= v14 {
			if int32(0) < l1 {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0+v14<<(uint(int32(4))%32)+l1*int32(100)-int32(12))))
				return v32
			} else {
				v34 = F_SystemAttributeDefinition(m, l1)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+68))
					return v38
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_SPI_gettypeid[0])) = int32(-9)
			return int32(0)
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_SPI_gettypeid[0])) = int32(-9)
		return int32(0)
	}
}
func F__SPI_rollback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v236 int64
	_ = v236
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(192)
	m.G0 = v10
	v15 = v2
	v16 = v2
	v17 = v2
	v18 = v2
	v19 = int32(-1)
	goto L1
L1:
	;
	goto L4
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	goto L2
L4:
	;
	if v19 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v235 = int32(m.ExcTag)
	v236 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v235 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_rollback[0]))
	v25 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_rollback[1]))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+40)))
	if v26 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v119 = v15
	v120 = v16
	v121 = v17
	v122 = v18
	goto L9
L9:
	;
	if v122 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v23
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v23
	v61 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_rollback[2]))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+28))
	goto L17
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v23
	F_errcode(m, int32(1282))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v23
	F_errmsg(m, int32(_a_F__SPI_rollback_0), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v23
	F_errfinish(m, int32(_a_F__SPI_rollback_1), int32(341), int32(_a_F__SPI_rollback_2))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	goto L3
L17:
	;
	if int32(1) < v62 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v23
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if l0 != 0 {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v23
	F_errcode(m, int32(1282))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v23
	F_errmsg(m, int32(_a_F__SPI_rollback_3), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v23
	F_errfinish(m, int32(_a_F__SPI_rollback_1), int32(347), int32(_a_F__SPI_rollback_2))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L6
	} else {
		goto L24
	}
L24:
	;
	goto L3
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v23
	v97 = v10 + int32(172)
	v99 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_rollback[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v99
	v102 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__SPI_rollback[4])))
	*(*uint8)(unsafe.Add(mBase, uint32(v97)+4)) = uint8(v102)
	v105 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__SPI_rollback[5])))
	*(*uint8)(unsafe.Add(mBase, uint32(v97)+5)) = uint8(v105)
	goto L28
L26:
	;
	goto L27
L27:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_rollback[6]))
	v111 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_rollback[7]))
	goto L29
L28:
	;
	goto L27
L29:
	;
	v113 = v10 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v113)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = v10 + int32(12)
	goto L32
L30:
	;
	v119 = v23
	v120 = v111
	v121 = v109
	v122 = int32(0)
	goto L9
L32:
	;
	goto L30
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F__SPI_rollback[7])) = v10 + int32(16)
	v130 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_rollback[1]))
	v131 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v130)+41)) = uint8(v131)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v119
	F_HoldPinnedPortals(m)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L6
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, _c_F__SPI_rollback[6])) = v121
	*(*int32)(unsafe.Add(mBase, _c_F__SPI_rollback[7])) = v120
	*(*int32)(unsafe.Add(mBase, _c_F__SPI_rollback[0])) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v119
	v189 = F_CopyErrorData(m)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L6
	} else {
		goto L44
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v119
	F_ForgetPortalSnapshots(m)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v119
	F_AbortCurrentTransaction(m)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v119
	F_StartTransactionCommand(m)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	if l0 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v119
	v157 = v10 + int32(172)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	*(*int32)(unsafe.Add(mBase, _c_F__SPI_rollback[3])) = v159
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+4)))
	*(*uint8)(unsafe.Add(mBase, _c_F__SPI_rollback[4])) = uint8(v162)
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+5)))
	*(*uint8)(unsafe.Add(mBase, _c_F__SPI_rollback[5])) = uint8(v165)
	goto L43
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, _c_F__SPI_rollback[0])) = v119
	v170 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_rollback[1]))
	v171 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v170)+41)) = uint8(v171)
	*(*int32)(unsafe.Add(mBase, _c_F__SPI_rollback[6])) = v121
	*(*int32)(unsafe.Add(mBase, _c_F__SPI_rollback[7])) = v120
	m.G0 = v10 + int32(192)
	return
L43:
	;
	goto L42
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v119
	F_FlushErrorState(m)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v119
	F_AbortCurrentTransaction(m)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v119
	F_StartTransactionCommand(m)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	if l0 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v119
	v210 = v10 + int32(172)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	*(*int32)(unsafe.Add(mBase, _c_F__SPI_rollback[3])) = v212
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+4)))
	*(*uint8)(unsafe.Add(mBase, _c_F__SPI_rollback[4])) = uint8(v215)
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+5)))
	*(*uint8)(unsafe.Add(mBase, _c_F__SPI_rollback[5])) = uint8(v218)
	goto L51
L49:
	;
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, _c_F__SPI_rollback[0])) = v119
	v223 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_rollback[1]))
	v224 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v223)+41)) = uint8(v224)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v119
	F_ReThrowError(m, v189)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L6
	} else {
		goto L52
	}
L51:
	;
	goto L50
L52:
	;
	goto L5
L53:
	;
	v240 = int32(v236)
	m.G0 = v10
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v240)))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	if v10+int32(12) == v246 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	m.ExcPending = 1
	goto L62
L55:
	;
	if v250 != 0 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v243)+4))
	v250 = v248
	goto L58
L57:
	;
	v250 = int32(0)
	goto L58
L58:
	;
	goto L55
L59:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v10)+188))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v10)+184))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v10)+180))
	v15 = v251
	v16 = v252
	v17 = v253
	v18 = v242
	v19 = v250
	goto L1
L60:
	;
	goto L61
L61:
	;
	F___wasm_longjmp(m, v243, v242)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	return
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
