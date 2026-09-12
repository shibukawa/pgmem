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
		*(*int32)(unsafe.Add(mBase, _consts[351])) = int32(-6)
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _consts[346]))
		if v12 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _consts[351])) = int32(-4)
			return int32(0)
		} else {
			v20 = int32(4470400)
			v21 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v23
			v25 = F_heap_copytuple(m, l0)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v21
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = int32(-6)
	if l0 == int32(0) {
		v74 = v10
		m.G0 = v8 + int32(48)
		return v74
	} else {
		if l1 == int32(0) {
			v74 = v10
			m.G0 = v8 + int32(48)
			return v74
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _consts[346]))
			if v16 == int32(0) {
				v74 = int32(-4)
				m.G0 = v8 + int32(48)
				return v74
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, _consts[37]))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
				v24 = *(*int32)(unsafe.Add(mBase, _consts[346]))
				*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v22
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
				v28 = v8 + int32(28)
				v29 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v28))) = v29
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v26
				*(*int64)(unsafe.Add(mBase, uint32(v8)+20)) = v29
				*(*int64)(unsafe.Add(mBase, uint32(v8)+36)) = v29
				*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(2048)
				*(*int64)(unsafe.Add(mBase, uint32(v8)+12)) = v29
				*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(569278163)
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				if v45 != 0 {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v46
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v48
				} else {
				}
				F__SPI_prepare_oneshot_plan(m, l0, v8+int32(8))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					v58 = int32(0)
					v61 = F__SPI_execute_plan(m, v8+int32(8), l1, v58, v58, int32(1))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						v65 = *(*int32)(unsafe.Add(mBase, _consts[346]))
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v66
						*(*int32)(unsafe.Add(mBase, uint32(v65)+12)) = int32(0)
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v65)+24))
						F_MemoryContextReset(m, v70)
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							v74 = v61
							m.G0 = v8 + int32(48)
							return v74
						}
					}
				}
			}
		}
	}
}
func F_SPI_execute_plan_extended(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	v5 = int32(-6)
	if l0 == int32(0) {
		v46 = v5
		return v46
	} else {
		if l1 == int32(0) {
			v46 = v5
			return v46
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v10 != int32(569278163) {
				v46 = v5
				return v46
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, _consts[346]))
				if v14 == int32(0) {
					return int32(-4)
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, _consts[37]))
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
					v23 = *(*int32)(unsafe.Add(mBase, _consts[346]))
					*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v21
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v26
					v28 = int32(0)
					v31 = F__SPI_execute_plan(m, l0, l1, v28, v28, int32(1))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, _consts[346]))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v38
						*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = int32(0)
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
						F_MemoryContextReset(m, v42)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							v46 = v31
							return v46
						}
					}
				}
			}
		}
	}
}
func F_SPI_gettypeid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	*(*int32)(unsafe.Add(mBase, _consts[351])) = int32(0)
	if l1 < int32(-6) {
		*(*int32)(unsafe.Add(mBase, _consts[351])) = int32(-9)
		return int32(0)
	} else {
		if l1 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _consts[351])) = int32(-9)
			return int32(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if l1 <= v11 {
				if int32(0) < l1 {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0+v11<<(uint(int32(4))%32)+l1*int32(100)-int32(12))))
					return v29
				} else {
					v32 = F_SystemAttributeDefinition(m, base.I32_extend16_s(l1))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)+68))
						return v36
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[351])) = int32(-9)
				return int32(0)
			}
		}
	}
}
func F__SPI_rollback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int64
	_ = v286
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v18 = v2
	v19 = v2
	v20 = v2
	v21 = v2
	v22 = v2
	v23 = v2
	v24 = v13
	v25 = int32(-1)
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
	if v25 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v285 = int32(m.ExcTag)
	v286 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v285 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L7:
	;
	v29 = v24 - int32(16)
	m.G0 = v29
	v32 = v29 - int32(160)
	m.G0 = v32
	v35 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v37 = *(*int32)(unsafe.Add(mBase, _consts[346]))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+40)))
	if v38 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v147 = v18
	v148 = v19
	v149 = v20
	v150 = v21
	v151 = v22
	v152 = v23
	v153 = v24
	goto L9
L9:
	;
	if v152 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v29
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		v284 = v32
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v29
	v83 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+28))
	goto L17
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v29
	F_errcode(m, int32(1282))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		v284 = v32
		goto L6
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v29
	F_errmsg(m, int32(257157), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		v284 = v32
		goto L6
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v29
	F_errfinish(m, int32(488777), int32(341), int32(313095))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		v284 = v32
		goto L6
	} else {
		goto L16
	}
L16:
	;
	goto L3
L17:
	;
	if int32(1) < v84 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v29
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		v284 = v32
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
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v29
	F_errcode(m, int32(1282))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		v284 = v32
		goto L6
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v29
	F_errmsg(m, int32(336878), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		v284 = v32
		goto L6
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v29
	F_errfinish(m, int32(488777), int32(347), int32(313095))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		v284 = v32
		goto L6
	} else {
		goto L24
	}
L24:
	;
	goto L3
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v29
	v129 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v129
	v132 = int32(*(*uint8)(unsafe.Add(mBase, _consts[156])))
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+4)) = uint8(v132)
	v135 = int32(*(*uint8)(unsafe.Add(mBase, _consts[111])))
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+5)) = uint8(v135)
	goto L28
L26:
	;
	goto L27
L27:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	v141 = *(*int32)(unsafe.Add(mBase, _consts[294]))
	goto L29
L28:
	;
	goto L27
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v13 + int32(8)
	goto L32
L30:
	;
	v147 = v32
	v148 = v29
	v149 = v35
	v150 = v141
	v151 = v139
	v152 = int32(0)
	v153 = v32
	goto L9
L32:
	;
	goto L30
L33:
	;
	*(*int32)(unsafe.Add(mBase, _consts[294])) = v147
	v159 = *(*int32)(unsafe.Add(mBase, _consts[346]))
	v160 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v159)+41)) = uint8(v160)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v148
	F_HoldPinnedPortals(m)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		v284 = v153
		goto L6
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, _consts[77])) = v151
	*(*int32)(unsafe.Add(mBase, _consts[294])) = v150
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v148
	v228 = F_CopyErrorData(m)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		v284 = v153
		goto L6
	} else {
		goto L44
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v148
	F_ForgetPortalSnapshots(m)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		v284 = v153
		goto L6
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v148
	F_AbortCurrentTransaction(m)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		v284 = v153
		goto L6
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v148
	F_StartTransactionCommand(m)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		v284 = v153
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
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v148
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, _consts[110])) = v196
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+4)))
	*(*uint8)(unsafe.Add(mBase, _consts[156])) = uint8(v199)
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+5)))
	*(*uint8)(unsafe.Add(mBase, _consts[111])) = uint8(v202)
	goto L43
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v149
	v207 = *(*int32)(unsafe.Add(mBase, _consts[346]))
	v208 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v207)+41)) = uint8(v208)
	*(*int32)(unsafe.Add(mBase, _consts[77])) = v151
	*(*int32)(unsafe.Add(mBase, _consts[294])) = v150
	m.G0 = v13 + int32(32)
	return
L43:
	;
	goto L42
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v148
	F_FlushErrorState(m)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		v284 = v153
		goto L6
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v148
	F_AbortCurrentTransaction(m)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		v284 = v153
		goto L6
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v148
	F_StartTransactionCommand(m)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		v284 = v153
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
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v148
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, _consts[110])) = v257
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+4)))
	*(*uint8)(unsafe.Add(mBase, _consts[156])) = uint8(v260)
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+5)))
	*(*uint8)(unsafe.Add(mBase, _consts[111])) = uint8(v263)
	goto L51
L49:
	;
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v149
	v268 = *(*int32)(unsafe.Add(mBase, _consts[346]))
	v269 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v268)+41)) = uint8(v269)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v148
	F_ReThrowError(m, v228)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		v284 = v153
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
	v290 = int32(v286)
	m.G0 = v284
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v290)+4))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	if v13+int32(8) == v297 {
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
	if v300 != 0 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v293)+4))
	v300 = v299
	goto L58
L57:
	;
	v300 = int32(0)
	goto L58
L58:
	;
	goto L55
L59:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v18 = v302
	v19 = v301
	v20 = v303
	v21 = v304
	v22 = v305
	v23 = v292
	v24 = v284
	v25 = v300
	goto L1
L60:
	;
	goto L61
L61:
	;
	F___wasm_longjmp(m, v293, v292)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
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
