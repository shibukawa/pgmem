package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ParameterAclLookup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = F_convert_GUC_name_for_parameter_acl(m, l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = F_cstring_to_text(m, v11)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = int32(0)
			v20 = F_GetSysCacheOid(m, int32(43), v15, v17, v17, v17)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				if l1 != 0 {
					F_pfree(m, v11)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(16)
						return v20
					}
				} else {
					if v20 != 0 {
						F_pfree(m, v11)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(16)
							return v20
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(67137668))
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
								F_errmsg(m, int32(72385), v8)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(494374), int32(50), int32(231858))
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
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
func F_PostgresMainLongJmp(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v157 int32
	_ = v157
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	v1 = int32(0)
	v3 = int32(4481692)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v5 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v1
	*(*int32)(unsafe.Add(mBase, _consts[630])) = v1
	*(*int32)(unsafe.Add(mBase, _consts[842])) = v1
	*(*uint8)(unsafe.Add(mBase, _consts[843])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[844])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[845])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[846])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[847])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[848])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[849])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[850])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[851])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[852])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[853])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[854])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[855])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[856])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[857])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[858])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[859])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[860])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[861])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[862])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[863])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[864])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[865])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[866])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[867])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[868])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[869])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[870])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[871])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[872])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[873])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[874])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[875])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[876])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[877])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[878])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[879])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[880])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[881])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[882])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[883])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[884])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[885])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[886])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[887])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _consts[888])) = uint8(v1)
	goto L1
L1:
	;
	v157 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[764])) = v157
	*(*uint8)(unsafe.Add(mBase, _consts[889])) = uint8(v157)
	*(*uint8)(unsafe.Add(mBase, _consts[826])) = uint8(v157)
	*(*uint8)(unsafe.Add(mBase, _consts[890])) = uint8(v157)
	v169 = *(*int32)(unsafe.Add(mBase, _consts[701]))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	m.T0[v170].(func(*base.Module))(m)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return
L3:
	;
	F_EmitErrorReport(m)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(0)
	F_AbortCurrentTransaction(m)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, _consts[686])))
	if v181 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_LWLockReleaseAll(m)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L2
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v252 = m.G0
	v254 = v252 - int32(32)
	m.G0 = v254
	v259 = *(*int32)(unsafe.Add(mBase, _consts[891]))
	F_hash_seq_init(m, v254+int32(12), v259)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L2
	} else {
		goto L38
	}
L9:
	;
	goto L8
L10:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _consts[127]))
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = int32(0)
	F_pgaio_error_cleanup(m)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _consts[892]))
	if v195 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v206 = *(*int32)(unsafe.Add(mBase, _consts[667]))
	if v206 != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v195)+1168))
	if v198 < int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v195)+1168))
	v202 = F_close(m, v201)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v195)+1168)) = int32(-1)
	goto L16
L16:
	;
	goto L13
L17:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L2
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	F_ReplicationSlotCleanup(m, int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L2
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	v213 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[893])) = v213
	v216 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+24))
	goto L22
L22:
	;
	if base.B2i32(v217 != v213) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	F_ReleaseAuxProcessResources(m, int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L2
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v226 = *(*int32)(unsafe.Add(mBase, _consts[894]))
	if v226 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L25
L27:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L2
	} else {
		goto L37
	}
L28:
	;
	v228 = *(*int32)(unsafe.Add(mBase, _consts[708]))
	if v228 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v230 = *(*int32)(unsafe.Add(mBase, _consts[707]))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	if v231 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v230)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v230)+76)) = int32(1)
	if v232 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	goto L9
L33:
	;
	F_s_lock(m, v230+int32(76), int32(492410), int32(3869), int32(351475))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L2
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v242 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v230)+76)) = v242
	*(*int32)(unsafe.Add(mBase, uint32(v230)+4)) = v242
	goto L32
L36:
	;
	goto L35
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	v264 = F_hash_seq_search(m, v254+int32(12))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	if v264 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v266 = v264
	goto L43
L41:
	;
	goto L42
L42:
	;
	m.G0 = v254 + int32(32)
	v287 = *(*int32)(unsafe.Add(mBase, _consts[667]))
	if v287 != 0 {
		goto L51
	} else {
		goto L52
	}
L43:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v266)+64))
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268)+85)))
	if v269 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L42
L45:
	;
	v272 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v268)+84)) = uint8(v272)
	F_PortalDrop(m, v268, v272)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L2
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v279 = F_hash_seq_search(m, v254+int32(12))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L2
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	if v279 != 0 {
		v266 = v279
		goto L43
	} else {
		goto L50
	}
L50:
	;
	goto L44
L51:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L2
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	F_ReplicationSlotCleanup(m, int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L2
	} else {
		goto L55
	}
L54:
	;
	goto L53
L55:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, _consts[895])))
	if v294 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v296 = *(*int32)(unsafe.Add(mBase, _consts[896]))
	m.T0[v296].(func(*base.Module))(m)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L2
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v301 = *(*int32)(unsafe.Add(mBase, _consts[897]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v301
	F_FlushErrorState(m)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L2
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, _consts[898])))
	if v306 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v308 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[899])) = uint8(v308)
	goto L63
L62:
	;
	goto L63
L63:
	;
	v311 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[900])) = uint8(v311)
	v314 = int32(*(*uint8)(unsafe.Add(mBase, _consts[901])))
	if v314 == v311 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v317 = int32(4481692)
	v319 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v319 - int32(1)
	v324 = int32(*(*uint8)(unsafe.Add(mBase, _consts[899])))
	if v324 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L2
	} else {
		goto L70
	}
L67:
	;
	v328 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[902])) = uint8(v328)
	goto L69
L68:
	;
	goto L69
L69:
	;
	return
L70:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	F_errmsg(m, int32(67733), int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(491003), int32(5008), int32(236239))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_PrepareSortSupportFromIndexRel(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	v2 = l1
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+10)))
	if v12 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+84))
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v20
			F_errmsg_internal(m, int32(58914), v9)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				F_errfinish(m, int32(489764), int32(170), int32(306526))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v30 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+10)))
		v34 = v30<<(uint(int32(2))%32) - int32(4)
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v34+v35)))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v38+v34)))
		*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)) = uint8(v2)
		F_FinishSortSupportFunction(m, v37, v40, l2)
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return
		} else {
			m.G0 = v9 + int32(16)
			return
		}
	}
}
func F_PrepareSortSupportFromOrderingOp(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v14 = F_get_ordering_op_properties(m, l0, v6+int32(12), v6+int32(8), v6+int32(4))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		if v14 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(208253), v6)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					F_errfinish(m, int32(489764), int32(146), int32(237747))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)) = uint8(base.B2i32(v31 == int32(5)))
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			F_FinishSortSupportFunction(m, v35, v36, l1)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				m.G0 = v6 + int32(16)
				return
			}
		}
	}
}
func F_p_isURLPath(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v59 int32
	_ = v59
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	v2 = int32(0)
	v7 = F_palloc0(m, int32(48))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v13 + v15
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v18 - v20
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+16)) = uint8(v23)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v25 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v25 + v27<<(uint(int32(2))%32)
	goto L5
L4:
	;
	goto L5
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v32 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v32 + v34<<(uint(int32(2))%32)
	goto L8
L7:
	;
	goto L8
L8:
	;
	v40 = F_palloc(m, int32(32))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v42 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v40))) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v40)+24)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v40)+16)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v40)+8)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = int32(0)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
	v55 = F_palloc(m, int32(32))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v53 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v55)+24)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v55)+20)) = int32(57)
	F_check_stack_depth(m)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L15
	}
L12:
	;
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v53)))
	*(*int64)(unsafe.Add(mBase, uint32(v55))) = v57
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+16)) = v59
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v53)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v55)+8)) = v61
	goto L11
L13:
	;
	goto L14
L14:
	;
	v63 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v55))) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v55)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v55)+8)) = v63
	goto L11
L15:
	;
	v77 = F_TParserGet(m, v7)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
	if v111 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	if v77 == int32(0) {
		v110 = v2
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
	if v81 != int32(18) {
		v110 = v2
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v7)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = v85 + v86
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v90 + v91
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v7)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = v95 + v96
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+16))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v99)+16)) = v100 + v101
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v104)+8)) = v106
	v110 = int32(1)
	goto L16
L20:
	;
	v113 = v111
	goto L23
L21:
	;
	goto L22
L22:
	;
	F_pfree(m, v7)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L27
	}
L23:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113)+24))
	F_pfree(m, v113)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	goto L22
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v117
	if v117 != 0 {
		v113 = v117
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	return v110
}
func F_packArcInfoCmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v6 < v7 {
		return int32(-1)
	} else {
		v11 = int32(1)
		if v7 < v6 {
			v28 = v11
			return v28
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			if v13 < v14 {
				return int32(-1)
			} else {
				if v14 < v13 {
					v28 = v11
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					if v20 < v21 {
						v28 = int32(-1)
					} else {
						v28 = base.B2i32(v21 < v20)
					}
				}
				return v28
			}
		}
	}
}
func F_pagetable_delete(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var __phi52 int32
	_ = __phi52
	var v53 int32
	_ = v53
	var __phi53 int32
	_ = __phi53
	var v55 int32
	_ = v55
	var __phi55 int32
	_ = __phi55
	var v56 int32
	_ = v56
	var __phi56 int32
	_ = __phi56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	v8 = int32(16)
	v12 = (int32(base.Ui32(l1)>>(uint(v8)%32)) ^ l1) * int32(-2048144789)
	v17 = (int32(base.Ui32(v12)>>(uint(int32(13))%32)) ^ v12) * int32(-1028477387)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v25 = int32(base.Ui32(v17)>>(uint(v8)%32)) ^ v17
	goto L1
L1:
	;
	v30 = v25 & v22
	v33 = v21 + v30*int32(48)
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
	switch v34 {
	case 0:
		v112 = int32(0)
		goto L4
	case 1:
		goto L5
	default:
		goto L3
	}
L3:
	;
	v25 = v30 + int32(1)
	goto L1
L4:
	;
	return v112
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if v35 != l1 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v38 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v37 - v38
	v44 = v22 & (v30 + v38)
	v47 = v21 + v44*int32(48)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+4)))
	if v48 != v38 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v104 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v100)+4)) = uint8(v104)
	v112 = v38
	goto L4
L8:
	;
	v100 = v33
	goto L7
L9:
	;
	goto L10
L10:
	;
	__phi52 = v44
	__phi53 = v33
	__phi55 = v47
	__phi56 = v22
	v52 = __phi52
	v53 = __phi53
	v55 = __phi55
	v56 = __phi56
	goto L11
L11:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v59 = int32(16)
	v63 = (int32(base.Ui32(v58)>>(uint(v59)%32)) ^ v58) * int32(-2048144789)
	v68 = (int32(base.Ui32(v63)>>(uint(int32(13))%32)) ^ v63) * int32(-1028477387)
	if v52 == (int32(base.Ui32(v68)>>(uint(v59)%32))^v68)&v56 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v100 = v55
	goto L7
L13:
	;
	v100 = v53
	goto L7
L14:
	;
	goto L15
L15:
	;
	v74 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
	*(*int64)(unsafe.Add(mBase, uint32(v53))) = v74
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v55)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+40)) = v76
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v55)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+32)) = v78
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v55)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+24)) = v80
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v55)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+16)) = v82
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v55)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+8)) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v88 = int32(1)
	v90 = v87 & (v52 + v88)
	v93 = v86 + v90*int32(48)
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+4)))
	if v94 == v88 {
		__phi52 = v90
		__phi53 = v55
		__phi55 = v93
		__phi56 = v87
		v52 = __phi52
		v53 = __phi53
		v55 = __phi55
		v56 = __phi56
		goto L11
	} else {
		goto L16
	}
L16:
	;
	goto L12
}
func F_pagetable_insert(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v49 int64
	_ = v49
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v70 int64
	_ = v70
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int64
	_ = v106
	var v116 int64
	_ = v116
	var v124 int32
	_ = v124
	var v131 float64
	_ = v131
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v150 int32
	_ = v150
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int64
	_ = v249
	var v251 int64
	_ = v251
	var v253 int64
	_ = v253
	var v255 int64
	_ = v255
	var v257 int64
	_ = v257
	var v259 int64
	_ = v259
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v438 int64
	_ = v438
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int64
	_ = v491
	var v493 int64
	_ = v493
	var v495 int64
	_ = v495
	var v497 int64
	_ = v497
	var v499 int64
	_ = v499
	var v501 int64
	_ = v501
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v532 int64
	_ = v532
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v564 int32
	_ = v564
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v586 int32
	_ = v586
	var v595 int32
	_ = v595
	var v608 int32
	_ = v608
	var v614 int32
	_ = v614
	v17 = int32(16)
	v21 = (int32(base.Ui32(l1)>>(uint(v17)%32)) ^ l1) * int32(-2048144789)
	v26 = (int32(base.Ui32(v21)>>(uint(int32(13))%32)) ^ v21) * int32(-1028477387)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v35 = v30
	v42 = v31
	goto L1
L1:
	;
	if base.Ui32(v35) <= base.Ui32(v42) {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	v614 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v614
	v35 = v614
	v42 = v608
	goto L1
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v586))) = l1
	v595 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v595)
	return v586
L5:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v573 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v572 + v573
	*(*uint8)(unsafe.Add(mBase, uint32(v564)+4)) = uint8(v573)
	v586 = v564
	goto L4
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L26
	} else {
		goto L109
	}
L7:
	;
	v49 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v49 == int64(4294967296) {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v357 = int32(0)
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v360 = (int32(base.Ui32(v26)>>(uint(v17)%32)) ^ v26) & v359
	v363 = v358 + v360*int32(48)
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363)+4)))
	if v364 == v357 {
		v564 = v363
		goto L5
	} else {
		goto L76
	}
L10:
	;
	v52 = int32(0)
	v54 = int64(2)
	v56 = v49 << (uint(int64(1)) % 64)
	if base.Ui64(v56) <= base.Ui64(v54) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L9
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L26
	} else {
		goto L73
	}
L13:
	;
	v59 = v54
	goto L15
L14:
	;
	v59 = v56
	goto L15
L15:
	;
	v60 = int64(1)
	if v59&(v59-v60) == int64(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v70 = v59
	goto L18
L17:
	;
	v70 = v60 << (uint(int64(64)-base.I64_clz(v59)) % 64)
	goto L18
L18:
	;
	if base.Ui64(v70*int64(48)) < base.Ui64(int64(2147483647)) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v76 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v79 = base.I32_wrap_i64(v70) * int32(48)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+112))
	if v81 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	goto L21
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L26
	} else {
		goto L70
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v104
	v106 = int64(1)
	if v70&(v70-v106) == int64(0) {
		goto L30
	} else {
		goto L31
	}
L23:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v86 = F_MemoryContextAllocExtended(m, v84, v79, int32(5))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v80)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+100)) = v90
	v95 = F_dsa_allocate_extended(m, v81, v79|int32(4), int32(5))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L26
	} else {
		goto L28
	}
L26:
	;
	return int32(0)
L27:
	;
	v104 = v86
	goto L22
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+96)) = v95
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v80)+112))
	v99 = F_dsa_get_address(m, v98, v95)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v104 = v99 + int32(4)
	goto L22
L30:
	;
	v116 = v70
	goto L32
L31:
	;
	v116 = v106 << (uint(int64(64)-base.I64_clz(v70)) % 64)
	goto L32
L32:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v116*int64(48)) {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v116
	v124 = base.I32_wrap_i64(v116) - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v124
	v131 = base.F64_mul(base.F64_convert_i64_u(v116), float64(0.9))
	if base.F64_lt(v131, float64(4.294967296e+09))&base.F64_ge(v131, float64(0)) != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v116 == int64(4294967296) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v137 = base.I32_trunc_f64_u(v131)
	v139 = v137
	goto L34
L36:
	;
	goto L37
L37:
	;
	v139 = int32(0)
	goto L34
L38:
	;
	v140 = int32(-85899346)
	goto L40
L39:
	;
	v140 = v139
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v140
	if v76 != int64(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v150 = v52
	goto L45
L42:
	;
	goto L43
L43:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v303)+112))
	if v304 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L44:
	;
	v195 = v188
	v198 = v52
	goto L50
L45:
	;
	v162 = v75 + v150*int32(48)
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+4)))
	if v163 != int32(1) {
		v188 = v150
		goto L44
	} else {
		goto L47
	}
L46:
	;
	v188 = int32(0)
	goto L44
L47:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v167 = int32(16)
	v171 = (int32(base.Ui32(v166)>>(uint(v167)%32)) ^ v166) * int32(-2048144789)
	v176 = (int32(base.Ui32(v171)>>(uint(int32(13))%32)) ^ v171) * int32(-1028477387)
	if (int32(base.Ui32(v176)>>(uint(v167)%32))^v176)&v124 == v150 {
		v188 = v150
		goto L44
	} else {
		goto L48
	}
L48:
	;
	v183 = v150 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v183)) < base.Ui64(v76) {
		v150 = v183
		goto L45
	} else {
		goto L49
	}
L49:
	;
	goto L46
L50:
	;
	v207 = v75 + v195*int32(48)
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+4)))
	if v208 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L43
L52:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
	v212 = int32(16)
	v216 = (int32(base.Ui32(v211)>>(uint(v212)%32)) ^ v211) * int32(-2048144789)
	v221 = (int32(base.Ui32(v216)>>(uint(int32(13))%32)) ^ v216) * int32(-1028477387)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v229 = int32(base.Ui32(v221)>>(uint(v212)%32)) ^ v221
	goto L55
L53:
	;
	goto L54
L54:
	;
	v278 = v195 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v278)) < base.Ui64(v76) {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v242 = v229 & v225
	v247 = v104 + v242*int32(48)
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247)+4)))
	if v248 != 0 {
		v229 = v242 + int32(1)
		goto L55
	} else {
		goto L57
	}
L56:
	;
	v249 = *(*int64)(unsafe.Add(mBase, uint32(v207)))
	*(*int64)(unsafe.Add(mBase, uint32(v247))) = v249
	v251 = *(*int64)(unsafe.Add(mBase, uint32(v207)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v247)+40)) = v251
	v253 = *(*int64)(unsafe.Add(mBase, uint32(v207)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v247)+32)) = v253
	v255 = *(*int64)(unsafe.Add(mBase, uint32(v207)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v247)+24)) = v255
	v257 = *(*int64)(unsafe.Add(mBase, uint32(v207)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v247)+16)) = v257
	v259 = *(*int64)(unsafe.Add(mBase, uint32(v207)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v247)+8)) = v259
	goto L54
L57:
	;
	goto L56
L58:
	;
	v282 = v278
	goto L60
L59:
	;
	v282 = int32(0)
	goto L60
L60:
	;
	v284 = v198 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v284)) < base.Ui64(v76) {
		v195 = v282
		v198 = v284
		goto L50
	} else {
		goto L61
	}
L61:
	;
	goto L51
L62:
	;
	F_pfree(m, v75)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L26
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v303)+100))
	if v309 != 0 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	goto L11
L66:
	;
	F_dsa_free(m, v304, v309)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L26
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	goto L11
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v303)+100)) = int32(0)
	goto L68
L70:
	;
	F_errmsg_internal(m, int32(398249), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L26
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(324785), int32(327), int32(339174))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L26
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	F_errmsg_internal(m, int32(398249), int32(0))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L26
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(324785), int32(327), int32(339174))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L26
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	v370 = v360
	v371 = v357
	v375 = v363
	goto L77
L77:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v375)))
	if l1 == v383 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v564 = v541
	goto L5
L79:
	;
	v385 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v385)
	return v375
L80:
	;
	goto L81
L81:
	;
	v389 = v370 + int32(1)
	v390 = int32(16)
	v394 = (int32(base.Ui32(v383)>>(uint(v390)%32)) ^ v383) * int32(-2048144789)
	v399 = (int32(base.Ui32(v394)>>(uint(int32(13))%32)) ^ v394) * int32(-1028477387)
	v403 = (int32(base.Ui32(v399)>>(uint(v390)%32)) ^ v399) & v359
	if base.Ui32(v370) < base.Ui32(v403) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v407 = v370 + v405
	goto L84
L83:
	;
	v407 = v370
	goto L84
L84:
	;
	if base.Ui32(v407-v403) < base.Ui32(v371) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v411 = v359 & v389
	v414 = v358 + v411*int32(48)
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v414)+4)))
	if v415 != 0 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	v527 = v371 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v527) {
		goto L104
	} else {
		goto L105
	}
L88:
	;
	v421 = v411
	v423 = int32(0)
	goto L91
L89:
	;
	v455 = v414
	v456 = v411
	goto L90
L90:
	;
	if v370 != v456 {
		goto L98
	} else {
		goto L99
	}
L91:
	;
	v433 = v423 + int32(1)
	if int32(151) <= v433 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v455 = v449
	v456 = v446
	goto L90
L93:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v438 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v436), base.F64_convert_i64_u(v438)), float64(0.1)) != 0 {
		v608 = v436
		goto L3
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v446 = (v421 + int32(1)) & v359
	v449 = v358 + v446*int32(48)
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v449)+4)))
	if v450 != 0 {
		v421 = v446
		v423 = v433
		goto L91
	} else {
		goto L97
	}
L96:
	;
	goto L95
L97:
	;
	goto L92
L98:
	;
	v472 = v455
	v473 = v456
	goto L101
L99:
	;
	goto L100
L100:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v521 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v520 + v521
	*(*uint8)(unsafe.Add(mBase, uint32(v375)+4)) = uint8(v521)
	v586 = v375
	goto L4
L101:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v487 = v484 & (v473 - int32(1))
	v490 = v358 + v487*int32(48)
	v491 = *(*int64)(unsafe.Add(mBase, uint32(v490)))
	*(*int64)(unsafe.Add(mBase, uint32(v472))) = v491
	v493 = *(*int64)(unsafe.Add(mBase, uint32(v490)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v472)+40)) = v493
	v495 = *(*int64)(unsafe.Add(mBase, uint32(v490)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v472)+32)) = v495
	v497 = *(*int64)(unsafe.Add(mBase, uint32(v490)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v472)+24)) = v497
	v499 = *(*int64)(unsafe.Add(mBase, uint32(v490)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v472)+16)) = v499
	v501 = *(*int64)(unsafe.Add(mBase, uint32(v490)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v472)+8)) = v501
	if v370 != v487 {
		v472 = v490
		v473 = v487
		goto L101
	} else {
		goto L103
	}
L102:
	;
	goto L100
L103:
	;
	goto L102
L104:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v532 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v530), base.F64_convert_i64_u(v532)), float64(0.1)) != 0 {
		v608 = v530
		goto L3
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v538 = v359 & v389
	v541 = v358 + v538*int32(48)
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541)+4)))
	if v542 != 0 {
		v370 = v538
		v371 = v527
		v375 = v541
		goto L77
	} else {
		goto L108
	}
L107:
	;
	goto L106
L108:
	;
	goto L78
L109:
	;
	F_errmsg_internal(m, int32(459814), int32(0))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L26
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(324785), int32(630), int32(310110))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L26
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pairingheap_SpGistSearchItem_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 float64
	_ = v42
	var v48 int64
	_ = v48
	var v53 float64
	_ = v53
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v108 int32
	_ = v108
	v11 = int32(1)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+34)))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v13 == v11 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return int32(0)
L2:
	;
	return v108
L3:
	;
	v108 = int32(-1)
	goto L2
L4:
	;
	v81 = int32(1)
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+35)))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+35)))
	if v83 == v81 {
		goto L26
	} else {
		goto L27
	}
L5:
	;
	if v12&int32(1) != 0 {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	if v12&int32(1) != 0 {
		v108 = v11
		goto L2
	} else {
		goto L9
	}
L8:
	;
	goto L3
L9:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+112))
	if v20 <= int32(0) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v23 = int32(40)
	v31 = int32(0)
	goto L11
L11:
	;
	v39 = v31 << (uint(int32(3)) % 32)
	v40 = l1 + v23 + v39
	v42 = *(*float64)(unsafe.Add(mBase, uint32(v39+(l0+v23))))
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v42)&int64(9223372036854775807)) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L4
L13:
	;
	v69 = v31 + int32(1)
	if v69 != v20 {
		v31 = v69
		goto L11
	} else {
		goto L25
	}
L14:
	;
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v40)))
	if base.Ui64(v48&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L3
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v53 = *(*float64)(unsafe.Add(mBase, uint32(v40)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v53)&int64(9223372036854775807)) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L13
L18:
	;
	return int32(1)
L19:
	;
	goto L20
L20:
	;
	if base.F64_eq(v42, v53) != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	if base.F64_lt(v42, v53) != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v65 = int32(1)
	goto L24
L23:
	;
	v65 = int32(-1)
	goto L24
L24:
	;
	return v65
L25:
	;
	goto L12
L26:
	;
	if v82&int32(1) == int32(0) {
		v108 = v81
		goto L2
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	if v82&int32(1) == int32(0) {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	goto L1
L30:
	;
	goto L3
}
func F_palloc_extended(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+4)) = uint8(v3)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v10 = m.T0[v9].(func(*base.Module, int32, int32, int32) int32)(m, v5, l0, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if l1&int32(4) == int32(0) {
			return v10
		} else {
			if v10 == int32(0) {
				return v10
			} else {
				if base.Ui32(int32(1024)) < base.Ui32(l0) {
					v43 = F__emscripten_memset_bulkmem(m, v10, base.I32_extend8_s(int32(0)), l0)
					mBase = m.M
					return v10
				} else {
					if l0&int32(3) != 0 {
						v43 = F__emscripten_memset_bulkmem(m, v10, base.I32_extend8_s(int32(0)), l0)
						mBase = m.M
						return v10
					} else {
						v24 = l0 + v10
						if base.Ui32(v24) <= base.Ui32(v10) {
							return v10
						} else {
							v30 = v10 + int32(4)
							if base.Ui32(v30) < base.Ui32(v24) {
								v32 = v24
							} else {
								v32 = v30
							}
							v39 = F__emscripten_memset_bulkmem(m, v10, base.I32_extend8_s(int32(0)), (v10^int32(-1)+v32)&int32(-4)+int32(4))
							mBase = m.M
							return v39
						}
					}
				}
			}
		}
	}
}
func F_parseXidFromText(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v11 = F_strlen(m, l0)
	mBase = m.M
	if v11 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L20
	} else {
		goto L36
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L20
	} else {
		goto L32
	}
L3:
	;
	if v55 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	v55 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v17 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v18 = v10
	v19 = l0
	v20 = v11
	v21 = v17
	goto L11
L8:
	;
	v43 = l0
	v47 = int32(0)
	goto L9
L9:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	v55 = v47 - v48
	goto L3
L10:
	;
	v43 = v38
	v47 = v40
	goto L9
L11:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v21 != v23 {
		v38 = v19
		v40 = v21
		goto L10
	} else {
		goto L13
	}
L12:
	;
	v38 = v32
	v40 = int32(0)
	goto L10
L13:
	;
	if v23 == int32(0) {
		v38 = v19
		v40 = v21
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v28 = v20 - int32(1)
	if v28 == int32(0) {
		v38 = v19
		v40 = v21
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v31 = int32(1)
	v32 = v19 + v31
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v33 != 0 {
		v18 = v18 + v31
		v19 = v32
		v20 = v28
		v21 = v33
		goto L11
	} else {
		goto L16
	}
L16:
	;
	goto L12
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v6 + int32(-4)
	v61 = v11 + v10
	v65 = F_sscanf(m, v61, int32(59295), v6+int32(-32))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L20
	} else {
		goto L28
	}
L20:
	;
	return int32(0)
L21:
	;
	if v65 != int32(1) {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v71 = int32(10)
	v72 = F___strchrnul(m, v61, v71)
	mBase = m.M
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v74 == v71 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v78 == int32(0) {
		goto L1
	} else {
		goto L27
	}
L24:
	;
	v78 = v72
	goto L26
L25:
	;
	v78 = int32(0)
	goto L26
L26:
	;
	goto L23
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v78 + int32(1)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v8)+60))
	m.G0 = v8 - int32(-64)
	return v84
L28:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L20
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = l2
	F_errmsg(m, int32(698005), v6+int32(-16))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L20
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(492015), int32(1339), int32(63718))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L20
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
	F_errcode(m, int32(33685634))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L20
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l2
	F_errmsg(m, int32(698005), v6+int32(-48))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L20
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(492015), int32(1344), int32(63718))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L20
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L20
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l2
	F_errmsg(m, int32(698005), v8)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L20
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(492015), int32(1349), int32(63718))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L20
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_parse_format(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v184 int32
	_ = v184
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
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v415 int32
	_ = v415
	var v426 int32
	_ = v426
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v601 int32
	_ = v601
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v669 int32
	_ = v669
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v692 int32
	_ = v692
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v820 int32
	_ = v820
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v836 int32
	_ = v836
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v852 int32
	_ = v852
	var v857 int32
	_ = v857
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v873 int32
	_ = v873
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v900 int32
	_ = v900
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v916 int32
	_ = v916
	var v921 int32
	_ = v921
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v932 int32
	_ = v932
	var v937 int32
	_ = v937
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v952 int32
	_ = v952
	var v957 int32
	_ = v957
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v968 int32
	_ = v968
	var v972 int32
	_ = v972
	var v977 int32
	_ = v977
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v20 == int32(0) {
		v724 = l0
		goto L15
	} else {
		goto L16
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L84
	} else {
		goto L283
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L84
	} else {
		goto L278
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L84
	} else {
		goto L274
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L84
	} else {
		goto L270
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L84
	} else {
		goto L266
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L84
	} else {
		goto L262
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L84
	} else {
		goto L258
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L84
	} else {
		goto L254
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L84
	} else {
		goto L250
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L84
	} else {
		goto L246
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L84
	} else {
		goto L242
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L84
	} else {
		goto L238
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L84
	} else {
		goto L234
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L84
	} else {
		goto L230
	}
L15:
	;
	v739 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v724)+6)) = uint8(v739)
	v741 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v724))) = uint8(v741)
	m.G0 = v18 + int32(16)
	return
L16:
	;
	v28 = l5 & int32(1)
	v31 = l0
	v32 = l1
	v38 = v20
	goto L17
L17:
	;
	if v28 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v724 = v708
	goto L15
L19:
	;
	v723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v709))))
	if v723 != 0 {
		v31 = v708
		v32 = v709
		v38 = v723
		goto L17
	} else {
		goto L229
	}
L20:
	;
	if base.Ui32((v143-int32(126))&int32(255)) < base.Ui32(int32(163)) {
		goto L55
	} else {
		goto L56
	}
L21:
	;
	v137 = v32
	v143 = v38
	v147 = int32(0)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v49 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v137 = v32
	v143 = v38
	v147 = int32(0)
	goto L20
L25:
	;
	goto L26
L26:
	;
	v57 = v49
	v58 = l3
	goto L29
L27:
	;
	if v129&int32(255) == int32(0) {
		v708 = v31
		v709 = v126
		goto L19
	} else {
		goto L50
	}
L28:
	;
	v123 = v32 + v71
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	v126 = v123
	v129 = v124
	v131 = v125
	goto L27
L29:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	if v68 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v126 = v32
	v129 = v38
	v131 = int32(0)
	goto L27
L31:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v71 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	goto L33
L33:
	;
	v120 = v58 + int32(16)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	if v121 != 0 {
		v57 = v121
		v58 = v120
		goto L29
	} else {
		goto L49
	}
L34:
	;
	if v115 == int32(0) {
		goto L28
	} else {
		goto L48
	}
L35:
	;
	v115 = int32(0)
	goto L34
L36:
	;
	goto L37
L37:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v77 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v78 = v32
	v79 = v57
	v80 = v71
	v81 = v77
	goto L42
L39:
	;
	v103 = v57
	v107 = int32(0)
	goto L40
L40:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	v115 = v107 - v108
	goto L34
L41:
	;
	v103 = v98
	v107 = v100
	goto L40
L42:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v81 != v83 {
		v98 = v79
		v100 = v81
		goto L41
	} else {
		goto L44
	}
L43:
	;
	v98 = v92
	v100 = int32(0)
	goto L41
L44:
	;
	if v83 == int32(0) {
		v98 = v79
		v100 = v81
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v88 = v80 - int32(1)
	if v88 == int32(0) {
		v98 = v79
		v100 = v81
		goto L41
	} else {
		goto L46
	}
L46:
	;
	v91 = int32(1)
	v92 = v79 + v91
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
	if v93 != 0 {
		v78 = v78 + v91
		v79 = v92
		v80 = v88
		v81 = v93
		goto L42
	} else {
		goto L47
	}
L47:
	;
	goto L43
L48:
	;
	goto L33
L49:
	;
	goto L30
L50:
	;
	v137 = v126
	v143 = v129
	v147 = v131
	goto L20
L51:
	;
	v708 = v31 + int32(12)
	v709 = v692
	goto L19
L52:
	;
	if v255 != int32(92) {
		goto L193
	} else {
		goto L194
	}
L53:
	;
	if v255 == int32(32) {
		goto L189
	} else {
		goto L190
	}
L54:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+6)) = uint8(v147)
	v290 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v290)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v174
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
	if l5&int32(2) == int32(0) {
		goto L91
	} else {
		goto L92
	}
L55:
	;
	v253 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v253
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	if v255 == v253 {
		v724 = v31
		goto L15
	} else {
		goto L77
	}
L56:
	;
	v158 = v143 & int32(255)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l4-int32(128)+v158<<(uint(int32(2))%32))))
	if v162 < int32(0) {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v167 = l2 + v162*int32(20)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v173 = v168
	v174 = v167
	goto L58
L58:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
	if v184 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L55
L60:
	;
	if v228 == int32(0) {
		goto L54
	} else {
		goto L74
	}
L61:
	;
	v228 = int32(0)
	goto L60
L62:
	;
	goto L63
L63:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	if v190 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v191 = v137
	v192 = v173
	v193 = v184
	v194 = v190
	goto L68
L65:
	;
	v216 = v173
	v220 = int32(0)
	goto L66
L66:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216))))
	v228 = v220 - v221
	goto L60
L67:
	;
	v216 = v211
	v220 = v213
	goto L66
L68:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v194 != v196 {
		v211 = v192
		v213 = v194
		goto L67
	} else {
		goto L70
	}
L69:
	;
	v211 = v205
	v213 = int32(0)
	goto L67
L70:
	;
	if v196 == int32(0) {
		v211 = v192
		v213 = v194
		goto L67
	} else {
		goto L71
	}
L71:
	;
	v201 = v193 - int32(1)
	if v201 == int32(0) {
		v211 = v192
		v213 = v194
		goto L67
	} else {
		goto L72
	}
L72:
	;
	v204 = int32(1)
	v205 = v192 + v204
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+1)))
	if v206 != 0 {
		v191 = v191 + v204
		v192 = v205
		v193 = v201
		v194 = v206
		goto L68
	} else {
		goto L73
	}
L73:
	;
	goto L69
L74:
	;
	v232 = v174 + int32(20)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)))
	if v233 == int32(0) {
		goto L55
	} else {
		goto L75
	}
L75:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
	if v158 == v236 {
		v173 = v233
		v174 = v232
		goto L58
	} else {
		goto L76
	}
L76:
	;
	goto L59
L77:
	;
	if base.Ui32(l5) < base.Ui32(int32(4)) {
		goto L52
	} else {
		goto L78
	}
L78:
	;
	if v255 == int32(34) {
		goto L52
	} else {
		goto L79
	}
L79:
	;
	if base.Ui32(v255) <= base.Ui32(int32(63)) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v255))%64)&int64(864955565296582657) != int64(0) {
		goto L53
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	goto L82
L84:
	;
	return
L85:
	;
	F_errcode(m, int32(117440642))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v276 = F_pg_mblen_cstr(m, v137)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L84
	} else {
		goto L87
	}
L87:
	;
	v278 = F_pnstrdup(m, v137, v276)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L84
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v278
	F_errmsg(m, int32(706778), v18)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L84
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(495035), int32(1446), int32(111417))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L84
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	v481 = v137 + v293
	if v28 == int32(0) {
		v692 = v481
		goto L51
	} else {
		goto L164
	}
L92:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v174)+8))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	if v297&int32(16384) != 0 {
		goto L97
	} else {
		goto L98
	}
L93:
	;
	if v469&int32(1024) == int32(0) {
		goto L91
	} else {
		goto L162
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6)+12)) = v467
	v469 = v467
	goto L93
L95:
	;
	if v297&int32(4080) != 0 {
		goto L2
	} else {
		goto L161
	}
L96:
	;
	F_errmsg(m, int32(415483), int32(0))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L84
	} else {
		goto L159
	}
L97:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L84
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	switch v296 - int32(1) {
	case 0:
		v380 = v297
		goto L114
	case 1:
		goto L117
	case 2:
		goto L118
	case 3:
		goto L116
	default:
		v469 = v297
		goto L93
	case 5:
		goto L115
	case 6:
		goto L95
	case 7:
		goto L113
	case 8, 9:
		goto L106
	case 10:
		goto L111
	case 11:
		goto L110
	case 12:
		goto L108
	case 13, 29:
		goto L107
	case 14:
		goto L109
	case 16:
		goto L112
	case 18:
		goto L105
	}
L100:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L84
	} else {
		goto L101
	}
L101:
	;
	if v296 == int32(7) {
		goto L96
	} else {
		goto L102
	}
L102:
	;
	F_errmsg(m, int32(446997), int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L84
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(495035), int32(1197), int32(363160))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L84
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	if v297&int32(2) != 0 {
		goto L3
	} else {
		goto L158
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6)+32)) = int32(1)
	v469 = v297
	goto L93
L107:
	;
	if v297&int32(1024) != 0 {
		goto L4
	} else {
		goto L157
	}
L108:
	;
	if v297&int32(832) != 0 {
		goto L5
	} else {
		goto L156
	}
L109:
	;
	if v297&int32(64) != 0 {
		goto L6
	} else {
		goto L155
	}
L110:
	;
	if v297&int32(64) != 0 {
		goto L7
	} else {
		goto L153
	}
L111:
	;
	if v297&int32(64) != 0 {
		goto L8
	} else {
		goto L151
	}
L112:
	;
	if v297&int32(64) != 0 {
		goto L10
	} else {
		goto L143
	}
L113:
	;
	v467 = v297 | int32(32)
	goto L94
L114:
	;
	if v380&int32(2) != 0 {
		goto L12
	} else {
		goto L141
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6)+32)) = int32(1)
	v378 = v297 | int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l6)+12)) = v378
	v380 = v378
	goto L114
L116:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v368 != 0 {
		goto L135
	} else {
		goto L136
	}
L117:
	;
	if v297&int32(128) != 0 {
		goto L13
	} else {
		goto L126
	}
L118:
	;
	if v297&int32(128) != 0 {
		goto L14
	} else {
		goto L119
	}
L119:
	;
	if v297&int32(2048) != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l6)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l6)+20)) = v324 + int32(1)
	v469 = v297
	goto L93
L121:
	;
	goto L122
L122:
	;
	if v297&int32(2) != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l6)+4)) = v330 + int32(1)
	v469 = v297
	goto L93
L124:
	;
	goto L125
L125:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v334 + int32(1)
	v469 = v297
	goto L93
L126:
	;
	if v297&int32(10) != 0 {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	if v349&int32(2) == int32(0) {
		goto L132
	} else {
		goto L133
	}
L128:
	;
	v349 = v297
	goto L127
L129:
	;
	goto L130
L130:
	;
	v343 = v297 | int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l6)+12)) = v343
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	*(*int32)(unsafe.Add(mBase, uint32(l6)+24)) = v345 + int32(1)
	v349 = v343
	goto L127
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6)+28)) = v364 + v365
	v469 = v349
	goto L93
L132:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v356 = v354 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v356
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v364 = v356
	v365 = v358
	goto L131
L133:
	;
	goto L134
L134:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v361 = v359 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l6)+4)) = v361
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v364 = v363
	v365 = v361
	goto L131
L135:
	;
	v469 = v297
	goto L93
L136:
	;
	goto L137
L137:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v369|v297&int32(8) != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v469 = v297
	goto L93
L139:
	;
	goto L140
L140:
	;
	v467 = v297 | int32(16)
	goto L94
L141:
	;
	if v380&int32(2048) != 0 {
		goto L11
	} else {
		goto L142
	}
L142:
	;
	v467 = v380 | int32(2)
	goto L94
L143:
	;
	if v297&int32(896) != 0 {
		goto L9
	} else {
		goto L144
	}
L144:
	;
	if v297&int32(2) == int32(0) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6)+8)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l6)+32)) = int32(1)
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	*(*int32)(unsafe.Add(mBase, uint32(l6)+16)) = v401
	v467 = v297 | int32(64)
	goto L94
L146:
	;
	goto L147
L147:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	if v405 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v469 = v297
	goto L93
L149:
	;
	goto L150
L150:
	;
	v406 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l6)+32)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(l6)+8)) = v406
	v467 = v297 | int32(64)
	goto L94
L151:
	;
	v415 = v297 | int32(256)
	*(*int32)(unsafe.Add(mBase, uint32(l6)+12)) = v415
	if v297&int32(2) == int32(0) {
		v469 = v415
		goto L93
	} else {
		goto L152
	}
L152:
	;
	v467 = v297 | int32(8448)
	goto L94
L153:
	;
	v426 = v297 | int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(l6)+12)) = v426
	if v297&int32(2) == int32(0) {
		v469 = v426
		goto L93
	} else {
		goto L154
	}
L154:
	;
	v467 = v297 | int32(4608)
	goto L94
L155:
	;
	v467 = v297 | int32(768)
	goto L94
L156:
	;
	v467 = v297 | int32(128)
	goto L94
L157:
	;
	v467 = v297 | int32(1024)
	goto L94
L158:
	;
	v467 = v297 | int32(2048)
	goto L94
L159:
	;
	F_errfinish(m, int32(495035), int32(1347), int32(363160))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L84
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L161:
	;
	v467 = v297 | int32(16384)
	goto L94
L162:
	;
	if v469&int32(-1057) != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	goto L91
L164:
	;
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481))))
	if v484 == int32(0) {
		v692 = v481
		goto L51
	} else {
		goto L165
	}
L165:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v487 == int32(0) {
		v692 = v481
		goto L51
	} else {
		goto L166
	}
L166:
	;
	v494 = v487
	v495 = l3
	goto L167
L167:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v495)+12))
	if v505 == int32(2) {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+6)))
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v495)+8)))
	v560 = v558 | v559
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+6)) = uint8(v560)
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v495)+4))
	v692 = v481 + v562
	goto L51
L169:
	;
	goto L168
L170:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v495)+4))
	if v508 == int32(0) {
		goto L174
	} else {
		goto L175
	}
L171:
	;
	goto L172
L172:
	;
	v556 = v495 + int32(16)
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v556)))
	if v557 != 0 {
		v494 = v557
		v495 = v556
		goto L167
	} else {
		goto L188
	}
L173:
	;
	if v552 == int32(0) {
		goto L169
	} else {
		goto L187
	}
L174:
	;
	v552 = int32(0)
	goto L173
L175:
	;
	goto L176
L176:
	;
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481))))
	if v514 != 0 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v515 = v481
	v516 = v494
	v517 = v508
	v518 = v514
	goto L181
L178:
	;
	v540 = v494
	v544 = int32(0)
	goto L179
L179:
	;
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v540))))
	v552 = v544 - v545
	goto L173
L180:
	;
	v540 = v535
	v544 = v537
	goto L179
L181:
	;
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
	if v518 != v520 {
		v535 = v516
		v537 = v518
		goto L180
	} else {
		goto L183
	}
L182:
	;
	v535 = v529
	v537 = int32(0)
	goto L180
L183:
	;
	if v520 == int32(0) {
		v535 = v516
		v537 = v518
		goto L180
	} else {
		goto L184
	}
L184:
	;
	v525 = v517 - int32(1)
	if v525 == int32(0) {
		v535 = v516
		v537 = v518
		goto L180
	} else {
		goto L185
	}
L185:
	;
	v528 = int32(1)
	v529 = v516 + v528
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+1)))
	if v530 != 0 {
		v515 = v515 + v528
		v516 = v529
		v517 = v525
		v518 = v530
		goto L181
	} else {
		goto L186
	}
L186:
	;
	goto L182
L187:
	;
	goto L172
L188:
	;
	v692 = v481
	goto L51
L189:
	;
	v568 = int32(5)
	goto L191
L190:
	;
	v568 = int32(4)
	goto L191
L191:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v568)
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	v571 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v571
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+2)) = uint8(v571)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)) = uint8(v570)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+6)) = uint8(v571)
	v692 = v137 + int32(1)
	goto L51
L192:
	;
	v642 = F_pg_mblen_cstr(m, v640)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L84
	} else {
		goto L216
	}
L193:
	;
	if v255 != int32(34) {
		v640 = v137
		goto L192
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v635 = v137 + int32(1)
	v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v635))))
	if v636 == int32(34) {
		goto L213
	} else {
		goto L214
	}
L196:
	;
	v586 = v31
	v591 = v137 + int32(1)
	goto L197
L197:
	;
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v591))))
	if v601 != int32(92) {
		goto L200
	} else {
		goto L201
	}
L199:
	;
	v616 = F_pg_mblen_cstr(m, v615)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L84
	} else {
		goto L208
	}
L200:
	;
	if v601 == int32(0) {
		v724 = v586
		goto L15
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	v611 = v591 + int32(1)
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611))))
	if v612 != 0 {
		goto L205
	} else {
		goto L206
	}
L203:
	;
	if v601 != int32(34) {
		v615 = v591
		goto L199
	} else {
		goto L204
	}
L204:
	;
	v708 = v586
	v709 = v591 + int32(1)
	goto L19
L205:
	;
	v613 = v611
	goto L207
L206:
	;
	v613 = v591
	goto L207
L207:
	;
	v615 = v613
	goto L199
L208:
	;
	v618 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v586))) = uint8(v618)
	v621 = v586 + int32(1)
	if v616 != 0 {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	v625 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v623+v616))) = uint8(v625)
	*(*uint8)(unsafe.Add(mBase, uint32(v586)+6)) = uint8(v625)
	*(*int32)(unsafe.Add(mBase, uint32(v586)+8)) = v625
	v586 = v586 + int32(12)
	v591 = v616 + v615
	goto L197
L210:
	;
	v622 = F__emscripten_memcpy_bulkmem(m, v621, v615, v616)
	mBase = m.M
	v623 = v622
	goto L212
L211:
	;
	v623 = v621
	goto L212
L212:
	;
	goto L209
L213:
	;
	v639 = v635
	goto L215
L214:
	;
	v639 = v137
	goto L215
L215:
	;
	v640 = v639
	goto L192
L216:
	;
	v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v640))))
	if v28 == int32(0) {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v677)
	v680 = v31 + int32(1)
	if v642 != 0 {
		goto L226
	} else {
		goto L227
	}
L218:
	;
	v669 = int32(5)
	if base.Ui32(v644-int32(9)) < base.Ui32(v669) {
		v677 = v669
		goto L217
	} else {
		goto L223
	}
L219:
	;
	if base.Ui32(int32(93)) < base.Ui32((v644-int32(33))&int32(255)) {
		goto L218
	} else {
		goto L220
	}
L220:
	;
	if base.Ui32(int32(229)) < base.Ui32((v644&int32(223)-int32(91))&int32(255)) {
		goto L218
	} else {
		goto L221
	}
L221:
	;
	if base.Ui32((v644-int32(58))&int32(255)) < base.Ui32(int32(246)) {
		v677 = int32(4)
		goto L217
	} else {
		goto L222
	}
L222:
	;
	goto L218
L223:
	;
	if v644 == int32(32) {
		v677 = v669
		goto L217
	} else {
		goto L224
	}
L224:
	;
	v677 = int32(3)
	goto L217
L225:
	;
	v684 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v682+v642))) = uint8(v684)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+6)) = uint8(v684)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v684
	v692 = v640 + v642
	goto L51
L226:
	;
	v681 = F__emscripten_memcpy_bulkmem(m, v680, v640, v642)
	mBase = m.M
	v682 = v681
	goto L228
L227:
	;
	v682 = v680
	goto L228
L228:
	;
	goto L225
L229:
	;
	goto L18
L230:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L84
	} else {
		goto L231
	}
L231:
	;
	F_errmsg(m, int32(711674), int32(0))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L84
	} else {
		goto L232
	}
L232:
	;
	F_errfinish(m, int32(495035), int32(1205), int32(363160))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L84
	} else {
		goto L233
	}
L233:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L234:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L84
	} else {
		goto L235
	}
L235:
	;
	F_errmsg(m, int32(711700), int32(0))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L84
	} else {
		goto L236
	}
L236:
	;
	F_errfinish(m, int32(495035), int32(1221), int32(363160))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L84
	} else {
		goto L237
	}
L237:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L238:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L84
	} else {
		goto L239
	}
L239:
	;
	F_errmsg(m, int32(118697), int32(0))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L84
	} else {
		goto L240
	}
L240:
	;
	F_errfinish(m, int32(495035), int32(1248), int32(363160))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L84
	} else {
		goto L241
	}
L241:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L242:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L84
	} else {
		goto L243
	}
L243:
	;
	F_errmsg(m, int32(221510), int32(0))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L84
	} else {
		goto L244
	}
L244:
	;
	F_errfinish(m, int32(495035), int32(1252), int32(363160))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L84
	} else {
		goto L245
	}
L245:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L246:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L84
	} else {
		goto L247
	}
L247:
	;
	F_errmsg(m, int32(415440), int32(0))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L84
	} else {
		goto L248
	}
L248:
	;
	F_errfinish(m, int32(495035), int32(1264), int32(363160))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L84
	} else {
		goto L249
	}
L249:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L250:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L84
	} else {
		goto L251
	}
L251:
	;
	F_errmsg(m, int32(221639), int32(0))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L84
	} else {
		goto L252
	}
L252:
	;
	F_errfinish(m, int32(495035), int32(1268), int32(363160))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L84
	} else {
		goto L253
	}
L253:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L254:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L84
	} else {
		goto L255
	}
L255:
	;
	F_errmsg(m, int32(221720), int32(0))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L84
	} else {
		goto L256
	}
L256:
	;
	F_errfinish(m, int32(495035), int32(1288), int32(363160))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L84
	} else {
		goto L257
	}
L257:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L258:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L84
	} else {
		goto L259
	}
L259:
	;
	F_errmsg(m, int32(221687), int32(0))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L84
	} else {
		goto L260
	}
L260:
	;
	F_errfinish(m, int32(495035), int32(1298), int32(363160))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L84
	} else {
		goto L261
	}
L261:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L262:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L84
	} else {
		goto L263
	}
L263:
	;
	F_errmsg(m, int32(221801), int32(0))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L84
	} else {
		goto L264
	}
L264:
	;
	F_errfinish(m, int32(495035), int32(1308), int32(363160))
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L84
	} else {
		goto L265
	}
L265:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L266:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L84
	} else {
		goto L267
	}
L267:
	;
	F_errmsg(m, int32(221753), int32(0))
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L84
	} else {
		goto L268
	}
L268:
	;
	F_errfinish(m, int32(495035), int32(1317), int32(363160))
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L84
	} else {
		goto L269
	}
L269:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L270:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L84
	} else {
		goto L271
	}
L271:
	;
	F_errmsg(m, int32(415461), int32(0))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L84
	} else {
		goto L272
	}
L272:
	;
	F_errfinish(m, int32(495035), int32(1326), int32(363160))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L84
	} else {
		goto L273
	}
L273:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L274:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L84
	} else {
		goto L275
	}
L275:
	;
	F_errmsg(m, int32(221510), int32(0))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L84
	} else {
		goto L276
	}
L276:
	;
	F_errfinish(m, int32(495035), int32(1339), int32(363160))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L84
	} else {
		goto L277
	}
L277:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L278:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L84
	} else {
		goto L279
	}
L279:
	;
	F_errmsg(m, int32(125588), int32(0))
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L84
	} else {
		goto L280
	}
L280:
	;
	F_errdetail(m, int32(571203), int32(0))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L84
	} else {
		goto L281
	}
L281:
	;
	F_errfinish(m, int32(495035), int32(1354), int32(363160))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L84
	} else {
		goto L282
	}
L282:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L283:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L84
	} else {
		goto L284
	}
L284:
	;
	F_errmsg(m, int32(125548), int32(0))
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L84
	} else {
		goto L285
	}
L285:
	;
	F_errdetail(m, int32(650911), int32(0))
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L84
	} else {
		goto L286
	}
L286:
	;
	F_errfinish(m, int32(495035), int32(1364), int32(363160))
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L84
	} else {
		goto L287
	}
L287:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_parse_scram_secret(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v135 int64
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v223 int32
	_ = v223
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v308 int32
	_ = v308
	var v324 int32
	_ = v324
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v424 int32
	_ = v424
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v509 int32
	_ = v509
	var v525 int32
	_ = v525
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v604 int32
	_ = v604
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v624 int32
	_ = v624
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v709 int32
	_ = v709
	var v725 int32
	_ = v725
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v750 int32
	_ = v750
	var v759 int32
	_ = v759
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = F_pstrdup(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v17
	v23 = v15 + int32(12)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v38 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	v27 = F_strcspn(m, v26, int32(671510))
	mBase = m.M
	v28 = v27 + v26
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v29 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	goto L3
L7:
	;
	v30 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v30)
	v35 = v28 + int32(1)
	goto L9
L8:
	;
	v35 = int32(0)
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v35
	goto L6
L10:
	;
	m.G0 = v15 + int32(16)
	return v759
L11:
	;
	v750 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v750
	v759 = v750
	goto L10
L12:
	;
	v42 = v15 + int32(12)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v45 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v57 == int32(0) {
		goto L11
	} else {
		goto L20
	}
L14:
	;
	v46 = F_strcspn(m, v45, int32(544243))
	mBase = m.M
	v47 = v46 + v45
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v48 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	goto L13
L17:
	;
	v49 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v49)
	v54 = v47 + int32(1)
	goto L19
L18:
	;
	v54 = int32(0)
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v54
	goto L16
L20:
	;
	v61 = v15 + int32(12)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v64 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v76 == int32(0) {
		goto L11
	} else {
		goto L28
	}
L22:
	;
	v65 = F_strcspn(m, v64, int32(671510))
	mBase = m.M
	v66 = v65 + v64
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v67 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	v68 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v68)
	v73 = v66 + int32(1)
	goto L27
L26:
	;
	v73 = int32(0)
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v73
	goto L24
L28:
	;
	v80 = v15 + int32(12)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	if v83 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v95 == int32(0) {
		goto L11
	} else {
		goto L36
	}
L30:
	;
	v84 = F_strcspn(m, v83, int32(544243))
	mBase = m.M
	v85 = v84 + v83
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	if v86 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	goto L29
L33:
	;
	v87 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v87)
	v92 = v85 + int32(1)
	goto L35
L34:
	;
	v92 = int32(0)
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v92
	goto L32
L36:
	;
	v98 = int32(547004)
	v101 = int32(*(*uint8)(unsafe.Add(mBase, _consts[437])))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v102 == int32(0) {
		v121 = v101
		v122 = v102
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v122-v121 != 0 {
		goto L11
	} else {
		goto L45
	}
L38:
	;
	goto L37
L39:
	;
	if v101 != v102 {
		v121 = v101
		v122 = v102
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v106 = v26
	v107 = v98
	goto L41
L41:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
	if v111 == int32(0) {
		v121 = v110
		v122 = v111
		goto L38
	} else {
		goto L43
	}
L42:
	;
	v121 = v110
	v122 = v111
	goto L38
L43:
	;
	v114 = int32(1)
	if v110 == v111 {
		v106 = v106 + v114
		v107 = v107 + v114
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(32)
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	v135 = F_strtox_2(m, v45, v15+int32(8), int32(10), int64(2147483648))
	mBase = m.M
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = base.I32_wrap_i64(v135)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	if v139 != 0 {
		goto L11
	} else {
		goto L47
	}
L47:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v141 != 0 {
		goto L11
	} else {
		goto L48
	}
L48:
	;
	v142 = F_strlen(m, v64)
	mBase = m.M
	v146 = v142 * int32(3) >> (uint(int32(2)) % 32)
	goto L49
L49:
	;
	v147 = F_palloc(m, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v149 = F_strlen(m, v64)
	mBase = m.M
	v150 = int32(0)
	v157 = v64 + v149
	if base.Ui32(v64) < base.Ui32(v157) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	if v337 < int32(0) {
		goto L11
	} else {
		goto L95
	}
L52:
	;
	v324 = F___memset(m, v147, int32(0), v146)
	mBase = m.M
	v337 = int32(-1)
	goto L51
L53:
	;
	v159 = v64
	v163 = v150
	v164 = v147
	v165 = v150
	v167 = v150
	goto L56
L54:
	;
	v308 = v147
	goto L55
L55:
	;
	v337 = v308 - v147
	goto L51
L56:
	;
	v171 = v159 + int32(1)
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	if v172 != int32(61) {
		goto L64
	} else {
		goto L65
	}
L57:
	;
	if v296 != 0 {
		goto L52
	} else {
		goto L94
	}
L58:
	;
	if v294 != v157 {
		v159 = v294
		v163 = v296
		v164 = v297
		v165 = v298
		v167 = v300
		goto L56
	} else {
		goto L93
	}
L59:
	;
	if v146 < v164-v147+int32(1) {
		goto L52
	} else {
		goto L80
	}
L60:
	;
	v250 = int32(2)
	v251 = v171
	v254 = v167 << (uint(int32(6)) % 32)
	goto L59
L61:
	;
	v239 = v236 + v235<<(uint(int32(6))%32)
	v241 = v232 + int32(1)
	if v241 == int32(4) {
		v250 = v233
		v251 = v234
		v254 = v239
		goto L59
	} else {
		goto L79
	}
L62:
	;
	v232 = int32(3)
	v233 = int32(1)
	v234 = v159 + int32(2)
	v235 = v190
	v236 = v185
	goto L61
L63:
	;
	if base.Ui32(int32(125)) < base.Ui32((v209-int32(1))&int32(255)) {
		goto L52
	} else {
		goto L77
	}
L64:
	;
	v176 = v172 - int32(9)
	if base.Ui32(int32(23)) < base.Ui32(v176) {
		v209 = v172
		v210 = v163
		v211 = v165
		v212 = v171
		v213 = v167
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v185 = int32(0)
	if v165 != 0 {
		v232 = v163
		v233 = v165
		v234 = v171
		v235 = v167
		v236 = v185
		goto L61
	} else {
		goto L69
	}
L67:
	;
	if int32(1)<<(uint(v176)%32)&int32(8388627) == int32(0) {
		v209 = v172
		v210 = v163
		v211 = v165
		v212 = v171
		v213 = v167
		goto L63
	} else {
		goto L68
	}
L68:
	;
	goto L52
L69:
	;
	switch v163 - int32(2) {
	case 0:
		goto L70
	case 1:
		goto L60
	default:
		goto L52
	}
L70:
	;
	if v171 == v157 {
		goto L52
	} else {
		goto L71
	}
L71:
	;
	v190 = v167 << (uint(int32(6)) % 32)
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171))))
	if v191 == int32(61) {
		goto L62
	} else {
		goto L72
	}
L72:
	;
	v195 = v191 - int32(9)
	if int32(1)<<(uint(v195)%32)&int32(8388627) != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v203 = base.B2i32(base.Ui32(v195) <= base.Ui32(int32(23)))
	goto L75
L74:
	;
	v203 = int32(0)
	goto L75
L75:
	;
	if v203 != 0 {
		goto L52
	} else {
		goto L76
	}
L76:
	;
	v209 = v191
	v210 = int32(3)
	v211 = int32(1)
	v212 = v159 + int32(2)
	v213 = v190
	goto L63
L77:
	;
	v223 = int32(*(*int8)(unsafe.Add(mBase, uint32(v209)+uint32(_consts[438]))))
	if v223 < int32(0) {
		goto L52
	} else {
		goto L78
	}
L78:
	;
	v232 = v210
	v233 = v211
	v234 = v212
	v235 = v213
	v236 = v223
	goto L61
L79:
	;
	v294 = v234
	v296 = v241
	v297 = v164
	v298 = v233
	v300 = v239
	goto L58
L80:
	;
	v260 = int32(base.Ui32(v254) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v164))) = uint8(v260)
	v263 = v164 + int32(1)
	if base.Ui32(v250) < base.Ui32(int32(2)) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v267 = v250
	goto L83
L82:
	;
	v267 = int32(0)
	goto L83
L83:
	;
	if v267 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	if v146 < v263-v147+int32(1) {
		goto L52
	} else {
		goto L87
	}
L85:
	;
	v279 = v263
	goto L86
L86:
	;
	v280 = int32(0)
	if v250 == v280 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	v275 = int32(base.Ui32(v254) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v164)+1)) = uint8(v275)
	v279 = v164 + int32(2)
	goto L86
L88:
	;
	v294 = v251
	v296 = int32(0)
	v297 = v292
	v298 = v250
	v300 = v280
	goto L58
L89:
	;
	if v146 < v279-v147+int32(1) {
		goto L52
	} else {
		goto L92
	}
L90:
	;
	if base.Ui32(int32(3)) <= base.Ui32(v250) {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v292 = v279
	goto L88
L92:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v279))) = uint8(v254)
	v292 = v279 + int32(1)
	goto L88
L93:
	;
	goto L57
L94:
	;
	v308 = v297
	goto L55
L95:
	;
	v340 = F_pstrdup(m, v64)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v340
	v343 = F_strlen(m, v83)
	mBase = m.M
	v347 = v343 * int32(3) >> (uint(int32(2)) % 32)
	goto L97
L97:
	;
	v348 = F_palloc(m, v347)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v350 = F_strlen(m, v83)
	mBase = m.M
	v351 = int32(0)
	v358 = v83 + v350
	if base.Ui32(v83) < base.Ui32(v358) {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v538 != v539 {
		goto L11
	} else {
		goto L143
	}
L100:
	;
	v525 = F___memset(m, v348, int32(0), v347)
	mBase = m.M
	v538 = int32(-1)
	goto L99
L101:
	;
	v360 = v83
	v364 = v351
	v365 = v348
	v366 = v351
	v368 = v351
	goto L104
L102:
	;
	v509 = v348
	goto L103
L103:
	;
	v538 = v509 - v348
	goto L99
L104:
	;
	v372 = v360 + int32(1)
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360))))
	if v373 != int32(61) {
		goto L112
	} else {
		goto L113
	}
L105:
	;
	if v497 != 0 {
		goto L100
	} else {
		goto L142
	}
L106:
	;
	if v495 != v358 {
		v360 = v495
		v364 = v497
		v365 = v498
		v366 = v499
		v368 = v501
		goto L104
	} else {
		goto L141
	}
L107:
	;
	if v347 < v365-v348+int32(1) {
		goto L100
	} else {
		goto L128
	}
L108:
	;
	v451 = int32(2)
	v452 = v372
	v455 = v368 << (uint(int32(6)) % 32)
	goto L107
L109:
	;
	v440 = v437 + v436<<(uint(int32(6))%32)
	v442 = v433 + int32(1)
	if v442 == int32(4) {
		v451 = v434
		v452 = v435
		v455 = v440
		goto L107
	} else {
		goto L127
	}
L110:
	;
	v433 = int32(3)
	v434 = int32(1)
	v435 = v360 + int32(2)
	v436 = v391
	v437 = v386
	goto L109
L111:
	;
	if base.Ui32(int32(125)) < base.Ui32((v410-int32(1))&int32(255)) {
		goto L100
	} else {
		goto L125
	}
L112:
	;
	v377 = v373 - int32(9)
	if base.Ui32(int32(23)) < base.Ui32(v377) {
		v410 = v373
		v411 = v364
		v412 = v366
		v413 = v372
		v414 = v368
		goto L111
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v386 = int32(0)
	if v366 != 0 {
		v433 = v364
		v434 = v366
		v435 = v372
		v436 = v368
		v437 = v386
		goto L109
	} else {
		goto L117
	}
L115:
	;
	if int32(1)<<(uint(v377)%32)&int32(8388627) == int32(0) {
		v410 = v373
		v411 = v364
		v412 = v366
		v413 = v372
		v414 = v368
		goto L111
	} else {
		goto L116
	}
L116:
	;
	goto L100
L117:
	;
	switch v364 - int32(2) {
	case 0:
		goto L118
	case 1:
		goto L108
	default:
		goto L100
	}
L118:
	;
	if v372 == v358 {
		goto L100
	} else {
		goto L119
	}
L119:
	;
	v391 = v368 << (uint(int32(6)) % 32)
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372))))
	if v392 == int32(61) {
		goto L110
	} else {
		goto L120
	}
L120:
	;
	v396 = v392 - int32(9)
	if int32(1)<<(uint(v396)%32)&int32(8388627) != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v404 = base.B2i32(base.Ui32(v396) <= base.Ui32(int32(23)))
	goto L123
L122:
	;
	v404 = int32(0)
	goto L123
L123:
	;
	if v404 != 0 {
		goto L100
	} else {
		goto L124
	}
L124:
	;
	v410 = v392
	v411 = int32(3)
	v412 = int32(1)
	v413 = v360 + int32(2)
	v414 = v391
	goto L111
L125:
	;
	v424 = int32(*(*int8)(unsafe.Add(mBase, uint32(v410)+uint32(_consts[438]))))
	if v424 < int32(0) {
		goto L100
	} else {
		goto L126
	}
L126:
	;
	v433 = v411
	v434 = v412
	v435 = v413
	v436 = v414
	v437 = v424
	goto L109
L127:
	;
	v495 = v435
	v497 = v442
	v498 = v365
	v499 = v434
	v501 = v440
	goto L106
L128:
	;
	v461 = int32(base.Ui32(v455) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v365))) = uint8(v461)
	v464 = v365 + int32(1)
	if base.Ui32(v451) < base.Ui32(int32(2)) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v468 = v451
	goto L131
L130:
	;
	v468 = int32(0)
	goto L131
L131:
	;
	if v468 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	if v347 < v464-v348+int32(1) {
		goto L100
	} else {
		goto L135
	}
L133:
	;
	v480 = v464
	goto L134
L134:
	;
	v481 = int32(0)
	if v451 == v481 {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	v476 = int32(base.Ui32(v455) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v365)+1)) = uint8(v476)
	v480 = v365 + int32(2)
	goto L134
L136:
	;
	v495 = v452
	v497 = int32(0)
	v498 = v493
	v499 = v451
	v501 = v481
	goto L106
L137:
	;
	if v347 < v480-v348+int32(1) {
		goto L100
	} else {
		goto L140
	}
L138:
	;
	if base.Ui32(int32(3)) <= base.Ui32(v451) {
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v493 = v480
	goto L136
L140:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v480))) = uint8(v455)
	v493 = v480 + int32(1)
	goto L136
L141:
	;
	goto L105
L142:
	;
	v509 = v498
	goto L103
L143:
	;
	if v538 != 0 {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	v543 = F_strlen(m, v95)
	mBase = m.M
	v547 = v543 * int32(3) >> (uint(int32(2)) % 32)
	goto L148
L145:
	;
	v541 = F__emscripten_memcpy_bulkmem(m, l5, v348, v538)
	mBase = m.M
	goto L147
L146:
	;
	goto L147
L147:
	;
	goto L144
L148:
	;
	v548 = F_palloc(m, v547)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	v550 = F_strlen(m, v95)
	mBase = m.M
	v551 = int32(0)
	v558 = v95 + v550
	if base.Ui32(v95) < base.Ui32(v558) {
		goto L152
	} else {
		goto L153
	}
L150:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v738 != v739 {
		goto L11
	} else {
		goto L194
	}
L151:
	;
	v725 = F___memset(m, v548, int32(0), v547)
	mBase = m.M
	v738 = int32(-1)
	goto L150
L152:
	;
	v560 = v95
	v564 = v551
	v565 = v548
	v566 = v551
	v568 = v551
	goto L155
L153:
	;
	v709 = v548
	goto L154
L154:
	;
	v738 = v709 - v548
	goto L150
L155:
	;
	v572 = v560 + int32(1)
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v560))))
	if v573 != int32(61) {
		goto L163
	} else {
		goto L164
	}
L156:
	;
	if v697 != 0 {
		goto L151
	} else {
		goto L193
	}
L157:
	;
	if v695 != v558 {
		v560 = v695
		v564 = v697
		v565 = v698
		v566 = v699
		v568 = v701
		goto L155
	} else {
		goto L192
	}
L158:
	;
	if v547 < v565-v548+int32(1) {
		goto L151
	} else {
		goto L179
	}
L159:
	;
	v651 = int32(2)
	v652 = v572
	v655 = v568 << (uint(int32(6)) % 32)
	goto L158
L160:
	;
	v640 = v637 + v636<<(uint(int32(6))%32)
	v642 = v633 + int32(1)
	if v642 == int32(4) {
		v651 = v634
		v652 = v635
		v655 = v640
		goto L158
	} else {
		goto L178
	}
L161:
	;
	v633 = int32(3)
	v634 = int32(1)
	v635 = v560 + int32(2)
	v636 = v591
	v637 = v586
	goto L160
L162:
	;
	if base.Ui32(int32(125)) < base.Ui32((v610-int32(1))&int32(255)) {
		goto L151
	} else {
		goto L176
	}
L163:
	;
	v577 = v573 - int32(9)
	if base.Ui32(int32(23)) < base.Ui32(v577) {
		v610 = v573
		v611 = v564
		v612 = v566
		v613 = v572
		v614 = v568
		goto L162
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v586 = int32(0)
	if v566 != 0 {
		v633 = v564
		v634 = v566
		v635 = v572
		v636 = v568
		v637 = v586
		goto L160
	} else {
		goto L168
	}
L166:
	;
	if int32(1)<<(uint(v577)%32)&int32(8388627) == int32(0) {
		v610 = v573
		v611 = v564
		v612 = v566
		v613 = v572
		v614 = v568
		goto L162
	} else {
		goto L167
	}
L167:
	;
	goto L151
L168:
	;
	switch v564 - int32(2) {
	case 0:
		goto L169
	case 1:
		goto L159
	default:
		goto L151
	}
L169:
	;
	if v572 == v558 {
		goto L151
	} else {
		goto L170
	}
L170:
	;
	v591 = v568 << (uint(int32(6)) % 32)
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572))))
	if v592 == int32(61) {
		goto L161
	} else {
		goto L171
	}
L171:
	;
	v596 = v592 - int32(9)
	if int32(1)<<(uint(v596)%32)&int32(8388627) != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v604 = base.B2i32(base.Ui32(v596) <= base.Ui32(int32(23)))
	goto L174
L173:
	;
	v604 = int32(0)
	goto L174
L174:
	;
	if v604 != 0 {
		goto L151
	} else {
		goto L175
	}
L175:
	;
	v610 = v592
	v611 = int32(3)
	v612 = int32(1)
	v613 = v560 + int32(2)
	v614 = v591
	goto L162
L176:
	;
	v624 = int32(*(*int8)(unsafe.Add(mBase, uint32(v610)+uint32(_consts[438]))))
	if v624 < int32(0) {
		goto L151
	} else {
		goto L177
	}
L177:
	;
	v633 = v611
	v634 = v612
	v635 = v613
	v636 = v614
	v637 = v624
	goto L160
L178:
	;
	v695 = v635
	v697 = v642
	v698 = v565
	v699 = v634
	v701 = v640
	goto L157
L179:
	;
	v661 = int32(base.Ui32(v655) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v565))) = uint8(v661)
	v664 = v565 + int32(1)
	if base.Ui32(v651) < base.Ui32(int32(2)) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v668 = v651
	goto L182
L181:
	;
	v668 = int32(0)
	goto L182
L182:
	;
	if v668 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	if v547 < v664-v548+int32(1) {
		goto L151
	} else {
		goto L186
	}
L184:
	;
	v680 = v664
	goto L185
L185:
	;
	v681 = int32(0)
	if v651 == v681 {
		goto L188
	} else {
		goto L189
	}
L186:
	;
	v676 = int32(base.Ui32(v655) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v565)+1)) = uint8(v676)
	v680 = v565 + int32(2)
	goto L185
L187:
	;
	v695 = v652
	v697 = int32(0)
	v698 = v693
	v699 = v651
	v701 = v681
	goto L157
L188:
	;
	if v547 < v680-v548+int32(1) {
		goto L151
	} else {
		goto L191
	}
L189:
	;
	if base.Ui32(int32(3)) <= base.Ui32(v651) {
		goto L188
	} else {
		goto L190
	}
L190:
	;
	v693 = v680
	goto L187
L191:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v680))) = uint8(v655)
	v693 = v680 + int32(1)
	goto L187
L192:
	;
	goto L156
L193:
	;
	v709 = v698
	goto L154
L194:
	;
	if v738 != 0 {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	v759 = int32(1)
	goto L10
L196:
	;
	v741 = F__emscripten_memcpy_bulkmem(m, l6, v548, v738)
	mBase = m.M
	goto L198
L197:
	;
	goto L198
L198:
	;
	goto L195
}
func F_parser_coercion_errposition(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	if l1 < int32(0) {
		v6 = F_exprLocation(m, l2)
		v7 = v6
	} else {
		v7 = l1
	}
	F_parser_errposition(m, l0, v7)
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		return
	}
}
func F_pathkeys_contained_in(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	if l0 == l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(1)
L2:
	;
	goto L3
L3:
	;
	v12 = int32(0)
	goto L5
L4:
	;
	return v49
L5:
	;
	v16 = int32(0)
	if l0 == v16 {
		v26 = v16
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v49 = int32(0)
	goto L4
L7:
	;
	if l1 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v20 <= v12 {
		v26 = int32(0)
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v26 = v22 + v12<<(uint(int32(2))%32)
	goto L7
L10:
	;
	v33 = base.B2i32(v26 == int32(0))
	if v26 == int32(0) {
		v49 = v33
		goto L4
	} else {
		goto L15
	}
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v12 < v27 {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	return base.B2i32(v26 == int32(0))
L14:
	;
	goto L13
L15:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v39 = v36 + v12<<(uint(int32(2))%32)
	if v39 == int32(0) {
		v49 = v33
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v44 == v45 {
		v12 = v12 + int32(1)
		goto L5
	} else {
		goto L17
	}
L17:
	;
	goto L6
}
func F_pct_info_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	if v6 != v7 {
		if v6 < v7 {
			v12 = int32(-1)
		} else {
			v12 = int32(1)
		}
		return v12
	} else {
		v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		v15 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
		if v14 != v15 {
			if v14 < v15 {
				v20 = int32(-1)
			} else {
				v20 = int32(1)
			}
			v21 = v20
		} else {
			v21 = int32(0)
		}
		return v21
	}
}
func F_pglz_decompress(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v171 int32
	_ = v171
	var v187 int32
	_ = v187
	v13 = l2 + l3
	v14 = l0 + l1
	if base.Ui32(v14) <= base.Ui32(l0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v187
L2:
	;
	if l4 != 0 {
		goto L46
	} else {
		goto L47
	}
L3:
	;
	v159 = l0
	v160 = l2
	goto L2
L4:
	;
	if base.Ui32(v13) <= base.Ui32(l2) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v17 = l0
	v18 = l2
	goto L6
L6:
	;
	v30 = v17 + int32(1)
	if base.Ui32(v14) <= base.Ui32(v30) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v159 = v145
	v160 = v146
	goto L2
L8:
	;
	if base.Ui32(v14) <= base.Ui32(v145) {
		v159 = v145
		v160 = v146
		goto L2
	} else {
		goto L44
	}
L9:
	;
	v145 = v30
	v146 = v18
	goto L8
L10:
	;
	goto L11
L11:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v35 = v18
	v37 = v30
	v43 = v32
	v44 = int32(0)
	goto L12
L12:
	;
	if v43&int32(1) != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v145 = v122
	v146 = v134
	goto L8
L14:
	;
	if base.Ui32(int32(6)) < base.Ui32(v44) {
		v145 = v122
		v146 = v134
		goto L8
	} else {
		goto L41
	}
L15:
	;
	v48 = int32(-1)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	v53 = v49&int32(15) + int32(3)
	if v53 != int32(18) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	*(*uint8)(unsafe.Add(mBase, uint32(v35))) = uint8(v116)
	v118 = int32(1)
	v122 = v37 + v118
	v134 = v35 + v118
	goto L14
L18:
	;
	v63 = v53
	v64 = v37 + int32(2)
	goto L20
L19:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+2)))
	v63 = v58 + int32(18)
	v64 = v37 + int32(3)
	goto L20
L20:
	;
	if base.Ui32(v14) < base.Ui32(v64) {
		v187 = v48
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	v71 = v66 | v49<<(uint(int32(4))%32)&int32(3840)
	if v71 == int32(0) {
		v187 = v48
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v35-l2 < v71 {
		v187 = v48
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v76 = v13 - v35
	if v63 < v76 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v78 = v63
	goto L26
L25:
	;
	v78 = v76
	goto L26
L26:
	;
	if v71 < v78 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v81 = v35
	v83 = v71
	v85 = v78
	goto L30
L28:
	;
	v101 = v35
	v103 = v71
	v105 = v78
	goto L29
L29:
	;
	if v105 != 0 {
		goto L38
	} else {
		goto L39
	}
L30:
	;
	v92 = v85 - v83
	if v83 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v101 = v96
	v103 = v98
	v105 = v92
	goto L29
L32:
	;
	v96 = v95 + v83
	v98 = v83 << (uint(int32(1)) % 32)
	if v98 < v92 {
		v81 = v96
		v83 = v98
		v85 = v92
		goto L30
	} else {
		goto L36
	}
L33:
	;
	v94 = F__emscripten_memcpy_bulkmem(m, v81, v81-v83, v83)
	mBase = m.M
	v95 = v94
	goto L35
L34:
	;
	v95 = v81
	goto L35
L35:
	;
	goto L32
L36:
	;
	goto L31
L37:
	;
	v122 = v64
	v134 = v114 + v105
	goto L14
L38:
	;
	v113 = F__emscripten_memcpy_bulkmem(m, v101, v101-v103, v105)
	mBase = m.M
	v114 = v113
	goto L40
L39:
	;
	v114 = v101
	goto L40
L40:
	;
	goto L37
L41:
	;
	if base.Ui32(v14) <= base.Ui32(v122) {
		v145 = v122
		v146 = v134
		goto L8
	} else {
		goto L42
	}
L42:
	;
	v138 = int32(1)
	if base.Ui32(v134) < base.Ui32(v13) {
		v35 = v134
		v37 = v122
		v43 = int32(base.Ui32(v43&int32(254)) >> (uint(v138) % 32))
		v44 = v44 + v138
		goto L12
	} else {
		goto L43
	}
L43:
	;
	goto L13
L44:
	;
	if base.Ui32(v146) < base.Ui32(v13) {
		v17 = v145
		v18 = v146
		goto L6
	} else {
		goto L45
	}
L45:
	;
	goto L7
L46:
	;
	v171 = int32(-1)
	if v159 != v14 {
		v187 = v171
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v187 = v160 - l2
	goto L1
L49:
	;
	if v160 != v13 {
		v187 = v171
		goto L1
	} else {
		goto L50
	}
L50:
	;
	goto L48
}
func F_phraseto_tsquery_byid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(4)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v13
		v17 = F_text_to_cstring(m, v9)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v24 = F_parse_tsquery(m, v17, int32(1174), v6+int32(8), int32(1), int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v24
			}
		}
	}
}
func F_plan_elem_desc(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+10)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v14
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v12
	F_appendStringInfo(m, l0, int32(57693), v10)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return
	} else {
		F_appendStringInfoString(m, l0, int32(543756))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+10)))
			F_array_desc(m, l0, v26, int32(2), v28, int32(243), int32(0))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+10)))
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v33 + v34<<(uint(int32(1))%32)
				F_appendStringInfoString(m, l0, int32(6954))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					m.G0 = v10 + int32(16)
					return
				}
			}
		}
	}
}
func F_planstate_tree_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_check_stack_depth(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v13 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return v231
L4:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v46 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v16 <= int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v23 = int32(0)
	goto L7
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27+v23<<(uint(int32(2))%32))))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v33 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v32, l2)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L4
L9:
	;
	if v33 != 0 {
		v231 = int32(1)
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v36 = v23 + int32(1)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v36 < v37 {
		v23 = v36
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v55 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v49 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v46, l2)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v49 == int32(0) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	return int32(1)
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	switch v64 - int32(334) {
	case 0:
		goto L26
	case 1:
		goto L25
	default:
		goto L20
	case 3:
		goto L24
	case 4:
		goto L23
	case 13:
		goto L22
	case 21:
		goto L21
	}
L17:
	;
	v58 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v55, l2)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v58 == int32(0) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	return int32(1)
L20:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v195 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L21:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v162 == int32(0) {
		goto L20
	} else {
		goto L53
	}
L22:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v156 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v155, l2)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L51
	}
L23:
	;
	v133 = int32(0)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v134 <= v133 {
		goto L20
	} else {
		goto L45
	}
L24:
	;
	v111 = int32(0)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v112 <= v111 {
		goto L20
	} else {
		goto L39
	}
L25:
	;
	v89 = int32(0)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v90 <= v89 {
		goto L20
	} else {
		goto L33
	}
L26:
	;
	v67 = int32(0)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v68 <= v67 {
		goto L20
	} else {
		goto L27
	}
L27:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v76 = v67
	goto L28
L28:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v71+v76<<(uint(int32(2))%32))))
	v84 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v83, l2)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	goto L20
L30:
	;
	if v84 != 0 {
		v231 = int32(1)
		goto L3
	} else {
		goto L31
	}
L31:
	;
	v87 = v76 + int32(1)
	if v68 != v87 {
		v76 = v87
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v98 = v89
	goto L34
L34:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v93+v98<<(uint(int32(2))%32))))
	v106 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v105, l2)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L36
	}
L35:
	;
	goto L20
L36:
	;
	if v106 != 0 {
		v231 = int32(1)
		goto L3
	} else {
		goto L37
	}
L37:
	;
	v109 = v98 + int32(1)
	if v90 != v109 {
		v98 = v109
		goto L34
	} else {
		goto L38
	}
L38:
	;
	goto L35
L39:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v120 = v111
	goto L40
L40:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v115+v120<<(uint(int32(2))%32))))
	v128 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v127, l2)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	goto L20
L42:
	;
	if v128 != 0 {
		v231 = int32(1)
		goto L3
	} else {
		goto L43
	}
L43:
	;
	v131 = v120 + int32(1)
	if v112 != v131 {
		v120 = v131
		goto L40
	} else {
		goto L44
	}
L44:
	;
	goto L41
L45:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v142 = v133
	goto L46
L46:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v137+v142<<(uint(int32(2))%32))))
	v150 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v149, l2)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	goto L20
L48:
	;
	if v150 != 0 {
		v231 = int32(1)
		goto L3
	} else {
		goto L49
	}
L49:
	;
	v153 = v142 + int32(1)
	if v134 != v153 {
		v142 = v153
		goto L46
	} else {
		goto L50
	}
L50:
	;
	goto L47
L51:
	;
	if v156 == int32(0) {
		goto L20
	} else {
		goto L52
	}
L52:
	;
	return int32(1)
L53:
	;
	v165 = int32(0)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	if v166 <= v165 {
		goto L20
	} else {
		goto L54
	}
L54:
	;
	v173 = v165
	goto L55
L55:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v162)+12))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v177+v173<<(uint(int32(2))%32))))
	v182 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v181, l2)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	goto L20
L57:
	;
	if v182 != 0 {
		v231 = int32(1)
		goto L3
	} else {
		goto L58
	}
L58:
	;
	v185 = v173 + int32(1)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	if v185 < v186 {
		v173 = v185
		goto L55
	} else {
		goto L59
	}
L59:
	;
	goto L56
L60:
	;
	return int32(0)
L61:
	;
	goto L62
L62:
	;
	v200 = int32(0)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	if v201 <= v200 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	return int32(0)
L64:
	;
	goto L65
L65:
	;
	v209 = v200
	goto L66
L66:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v195)+12))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v213+v209<<(uint(int32(2))%32))))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+8))
	v219 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v218, l2)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L68
	}
L67:
	;
	v231 = v219
	goto L3
L68:
	;
	if v219 != 0 {
		v231 = v219
		goto L3
	} else {
		goto L69
	}
L69:
	;
	v222 = v209 + int32(1)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	if v222 < v223 {
		v209 = v222
		goto L66
	} else {
		goto L70
	}
L70:
	;
	goto L67
}
func F_porter_ISO_8859_1_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v227 int32
	_ = v227
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v284 int32
	_ = v284
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v341 int32
	_ = v341
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v469 int32
	_ = v469
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v624 int32
	_ = v624
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v725 int32
	_ = v725
	var v736 int32
	_ = v736
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v768 int32
	_ = v768
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v917 int32
	_ = v917
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1041 int32
	_ = v1041
	var v1045 int32
	_ = v1045
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1073 int32
	_ = v1073
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1086 int32
	_ = v1086
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1137 int32
	_ = v1137
	var v1151 int32
	_ = v1151
	var v1155 int32
	_ = v1155
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1184 int32
	_ = v1184
	var v1188 int32
	_ = v1188
	var v1201 int32
	_ = v1201
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1217 int32
	_ = v1217
	var v1221 int32
	_ = v1221
	var v1225 int32
	_ = v1225
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1246 int32
	_ = v1246
	var v1248 int32
	_ = v1248
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1262 int32
	_ = v1262
	var v1265 int32
	_ = v1265
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1283 int32
	_ = v1283
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v10 == v12 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v1283
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	v37 = v10
	goto L8
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v10))))
	if v16 != int32(121) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v19 = int32(1)
	v20 = v10 + v19
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v20
	v25 = F_slice_from_s(m, l0, v19, int32(2162741))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	if v25 < int32(0) {
		v1283 = v25
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = int32(1)
	goto L2
L8:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v51 < v50 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v125)+4)) = v95
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v125))) = v127
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v138 < v129 {
		goto L35
	} else {
		goto L36
	}
L10:
	;
	goto L9
L11:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v94 != 0 {
		goto L26
	} else {
		goto L27
	}
L12:
	;
	v53 = v50
	goto L14
L13:
	;
	v53 = v51
	goto L14
L14:
	;
	goto L16
L15:
	;
	v94 = v90
	goto L11
L16:
	;
	if v50 == v53 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v90 = int32(0)
	goto L15
L18:
	;
	v94 = int32(-1)
	goto L11
L19:
	;
	goto L20
L20:
	;
	v65 = int32(1)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66+v50))))
	if int32(121) < v68 {
		v90 = v65
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v70 = v68 - int32(97)
	if v70 < int32(0) {
		v90 = v65
		goto L15
	} else {
		goto L22
	}
L22:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v70)>>(uint(int32(3))%32)))+uint32(_consts[1290]))))
	if int32(base.Ui32(v76)>>(uint(v70&int32(7))%32))&int32(1) == int32(0) {
		v90 = v65
		goto L15
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v50 + int32(1)
	goto L24
L24:
	;
	goto L17
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v37
	v111 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v96 + v111
	v116 = F_slice_from_s(m, l0, v111, int32(2162746))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L5
	} else {
		goto L31
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v37
	if v95 <= v37 {
		goto L10
	} else {
		goto L30
	}
L27:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v96
	if v95 == v96 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99+v96))))
	if v101 == int32(121) {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	goto L26
L30:
	;
	v108 = v37 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v108
	v37 = v108
	goto L8
L31:
	;
	if v116 < int32(0) {
		v1283 = v116
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v120)+8)) = int32(1)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v37 = v123
	goto L8
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v129
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v359
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v359
	if v359 <= v129 {
		goto L96
	} else {
		goto L97
	}
L34:
	;
	if v178 < int32(0) {
		goto L33
	} else {
		goto L49
	}
L35:
	;
	v140 = v129
	goto L37
L36:
	;
	v140 = v138
	goto L37
L37:
	;
	v147 = v129
	goto L39
L38:
	;
	v178 = v158
	goto L34
L39:
	;
	if v147 == v140 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v178 = int32(-1)
	goto L34
L42:
	;
	goto L43
L43:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v147))))
	if int32(121) < v153 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v170 = v147 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v170
	v147 = v170
	goto L39
L45:
	;
	v155 = v153 - int32(97)
	if v155 < int32(0) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v158 = int32(1)
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v155)>>(uint(int32(3))%32)))+uint32(_consts[1290]))))
	if int32(base.Ui32(v162)>>(uint(v155&int32(7))%32))&v158 != 0 {
		goto L38
	} else {
		goto L47
	}
L47:
	;
	goto L44
L49:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v182 = v181 + v178
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v182
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v193 < v182 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if v236 < int32(0) {
		goto L33
	} else {
		goto L64
	}
L51:
	;
	v195 = v182
	goto L53
L52:
	;
	v195 = v193
	goto L53
L53:
	;
	v202 = v182
	goto L55
L54:
	;
	v236 = int32(1)
	goto L50
L55:
	;
	if v202 == v195 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v236 = int32(-1)
	goto L50
L58:
	;
	goto L59
L59:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208+v202))))
	if int32(121) < v210 {
		goto L54
	} else {
		goto L60
	}
L60:
	;
	v212 = v210 - int32(97)
	if v212 < int32(0) {
		goto L54
	} else {
		goto L61
	}
L61:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v212)>>(uint(int32(3))%32)))+uint32(_consts[1290]))))
	if int32(base.Ui32(v218)>>(uint(v212&int32(7))%32))&int32(1) == int32(0) {
		goto L54
	} else {
		goto L62
	}
L62:
	;
	v227 = v202 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v227
	v202 = v227
	goto L55
L64:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v240 = v239 + v236
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v240
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v242)+4)) = v240
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v252 < v251 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	if v292 < int32(0) {
		goto L33
	} else {
		goto L80
	}
L66:
	;
	v254 = v251
	goto L68
L67:
	;
	v254 = v252
	goto L68
L68:
	;
	v261 = v251
	goto L70
L69:
	;
	v292 = v272
	goto L65
L70:
	;
	if v261 == v254 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v292 = int32(-1)
	goto L65
L73:
	;
	goto L74
L74:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265+v261))))
	if int32(121) < v267 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v284 = v261 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v284
	v261 = v284
	goto L70
L76:
	;
	v269 = v267 - int32(97)
	if v269 < int32(0) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v272 = int32(1)
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v269)>>(uint(int32(3))%32)))+uint32(_consts[1290]))))
	if int32(base.Ui32(v276)>>(uint(v269&int32(7))%32))&v272 != 0 {
		goto L69
	} else {
		goto L78
	}
L78:
	;
	goto L75
L80:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v296 = v295 + v292
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v296
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v307 < v296 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	if v350 < int32(0) {
		goto L33
	} else {
		goto L95
	}
L82:
	;
	v309 = v296
	goto L84
L83:
	;
	v309 = v307
	goto L84
L84:
	;
	v316 = v296
	goto L86
L85:
	;
	v350 = int32(1)
	goto L81
L86:
	;
	if v316 == v309 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v350 = int32(-1)
	goto L81
L89:
	;
	goto L90
L90:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322+v316))))
	if int32(121) < v324 {
		goto L85
	} else {
		goto L91
	}
L91:
	;
	v326 = v324 - int32(97)
	if v326 < int32(0) {
		goto L85
	} else {
		goto L92
	}
L92:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v326)>>(uint(int32(3))%32)))+uint32(_consts[1290]))))
	if int32(base.Ui32(v332)>>(uint(v326&int32(7))%32))&int32(1) == int32(0) {
		goto L85
	} else {
		goto L93
	}
L93:
	;
	v341 = v316 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v341
	v316 = v341
	goto L86
L95:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v353))) = v354 + v350
	goto L33
L96:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v397
	v401 = v397 - int32(1)
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v401 <= v402 {
		goto L110
	} else {
		goto L111
	}
L97:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+v359-int32(1)))))
	if v367 != int32(115) {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v372 = F_find_among_b(m, l0, int32(4193920), int32(4))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L5
	} else {
		goto L99
	}
L99:
	;
	if v372 == int32(0) {
		goto L96
	} else {
		goto L100
	}
L100:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v376
	switch v372 - int32(1) {
	case 0:
		goto L103
	case 1:
		goto L102
	case 2:
		goto L101
	default:
		goto L96
	}
L101:
	;
	v392 = F_slice_del(m, l0)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L5
	} else {
		goto L108
	}
L102:
	;
	v388 = F_slice_from_s(m, l0, int32(1), int32(2162750))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L5
	} else {
		goto L106
	}
L103:
	;
	v382 = F_slice_from_s(m, l0, int32(2), int32(2162748))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L5
	} else {
		goto L104
	}
L104:
	;
	if int32(0) <= v382 {
		goto L96
	} else {
		goto L105
	}
L105:
	;
	v1283 = v382
	goto L1
L106:
	;
	if int32(0) <= v388 {
		goto L96
	} else {
		goto L107
	}
L107:
	;
	v1283 = v388
	goto L1
L108:
	;
	if v392 < int32(0) {
		v1283 = v392
		goto L1
	} else {
		goto L109
	}
L109:
	;
	goto L96
L110:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v710
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v710
	v713 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v710 <= v713 {
		goto L189
	} else {
		goto L190
	}
L111:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404+v401))))
	switch v406 - int32(100) {
	case 0, 3:
		goto L112
	default:
		goto L110
	}
L112:
	;
	v411 = F_find_among_b(m, l0, int32(4194000), int32(3))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L5
	} else {
		goto L113
	}
L113:
	;
	if v411 == int32(0) {
		goto L110
	} else {
		goto L114
	}
L114:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v415
	switch v411 - int32(1) {
	case 0:
		goto L116
	case 1:
		goto L115
	default:
		goto L110
	}
L115:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v443 = v436
	goto L122
L116:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v419)+4))
	if v415 < v420 {
		goto L110
	} else {
		goto L117
	}
L117:
	;
	v424 = F_slice_from_s(m, l0, int32(2), int32(2162761))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L5
	} else {
		goto L118
	}
L118:
	;
	if int32(0) <= v424 {
		goto L110
	} else {
		goto L119
	}
L119:
	;
	v1283 = v424
	goto L1
L120:
	;
	if v477 < int32(0) {
		goto L110
	} else {
		goto L132
	}
L121:
	;
	v477 = v457
	goto L120
L122:
	;
	if v443 <= v437 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v477 = int32(-1)
	goto L120
L125:
	;
	goto L126
L126:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448+v443-int32(1)))))
	if int32(121) < v452 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v469 = v443 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v469
	v443 = v469
	goto L122
L128:
	;
	v454 = v452 - int32(97)
	if v454 < int32(0) {
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v457 = int32(1)
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v454)>>(uint(int32(3))%32)))+uint32(_consts[1290]))))
	if int32(base.Ui32(v461)>>(uint(v454&int32(7))%32))&v457 != 0 {
		goto L121
	} else {
		goto L130
	}
L130:
	;
	goto L127
L132:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v480 + (v415 - v428)
	v484 = F_slice_del(m, l0)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L5
	} else {
		goto L133
	}
L133:
	;
	if v484 < int32(0) {
		v1283 = v484
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v490 = v488 - v489
	v492 = v488 - int32(1)
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v492 <= v493 {
		v536 = v488
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v537)+4))
	if v536 != v538 {
		goto L110
	} else {
		goto L147
	}
L136:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v495+v492))))
	if v497&int32(224) != int32(96) {
		v536 = v488
		goto L135
	} else {
		goto L137
	}
L137:
	;
	if int32(1)<<(uint(v497)%32)&int32(68514004) == int32(0) {
		v536 = v488
		goto L135
	} else {
		goto L138
	}
L138:
	;
	v510 = F_find_among_b(m, l0, int32(4194064), int32(13))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L5
	} else {
		goto L139
	}
L139:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v513 = v512 + v490
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v513
	switch v510 - int32(1) {
	case 0:
		goto L141
	case 1:
		goto L140
	case 2:
		v536 = v513
		goto L135
	default:
		goto L110
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v513
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v513 <= v525 {
		goto L110
	} else {
		goto L144
	}
L141:
	;
	v519 = F_insert_s(m, l0, v513, v513, int32(1), int32(2162763))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L5
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v513
	if int32(0) <= v519 {
		goto L110
	} else {
		goto L143
	}
L143:
	;
	v1283 = v519
	goto L1
L144:
	;
	v528 = v513 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v528
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v528
	v531 = F_slice_del(m, l0)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L5
	} else {
		goto L145
	}
L145:
	;
	if int32(0) <= v531 {
		goto L110
	} else {
		goto L146
	}
L146:
	;
	v1283 = v531
	goto L1
L147:
	;
	v540 = int32(0)
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L151
L148:
	;
	if v694 == int32(0) {
		goto L110
	} else {
		goto L186
	}
L149:
	;
	if v589 != 0 {
		v694 = v540
		goto L148
	} else {
		goto L161
	}
L150:
	;
	v589 = v586
	goto L149
L151:
	;
	if v548 <= v549 {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	v586 = int32(0)
	goto L150
L153:
	;
	v589 = int32(-1)
	goto L149
L154:
	;
	goto L155
L155:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v560+v548-int32(1)))))
	if int32(121) < v564 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v548 - int32(1)
	goto L160
L157:
	;
	v566 = v564 - int32(89)
	if v566 < int32(0) {
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v569 = int32(1)
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v566)>>(uint(int32(3))%32)))+uint32(_consts[1291]))))
	if int32(base.Ui32(v573)>>(uint(v566&int32(7))%32))&v569 != 0 {
		v586 = v569
		goto L150
	} else {
		goto L159
	}
L159:
	;
	goto L156
L160:
	;
	goto L152
L161:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v599 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L164
L162:
	;
	if v642 != 0 {
		v694 = v540
		goto L148
	} else {
		goto L173
	}
L163:
	;
	v642 = v638
	goto L162
L164:
	;
	if v598 <= v599 {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	v638 = int32(0)
	goto L163
L166:
	;
	v642 = int32(-1)
	goto L162
L167:
	;
	goto L168
L168:
	;
	v611 = int32(1)
	v612 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612+v598-v611))))
	if int32(121) < v616 {
		v638 = v611
		goto L163
	} else {
		goto L169
	}
L169:
	;
	v618 = v616 - int32(97)
	if v618 < int32(0) {
		v638 = v611
		goto L163
	} else {
		goto L170
	}
L170:
	;
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v618)>>(uint(int32(3))%32)))+uint32(_consts[1290]))))
	if int32(base.Ui32(v624)>>(uint(v618&int32(7))%32))&int32(1) == int32(0) {
		v638 = v611
		goto L163
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v598 - int32(1)
	goto L172
L172:
	;
	goto L165
L173:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v651 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L176
L174:
	;
	v694 = base.B2i32(v691 == int32(0))
	goto L148
L175:
	;
	v691 = v688
	goto L174
L176:
	;
	if v650 <= v651 {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	v688 = int32(0)
	goto L175
L178:
	;
	v691 = int32(-1)
	goto L174
L179:
	;
	goto L180
L180:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v662+v650-int32(1)))))
	if int32(121) < v666 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v650 - int32(1)
	goto L185
L182:
	;
	v668 = v666 - int32(97)
	if v668 < int32(0) {
		goto L181
	} else {
		goto L183
	}
L183:
	;
	v671 = int32(1)
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v668)>>(uint(int32(3))%32)))+uint32(_consts[1290]))))
	if int32(base.Ui32(v675)>>(uint(v668&int32(7))%32))&v671 != 0 {
		v688 = v671
		goto L175
	} else {
		goto L184
	}
L184:
	;
	goto L181
L185:
	;
	goto L177
L186:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v698 = v697 + v490
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v698
	v702 = F_insert_s(m, l0, v698, v698, int32(1), int32(2162764))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L5
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v698
	if v702 < int32(0) {
		v1283 = v702
		goto L1
	} else {
		goto L188
	}
L188:
	;
	goto L110
L189:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v789
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v789
	v792 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v789-int32(2) <= v792 {
		goto L207
	} else {
		goto L208
	}
L190:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v715+v710-int32(1)))))
	if v719|int32(32) != int32(121) {
		goto L189
	} else {
		goto L191
	}
L191:
	;
	v725 = v710 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v725
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v725
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v742 = v725
	goto L194
L192:
	;
	if v776 < int32(0) {
		goto L189
	} else {
		goto L204
	}
L193:
	;
	v776 = v756
	goto L192
L194:
	;
	if v742 <= v736 {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v776 = int32(-1)
	goto L192
L197:
	;
	goto L198
L198:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v747+v742-int32(1)))))
	if int32(121) < v751 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v768 = v742 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v768
	v742 = v768
	goto L194
L200:
	;
	v753 = v751 - int32(97)
	if v753 < int32(0) {
		goto L199
	} else {
		goto L201
	}
L201:
	;
	v756 = int32(1)
	v760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v753)>>(uint(int32(3))%32)))+uint32(_consts[1290]))))
	if int32(base.Ui32(v760)>>(uint(v753&int32(7))%32))&v756 != 0 {
		goto L193
	} else {
		goto L202
	}
L202:
	;
	goto L199
L204:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v779 - v776
	v784 = F_slice_from_s(m, l0, int32(1), int32(2162802))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L5
	} else {
		goto L205
	}
L205:
	;
	if v784 < int32(0) {
		v1283 = v784
		goto L1
	} else {
		goto L206
	}
L206:
	;
	goto L189
L207:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v904
	v906 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v904
	v909 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v904-int32(2) <= v909 {
		v961 = v906
		goto L253
	} else {
		goto L254
	}
L208:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v800 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v796+v789-int32(1)))))
	if v800&int32(224) != int32(96) {
		goto L207
	} else {
		goto L209
	}
L209:
	;
	if int32(1)<<(uint(v800)%32)&int32(815616) == int32(0) {
		goto L207
	} else {
		goto L210
	}
L210:
	;
	v813 = F_find_among_b(m, l0, int32(4194336), int32(20))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L5
	} else {
		goto L211
	}
L211:
	;
	if v813 == int32(0) {
		goto L207
	} else {
		goto L212
	}
L212:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v817
	v819 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v819)+4))
	if v817 < v820 {
		goto L207
	} else {
		goto L213
	}
L213:
	;
	switch v813 - int32(1) {
	case 0:
		goto L226
	case 1:
		goto L225
	case 2:
		goto L224
	case 3:
		goto L223
	case 4:
		goto L222
	case 5:
		goto L221
	case 6:
		goto L220
	case 7:
		goto L219
	case 8:
		goto L218
	case 9:
		goto L217
	case 10:
		goto L216
	case 11:
		goto L215
	case 12:
		goto L214
	default:
		goto L207
	}
L214:
	;
	v898 = F_slice_from_s(m, l0, int32(3), int32(2162840))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L5
	} else {
		goto L251
	}
L215:
	;
	v892 = F_slice_from_s(m, l0, int32(3), int32(2162837))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L5
	} else {
		goto L249
	}
L216:
	;
	v886 = F_slice_from_s(m, l0, int32(3), int32(2162834))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L5
	} else {
		goto L247
	}
L217:
	;
	v880 = F_slice_from_s(m, l0, int32(3), int32(2162831))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L5
	} else {
		goto L245
	}
L218:
	;
	v874 = F_slice_from_s(m, l0, int32(2), int32(2162829))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L5
	} else {
		goto L243
	}
L219:
	;
	v868 = F_slice_from_s(m, l0, int32(3), int32(2162826))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L5
	} else {
		goto L241
	}
L220:
	;
	v862 = F_slice_from_s(m, l0, int32(3), int32(2162823))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L5
	} else {
		goto L239
	}
L221:
	;
	v856 = F_slice_from_s(m, l0, int32(1), int32(2162822))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L5
	} else {
		goto L237
	}
L222:
	;
	v850 = F_slice_from_s(m, l0, int32(3), int32(2162819))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L5
	} else {
		goto L235
	}
L223:
	;
	v844 = F_slice_from_s(m, l0, int32(4), int32(2162815))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L5
	} else {
		goto L233
	}
L224:
	;
	v838 = F_slice_from_s(m, l0, int32(4), int32(2162811))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L5
	} else {
		goto L231
	}
L225:
	;
	v832 = F_slice_from_s(m, l0, int32(4), int32(2162807))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L5
	} else {
		goto L229
	}
L226:
	;
	v826 = F_slice_from_s(m, l0, int32(4), int32(2162803))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L5
	} else {
		goto L227
	}
L227:
	;
	if int32(0) <= v826 {
		goto L207
	} else {
		goto L228
	}
L228:
	;
	v1283 = v826
	goto L1
L229:
	;
	if int32(0) <= v832 {
		goto L207
	} else {
		goto L230
	}
L230:
	;
	v1283 = v832
	goto L1
L231:
	;
	if int32(0) <= v838 {
		goto L207
	} else {
		goto L232
	}
L232:
	;
	v1283 = v838
	goto L1
L233:
	;
	if int32(0) <= v844 {
		goto L207
	} else {
		goto L234
	}
L234:
	;
	v1283 = v844
	goto L1
L235:
	;
	if int32(0) <= v850 {
		goto L207
	} else {
		goto L236
	}
L236:
	;
	v1283 = v850
	goto L1
L237:
	;
	if int32(0) <= v856 {
		goto L207
	} else {
		goto L238
	}
L238:
	;
	v1283 = v856
	goto L1
L239:
	;
	if int32(0) <= v862 {
		goto L207
	} else {
		goto L240
	}
L240:
	;
	v1283 = v862
	goto L1
L241:
	;
	if int32(0) <= v868 {
		goto L207
	} else {
		goto L242
	}
L242:
	;
	v1283 = v868
	goto L1
L243:
	;
	if int32(0) <= v874 {
		goto L207
	} else {
		goto L244
	}
L244:
	;
	v1283 = v874
	goto L1
L245:
	;
	if int32(0) <= v880 {
		goto L207
	} else {
		goto L246
	}
L246:
	;
	v1283 = v880
	goto L1
L247:
	;
	if int32(0) <= v886 {
		goto L207
	} else {
		goto L248
	}
L248:
	;
	v1283 = v886
	goto L1
L249:
	;
	if int32(0) <= v892 {
		goto L207
	} else {
		goto L250
	}
L250:
	;
	v1283 = v892
	goto L1
L251:
	;
	if v898 < int32(0) {
		v1283 = v898
		goto L1
	} else {
		goto L252
	}
L252:
	;
	goto L207
L253:
	;
	if v961 < int32(0) {
		v1283 = v961
		goto L1
	} else {
		goto L270
	}
L254:
	;
	v913 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v917 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v913+v904-int32(1)))))
	if v917&int32(224) != int32(96) {
		v961 = v906
		goto L253
	} else {
		goto L255
	}
L255:
	;
	if int32(1)<<(uint(v917)%32)&int32(528928) == int32(0) {
		v961 = v906
		goto L253
	} else {
		goto L256
	}
L256:
	;
	v930 = F_find_among_b(m, l0, int32(4194736), int32(7))
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L5
	} else {
		goto L257
	}
L257:
	;
	if v930 == int32(0) {
		v961 = v906
		goto L253
	} else {
		goto L258
	}
L258:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v934
	v936 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v936)+4))
	if v934 < v937 {
		v961 = v906
		goto L253
	} else {
		goto L259
	}
L259:
	;
	switch v930 - int32(1) {
	case 0:
		goto L263
	case 1:
		goto L262
	case 2:
		goto L261
	default:
		goto L260
	}
L260:
	;
	v961 = int32(1)
	goto L253
L261:
	;
	v953 = F_slice_del(m, l0)
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L5
	} else {
		goto L268
	}
L262:
	;
	v949 = F_slice_from_s(m, l0, int32(2), int32(2162949))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L5
	} else {
		goto L266
	}
L263:
	;
	v943 = F_slice_from_s(m, l0, int32(2), int32(2162947))
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L5
	} else {
		goto L264
	}
L264:
	;
	if int32(0) <= v943 {
		goto L260
	} else {
		goto L265
	}
L265:
	;
	v961 = v943
	goto L253
L266:
	;
	if int32(0) <= v949 {
		goto L260
	} else {
		goto L267
	}
L267:
	;
	v961 = v949
	goto L253
L268:
	;
	if v953 < int32(0) {
		v961 = v953
		goto L253
	} else {
		goto L269
	}
L269:
	;
	goto L260
L270:
	;
	v964 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v964
	v966 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v964
	v970 = v964 - int32(1)
	v971 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v970 <= v971 {
		v1027 = v966
		goto L271
	} else {
		goto L272
	}
L271:
	;
	if v1027 < int32(0) {
		v1283 = v1027
		goto L1
	} else {
		goto L287
	}
L272:
	;
	v973 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v973+v970))))
	if v975&int32(224) != int32(96) {
		v1027 = v966
		goto L271
	} else {
		goto L273
	}
L273:
	;
	if int32(1)<<(uint(v975)%32)&int32(3961384) == int32(0) {
		v1027 = v966
		goto L271
	} else {
		goto L274
	}
L274:
	;
	v988 = F_find_among_b(m, l0, int32(4194880), int32(19))
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L5
	} else {
		goto L275
	}
L275:
	;
	if v988 == int32(0) {
		v1027 = v966
		goto L271
	} else {
		goto L276
	}
L276:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v992
	v994 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v994)))
	if v992 < v995 {
		v1027 = v966
		goto L271
	} else {
		goto L277
	}
L277:
	;
	switch v988 - int32(1) {
	case 0:
		goto L280
	case 1:
		goto L279
	default:
		goto L278
	}
L278:
	;
	v1027 = int32(1)
	goto L271
L279:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v992 <= v1003 {
		v1027 = v966
		goto L271
	} else {
		goto L283
	}
L280:
	;
	v999 = F_slice_del(m, l0)
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L5
	} else {
		goto L281
	}
L281:
	;
	if int32(0) <= v999 {
		goto L278
	} else {
		goto L282
	}
L282:
	;
	v1027 = v999
	goto L271
L283:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1007 = int32(1)
	v1009 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1005+v992-v1007))))
	if base.Ui32(v1007) < base.Ui32((v1009-int32(115))&int32(255)) {
		v1027 = v966
		goto L271
	} else {
		goto L284
	}
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v992 - int32(1)
	v1019 = F_slice_del(m, l0)
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L5
	} else {
		goto L285
	}
L285:
	;
	if v1019 < int32(0) {
		v1027 = v1019
		goto L271
	} else {
		goto L286
	}
L286:
	;
	goto L278
L287:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1030
	v1032 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1030
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1030 <= v1035 {
		v1221 = v1032
		goto L288
	} else {
		goto L289
	}
L288:
	;
	if v1221 < int32(0) {
		v1283 = v1221
		goto L1
	} else {
		goto L338
	}
L289:
	;
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1037+v1030-int32(1)))))
	if v1041 != int32(101) {
		v1221 = v1032
		goto L288
	} else {
		goto L290
	}
L290:
	;
	v1045 = v1030 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1045
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1045
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v1048)))
	if v1030 <= v1049 {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v1048)+4))
	if v1030 <= v1051 {
		v1221 = v1032
		goto L288
	} else {
		goto L294
	}
L292:
	;
	goto L293
L293:
	;
	v1213 = F_slice_del(m, l0)
	mBase = m.M
	v1214 = m.ExcPending
	if v1214 != 0 {
		goto L5
	} else {
		goto L334
	}
L294:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L298
L295:
	;
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1207 + (v1045 - v1053)
	goto L293
L296:
	;
	if v1102 != 0 {
		goto L295
	} else {
		goto L308
	}
L297:
	;
	v1102 = v1099
	goto L296
L298:
	;
	if v1061 <= v1062 {
		goto L300
	} else {
		goto L301
	}
L299:
	;
	v1099 = int32(0)
	goto L297
L300:
	;
	v1102 = int32(-1)
	goto L296
L301:
	;
	goto L302
L302:
	;
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1077 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1073+v1061-int32(1)))))
	if int32(121) < v1077 {
		goto L303
	} else {
		goto L304
	}
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1061 - int32(1)
	goto L307
L304:
	;
	v1079 = v1077 - int32(89)
	if v1079 < int32(0) {
		goto L303
	} else {
		goto L305
	}
L305:
	;
	v1082 = int32(1)
	v1086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1079)>>(uint(int32(3))%32)))+uint32(_consts[1291]))))
	if int32(base.Ui32(v1086)>>(uint(v1079&int32(7))%32))&v1082 != 0 {
		v1099 = v1082
		goto L297
	} else {
		goto L306
	}
L306:
	;
	goto L303
L307:
	;
	goto L299
L308:
	;
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L311
L309:
	;
	if v1155 != 0 {
		goto L295
	} else {
		goto L320
	}
L310:
	;
	v1155 = v1151
	goto L309
L311:
	;
	if v1111 <= v1112 {
		goto L313
	} else {
		goto L314
	}
L312:
	;
	v1151 = int32(0)
	goto L310
L313:
	;
	v1155 = int32(-1)
	goto L309
L314:
	;
	goto L315
L315:
	;
	v1124 = int32(1)
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1125+v1111-v1124))))
	if int32(121) < v1129 {
		v1151 = v1124
		goto L310
	} else {
		goto L316
	}
L316:
	;
	v1131 = v1129 - int32(97)
	if v1131 < int32(0) {
		v1151 = v1124
		goto L310
	} else {
		goto L317
	}
L317:
	;
	v1137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1131)>>(uint(int32(3))%32)))+uint32(_consts[1290]))))
	if int32(base.Ui32(v1137)>>(uint(v1131&int32(7))%32))&int32(1) == int32(0) {
		v1151 = v1124
		goto L310
	} else {
		goto L318
	}
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1111 - int32(1)
	goto L319
L319:
	;
	goto L312
L320:
	;
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L323
L321:
	;
	if v1204 == int32(0) {
		v1221 = v1032
		goto L288
	} else {
		goto L333
	}
L322:
	;
	v1204 = v1201
	goto L321
L323:
	;
	if v1163 <= v1164 {
		goto L325
	} else {
		goto L326
	}
L324:
	;
	v1201 = int32(0)
	goto L322
L325:
	;
	v1204 = int32(-1)
	goto L321
L326:
	;
	goto L327
L327:
	;
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175+v1163-int32(1)))))
	if int32(121) < v1179 {
		goto L328
	} else {
		goto L329
	}
L328:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1163 - int32(1)
	goto L332
L329:
	;
	v1181 = v1179 - int32(97)
	if v1181 < int32(0) {
		goto L328
	} else {
		goto L330
	}
L330:
	;
	v1184 = int32(1)
	v1188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1181)>>(uint(int32(3))%32)))+uint32(_consts[1290]))))
	if int32(base.Ui32(v1188)>>(uint(v1181&int32(7))%32))&v1184 != 0 {
		v1201 = v1184
		goto L322
	} else {
		goto L331
	}
L331:
	;
	goto L328
L332:
	;
	goto L324
L333:
	;
	goto L295
L334:
	;
	if int32(0) <= v1213 {
		goto L335
	} else {
		goto L336
	}
L335:
	;
	v1217 = int32(1)
	goto L337
L336:
	;
	v1217 = v1213
	goto L337
L337:
	;
	v1221 = v1217
	goto L288
L338:
	;
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1225
	v1227 = F_r_Step_5b(m, l0)
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L5
	} else {
		goto L339
	}
L339:
	;
	if v1227 < int32(0) {
		v1283 = v1227
		goto L1
	} else {
		goto L340
	}
L340:
	;
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1231
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v1233)+8))
	if v1234 == int32(0) {
		goto L341
	} else {
		goto L342
	}
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1231
	v1283 = int32(1)
	goto L1
L342:
	;
	goto L343
L343:
	;
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1244 < v1243 {
		goto L345
	} else {
		goto L346
	}
L344:
	;
	v1283 = v1270
	goto L1
L345:
	;
	v1246 = v1243
	goto L347
L346:
	;
	v1246 = v1244
	goto L347
L347:
	;
	v1248 = v1243
	goto L348
L348:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1248
	if v1248 != v1244 {
		goto L351
	} else {
		goto L352
	}
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1248
	v1265 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1248 + v1265
	v1270 = F_slice_from_s(m, l0, v1265, int32(2162747))
	mBase = m.M
	v1271 = m.ExcPending
	if v1271 != 0 {
		goto L5
	} else {
		goto L356
	}
L350:
	;
	goto L349
L351:
	;
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1255+v1248))))
	if v1257 == int32(89) {
		goto L350
	} else {
		goto L354
	}
L352:
	;
	goto L353
L353:
	;
	if v1248 == v1246 {
		goto L341
	} else {
		goto L355
	}
L354:
	;
	goto L353
L355:
	;
	v1262 = v1248 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1262
	v1248 = v1262
	goto L348
L356:
	;
	if int32(0) <= v1270 {
		goto L343
	} else {
		goto L357
	}
L357:
	;
	goto L344
}
func F_porter_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v114 int32
	_ = v114
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v152 int32
	_ = v152
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
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v322 int32
	_ = v322
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v393 int32
	_ = v393
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v444 int32
	_ = v444
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v470 int32
	_ = v470
	var v478 int32
	_ = v478
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v518 int32
	_ = v518
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v569 int32
	_ = v569
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v595 int32
	_ = v595
	var v602 int32
	_ = v602
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v640 int32
	_ = v640
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v691 int32
	_ = v691
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v717 int32
	_ = v717
	var v725 int32
	_ = v725
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v841 int32
	_ = v841
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v901 int32
	_ = v901
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v914 int32
	_ = v914
	var v920 int32
	_ = v920
	var v929 int32
	_ = v929
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1014 int32
	_ = v1014
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1044 int32
	_ = v1044
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1145 int32
	_ = v1145
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1164 int32
	_ = v1164
	var v1181 int32
	_ = v1181
	var v1188 int32
	_ = v1188
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1234 int32
	_ = v1234
	var v1236 int32
	_ = v1236
	var v1238 int32
	_ = v1238
	var v1240 int32
	_ = v1240
	var v1253 int32
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1275 int32
	_ = v1275
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1288 int32
	_ = v1288
	var v1294 int32
	_ = v1294
	var v1310 int32
	_ = v1310
	var v1317 int32
	_ = v1317
	var v1331 int32
	_ = v1331
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1363 int32
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1367 int32
	_ = v1367
	var v1369 int32
	_ = v1369
	var v1382 int32
	_ = v1382
	var v1384 int32
	_ = v1384
	var v1386 int32
	_ = v1386
	var v1404 int32
	_ = v1404
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1417 int32
	_ = v1417
	var v1423 int32
	_ = v1423
	var v1440 int32
	_ = v1440
	var v1447 int32
	_ = v1447
	var v1450 int32
	_ = v1450
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1466 int32
	_ = v1466
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1475 int32
	_ = v1475
	var v1481 int32
	_ = v1481
	var v1497 int32
	_ = v1497
	var v1501 int32
	_ = v1501
	var v1510 int32
	_ = v1510
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1521 int32
	_ = v1521
	var v1523 int32
	_ = v1523
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1533 int32
	_ = v1533
	var v1535 int32
	_ = v1535
	var v1548 int32
	_ = v1548
	var v1550 int32
	_ = v1550
	var v1552 int32
	_ = v1552
	var v1570 int32
	_ = v1570
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1583 int32
	_ = v1583
	var v1589 int32
	_ = v1589
	var v1598 int32
	_ = v1598
	var v1613 int32
	_ = v1613
	var v1616 int32
	_ = v1616
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1626 int32
	_ = v1626
	var v1629 int32
	_ = v1629
	var v1633 int32
	_ = v1633
	var v1637 int32
	_ = v1637
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1654 int32
	_ = v1654
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1675 int32
	_ = v1675
	var v1676 int32
	_ = v1676
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1699 int32
	_ = v1699
	var v1700 int32
	_ = v1700
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1741 int32
	_ = v1741
	var v1743 int32
	_ = v1743
	var v1746 int32
	_ = v1746
	var v1750 int32
	_ = v1750
	var v1754 int32
	_ = v1754
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1771 int32
	_ = v1771
	var v1773 int32
	_ = v1773
	var v1774 int32
	_ = v1774
	var v1780 int32
	_ = v1780
	var v1781 int32
	_ = v1781
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1790 int32
	_ = v1790
	var v1791 int32
	_ = v1791
	var v1797 int32
	_ = v1797
	var v1801 int32
	_ = v1801
	var v1803 int32
	_ = v1803
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1810 int32
	_ = v1810
	var v1812 int32
	_ = v1812
	var v1825 int32
	_ = v1825
	var v1826 int32
	_ = v1826
	var v1829 int32
	_ = v1829
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1840 int32
	_ = v1840
	var v1842 int32
	_ = v1842
	var v1844 int32
	_ = v1844
	var v1846 int32
	_ = v1846
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1863 int32
	_ = v1863
	var v1867 int32
	_ = v1867
	var v1869 int32
	_ = v1869
	var v1872 int32
	_ = v1872
	var v1874 int32
	_ = v1874
	var v1878 int32
	_ = v1878
	var v1882 int32
	_ = v1882
	var v1885 int32
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1888 int32
	_ = v1888
	var v1890 int32
	_ = v1890
	var v1904 int32
	_ = v1904
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1928 int32
	_ = v1928
	var v1930 int32
	_ = v1930
	var v1936 int32
	_ = v1936
	var v1938 int32
	_ = v1938
	var v1940 int32
	_ = v1940
	var v1942 int32
	_ = v1942
	var v1955 int32
	_ = v1955
	var v1957 int32
	_ = v1957
	var v1959 int32
	_ = v1959
	var v1977 int32
	_ = v1977
	var v1985 int32
	_ = v1985
	var v1986 int32
	_ = v1986
	var v1990 int32
	_ = v1990
	var v1996 int32
	_ = v1996
	var v2013 int32
	_ = v2013
	var v2020 int32
	_ = v2020
	var v2034 int32
	_ = v2034
	var v2037 int32
	_ = v2037
	var v2038 int32
	_ = v2038
	var v2055 int32
	_ = v2055
	var v2056 int32
	_ = v2056
	var v2058 int32
	_ = v2058
	var v2060 int32
	_ = v2060
	var v2066 int32
	_ = v2066
	var v2068 int32
	_ = v2068
	var v2070 int32
	_ = v2070
	var v2072 int32
	_ = v2072
	var v2085 int32
	_ = v2085
	var v2087 int32
	_ = v2087
	var v2089 int32
	_ = v2089
	var v2107 int32
	_ = v2107
	var v2115 int32
	_ = v2115
	var v2116 int32
	_ = v2116
	var v2120 int32
	_ = v2120
	var v2126 int32
	_ = v2126
	var v2142 int32
	_ = v2142
	var v2149 int32
	_ = v2149
	var v2163 int32
	_ = v2163
	var v2166 int32
	_ = v2166
	var v2167 int32
	_ = v2167
	var v2184 int32
	_ = v2184
	var v2185 int32
	_ = v2185
	var v2187 int32
	_ = v2187
	var v2189 int32
	_ = v2189
	var v2195 int32
	_ = v2195
	var v2197 int32
	_ = v2197
	var v2199 int32
	_ = v2199
	var v2201 int32
	_ = v2201
	var v2214 int32
	_ = v2214
	var v2216 int32
	_ = v2216
	var v2218 int32
	_ = v2218
	var v2236 int32
	_ = v2236
	var v2244 int32
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2249 int32
	_ = v2249
	var v2255 int32
	_ = v2255
	var v2272 int32
	_ = v2272
	var v2279 int32
	_ = v2279
	var v2282 int32
	_ = v2282
	var v2288 int32
	_ = v2288
	var v2289 int32
	_ = v2289
	var v2292 int32
	_ = v2292
	var v2297 int32
	_ = v2297
	var v2300 int32
	_ = v2300
	var v2302 int32
	_ = v2302
	var v2303 int32
	_ = v2303
	var v2306 int32
	_ = v2306
	var v2308 int32
	_ = v2308
	var v2309 int32
	_ = v2309
	var v2318 int32
	_ = v2318
	var v2320 int32
	_ = v2320
	var v2326 int32
	_ = v2326
	var v2327 int32
	_ = v2327
	var v2330 int32
	_ = v2330
	var v2339 int32
	_ = v2339
	var v2341 int32
	_ = v2341
	var v2346 int32
	_ = v2346
	var v2348 int32
	_ = v2348
	var v2355 int32
	_ = v2355
	var v2358 int32
	_ = v2358
	var v2362 int32
	_ = v2362
	var v2369 int32
	_ = v2369
	var v2370 int32
	_ = v2370
	var v2384 int32
	_ = v2384
	var v2389 int32
	_ = v2389
	var v2394 int32
	_ = v2394
	var v2395 int32
	_ = v2395
	var v2407 int32
	_ = v2407
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v10 == v12 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v2407
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	v37 = v10
	goto L8
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v10))))
	if v16 != int32(121) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v19 = int32(1)
	v20 = v10 + v19
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v20
	v25 = F_slice_from_s(m, l0, v19, int32(2190533))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	if v25 < int32(0) {
		v2407 = v25
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = int32(1)
	goto L2
L8:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L15
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v245))) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v245)+4)) = v246
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v271 = v249
	goto L70
L10:
	;
	goto L9
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v37
	v231 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v162 + v231
	v236 = F_slice_from_s(m, l0, v231, int32(2190538))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L5
	} else {
		goto L65
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v37
	goto L46
L13:
	;
	if v159 != 0 {
		goto L37
	} else {
		goto L38
	}
L14:
	;
	v159 = v152
	goto L13
L15:
	;
	if v54 <= v53 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v152 = int32(0)
	goto L14
L17:
	;
	v159 = int32(-1)
	goto L13
L18:
	;
	goto L19
L19:
	;
	v70 = int32(1)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+v55))))
	if base.Ui32(v72) < base.Ui32(int32(192)) {
		v129 = v72
		v130 = v70
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if int32(121) < v129 {
		v152 = v130
		goto L14
	} else {
		goto L33
	}
L21:
	;
	v76 = v53 + int32(1)
	if v76 == v54 {
		v129 = v72
		v130 = v70
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76+v55))))
	v81 = v79 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v72) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+v55))))
	v97 = v95 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v72) {
		goto L29
	} else {
		goto L30
	}
L24:
	;
	v85 = v53 + int32(2)
	if v85 != v54 {
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v129 = v72<<(uint(int32(6))%32)&int32(1984) | v81
	v130 = int32(2)
	goto L20
L27:
	;
	goto L26
L28:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+v101))))
	v129 = v114&int32(63) | (v72<<(uint(int32(18))%32)&int32(1835008) | v81<<(uint(int32(12))%32) | v97<<(uint(int32(6))%32))
	v130 = int32(4)
	goto L20
L29:
	;
	v101 = v53 + int32(3)
	if v101 != v54 {
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v129 = v72<<(uint(int32(12))%32)&int32(61440) | v81<<(uint(int32(6))%32) | v97
	v130 = int32(3)
	goto L20
L32:
	;
	goto L31
L33:
	;
	v134 = v129 - int32(97)
	if v134 < int32(0) {
		v152 = v130
		goto L14
	} else {
		goto L34
	}
L34:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v134)>>(uint(int32(3))%32)))+uint32(_consts[1305]))))
	if int32(base.Ui32(v140)>>(uint(v134&int32(7))%32))&int32(1) == int32(0) {
		v152 = v130
		goto L14
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v130 + v53
	goto L36
L36:
	;
	goto L16
L37:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v171 = v160
	v172 = v161
	goto L12
L38:
	;
	goto L39
L39:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v162
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v165 == v162 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v171 = v162
	v172 = v164
	goto L12
L41:
	;
	goto L42
L42:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164+v162))))
	if v168 == int32(121) {
		goto L11
	} else {
		goto L43
	}
L43:
	;
	v171 = v165
	v172 = v164
	goto L12
L44:
	;
	if v226 < int32(0) {
		goto L10
	} else {
		goto L64
	}
L46:
	;
	goto L47
L47:
	;
	goto L48
L48:
	;
	v181 = v37
	v183 = int32(1)
	goto L51
L50:
	;
	v226 = v211
	goto L44
L51:
	;
	if v171 <= v181 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L50
L53:
	;
	v226 = int32(-1)
	goto L44
L54:
	;
	goto L55
L55:
	;
	v188 = v181 + int32(1)
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172+v181))))
	if base.Ui32(v190) < base.Ui32(int32(192)) {
		v211 = v188
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v212 = int32(1)
	if v212 < v183 {
		v181 = v211
		v183 = v183 - v212
		goto L51
	} else {
		goto L63
	}
L57:
	;
	if v171 <= v188 {
		v211 = v188
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v197 = v188
	goto L59
L59:
	;
	v200 = int32(*(*int8)(unsafe.Add(mBase, uint32(v172+v197))))
	if int32(-65) < v200 {
		v211 = v197
		goto L56
	} else {
		goto L61
	}
L60:
	;
	v211 = v171
	goto L56
L61:
	;
	v204 = v197 + int32(1)
	if v204 != v171 {
		v197 = v204
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	goto L52
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v226
	v37 = v226
	goto L8
L65:
	;
	if v236 < int32(0) {
		v2407 = v236
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v240)+8)) = int32(1)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v37 = v243
	goto L8
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v249
	v745 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v745
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v745
	if v745 <= v249 {
		goto L170
	} else {
		goto L171
	}
L68:
	;
	if v366 < int32(0) {
		goto L67
	} else {
		goto L93
	}
L69:
	;
	v366 = v338
	goto L68
L70:
	;
	if v262 <= v271 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v366 = int32(-1)
	goto L68
L73:
	;
	goto L74
L74:
	;
	v278 = int32(1)
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271+v263))))
	if base.Ui32(v280) < base.Ui32(int32(192)) {
		v337 = v280
		v338 = v278
		goto L75
	} else {
		goto L76
	}
L75:
	;
	if int32(121) < v337 {
		goto L88
	} else {
		goto L89
	}
L76:
	;
	v284 = v271 + int32(1)
	if v284 == v262 {
		v337 = v280
		v338 = v278
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284+v263))))
	v289 = v287 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v280) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293+v263))))
	v305 = v303 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v280) {
		goto L84
	} else {
		goto L85
	}
L79:
	;
	v293 = v271 + int32(2)
	if v293 != v262 {
		goto L78
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v337 = v280<<(uint(int32(6))%32)&int32(1984) | v289
	v338 = int32(2)
	goto L75
L82:
	;
	goto L81
L83:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263+v309))))
	v337 = v322&int32(63) | (v280<<(uint(int32(18))%32)&int32(1835008) | v289<<(uint(int32(12))%32) | v305<<(uint(int32(6))%32))
	v338 = int32(4)
	goto L75
L84:
	;
	v309 = v271 + int32(3)
	if v309 != v262 {
		goto L83
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v337 = v280<<(uint(int32(12))%32)&int32(61440) | v289<<(uint(int32(6))%32) | v305
	v338 = int32(3)
	goto L75
L87:
	;
	goto L86
L88:
	;
	v355 = v338 + v271
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v355
	v271 = v355
	goto L70
L89:
	;
	v342 = v337 - int32(97)
	if v342 < int32(0) {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v342)>>(uint(int32(3))%32)))+uint32(_consts[1305]))))
	if int32(base.Ui32(v348)>>(uint(v342&int32(7))%32))&int32(1) != 0 {
		goto L69
	} else {
		goto L91
	}
L91:
	;
	goto L88
L93:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v370 = v369 + v366
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v370
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v393 = v370
	goto L96
L94:
	;
	if v489 < int32(0) {
		goto L67
	} else {
		goto L118
	}
L95:
	;
	v489 = v460
	goto L94
L96:
	;
	if v384 <= v393 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v489 = int32(-1)
	goto L94
L99:
	;
	goto L100
L100:
	;
	v400 = int32(1)
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393+v385))))
	if base.Ui32(v402) < base.Ui32(int32(192)) {
		v459 = v402
		v460 = v400
		goto L101
	} else {
		goto L102
	}
L101:
	;
	if int32(121) < v459 {
		goto L95
	} else {
		goto L114
	}
L102:
	;
	v406 = v393 + int32(1)
	if v406 == v384 {
		v459 = v402
		v460 = v400
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406+v385))))
	v411 = v409 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v402) {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415+v385))))
	v427 = v425 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v402) {
		goto L110
	} else {
		goto L111
	}
L105:
	;
	v415 = v393 + int32(2)
	if v415 != v384 {
		goto L104
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v459 = v402<<(uint(int32(6))%32)&int32(1984) | v411
	v460 = int32(2)
	goto L101
L108:
	;
	goto L107
L109:
	;
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385+v431))))
	v459 = v444&int32(63) | (v402<<(uint(int32(18))%32)&int32(1835008) | v411<<(uint(int32(12))%32) | v427<<(uint(int32(6))%32))
	v460 = int32(4)
	goto L101
L110:
	;
	v431 = v393 + int32(3)
	if v431 != v384 {
		goto L109
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v459 = v402<<(uint(int32(12))%32)&int32(61440) | v411<<(uint(int32(6))%32) | v427
	v460 = int32(3)
	goto L101
L113:
	;
	goto L112
L114:
	;
	v464 = v459 - int32(97)
	if v464 < int32(0) {
		goto L95
	} else {
		goto L115
	}
L115:
	;
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v464)>>(uint(int32(3))%32)))+uint32(_consts[1305]))))
	if int32(base.Ui32(v470)>>(uint(v464&int32(7))%32))&int32(1) == int32(0) {
		goto L95
	} else {
		goto L116
	}
L116:
	;
	v478 = v460 + v393
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v478
	v393 = v478
	goto L96
L118:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v493 = v492 + v489
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v493
	v495 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v495)+4)) = v493
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v518 = v508
	goto L121
L119:
	;
	if v613 < int32(0) {
		goto L67
	} else {
		goto L144
	}
L120:
	;
	v613 = v585
	goto L119
L121:
	;
	if v509 <= v518 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v613 = int32(-1)
	goto L119
L124:
	;
	goto L125
L125:
	;
	v525 = int32(1)
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518+v510))))
	if base.Ui32(v527) < base.Ui32(int32(192)) {
		v584 = v527
		v585 = v525
		goto L126
	} else {
		goto L127
	}
L126:
	;
	if int32(121) < v584 {
		goto L139
	} else {
		goto L140
	}
L127:
	;
	v531 = v518 + int32(1)
	if v531 == v509 {
		v584 = v527
		v585 = v525
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v531+v510))))
	v536 = v534 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v527) {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	v550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v540+v510))))
	v552 = v550 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v527) {
		goto L135
	} else {
		goto L136
	}
L130:
	;
	v540 = v518 + int32(2)
	if v540 != v509 {
		goto L129
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v584 = v527<<(uint(int32(6))%32)&int32(1984) | v536
	v585 = int32(2)
	goto L126
L133:
	;
	goto L132
L134:
	;
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510+v556))))
	v584 = v569&int32(63) | (v527<<(uint(int32(18))%32)&int32(1835008) | v536<<(uint(int32(12))%32) | v552<<(uint(int32(6))%32))
	v585 = int32(4)
	goto L126
L135:
	;
	v556 = v518 + int32(3)
	if v556 != v509 {
		goto L134
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v584 = v527<<(uint(int32(12))%32)&int32(61440) | v536<<(uint(int32(6))%32) | v552
	v585 = int32(3)
	goto L126
L138:
	;
	goto L137
L139:
	;
	v602 = v585 + v518
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v602
	v518 = v602
	goto L121
L140:
	;
	v589 = v584 - int32(97)
	if v589 < int32(0) {
		goto L139
	} else {
		goto L141
	}
L141:
	;
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v589)>>(uint(int32(3))%32)))+uint32(_consts[1305]))))
	if int32(base.Ui32(v595)>>(uint(v589&int32(7))%32))&int32(1) != 0 {
		goto L120
	} else {
		goto L142
	}
L142:
	;
	goto L139
L144:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v617 = v616 + v613
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v617
	v631 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v640 = v617
	goto L147
L145:
	;
	if v736 < int32(0) {
		goto L67
	} else {
		goto L169
	}
L146:
	;
	v736 = v707
	goto L145
L147:
	;
	if v631 <= v640 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v736 = int32(-1)
	goto L145
L150:
	;
	goto L151
L151:
	;
	v647 = int32(1)
	v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v640+v632))))
	if base.Ui32(v649) < base.Ui32(int32(192)) {
		v706 = v649
		v707 = v647
		goto L152
	} else {
		goto L153
	}
L152:
	;
	if int32(121) < v706 {
		goto L146
	} else {
		goto L165
	}
L153:
	;
	v653 = v640 + int32(1)
	if v653 == v631 {
		v706 = v649
		v707 = v647
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v653+v632))))
	v658 = v656 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v649) {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v662+v632))))
	v674 = v672 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v649) {
		goto L161
	} else {
		goto L162
	}
L156:
	;
	v662 = v640 + int32(2)
	if v662 != v631 {
		goto L155
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v706 = v649<<(uint(int32(6))%32)&int32(1984) | v658
	v707 = int32(2)
	goto L152
L159:
	;
	goto L158
L160:
	;
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v632+v678))))
	v706 = v691&int32(63) | (v649<<(uint(int32(18))%32)&int32(1835008) | v658<<(uint(int32(12))%32) | v674<<(uint(int32(6))%32))
	v707 = int32(4)
	goto L152
L161:
	;
	v678 = v640 + int32(3)
	if v678 != v631 {
		goto L160
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	v706 = v649<<(uint(int32(12))%32)&int32(61440) | v658<<(uint(int32(6))%32) | v674
	v707 = int32(3)
	goto L152
L164:
	;
	goto L163
L165:
	;
	v711 = v706 - int32(97)
	if v711 < int32(0) {
		goto L146
	} else {
		goto L166
	}
L166:
	;
	v717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v711)>>(uint(int32(3))%32)))+uint32(_consts[1305]))))
	if int32(base.Ui32(v717)>>(uint(v711&int32(7))%32))&int32(1) == int32(0) {
		goto L146
	} else {
		goto L167
	}
L167:
	;
	v725 = v707 + v640
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v725
	v640 = v725
	goto L147
L169:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v740 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v739))) = v740 + v736
	goto L67
L170:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v783
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v783
	v787 = v783 - int32(1)
	v788 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v787 <= v788 {
		goto L184
	} else {
		goto L185
	}
L171:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v749+v745-int32(1)))))
	if v753 != int32(115) {
		goto L170
	} else {
		goto L172
	}
L172:
	;
	v758 = F_find_among_b(m, l0, int32(4295296), int32(4))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L5
	} else {
		goto L173
	}
L173:
	;
	if v758 == int32(0) {
		goto L170
	} else {
		goto L174
	}
L174:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v762
	switch v758 - int32(1) {
	case 0:
		goto L177
	case 1:
		goto L176
	case 2:
		goto L175
	default:
		goto L170
	}
L175:
	;
	v778 = F_slice_del(m, l0)
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L5
	} else {
		goto L182
	}
L176:
	;
	v774 = F_slice_from_s(m, l0, int32(1), int32(2190542))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L5
	} else {
		goto L180
	}
L177:
	;
	v768 = F_slice_from_s(m, l0, int32(2), int32(2190540))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L5
	} else {
		goto L178
	}
L178:
	;
	if int32(0) <= v768 {
		goto L170
	} else {
		goto L179
	}
L179:
	;
	v2407 = v768
	goto L1
L180:
	;
	if int32(0) <= v774 {
		goto L170
	} else {
		goto L181
	}
L181:
	;
	v2407 = v774
	goto L1
L182:
	;
	if v778 < int32(0) {
		v2407 = v778
		goto L1
	} else {
		goto L183
	}
L183:
	;
	goto L170
L184:
	;
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1466
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1466
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1466 <= v1469 {
		goto L317
	} else {
		goto L318
	}
L185:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v790+v787))))
	switch v792 - int32(100) {
	case 0, 3:
		goto L186
	default:
		goto L184
	}
L186:
	;
	v797 = F_find_among_b(m, l0, int32(4295376), int32(3))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L5
	} else {
		goto L187
	}
L187:
	;
	if v797 == int32(0) {
		goto L184
	} else {
		goto L188
	}
L188:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v801
	switch v797 - int32(1) {
	case 0:
		goto L190
	case 1:
		goto L189
	default:
		goto L184
	}
L189:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v831 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v832 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v841 = v831
	goto L196
L190:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v805)+4))
	if v801 < v806 {
		goto L184
	} else {
		goto L191
	}
L191:
	;
	v810 = F_slice_from_s(m, l0, int32(2), int32(2190553))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L5
	} else {
		goto L192
	}
L192:
	;
	if int32(0) <= v810 {
		goto L184
	} else {
		goto L193
	}
L193:
	;
	v2407 = v810
	goto L1
L194:
	;
	if v944 < int32(0) {
		goto L184
	} else {
		goto L213
	}
L195:
	;
	v944 = int32(-1)
	goto L194
L196:
	;
	if v841 <= v832 {
		goto L195
	} else {
		goto L198
	}
L198:
	;
	v849 = int32(1)
	v850 = v841 - v849
	v852 = int32(*(*int8)(unsafe.Add(mBase, uint32(v828+v850))))
	v854 = v852 & int32(255)
	if v850 == v832 {
		v909 = v854
		v910 = v849
		goto L199
	} else {
		goto L200
	}
L199:
	;
	if int32(121) < v909 {
		goto L208
	} else {
		goto L209
	}
L200:
	;
	if int32(0) <= v852 {
		v909 = v854
		v910 = v849
		goto L199
	} else {
		goto L201
	}
L201:
	;
	v860 = v854 & int32(63)
	v862 = v841 - int32(2)
	v864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v828+v862))))
	v866 = v864 << (uint(int32(6)) % 32)
	if base.B2i32(v862 != v832)&base.B2i32(base.Ui32(v864) < base.Ui32(int32(192))) == int32(0) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v909 = v866&int32(1984) | v860
	v910 = int32(2)
	goto L199
L203:
	;
	goto L204
L204:
	;
	v879 = v866&int32(4032) | v860
	v881 = v841 - int32(3)
	v883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v828+v881))))
	if base.B2i32(v881 != v832)&base.B2i32(base.Ui32(v883) < base.Ui32(int32(224))) == int32(0) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v909 = v883<<(uint(int32(12))%32)&int32(61440) | v879
	v910 = int32(3)
	goto L199
L206:
	;
	goto L207
L207:
	;
	v901 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v841+(v828-int32(4))))))
	v909 = v883<<(uint(int32(12))%32)&int32(258048) | v901&int32(7)<<(uint(int32(18))%32) | v879
	v910 = int32(4)
	goto L199
L208:
	;
	v929 = v841 - v910
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v929
	v841 = v929
	goto L196
L209:
	;
	v914 = v909 - int32(97)
	if v914 < int32(0) {
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v920 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v914)>>(uint(int32(3))%32)))+uint32(_consts[1305]))))
	if int32(base.Ui32(v920)>>(uint(v914&int32(7))%32))&int32(1) == int32(0) {
		goto L208
	} else {
		goto L211
	}
L211:
	;
	v944 = v910
	goto L194
L213:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v947 + (v801 - v814)
	v951 = F_slice_del(m, l0)
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L5
	} else {
		goto L214
	}
L214:
	;
	if v951 < int32(0) {
		v2407 = v951
		goto L1
	} else {
		goto L215
	}
L215:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v956 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v957 = v955 - v956
	v959 = v955 - int32(1)
	v960 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v959 <= v960 {
		v1054 = v955
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1055)+4))
	if v1054 != v1056 {
		goto L184
	} else {
		goto L248
	}
L217:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v964 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v962+v959))))
	if v964&int32(224) != int32(96) {
		v1054 = v955
		goto L216
	} else {
		goto L218
	}
L218:
	;
	if int32(1)<<(uint(v964)%32)&int32(68514004) == int32(0) {
		v1054 = v955
		goto L216
	} else {
		goto L219
	}
L219:
	;
	v977 = F_find_among_b(m, l0, int32(4295440), int32(13))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L5
	} else {
		goto L220
	}
L220:
	;
	v979 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v980 = v979 + v957
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v980
	switch v977 - int32(1) {
	case 0:
		goto L222
	case 1:
		goto L221
	case 2:
		v1054 = v980
		goto L216
	default:
		goto L184
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v980
	v992 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v993 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L227
L222:
	;
	v986 = F_insert_s(m, l0, v980, v980, int32(1), int32(2190555))
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L5
	} else {
		goto L223
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v980
	if int32(0) <= v986 {
		goto L184
	} else {
		goto L224
	}
L224:
	;
	v2407 = v986
	goto L1
L225:
	;
	if v1044 < int32(0) {
		goto L184
	} else {
		goto L245
	}
L227:
	;
	goto L228
L228:
	;
	goto L229
L229:
	;
	v1000 = v980
	v1002 = int32(1)
	goto L232
L231:
	;
	v1044 = v1026
	goto L225
L232:
	;
	if v1000 <= v993 {
		goto L234
	} else {
		goto L235
	}
L233:
	;
	goto L231
L234:
	;
	v1044 = int32(-1)
	goto L225
L235:
	;
	goto L236
L236:
	;
	v1007 = v1000 - int32(1)
	v1009 = int32(*(*int8)(unsafe.Add(mBase, uint32(v992+v1007))))
	if int32(0) <= v1009 {
		v1026 = v1007
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v1030 = int32(1)
	if v1030 < v1002 {
		v1000 = v1026
		v1002 = v1002 - v1030
		goto L232
	} else {
		goto L244
	}
L238:
	;
	if v1007 <= v993 {
		v1026 = v1007
		goto L237
	} else {
		goto L239
	}
L239:
	;
	v1014 = v1007
	goto L240
L240:
	;
	v1019 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v992+v1014))))
	if base.Ui32(int32(191)) < base.Ui32(v1019) {
		v1026 = v1014
		goto L237
	} else {
		goto L242
	}
L241:
	;
	v1026 = v993
	goto L237
L242:
	;
	v1023 = v1014 - int32(1)
	if v993 < v1023 {
		v1014 = v1023
		goto L240
	} else {
		goto L243
	}
L243:
	;
	goto L241
L244:
	;
	goto L233
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1044
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1044
	v1049 = F_slice_del(m, l0)
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L5
	} else {
		goto L246
	}
L246:
	;
	if int32(0) <= v1049 {
		goto L184
	} else {
		goto L247
	}
L247:
	;
	v2407 = v1049
	goto L1
L248:
	;
	v1058 = int32(0)
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L252
L249:
	;
	if v1450 == int32(0) {
		goto L184
	} else {
		goto L314
	}
L250:
	;
	if v1188 != 0 {
		v1450 = v1058
		goto L249
	} else {
		goto L269
	}
L251:
	;
	v1188 = v1181
	goto L250
L252:
	;
	if v1075 <= v1076 {
		v1181 = int32(-1)
		goto L251
	} else {
		goto L254
	}
L253:
	;
	v1181 = int32(0)
	goto L251
L254:
	;
	v1093 = int32(1)
	v1094 = v1075 - v1093
	v1096 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1072+v1094))))
	v1098 = v1096 & int32(255)
	if v1094 == v1076 {
		v1153 = v1098
		v1154 = v1093
		goto L255
	} else {
		goto L256
	}
L255:
	;
	if int32(121) < v1153 {
		goto L264
	} else {
		goto L265
	}
L256:
	;
	if int32(0) <= v1096 {
		v1153 = v1098
		v1154 = v1093
		goto L255
	} else {
		goto L257
	}
L257:
	;
	v1104 = v1098 & int32(63)
	v1106 = v1075 - int32(2)
	v1108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1072+v1106))))
	v1110 = v1108 << (uint(int32(6)) % 32)
	if base.B2i32(v1106 != v1076)&base.B2i32(base.Ui32(v1108) < base.Ui32(int32(192))) == int32(0) {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v1153 = v1110&int32(1984) | v1104
	v1154 = int32(2)
	goto L255
L259:
	;
	goto L260
L260:
	;
	v1123 = v1110&int32(4032) | v1104
	v1125 = v1075 - int32(3)
	v1127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1072+v1125))))
	if base.B2i32(v1125 != v1076)&base.B2i32(base.Ui32(v1127) < base.Ui32(int32(224))) == int32(0) {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	v1153 = v1127<<(uint(int32(12))%32)&int32(61440) | v1123
	v1154 = int32(3)
	goto L255
L262:
	;
	goto L263
L263:
	;
	v1145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075+(v1072-int32(4))))))
	v1153 = v1127<<(uint(int32(12))%32)&int32(258048) | v1145&int32(7)<<(uint(int32(18))%32) | v1123
	v1154 = int32(4)
	goto L255
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1075 - v1154
	goto L268
L265:
	;
	v1158 = v1153 - int32(89)
	if v1158 < int32(0) {
		goto L264
	} else {
		goto L266
	}
L266:
	;
	v1164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1158)>>(uint(int32(3))%32)))+uint32(_consts[1306]))))
	if int32(base.Ui32(v1164)>>(uint(v1158&int32(7))%32))&int32(1) == int32(0) {
		goto L264
	} else {
		goto L267
	}
L267:
	;
	v1188 = v1154
	goto L250
L268:
	;
	goto L253
L269:
	;
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L272
L270:
	;
	if v1317 != 0 {
		v1450 = v1058
		goto L249
	} else {
		goto L294
	}
L271:
	;
	v1317 = v1310
	goto L270
L272:
	;
	if v1205 <= v1206 {
		v1310 = int32(-1)
		goto L271
	} else {
		goto L274
	}
L273:
	;
	v1310 = int32(0)
	goto L271
L274:
	;
	v1223 = int32(1)
	v1224 = v1205 - v1223
	v1226 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1202+v1224))))
	v1228 = v1226 & int32(255)
	if v1224 == v1206 {
		v1283 = v1228
		v1284 = v1223
		goto L275
	} else {
		goto L276
	}
L275:
	;
	if int32(121) < v1283 {
		goto L284
	} else {
		goto L285
	}
L276:
	;
	if int32(0) <= v1226 {
		v1283 = v1228
		v1284 = v1223
		goto L275
	} else {
		goto L277
	}
L277:
	;
	v1234 = v1228 & int32(63)
	v1236 = v1205 - int32(2)
	v1238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1202+v1236))))
	v1240 = v1238 << (uint(int32(6)) % 32)
	if base.B2i32(v1236 != v1206)&base.B2i32(base.Ui32(v1238) < base.Ui32(int32(192))) == int32(0) {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v1283 = v1240&int32(1984) | v1234
	v1284 = int32(2)
	goto L275
L279:
	;
	goto L280
L280:
	;
	v1253 = v1240&int32(4032) | v1234
	v1255 = v1205 - int32(3)
	v1257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1202+v1255))))
	if base.B2i32(v1255 != v1206)&base.B2i32(base.Ui32(v1257) < base.Ui32(int32(224))) == int32(0) {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1283 = v1257<<(uint(int32(12))%32)&int32(61440) | v1253
	v1284 = int32(3)
	goto L275
L282:
	;
	goto L283
L283:
	;
	v1275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1205+(v1202-int32(4))))))
	v1283 = v1257<<(uint(int32(12))%32)&int32(258048) | v1275&int32(7)<<(uint(int32(18))%32) | v1253
	v1284 = int32(4)
	goto L275
L284:
	;
	v1317 = v1284
	goto L270
L285:
	;
	goto L286
L286:
	;
	v1288 = v1283 - int32(97)
	if v1288 < int32(0) {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v1317 = v1284
	goto L270
L288:
	;
	goto L289
L289:
	;
	v1294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1288)>>(uint(int32(3))%32)))+uint32(_consts[1305]))))
	if int32(base.Ui32(v1294)>>(uint(v1288&int32(7))%32))&int32(1) == int32(0) {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v1317 = v1284
	goto L270
L291:
	;
	goto L292
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1205 - v1284
	goto L293
L293:
	;
	goto L273
L294:
	;
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L297
L295:
	;
	v1450 = base.B2i32(v1447 == int32(0))
	goto L249
L296:
	;
	v1447 = v1440
	goto L295
L297:
	;
	if v1334 <= v1335 {
		v1440 = int32(-1)
		goto L296
	} else {
		goto L299
	}
L298:
	;
	v1440 = int32(0)
	goto L296
L299:
	;
	v1352 = int32(1)
	v1353 = v1334 - v1352
	v1355 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1331+v1353))))
	v1357 = v1355 & int32(255)
	if v1353 == v1335 {
		v1412 = v1357
		v1413 = v1352
		goto L300
	} else {
		goto L301
	}
L300:
	;
	if int32(121) < v1412 {
		goto L309
	} else {
		goto L310
	}
L301:
	;
	if int32(0) <= v1355 {
		v1412 = v1357
		v1413 = v1352
		goto L300
	} else {
		goto L302
	}
L302:
	;
	v1363 = v1357 & int32(63)
	v1365 = v1334 - int32(2)
	v1367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1331+v1365))))
	v1369 = v1367 << (uint(int32(6)) % 32)
	if base.B2i32(v1365 != v1335)&base.B2i32(base.Ui32(v1367) < base.Ui32(int32(192))) == int32(0) {
		goto L303
	} else {
		goto L304
	}
L303:
	;
	v1412 = v1369&int32(1984) | v1363
	v1413 = int32(2)
	goto L300
L304:
	;
	goto L305
L305:
	;
	v1382 = v1369&int32(4032) | v1363
	v1384 = v1334 - int32(3)
	v1386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1331+v1384))))
	if base.B2i32(v1384 != v1335)&base.B2i32(base.Ui32(v1386) < base.Ui32(int32(224))) == int32(0) {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	v1412 = v1386<<(uint(int32(12))%32)&int32(61440) | v1382
	v1413 = int32(3)
	goto L300
L307:
	;
	goto L308
L308:
	;
	v1404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1334+(v1331-int32(4))))))
	v1412 = v1386<<(uint(int32(12))%32)&int32(258048) | v1404&int32(7)<<(uint(int32(18))%32) | v1382
	v1413 = int32(4)
	goto L300
L309:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1334 - v1413
	goto L313
L310:
	;
	v1417 = v1412 - int32(97)
	if v1417 < int32(0) {
		goto L309
	} else {
		goto L311
	}
L311:
	;
	v1423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1417)>>(uint(int32(3))%32)))+uint32(_consts[1305]))))
	if int32(base.Ui32(v1423)>>(uint(v1417&int32(7))%32))&int32(1) == int32(0) {
		goto L309
	} else {
		goto L312
	}
L312:
	;
	v1447 = v1413
	goto L295
L313:
	;
	goto L298
L314:
	;
	v1453 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1454 = v1453 + v957
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1454
	v1458 = F_insert_s(m, l0, v1454, v1454, int32(1), int32(2190556))
	mBase = m.M
	v1459 = m.ExcPending
	if v1459 != 0 {
		goto L5
	} else {
		goto L315
	}
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1454
	if v1458 < int32(0) {
		v2407 = v1458
		goto L1
	} else {
		goto L316
	}
L316:
	;
	goto L184
L317:
	;
	v1626 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1626
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1626
	v1629 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1626-int32(2) <= v1629 {
		goto L342
	} else {
		goto L343
	}
L318:
	;
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1471+v1466-int32(1)))))
	if v1475|int32(32) != int32(121) {
		goto L317
	} else {
		goto L319
	}
L319:
	;
	v1481 = v1466 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1481
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1510 = v1481
	goto L322
L320:
	;
	if v1613 < int32(0) {
		goto L317
	} else {
		goto L339
	}
L321:
	;
	v1613 = int32(-1)
	goto L320
L322:
	;
	if v1510 <= v1501 {
		goto L321
	} else {
		goto L324
	}
L324:
	;
	v1518 = int32(1)
	v1519 = v1510 - v1518
	v1521 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1497+v1519))))
	v1523 = v1521 & int32(255)
	if v1519 == v1501 {
		v1578 = v1523
		v1579 = v1518
		goto L325
	} else {
		goto L326
	}
L325:
	;
	if int32(121) < v1578 {
		goto L334
	} else {
		goto L335
	}
L326:
	;
	if int32(0) <= v1521 {
		v1578 = v1523
		v1579 = v1518
		goto L325
	} else {
		goto L327
	}
L327:
	;
	v1529 = v1523 & int32(63)
	v1531 = v1510 - int32(2)
	v1533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1497+v1531))))
	v1535 = v1533 << (uint(int32(6)) % 32)
	if base.B2i32(v1531 != v1501)&base.B2i32(base.Ui32(v1533) < base.Ui32(int32(192))) == int32(0) {
		goto L328
	} else {
		goto L329
	}
L328:
	;
	v1578 = v1535&int32(1984) | v1529
	v1579 = int32(2)
	goto L325
L329:
	;
	goto L330
L330:
	;
	v1548 = v1535&int32(4032) | v1529
	v1550 = v1510 - int32(3)
	v1552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1497+v1550))))
	if base.B2i32(v1550 != v1501)&base.B2i32(base.Ui32(v1552) < base.Ui32(int32(224))) == int32(0) {
		goto L331
	} else {
		goto L332
	}
L331:
	;
	v1578 = v1552<<(uint(int32(12))%32)&int32(61440) | v1548
	v1579 = int32(3)
	goto L325
L332:
	;
	goto L333
L333:
	;
	v1570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1510+(v1497-int32(4))))))
	v1578 = v1552<<(uint(int32(12))%32)&int32(258048) | v1570&int32(7)<<(uint(int32(18))%32) | v1548
	v1579 = int32(4)
	goto L325
L334:
	;
	v1598 = v1510 - v1579
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1598
	v1510 = v1598
	goto L322
L335:
	;
	v1583 = v1578 - int32(97)
	if v1583 < int32(0) {
		goto L334
	} else {
		goto L336
	}
L336:
	;
	v1589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1583)>>(uint(int32(3))%32)))+uint32(_consts[1305]))))
	if int32(base.Ui32(v1589)>>(uint(v1583&int32(7))%32))&int32(1) == int32(0) {
		goto L334
	} else {
		goto L337
	}
L337:
	;
	v1613 = v1579
	goto L320
L339:
	;
	v1616 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1616 - v1613
	v1621 = F_slice_from_s(m, l0, int32(1), int32(2190594))
	mBase = m.M
	v1622 = m.ExcPending
	if v1622 != 0 {
		goto L5
	} else {
		goto L340
	}
L340:
	;
	if v1621 < int32(0) {
		v2407 = v1621
		goto L1
	} else {
		goto L341
	}
L341:
	;
	goto L317
L342:
	;
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1741
	v1743 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1741
	v1746 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1741-int32(2) <= v1746 {
		v1797 = v1743
		goto L388
	} else {
		goto L389
	}
L343:
	;
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1633+v1626-int32(1)))))
	if v1637&int32(224) != int32(96) {
		goto L342
	} else {
		goto L344
	}
L344:
	;
	if int32(1)<<(uint(v1637)%32)&int32(815616) == int32(0) {
		goto L342
	} else {
		goto L345
	}
L345:
	;
	v1650 = F_find_among_b(m, l0, int32(4295712), int32(20))
	mBase = m.M
	v1651 = m.ExcPending
	if v1651 != 0 {
		goto L5
	} else {
		goto L346
	}
L346:
	;
	if v1650 == int32(0) {
		goto L342
	} else {
		goto L347
	}
L347:
	;
	v1654 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1654
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1657 = *(*int32)(unsafe.Add(mBase, uint32(v1656)+4))
	if v1654 < v1657 {
		goto L342
	} else {
		goto L348
	}
L348:
	;
	switch v1650 - int32(1) {
	case 0:
		goto L361
	case 1:
		goto L360
	case 2:
		goto L359
	case 3:
		goto L358
	case 4:
		goto L357
	case 5:
		goto L356
	case 6:
		goto L355
	case 7:
		goto L354
	case 8:
		goto L353
	case 9:
		goto L352
	case 10:
		goto L351
	case 11:
		goto L350
	case 12:
		goto L349
	default:
		goto L342
	}
L349:
	;
	v1735 = F_slice_from_s(m, l0, int32(3), int32(2190632))
	mBase = m.M
	v1736 = m.ExcPending
	if v1736 != 0 {
		goto L5
	} else {
		goto L386
	}
L350:
	;
	v1729 = F_slice_from_s(m, l0, int32(3), int32(2190629))
	mBase = m.M
	v1730 = m.ExcPending
	if v1730 != 0 {
		goto L5
	} else {
		goto L384
	}
L351:
	;
	v1723 = F_slice_from_s(m, l0, int32(3), int32(2190626))
	mBase = m.M
	v1724 = m.ExcPending
	if v1724 != 0 {
		goto L5
	} else {
		goto L382
	}
L352:
	;
	v1717 = F_slice_from_s(m, l0, int32(3), int32(2190623))
	mBase = m.M
	v1718 = m.ExcPending
	if v1718 != 0 {
		goto L5
	} else {
		goto L380
	}
L353:
	;
	v1711 = F_slice_from_s(m, l0, int32(2), int32(2190621))
	mBase = m.M
	v1712 = m.ExcPending
	if v1712 != 0 {
		goto L5
	} else {
		goto L378
	}
L354:
	;
	v1705 = F_slice_from_s(m, l0, int32(3), int32(2190618))
	mBase = m.M
	v1706 = m.ExcPending
	if v1706 != 0 {
		goto L5
	} else {
		goto L376
	}
L355:
	;
	v1699 = F_slice_from_s(m, l0, int32(3), int32(2190615))
	mBase = m.M
	v1700 = m.ExcPending
	if v1700 != 0 {
		goto L5
	} else {
		goto L374
	}
L356:
	;
	v1693 = F_slice_from_s(m, l0, int32(1), int32(2190614))
	mBase = m.M
	v1694 = m.ExcPending
	if v1694 != 0 {
		goto L5
	} else {
		goto L372
	}
L357:
	;
	v1687 = F_slice_from_s(m, l0, int32(3), int32(2190611))
	mBase = m.M
	v1688 = m.ExcPending
	if v1688 != 0 {
		goto L5
	} else {
		goto L370
	}
L358:
	;
	v1681 = F_slice_from_s(m, l0, int32(4), int32(2190607))
	mBase = m.M
	v1682 = m.ExcPending
	if v1682 != 0 {
		goto L5
	} else {
		goto L368
	}
L359:
	;
	v1675 = F_slice_from_s(m, l0, int32(4), int32(2190603))
	mBase = m.M
	v1676 = m.ExcPending
	if v1676 != 0 {
		goto L5
	} else {
		goto L366
	}
L360:
	;
	v1669 = F_slice_from_s(m, l0, int32(4), int32(2190599))
	mBase = m.M
	v1670 = m.ExcPending
	if v1670 != 0 {
		goto L5
	} else {
		goto L364
	}
L361:
	;
	v1663 = F_slice_from_s(m, l0, int32(4), int32(2190595))
	mBase = m.M
	v1664 = m.ExcPending
	if v1664 != 0 {
		goto L5
	} else {
		goto L362
	}
L362:
	;
	if int32(0) <= v1663 {
		goto L342
	} else {
		goto L363
	}
L363:
	;
	v2407 = v1663
	goto L1
L364:
	;
	if int32(0) <= v1669 {
		goto L342
	} else {
		goto L365
	}
L365:
	;
	v2407 = v1669
	goto L1
L366:
	;
	if int32(0) <= v1675 {
		goto L342
	} else {
		goto L367
	}
L367:
	;
	v2407 = v1675
	goto L1
L368:
	;
	if int32(0) <= v1681 {
		goto L342
	} else {
		goto L369
	}
L369:
	;
	v2407 = v1681
	goto L1
L370:
	;
	if int32(0) <= v1687 {
		goto L342
	} else {
		goto L371
	}
L371:
	;
	v2407 = v1687
	goto L1
L372:
	;
	if int32(0) <= v1693 {
		goto L342
	} else {
		goto L373
	}
L373:
	;
	v2407 = v1693
	goto L1
L374:
	;
	if int32(0) <= v1699 {
		goto L342
	} else {
		goto L375
	}
L375:
	;
	v2407 = v1699
	goto L1
L376:
	;
	if int32(0) <= v1705 {
		goto L342
	} else {
		goto L377
	}
L377:
	;
	v2407 = v1705
	goto L1
L378:
	;
	if int32(0) <= v1711 {
		goto L342
	} else {
		goto L379
	}
L379:
	;
	v2407 = v1711
	goto L1
L380:
	;
	if int32(0) <= v1717 {
		goto L342
	} else {
		goto L381
	}
L381:
	;
	v2407 = v1717
	goto L1
L382:
	;
	if int32(0) <= v1723 {
		goto L342
	} else {
		goto L383
	}
L383:
	;
	v2407 = v1723
	goto L1
L384:
	;
	if int32(0) <= v1729 {
		goto L342
	} else {
		goto L385
	}
L385:
	;
	v2407 = v1729
	goto L1
L386:
	;
	if v1735 < int32(0) {
		v2407 = v1735
		goto L1
	} else {
		goto L387
	}
L387:
	;
	goto L342
L388:
	;
	if v1797 < int32(0) {
		v2407 = v1797
		goto L1
	} else {
		goto L405
	}
L389:
	;
	v1750 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1750+v1741-int32(1)))))
	if v1754&int32(224) != int32(96) {
		v1797 = v1743
		goto L388
	} else {
		goto L390
	}
L390:
	;
	if int32(1)<<(uint(v1754)%32)&int32(528928) == int32(0) {
		v1797 = v1743
		goto L388
	} else {
		goto L391
	}
L391:
	;
	v1767 = F_find_among_b(m, l0, int32(4296112), int32(7))
	mBase = m.M
	v1768 = m.ExcPending
	if v1768 != 0 {
		goto L5
	} else {
		goto L392
	}
L392:
	;
	if v1767 == int32(0) {
		v1797 = v1743
		goto L388
	} else {
		goto L393
	}
L393:
	;
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1771
	v1773 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1774 = *(*int32)(unsafe.Add(mBase, uint32(v1773)+4))
	if v1771 < v1774 {
		v1797 = v1743
		goto L388
	} else {
		goto L394
	}
L394:
	;
	switch v1767 - int32(1) {
	case 0:
		goto L398
	case 1:
		goto L397
	case 2:
		goto L396
	default:
		goto L395
	}
L395:
	;
	v1797 = int32(1)
	goto L388
L396:
	;
	v1790 = F_slice_del(m, l0)
	mBase = m.M
	v1791 = m.ExcPending
	if v1791 != 0 {
		goto L5
	} else {
		goto L403
	}
L397:
	;
	v1786 = F_slice_from_s(m, l0, int32(2), int32(2190741))
	mBase = m.M
	v1787 = m.ExcPending
	if v1787 != 0 {
		goto L5
	} else {
		goto L401
	}
L398:
	;
	v1780 = F_slice_from_s(m, l0, int32(2), int32(2190739))
	mBase = m.M
	v1781 = m.ExcPending
	if v1781 != 0 {
		goto L5
	} else {
		goto L399
	}
L399:
	;
	if int32(0) <= v1780 {
		goto L395
	} else {
		goto L400
	}
L400:
	;
	v1797 = v1780
	goto L388
L401:
	;
	if int32(0) <= v1786 {
		goto L395
	} else {
		goto L402
	}
L402:
	;
	v1797 = v1786
	goto L388
L403:
	;
	if v1790 < int32(0) {
		v1797 = v1790
		goto L388
	} else {
		goto L404
	}
L404:
	;
	goto L395
L405:
	;
	v1801 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1801
	v1803 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1801
	v1807 = v1801 - int32(1)
	v1808 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1807 <= v1808 {
		v1863 = v1803
		goto L406
	} else {
		goto L407
	}
L406:
	;
	if v1863 < int32(0) {
		v2407 = v1863
		goto L1
	} else {
		goto L422
	}
L407:
	;
	v1810 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1812 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1810+v1807))))
	if v1812&int32(224) != int32(96) {
		v1863 = v1803
		goto L406
	} else {
		goto L408
	}
L408:
	;
	if int32(1)<<(uint(v1812)%32)&int32(3961384) == int32(0) {
		v1863 = v1803
		goto L406
	} else {
		goto L409
	}
L409:
	;
	v1825 = F_find_among_b(m, l0, int32(4296256), int32(19))
	mBase = m.M
	v1826 = m.ExcPending
	if v1826 != 0 {
		goto L5
	} else {
		goto L410
	}
L410:
	;
	if v1825 == int32(0) {
		v1863 = v1803
		goto L406
	} else {
		goto L411
	}
L411:
	;
	v1829 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1829
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(v1831)))
	if v1829 < v1832 {
		v1863 = v1803
		goto L406
	} else {
		goto L412
	}
L412:
	;
	switch v1825 - int32(1) {
	case 0:
		goto L415
	case 1:
		goto L414
	default:
		goto L413
	}
L413:
	;
	v1863 = int32(1)
	goto L406
L414:
	;
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1829 <= v1840 {
		v1863 = v1803
		goto L406
	} else {
		goto L418
	}
L415:
	;
	v1836 = F_slice_del(m, l0)
	mBase = m.M
	v1837 = m.ExcPending
	if v1837 != 0 {
		goto L5
	} else {
		goto L416
	}
L416:
	;
	if int32(0) <= v1836 {
		goto L413
	} else {
		goto L417
	}
L417:
	;
	v1863 = v1836
	goto L406
L418:
	;
	v1842 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1844 = int32(1)
	v1846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1842+v1829-v1844))))
	if base.Ui32(v1844) < base.Ui32((v1846-int32(115))&int32(255)) {
		v1863 = v1803
		goto L406
	} else {
		goto L419
	}
L419:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1829 - int32(1)
	v1856 = F_slice_del(m, l0)
	mBase = m.M
	v1857 = m.ExcPending
	if v1857 != 0 {
		goto L5
	} else {
		goto L420
	}
L420:
	;
	if v1856 < int32(0) {
		v1863 = v1856
		goto L406
	} else {
		goto L421
	}
L421:
	;
	goto L413
L422:
	;
	v1867 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1867
	v1869 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1867
	v1872 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1867 <= v1872 {
		v2297 = v1869
		goto L423
	} else {
		goto L424
	}
L423:
	;
	if v2297 < int32(0) {
		v2407 = v2297
		goto L1
	} else {
		goto L500
	}
L424:
	;
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1874+v1867-int32(1)))))
	if v1878 != int32(101) {
		v2297 = v1869
		goto L423
	} else {
		goto L425
	}
L425:
	;
	v1882 = v1867 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1882
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1882
	v1885 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1886 = *(*int32)(unsafe.Add(mBase, uint32(v1885)))
	if v1867 <= v1886 {
		goto L426
	} else {
		goto L427
	}
L426:
	;
	v1888 = *(*int32)(unsafe.Add(mBase, uint32(v1885)+4))
	if v1867 <= v1888 {
		v2297 = v1869
		goto L423
	} else {
		goto L429
	}
L427:
	;
	goto L428
L428:
	;
	v2288 = F_slice_del(m, l0)
	mBase = m.M
	v2289 = m.ExcPending
	if v2289 != 0 {
		goto L5
	} else {
		goto L496
	}
L429:
	;
	v1890 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1904 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1907 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1908 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L433
L430:
	;
	v2282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2282 + (v1882 - v1890)
	goto L428
L431:
	;
	if v2020 != 0 {
		goto L430
	} else {
		goto L450
	}
L432:
	;
	v2020 = v2013
	goto L431
L433:
	;
	if v1907 <= v1908 {
		v2013 = int32(-1)
		goto L432
	} else {
		goto L435
	}
L434:
	;
	v2013 = int32(0)
	goto L432
L435:
	;
	v1925 = int32(1)
	v1926 = v1907 - v1925
	v1928 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1904+v1926))))
	v1930 = v1928 & int32(255)
	if v1926 == v1908 {
		v1985 = v1930
		v1986 = v1925
		goto L436
	} else {
		goto L437
	}
L436:
	;
	if int32(121) < v1985 {
		goto L445
	} else {
		goto L446
	}
L437:
	;
	if int32(0) <= v1928 {
		v1985 = v1930
		v1986 = v1925
		goto L436
	} else {
		goto L438
	}
L438:
	;
	v1936 = v1930 & int32(63)
	v1938 = v1907 - int32(2)
	v1940 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1904+v1938))))
	v1942 = v1940 << (uint(int32(6)) % 32)
	if base.B2i32(v1938 != v1908)&base.B2i32(base.Ui32(v1940) < base.Ui32(int32(192))) == int32(0) {
		goto L439
	} else {
		goto L440
	}
L439:
	;
	v1985 = v1942&int32(1984) | v1936
	v1986 = int32(2)
	goto L436
L440:
	;
	goto L441
L441:
	;
	v1955 = v1942&int32(4032) | v1936
	v1957 = v1907 - int32(3)
	v1959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1904+v1957))))
	if base.B2i32(v1957 != v1908)&base.B2i32(base.Ui32(v1959) < base.Ui32(int32(224))) == int32(0) {
		goto L442
	} else {
		goto L443
	}
L442:
	;
	v1985 = v1959<<(uint(int32(12))%32)&int32(61440) | v1955
	v1986 = int32(3)
	goto L436
L443:
	;
	goto L444
L444:
	;
	v1977 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1907+(v1904-int32(4))))))
	v1985 = v1959<<(uint(int32(12))%32)&int32(258048) | v1977&int32(7)<<(uint(int32(18))%32) | v1955
	v1986 = int32(4)
	goto L436
L445:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1907 - v1986
	goto L449
L446:
	;
	v1990 = v1985 - int32(89)
	if v1990 < int32(0) {
		goto L445
	} else {
		goto L447
	}
L447:
	;
	v1996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1990)>>(uint(int32(3))%32)))+uint32(_consts[1306]))))
	if int32(base.Ui32(v1996)>>(uint(v1990&int32(7))%32))&int32(1) == int32(0) {
		goto L445
	} else {
		goto L448
	}
L448:
	;
	v2020 = v1986
	goto L431
L449:
	;
	goto L434
L450:
	;
	v2034 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2037 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2038 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L453
L451:
	;
	if v2149 != 0 {
		goto L430
	} else {
		goto L475
	}
L452:
	;
	v2149 = v2142
	goto L451
L453:
	;
	if v2037 <= v2038 {
		v2142 = int32(-1)
		goto L452
	} else {
		goto L455
	}
L454:
	;
	v2142 = int32(0)
	goto L452
L455:
	;
	v2055 = int32(1)
	v2056 = v2037 - v2055
	v2058 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2034+v2056))))
	v2060 = v2058 & int32(255)
	if v2056 == v2038 {
		v2115 = v2060
		v2116 = v2055
		goto L456
	} else {
		goto L457
	}
L456:
	;
	if int32(121) < v2115 {
		goto L465
	} else {
		goto L466
	}
L457:
	;
	if int32(0) <= v2058 {
		v2115 = v2060
		v2116 = v2055
		goto L456
	} else {
		goto L458
	}
L458:
	;
	v2066 = v2060 & int32(63)
	v2068 = v2037 - int32(2)
	v2070 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2034+v2068))))
	v2072 = v2070 << (uint(int32(6)) % 32)
	if base.B2i32(v2068 != v2038)&base.B2i32(base.Ui32(v2070) < base.Ui32(int32(192))) == int32(0) {
		goto L459
	} else {
		goto L460
	}
L459:
	;
	v2115 = v2072&int32(1984) | v2066
	v2116 = int32(2)
	goto L456
L460:
	;
	goto L461
L461:
	;
	v2085 = v2072&int32(4032) | v2066
	v2087 = v2037 - int32(3)
	v2089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2034+v2087))))
	if base.B2i32(v2087 != v2038)&base.B2i32(base.Ui32(v2089) < base.Ui32(int32(224))) == int32(0) {
		goto L462
	} else {
		goto L463
	}
L462:
	;
	v2115 = v2089<<(uint(int32(12))%32)&int32(61440) | v2085
	v2116 = int32(3)
	goto L456
L463:
	;
	goto L464
L464:
	;
	v2107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2037+(v2034-int32(4))))))
	v2115 = v2089<<(uint(int32(12))%32)&int32(258048) | v2107&int32(7)<<(uint(int32(18))%32) | v2085
	v2116 = int32(4)
	goto L456
L465:
	;
	v2149 = v2116
	goto L451
L466:
	;
	goto L467
L467:
	;
	v2120 = v2115 - int32(97)
	if v2120 < int32(0) {
		goto L468
	} else {
		goto L469
	}
L468:
	;
	v2149 = v2116
	goto L451
L469:
	;
	goto L470
L470:
	;
	v2126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2120)>>(uint(int32(3))%32)))+uint32(_consts[1305]))))
	if int32(base.Ui32(v2126)>>(uint(v2120&int32(7))%32))&int32(1) == int32(0) {
		goto L471
	} else {
		goto L472
	}
L471:
	;
	v2149 = v2116
	goto L451
L472:
	;
	goto L473
L473:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2037 - v2116
	goto L474
L474:
	;
	goto L454
L475:
	;
	v2163 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L478
L476:
	;
	if v2279 == int32(0) {
		v2297 = v1869
		goto L423
	} else {
		goto L495
	}
L477:
	;
	v2279 = v2272
	goto L476
L478:
	;
	if v2166 <= v2167 {
		v2272 = int32(-1)
		goto L477
	} else {
		goto L480
	}
L479:
	;
	v2272 = int32(0)
	goto L477
L480:
	;
	v2184 = int32(1)
	v2185 = v2166 - v2184
	v2187 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2163+v2185))))
	v2189 = v2187 & int32(255)
	if v2185 == v2167 {
		v2244 = v2189
		v2245 = v2184
		goto L481
	} else {
		goto L482
	}
L481:
	;
	if int32(121) < v2244 {
		goto L490
	} else {
		goto L491
	}
L482:
	;
	if int32(0) <= v2187 {
		v2244 = v2189
		v2245 = v2184
		goto L481
	} else {
		goto L483
	}
L483:
	;
	v2195 = v2189 & int32(63)
	v2197 = v2166 - int32(2)
	v2199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2163+v2197))))
	v2201 = v2199 << (uint(int32(6)) % 32)
	if base.B2i32(v2197 != v2167)&base.B2i32(base.Ui32(v2199) < base.Ui32(int32(192))) == int32(0) {
		goto L484
	} else {
		goto L485
	}
L484:
	;
	v2244 = v2201&int32(1984) | v2195
	v2245 = int32(2)
	goto L481
L485:
	;
	goto L486
L486:
	;
	v2214 = v2201&int32(4032) | v2195
	v2216 = v2166 - int32(3)
	v2218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2163+v2216))))
	if base.B2i32(v2216 != v2167)&base.B2i32(base.Ui32(v2218) < base.Ui32(int32(224))) == int32(0) {
		goto L487
	} else {
		goto L488
	}
L487:
	;
	v2244 = v2218<<(uint(int32(12))%32)&int32(61440) | v2214
	v2245 = int32(3)
	goto L481
L488:
	;
	goto L489
L489:
	;
	v2236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2166+(v2163-int32(4))))))
	v2244 = v2218<<(uint(int32(12))%32)&int32(258048) | v2236&int32(7)<<(uint(int32(18))%32) | v2214
	v2245 = int32(4)
	goto L481
L490:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2166 - v2245
	goto L494
L491:
	;
	v2249 = v2244 - int32(97)
	if v2249 < int32(0) {
		goto L490
	} else {
		goto L492
	}
L492:
	;
	v2255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2249)>>(uint(int32(3))%32)))+uint32(_consts[1305]))))
	if int32(base.Ui32(v2255)>>(uint(v2249&int32(7))%32))&int32(1) == int32(0) {
		goto L490
	} else {
		goto L493
	}
L493:
	;
	v2279 = v2245
	goto L476
L494:
	;
	goto L479
L495:
	;
	goto L430
L496:
	;
	if int32(0) <= v2288 {
		goto L497
	} else {
		goto L498
	}
L497:
	;
	v2292 = int32(1)
	goto L499
L498:
	;
	v2292 = v2288
	goto L499
L499:
	;
	v2297 = v2292
	goto L423
L500:
	;
	v2300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2300
	v2302 = F_r_Step_5b(m, l0)
	mBase = m.M
	v2303 = m.ExcPending
	if v2303 != 0 {
		goto L5
	} else {
		goto L501
	}
L501:
	;
	if v2302 < int32(0) {
		v2407 = v2302
		goto L1
	} else {
		goto L502
	}
L502:
	;
	v2306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2306
	v2308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2309 = *(*int32)(unsafe.Add(mBase, uint32(v2308)+8))
	if v2309 == int32(0) {
		goto L503
	} else {
		goto L504
	}
L503:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2306
	v2407 = int32(1)
	goto L1
L504:
	;
	goto L505
L505:
	;
	v2318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2320 = v2318
	goto L507
L506:
	;
	v2407 = v2394
	goto L1
L507:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2320
	v2326 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2327 != v2320 {
		goto L510
	} else {
		goto L511
	}
L508:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2320
	v2389 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2320 + v2389
	v2394 = F_slice_from_s(m, l0, v2389, int32(2190539))
	mBase = m.M
	v2395 = m.ExcPending
	if v2395 != 0 {
		goto L5
	} else {
		goto L535
	}
L509:
	;
	goto L508
L510:
	;
	v2330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2320+v2326))))
	if v2330 == int32(89) {
		goto L509
	} else {
		goto L513
	}
L511:
	;
	goto L512
L512:
	;
	goto L516
L513:
	;
	goto L512
L514:
	;
	if v2384 < int32(0) {
		goto L503
	} else {
		goto L534
	}
L516:
	;
	goto L517
L517:
	;
	goto L518
L518:
	;
	v2339 = v2320
	v2341 = int32(1)
	goto L521
L520:
	;
	v2384 = v2369
	goto L514
L521:
	;
	if v2327 <= v2339 {
		goto L523
	} else {
		goto L524
	}
L522:
	;
	goto L520
L523:
	;
	v2384 = int32(-1)
	goto L514
L524:
	;
	goto L525
L525:
	;
	v2346 = v2339 + int32(1)
	v2348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2326+v2339))))
	if base.Ui32(v2348) < base.Ui32(int32(192)) {
		v2369 = v2346
		goto L526
	} else {
		goto L527
	}
L526:
	;
	v2370 = int32(1)
	if v2370 < v2341 {
		v2339 = v2369
		v2341 = v2341 - v2370
		goto L521
	} else {
		goto L533
	}
L527:
	;
	if v2327 <= v2346 {
		v2369 = v2346
		goto L526
	} else {
		goto L528
	}
L528:
	;
	v2355 = v2346
	goto L529
L529:
	;
	v2358 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2326+v2355))))
	if int32(-65) < v2358 {
		v2369 = v2355
		goto L526
	} else {
		goto L531
	}
L530:
	;
	v2369 = v2327
	goto L526
L531:
	;
	v2362 = v2355 + int32(1)
	if v2362 != v2327 {
		v2355 = v2362
		goto L529
	} else {
		goto L532
	}
L532:
	;
	goto L530
L533:
	;
	goto L522
L534:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2384
	v2320 = v2384
	goto L507
L535:
	;
	if int32(0) <= v2394 {
		goto L505
	} else {
		goto L536
	}
L536:
	;
	goto L506
}
func F_predicatelock_hash(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v4 = *(*int32)(unsafe.Add(mBase, _consts[815]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = F_get_hash_value(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		return v6 ^ v10<<(uint(int32(4))%32)
	}
}
func F_prefix_init(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
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
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v14 = v11 - int32(2)
	if base.Ui32(int32(8)) < base.Ui32(v14) {
		v29 = v4
	} else {
		if int32(base.Ui32(int32(487))>>(uint(v14)%32))&int32(1) == int32(0) {
			v29 = v4
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v14<<(uint(int32(2))%32))+uint32(_consts[1432])))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
			v29 = v28
		}
	}
	if int32(32) < v29 {
		v74 = int32(-12)
		m.G0 = v8 + int32(48)
		return v74
	} else {
		v33 = v29 + int32(2)
		v36 = F_pullf_read_max(m, l2, v33, v8+int32(44), v8)
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int32(0)
		} else {
			if v36 < int32(0) {
				v74 = v36
				m.G0 = v8 + int32(48)
				return v74
			} else {
				if v36 != v33 {
					F_px_debug(m, int32(461504), int32(0))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						v69 = int32(-100)
						v73 = F___memset(m, v8, int32(0), int32(34))
						mBase = m.M
						v74 = v69
						m.G0 = v8 + int32(48)
						return v74
					}
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
					v49 = v48 + v29
					v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49-int32(2)))))
					v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
					if v52 == v53 {
						v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49-int32(1)))))
						v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
						if v58 == v59 {
							v69 = int32(0)
							v73 = F___memset(m, v8, int32(0), int32(34))
							mBase = m.M
							v74 = v69
							m.G0 = v8 + int32(48)
							return v74
						} else {
							v62 = int32(0)
							F_px_debug(m, int32(27341), v62)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = int32(1)
								v69 = v62
								v73 = F___memset(m, v8, int32(0), int32(34))
								mBase = m.M
								v74 = v69
								m.G0 = v8 + int32(48)
								return v74
							}
						}
					} else {
						v62 = int32(0)
						F_px_debug(m, int32(27341), v62)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = int32(1)
							v69 = v62
							v73 = F___memset(m, v8, int32(0), int32(34))
							mBase = m.M
							v74 = v69
							m.G0 = v8 + int32(48)
							return v74
						}
					}
				}
			}
		}
	}
}
func F_prepare_probe_slot(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
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
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	m.T0[v16].(func(*base.Module, int32))(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l1 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v21 = int32(4487040)
	v22 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v25
	if v13 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v129 = F_ExecStoreMinimalTuple(m, v127, v12, int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L19
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v22
	v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
	v122 = v120 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)) = uint16(v122)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+6)) = uint16(v125)
	goto L18
L7:
	;
	v29 = int32(1)
	v31 = int32(0)
	if v13 != v29 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v39 = v31
	v40 = int32(0)
	goto L11
L9:
	;
	v83 = v31
	goto L10
L10:
	;
	if v13&v29 == int32(0) {
		goto L6
	} else {
		goto L16
	}
L11:
	;
	v49 = v39 << (uint(int32(2)) % 32)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49+v50)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
	v56 = m.T0[v55].(func(*base.Module, int32, int32, int32) int32)(m, v52, v24, v53+v39)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	v83 = v77
	goto L10
L13:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v58+v49))) = v56
	v62 = v39 | int32(1)
	v64 = v62 << (uint(int32(2)) % 32)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v64+v65)))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	v71 = m.T0[v70].(func(*base.Module, int32, int32, int32) int32)(m, v67, v24, v68+v62)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v73+v64))) = v71
	v76 = int32(2)
	v77 = v39 + v76
	v79 = v40 + v76
	if v79 != v13&int32(-2) {
		v39 = v77
		v40 = v79
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v96 = v83 << (uint(int32(2)) % 32)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94+v96)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v98)+20))
	v102 = m.T0[v101].(func(*base.Module, int32, int32, int32) int32)(m, v98, v24, v99+v83)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v104+v96))) = v102
	goto L6
L18:
	;
	return
L19:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v133 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+6)))
	if v133 < v132 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_slot_getsomeattrs_int(m, v12, v132)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v140 = v13 << (uint(int32(2)) % 32)
	if v140 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L22
L24:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	if v13 != 0 {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v141 = F__emscripten_memcpy_bulkmem(m, v137, v138, v140)
	mBase = m.M
	goto L27
L26:
	;
	goto L27
L27:
	;
	goto L24
L28:
	;
	v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
	v149 = v147 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)) = uint16(v149)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+6)) = uint16(v152)
	goto L32
L29:
	;
	v145 = F__emscripten_memcpy_bulkmem(m, v143, v144, v13)
	mBase = m.M
	goto L31
L30:
	;
	goto L31
L31:
	;
	goto L28
L32:
	;
	return
}
func F_preprocess_groupclause(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	v3 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if l1 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v195
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v10 <= int32(0) {
		v195 = v3
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v9)+124))
	if v38 != 0 {
		goto L13
	} else {
		goto L14
	}
L5:
	;
	v14 = int32(0)
	v16 = v3
	goto L6
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22+v14<<(uint(int32(2))%32))))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v9)+100))
	v28 = F_get_sortgroupref_clause(m, v26, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v195 = v32
	goto L1
L8:
	;
	return int32(0)
L9:
	;
	v32 = F_lappend(m, v16, v28)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v35 = v14 + int32(1)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v35 < v36 {
		v14 = v35
		v16 = v32
		goto L6
	} else {
		goto L11
	}
L11:
	;
	goto L7
L12:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	v191 = F_list_copy(m, v190)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L8
	} else {
		goto L58
	}
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if int32(0) < v39 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	goto L15
L15:
	;
	v185 = v9 + int32(100)
	goto L12
L16:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v9)+100))
	if v107 == int32(0) {
		v195 = v91
		goto L1
	} else {
		goto L35
	}
L17:
	;
	v44 = v3
	v48 = v3
	goto L20
L18:
	;
	goto L19
L19:
	;
	v185 = v9 + int32(100)
	goto L12
L20:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v9)+100))
	if v50 == int32(0) {
		v91 = v44
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v91 != 0 {
		goto L16
	} else {
		goto L34
	}
L22:
	;
	goto L21
L23:
	;
	v53 = int32(0)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v54 <= v53 {
		v91 = v44
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v48<<(uint(int32(2))%32))))
	v62 = v53
	goto L25
L25:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70+v62<<(uint(int32(2))%32))))
	v75 = F_equal(m, v74, v61)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L8
	} else {
		goto L27
	}
L26:
	;
	v83 = F_lappend(m, v44, v74)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L8
	} else {
		goto L32
	}
L27:
	;
	if v75 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v80 = v62 + int32(1)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v80 < v81 {
		v62 = v80
		goto L25
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	goto L26
L31:
	;
	v91 = v44
	goto L22
L32:
	;
	v86 = v48 + int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v86 < v87 {
		v44 = v83
		v48 = v86
		goto L20
	} else {
		goto L33
	}
L33:
	;
	v91 = v83
	goto L22
L34:
	;
	goto L19
L35:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v110 <= int32(0) {
		v195 = v91
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v116 = int32(0)
	v118 = v91
	goto L37
L37:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v124+v116<<(uint(int32(2))%32))))
	v129 = int32(0)
	if v118 == v129 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v195 = v175
	goto L1
L39:
	;
	if v167 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L40:
	;
	v167 = int32(0)
	goto L39
L41:
	;
	goto L42
L42:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	if v135 <= int32(0) {
		v160 = v129
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v167 = v160
	goto L39
L44:
	;
	v138 = int32(0)
	if v138 < v135 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v141 = v135
	goto L47
L46:
	;
	v141 = v138
	goto L47
L47:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
	v144 = int32(0)
	goto L48
L48:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v142+v144<<(uint(int32(2))%32))))
	v153 = base.B2i32(v152 == v128)
	if v152 == v128 {
		v160 = v153
		goto L43
	} else {
		goto L50
	}
L49:
	;
	v160 = v153
	goto L43
L50:
	;
	v155 = v144 + int32(1)
	if v155 != v141 {
		v144 = v155
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	if v170 == int32(0) {
		v185 = v9 + int32(100)
		goto L12
	} else {
		goto L55
	}
L53:
	;
	v175 = v118
	goto L54
L54:
	;
	v177 = v116 + int32(1)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v177 < v178 {
		v116 = v177
		v118 = v175
		goto L37
	} else {
		goto L57
	}
L55:
	;
	v173 = F_lappend(m, v118, v128)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L8
	} else {
		goto L56
	}
L56:
	;
	v175 = v173
	goto L54
L57:
	;
	goto L38
L58:
	;
	v195 = v191
	goto L1
}
func F_printSubscripts(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v12 = v10
	goto L3
L2:
	;
	v12 = int32(0)
	goto L3
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v13 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v16 <= int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v22 = v12
	v25 = int32(0)
	goto L7
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_appendStringInfoChar(m, v19, int32(91))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L4
L9:
	;
	return
L10:
	;
	if v22 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	F_get_rule_expr(m, v36, l1, int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	v55 = int32(0)
	goto L13
L13:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v28+v25<<(uint(int32(2))%32))))
	F_get_rule_expr(m, v56, l1, int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L9
	} else {
		goto L19
	}
L14:
	;
	F_appendStringInfoChar(m, v19, int32(58))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v44 = v22 + int32(4)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if base.Ui32(v44) < base.Ui32(v47+v48<<(uint(int32(2))%32)) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v53 = v44
	goto L18
L17:
	;
	v53 = int32(0)
	goto L18
L18:
	;
	v55 = v53
	goto L13
L19:
	;
	F_appendStringInfoChar(m, v19, int32(93))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	v64 = v25 + int32(1)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v64 < v65 {
		v22 = v55
		v25 = v64
		goto L7
	} else {
		goto L21
	}
L21:
	;
	goto L8
}
func F_printf_core(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v125 int32
	_ = v125
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v518 int32
	_ = v518
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v551 int32
	_ = v551
	var v574 int64
	_ = v574
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v614 int64
	_ = v614
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int64
	_ = v645
	var v649 int32
	_ = v649
	var v674 int64
	_ = v674
	var v676 int32
	_ = v676
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v690 int32
	_ = v690
	var v727 int64
	_ = v727
	var v731 int32
	_ = v731
	var v756 int64
	_ = v756
	var v758 int32
	_ = v758
	var v763 int32
	_ = v763
	var v770 int32
	_ = v770
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v805 int64
	_ = v805
	var v809 int64
	_ = v809
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int64
	_ = v823
	var v824 int32
	_ = v824
	var v831 int64
	_ = v831
	var v832 int32
	_ = v832
	var v838 int32
	_ = v838
	var v839 int64
	_ = v839
	var v840 int64
	_ = v840
	var v846 int32
	_ = v846
	var v851 int32
	_ = v851
	var v855 int64
	_ = v855
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v873 int32
	_ = v873
	var v878 int32
	_ = v878
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v906 int32
	_ = v906
	var v908 int64
	_ = v908
	var v915 int32
	_ = v915
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v942 int64
	_ = v942
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v956 int32
	_ = v956
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v972 int32
	_ = v972
	var v977 int32
	_ = v977
	var v992 int32
	_ = v992
	var v996 int32
	_ = v996
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1018 int32
	_ = v1018
	var v1023 int32
	_ = v1023
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1051 int32
	_ = v1051
	var v1055 int32
	_ = v1055
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1066 int32
	_ = v1066
	var v1077 int32
	_ = v1077
	var v1101 int32
	_ = v1101
	var v1103 int32
	_ = v1103
	var v1108 float64
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1126 int32
	_ = v1126
	var v1149 int32
	_ = v1149
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1170 int32
	_ = v1170
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1196 int32
	_ = v1196
	var v1227 int32
	_ = v1227
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1254 int32
	_ = v1254
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1265 int32
	_ = v1265
	var v1267 int32
	_ = v1267
	var v1271 int32
	_ = v1271
	var v1273 int32
	_ = v1273
	var v1278 int32
	_ = v1278
	var v1282 int32
	_ = v1282
	var v1284 int32
	_ = v1284
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1364 int32
	_ = v1364
	var v1418 int32
	_ = v1418
	v8 = int32(0)
	v28 = m.G0
	v30 = v28 + int32(-64)
	m.G0 = v30
	*(*int32)(unsafe.Add(mBase, uint32(v30)+60)) = l1
	v36 = v28 + int32(-24)
	v38 = l1
	v51 = v8
	v57 = v8
	goto L5
L1:
	;
	m.G0 = v30 - int32(-64)
	return v1418
L2:
	;
	v1418 = int32(-1)
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = v1364
	goto L2
L4:
	;
	v1364 = int32(61)
	goto L3
L5:
	;
	v66 = v38
	v72 = int32(0)
	v79 = v51
	v85 = v57
	goto L7
L6:
	;
	v1418 = int32(0)
	goto L1
L7:
	;
	if v79^int32(2147483647) < v72 {
		goto L4
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	v95 = v72 + v79
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v96 != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	goto L8
L11:
	;
	v1258 = v1248 - v1244
	if v1258 < v1241 {
		goto L288
	} else {
		goto L289
	}
L12:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+39)) = uint8(v1227)
	v1241 = int32(1)
	v1242 = v592
	v1244 = v28 + int32(-25)
	v1247 = v596
	v1248 = v36
	v1254 = v597
	goto L11
L13:
	;
	v1364 = int32(28)
	goto L3
L14:
	;
	v104 = v66
	v109 = v96
	goto L17
L15:
	;
	goto L16
L16:
	;
	if l0 != 0 {
		v1418 = v95
		goto L1
	} else {
		goto L272
	}
L17:
	;
	v125 = v109 & int32(255)
	if v125 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v1113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+1)))
	v104 = v104 + int32(1)
	v109 = v1113
	goto L17
L20:
	;
	v194 = v174 - v66
	v196 = v95 ^ int32(2147483647)
	if v196 < v194 {
		goto L4
	} else {
		goto L31
	}
L21:
	;
	v168 = v104
	v174 = v104
	goto L20
L22:
	;
	goto L23
L23:
	;
	if v125 != int32(37) {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v137 = v104
	v142 = v104
	goto L25
L25:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+1)))
	if v157 != int32(37) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v168 = v164
	v174 = v161
	goto L20
L27:
	;
	v168 = v142
	v174 = v137
	goto L20
L28:
	;
	goto L29
L29:
	;
	v161 = v137 + int32(1)
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+2)))
	v164 = v142 + int32(2)
	if v162 == int32(37) {
		v137 = v161
		v142 = v164
		goto L25
	} else {
		goto L30
	}
L30:
	;
	goto L26
L31:
	;
	if l0 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_out(m, l0, v66, v194)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	if v194 != 0 {
		v66 = v168
		v72 = v194
		v79 = v95
		goto L7
	} else {
		goto L37
	}
L35:
	;
	return int32(0)
L36:
	;
	goto L34
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+60)) = v168
	v204 = v168 + int32(1)
	v205 = int32(-1)
	v206 = int32(*(*int8)(unsafe.Add(mBase, uint32(v168)+1)))
	v208 = v206 - int32(48)
	if base.Ui32(int32(9)) < base.Ui32(v208) {
		v217 = v204
		v218 = v205
		v219 = v85
		goto L38
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+60)) = v217
	v221 = int32(0)
	v222 = int32(*(*int8)(unsafe.Add(mBase, uint32(v217))))
	v224 = v222 - int32(32)
	if base.Ui32(int32(31)) < base.Ui32(v224) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+2)))
	if v211 != int32(36) {
		v217 = v204
		v218 = v205
		v219 = v85
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v217 = v168 + int32(3)
	v218 = v208
	v219 = int32(1)
	goto L38
L41:
	;
	if v291 == int32(42) {
		goto L51
	} else {
		goto L52
	}
L42:
	;
	v282 = v217
	v284 = v221
	v291 = v222
	goto L41
L43:
	;
	goto L44
L44:
	;
	v228 = int32(1) << (uint(v224) % 32)
	if v228&int32(75913) == int32(0) {
		v282 = v217
		v284 = v221
		v291 = v222
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v234 = v228
	v240 = v217
	v244 = v221
	goto L46
L46:
	;
	v261 = v240 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+60)) = v261
	v263 = v234 | v244
	v264 = int32(*(*int8)(unsafe.Add(mBase, uint32(v240)+1)))
	v265 = int32(32)
	v266 = v264 - v265
	if base.Ui32(v265) <= base.Ui32(v266) {
		v282 = v261
		v284 = v263
		v291 = v264
		goto L41
	} else {
		goto L48
	}
L47:
	;
	v282 = v261
	v284 = v263
	v291 = v264
	goto L41
L48:
	;
	v270 = int32(1) << (uint(v266) % 32)
	if v270&int32(75913) != 0 {
		v234 = v270
		v240 = v261
		v244 = v263
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401))))
	if v408 != int32(46) {
		goto L82
	} else {
		goto L83
	}
L51:
	;
	v302 = int32(*(*int8)(unsafe.Add(mBase, uint32(v282)+1)))
	v304 = v302 - int32(48)
	if base.Ui32(int32(9)) < base.Ui32(v304) {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	goto L53
L53:
	;
	v351 = v28 + int32(-4)
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v351)))
	v358 = int32(*(*int8)(unsafe.Add(mBase, uint32(v357))))
	v360 = v358 - int32(48)
	if base.Ui32(int32(9)) < base.Ui32(v360) {
		goto L68
	} else {
		goto L69
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+60)) = v339
	if int32(0) <= v341 {
		v401 = v339
		v403 = v284
		v404 = v341
		v405 = v342
		goto L50
	} else {
		goto L66
	}
L55:
	;
	if v219 != 0 {
		goto L13
	} else {
		goto L62
	}
L56:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282)+2)))
	if v307 != int32(36) {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	if l0 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v339 = v282 + int32(3)
	v341 = v322
	v342 = int32(1)
	goto L54
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4+v304<<(uint(int32(2))%32)))) = int32(10)
	v322 = int32(0)
	goto L58
L60:
	;
	goto L61
L61:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l3+v304<<(uint(int32(3))%32))))
	v322 = v321
	goto L58
L62:
	;
	v327 = v282 + int32(1)
	if l0 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+60)) = v327
	v331 = int32(0)
	v401 = v327
	v403 = v284
	v404 = v331
	v405 = v331
	goto L50
L64:
	;
	goto L65
L65:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v333 + int32(4)
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v333)))
	v339 = v327
	v341 = v337
	v342 = int32(0)
	goto L54
L66:
	;
	v401 = v339
	v403 = v284 | int32(8192)
	v404 = int32(0) - v341
	v405 = v342
	goto L50
L67:
	;
	if v397 < int32(0) {
		goto L4
	} else {
		goto L80
	}
L68:
	;
	v397 = int32(0)
	goto L67
L69:
	;
	goto L70
L70:
	;
	v365 = int32(0)
	v366 = v360
	v367 = v357
	goto L71
L71:
	;
	if base.Ui32(v365) <= base.Ui32(int32(214748364)) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v397 = v382
	goto L67
L73:
	;
	v375 = v365 * int32(10)
	if base.Ui32(v375^int32(2147483647)) < base.Ui32(v366) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v382 = int32(-1)
	goto L75
L75:
	;
	v384 = v367 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v351))) = v384
	v386 = int32(*(*int8)(unsafe.Add(mBase, uint32(v367)+1)))
	v388 = v386 - int32(48)
	if base.Ui32(v388) < base.Ui32(int32(10)) {
		v365 = v382
		v366 = v388
		v367 = v384
		goto L71
	} else {
		goto L79
	}
L76:
	;
	v380 = int32(-1)
	goto L78
L77:
	;
	v380 = v366 + v375
	goto L78
L78:
	;
	v382 = v380
	goto L75
L79:
	;
	goto L72
L80:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v30)+60))
	v401 = v400
	v403 = v284
	v404 = v397
	v405 = v219
	goto L50
L81:
	;
	v512 = v507
	v518 = int32(0)
	goto L111
L82:
	;
	v507 = v401
	v509 = int32(-1)
	v510 = int32(0)
	goto L81
L83:
	;
	goto L84
L84:
	;
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401)+1)))
	if v412 == int32(42) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v415 = int32(*(*int8)(unsafe.Add(mBase, uint32(v401)+2)))
	v417 = v415 - int32(48)
	if base.Ui32(int32(9)) < base.Ui32(v417) {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	goto L87
L87:
	;
	v454 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+60)) = v401 + v454
	v459 = v28 + int32(-4)
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v459)))
	v466 = int32(*(*int8)(unsafe.Add(mBase, uint32(v465))))
	v468 = v466 - int32(48)
	if base.Ui32(int32(9)) < base.Ui32(v468) {
		goto L99
	} else {
		goto L100
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+60)) = v448
	v507 = v448
	v509 = v450
	v510 = base.B2i32(int32(0) <= v450)
	goto L81
L89:
	;
	if v405 != 0 {
		goto L13
	} else {
		goto L96
	}
L90:
	;
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401)+3)))
	if v420 != int32(36) {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	if l0 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v448 = v401 + int32(4)
	v450 = v437
	goto L88
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4+v417<<(uint(int32(2))%32)))) = int32(10)
	v437 = int32(0)
	goto L92
L94:
	;
	goto L95
L95:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l3+v417<<(uint(int32(3))%32))))
	v437 = v436
	goto L92
L96:
	;
	v439 = v401 + int32(2)
	v440 = int32(0)
	if l0 == v440 {
		v448 = v439
		v450 = v440
		goto L88
	} else {
		goto L97
	}
L97:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v443 + int32(4)
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v443)))
	v448 = v439
	v450 = v447
	goto L88
L98:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v30)+60))
	v507 = v506
	v509 = v505
	v510 = v454
	goto L81
L99:
	;
	v505 = int32(0)
	goto L98
L100:
	;
	goto L101
L101:
	;
	v473 = int32(0)
	v474 = v468
	v475 = v465
	goto L102
L102:
	;
	if base.Ui32(v473) <= base.Ui32(int32(214748364)) {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v505 = v490
	goto L98
L104:
	;
	v483 = v473 * int32(10)
	if base.Ui32(v483^int32(2147483647)) < base.Ui32(v474) {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	v490 = int32(-1)
	goto L106
L106:
	;
	v492 = v475 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v459))) = v492
	v494 = int32(*(*int8)(unsafe.Add(mBase, uint32(v475)+1)))
	v496 = v494 - int32(48)
	if base.Ui32(v496) < base.Ui32(int32(10)) {
		v473 = v490
		v474 = v496
		v475 = v492
		goto L102
	} else {
		goto L110
	}
L107:
	;
	v488 = int32(-1)
	goto L109
L108:
	;
	v488 = v474 + v483
	goto L109
L109:
	;
	v490 = v488
	goto L106
L110:
	;
	goto L103
L111:
	;
	v538 = int32(28)
	v539 = int32(*(*int8)(unsafe.Add(mBase, uint32(v512))))
	if base.Ui32(v539-int32(123)) < base.Ui32(int32(-58)) {
		v1364 = v538
		goto L3
	} else {
		goto L113
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+60)) = v545
	if v551 != int32(27) {
		goto L116
	} else {
		goto L117
	}
L113:
	;
	v544 = int32(1)
	v545 = v512 + v544
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v539+v518*int32(58))+uint32(_consts[1441]))))
	if base.Ui32((v551-v544)&int32(255)) < base.Ui32(int32(8)) {
		v512 = v545
		v518 = v551
		goto L111
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v588&int32(32) != 0 {
		goto L2
	} else {
		goto L130
	}
L116:
	;
	if v551 == int32(0) {
		v1364 = v538
		goto L3
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	if int32(0) <= v218 {
		v1364 = v538
		goto L3
	} else {
		goto L128
	}
L119:
	;
	if int32(0) <= v218 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	if l0 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	goto L122
L122:
	;
	if l0 == int32(0) {
		goto L10
	} else {
		goto L126
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4+v218<<(uint(int32(2))%32)))) = v551
	v38 = v545
	v51 = v95
	v57 = v405
	goto L5
L124:
	;
	goto L125
L125:
	;
	v574 = *(*int64)(unsafe.Add(mBase, uint32(l3+v218<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+48)) = v574
	goto L115
L126:
	;
	F_pop_arg(m, v28+int32(-16), v551, l2, l6)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L35
	} else {
		goto L127
	}
L127:
	;
	goto L115
L128:
	;
	v584 = int32(0)
	if l0 == v584 {
		v66 = v545
		v72 = v584
		v79 = v95
		v85 = v405
		goto L7
	} else {
		goto L129
	}
L129:
	;
	goto L115
L130:
	;
	v592 = v403 & int32(-65537)
	if v403&int32(8192) != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v595 = v592
	goto L133
L132:
	;
	v595 = v403
	goto L133
L133:
	;
	v596 = int32(0)
	v597 = int32(29649)
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512))))
	v599 = base.I32_extend8_s(v598)
	if v598&int32(15) == int32(3) {
		goto L151
	} else {
		goto L152
	}
L134:
	;
	if v510&base.B2i32(v509 < int32(0)) != 0 {
		goto L4
	} else {
		goto L269
	}
L135:
	;
	F_pad(m, l0, int32(32), v404, v1077, v595^int32(8192))
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L35
	} else {
		goto L265
	}
L136:
	;
	v972 = int32(0)
	v977 = v963
	goto L239
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+12)) = int32(0)
	*(*uint32)(unsafe.Add(mBase, uint32(v30)+8)) = uint32(v942)
	v956 = v28 + int32(-56)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+48)) = v956
	v961 = int32(-1)
	v963 = v956
	goto L136
L138:
	;
	if v509 != 0 {
		goto L235
	} else {
		goto L236
	}
L139:
	;
	v942 = *(*int64)(unsafe.Add(mBase, uint32(v30)+48))
	if v942 != int64(0) {
		goto L137
	} else {
		goto L234
	}
L140:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v30)+48))
	if v926 != 0 {
		goto L220
	} else {
		goto L221
	}
L141:
	;
	v925 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+48)))
	v1227 = v925
	goto L12
L142:
	;
	if v510&base.B2i32(v893 < int32(0)) != 0 {
		goto L4
	} else {
		goto L210
	}
L143:
	;
	if base.Ui64(v823) < base.Ui64(int64(4294967296)) {
		goto L198
	} else {
		goto L199
	}
L144:
	;
	v805 = *(*int64)(unsafe.Add(mBase, uint32(v30)+48))
	if v805 < int64(0) {
		goto L187
	} else {
		goto L188
	}
L145:
	;
	v727 = *(*int64)(unsafe.Add(mBase, uint32(v30)+48))
	if v727 != int64(0) {
		goto L177
	} else {
		goto L178
	}
L146:
	;
	v645 = *(*int64)(unsafe.Add(mBase, uint32(v30)+48))
	if v645 != int64(0) {
		goto L169
	} else {
		goto L170
	}
L147:
	;
	v633 = int32(8)
	if base.Ui32(v509) <= base.Ui32(v633) {
		goto L166
	} else {
		goto L167
	}
L148:
	;
	v616 = int32(0)
	switch v518 {
	case 0:
		goto L165
	case 1:
		goto L164
	case 2:
		goto L163
	case 3:
		goto L162
	case 4:
		goto L161
	default:
		v66 = v545
		v72 = v616
		v79 = v95
		v85 = v405
		goto L7
	case 6:
		goto L160
	case 7:
		goto L159
	}
L149:
	;
	v614 = *(*int64)(unsafe.Add(mBase, uint32(v30)+48))
	v822 = v596
	v823 = v614
	v824 = int32(29649)
	goto L143
L150:
	;
	switch v607 - int32(65) {
	case 0, 4, 5, 6:
		goto L134
	case 1, 3:
		v1241 = v509
		v1242 = v595
		v1244 = v66
		v1247 = v596
		v1248 = v36
		v1254 = v597
		goto L11
	case 2:
		goto L139
	default:
		goto L157
	}
L151:
	;
	v606 = v599 & int32(-45)
	goto L153
L152:
	;
	v606 = v599
	goto L153
L153:
	;
	if v518 != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v607 = v606
	goto L156
L155:
	;
	v607 = v599
	goto L156
L156:
	;
	switch v607 - int32(88) {
	case 0, 32:
		v640 = v607
		v641 = v509
		v642 = v595
		goto L146
	case 1, 2, 3, 4, 5, 6, 7, 8, 10, 16, 18, 19, 20, 21, 25, 26, 28, 30, 31:
		v1241 = v509
		v1242 = v595
		v1244 = v66
		v1247 = v596
		v1248 = v36
		v1254 = v597
		goto L11
	case 9, 13, 14, 15:
		goto L134
	case 11:
		goto L141
	case 12, 17:
		goto L144
	case 22:
		goto L148
	case 23:
		goto L145
	case 24:
		goto L147
	case 27:
		goto L140
	case 29:
		goto L149
	default:
		goto L150
	}
L157:
	;
	if v607 == int32(83) {
		goto L138
	} else {
		goto L158
	}
L158:
	;
	v1241 = v509
	v1242 = v595
	v1244 = v66
	v1247 = v596
	v1248 = v36
	v1254 = v597
	goto L11
L159:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v30)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v630))) = base.I64_extend_i32_s(v95)
	v66 = v545
	v72 = v616
	v79 = v95
	v85 = v405
	goto L7
L160:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v30)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v628))) = v95
	v66 = v545
	v72 = v616
	v79 = v95
	v85 = v405
	goto L7
L161:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v30)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v626))) = uint8(v95)
	v66 = v545
	v72 = v616
	v79 = v95
	v85 = v405
	goto L7
L162:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v30)+48))
	*(*uint16)(unsafe.Add(mBase, uint32(v624))) = uint16(v95)
	v66 = v545
	v72 = v616
	v79 = v95
	v85 = v405
	goto L7
L163:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v30)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v621))) = base.I64_extend_i32_s(v95)
	v66 = v545
	v72 = v616
	v79 = v95
	v85 = v405
	goto L7
L164:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v30)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v619))) = v95
	v66 = v545
	v72 = v616
	v79 = v95
	v85 = v405
	goto L7
L165:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v30)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v617))) = v95
	v66 = v545
	v72 = v616
	v79 = v95
	v85 = v405
	goto L7
L166:
	;
	v636 = v633
	goto L168
L167:
	;
	v636 = v509
	goto L168
L168:
	;
	v640 = int32(120)
	v641 = v636
	v642 = v595 | int32(8)
	goto L146
L169:
	;
	v649 = v36
	v674 = v645
	goto L172
L170:
	;
	v690 = v36
	goto L171
L171:
	;
	if v645 == int64(0) {
		v893 = v641
		v894 = v642
		v896 = v690
		v899 = v596
		v906 = v597
		v908 = v645
		goto L142
	} else {
		goto L175
	}
L172:
	;
	v676 = v649 - int32(1)
	v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v674)&int32(15))+uint32(_consts[1442]))))
	v683 = v682 | v640&int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v676))) = uint8(v683)
	if base.Ui64(int64(15)) < base.Ui64(v674) {
		v649 = v676
		v674 = int64(base.Ui64(v674) >> (uint(int64(4)) % 64))
		goto L172
	} else {
		goto L174
	}
L173:
	;
	v690 = v676
	goto L171
L174:
	;
	goto L173
L175:
	;
	if v642&int32(8) == int32(0) {
		v893 = v641
		v894 = v642
		v896 = v690
		v899 = v596
		v906 = v597
		v908 = v645
		goto L142
	} else {
		goto L176
	}
L176:
	;
	v893 = v641
	v894 = v642
	v896 = v690
	v899 = int32(2)
	v906 = int32(base.Ui32(v640)>>(uint(int32(4))%32)) + int32(29649)
	v908 = v645
	goto L142
L177:
	;
	v731 = v36
	v756 = v727
	goto L180
L178:
	;
	v770 = v36
	goto L179
L179:
	;
	if v595&int32(8) == int32(0) {
		v893 = v509
		v894 = v595
		v896 = v770
		v899 = v596
		v906 = v597
		v908 = v727
		goto L142
	} else {
		goto L183
	}
L180:
	;
	v758 = v731 - int32(1)
	v763 = base.I32_wrap_i64(v756)&int32(7) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v758))) = uint8(v763)
	if base.Ui64(int64(7)) < base.Ui64(v756) {
		v731 = v758
		v756 = int64(base.Ui64(v756) >> (uint(int64(3)) % 64))
		goto L180
	} else {
		goto L182
	}
L181:
	;
	v770 = v758
	goto L179
L182:
	;
	goto L181
L183:
	;
	v800 = v36 - v770
	if v800 < v509 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v804 = v509
	goto L186
L185:
	;
	v804 = v800 + int32(1)
	goto L186
L186:
	;
	v893 = v804
	v894 = v595
	v896 = v770
	v899 = v596
	v906 = v597
	v908 = v727
	goto L142
L187:
	;
	v809 = int64(0) - v805
	*(*int64)(unsafe.Add(mBase, uint32(v30)+48)) = v809
	v822 = int32(1)
	v823 = v809
	v824 = int32(29649)
	goto L143
L188:
	;
	goto L189
L189:
	;
	if v595&int32(2048) != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v822 = int32(1)
	v823 = v805
	v824 = int32(29650)
	goto L143
L191:
	;
	goto L192
L192:
	;
	v820 = v595 & int32(1)
	if v820 != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v821 = int32(29651)
	goto L195
L194:
	;
	v821 = int32(29649)
	goto L195
L195:
	;
	v822 = v820
	v823 = v805
	v824 = v821
	goto L143
L196:
	;
	v893 = v509
	v894 = v595
	v896 = v878
	v899 = v822
	v906 = v824
	v908 = v823
	goto L142
L197:
	;
	if v855 != int64(0) {
		goto L204
	} else {
		goto L205
	}
L198:
	;
	v851 = v36
	v855 = v823
	goto L197
L199:
	;
	goto L200
L200:
	;
	v831 = v823
	v832 = v36
	goto L201
L201:
	;
	v838 = v832 - int32(1)
	v839 = int64(10)
	v840 = base.I64_div_u_s(v831, v839)
	v846 = base.I32_wrap_i64(v831-v840*v839) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v838))) = uint8(v846)
	if base.Ui64(int64(42949672959)) < base.Ui64(v831) {
		v831 = v840
		v832 = v838
		goto L201
	} else {
		goto L203
	}
L202:
	;
	v851 = v838
	v855 = v840
	goto L197
L203:
	;
	goto L202
L204:
	;
	v860 = v851
	v861 = base.I32_wrap_i64(v855)
	goto L207
L205:
	;
	v878 = v851
	goto L206
L206:
	;
	goto L196
L207:
	;
	v866 = v860 - int32(1)
	v867 = int32(10)
	v868 = base.I32_div_u_s(v861, v867)
	v873 = v861 - v868*v867 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v866))) = uint8(v873)
	if base.Ui32(int32(9)) < base.Ui32(v861) {
		v860 = v866
		v861 = v868
		goto L207
	} else {
		goto L209
	}
L208:
	;
	v878 = v866
	goto L206
L209:
	;
	goto L208
L210:
	;
	if v510 != 0 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v915 = v894 & int32(-65537)
	goto L213
L212:
	;
	v915 = v894
	goto L213
L213:
	;
	if v908 != int64(0) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v922 = base.B2i32(v908 == int64(0)) + (v36 - v896)
	if v922 < v893 {
		goto L217
	} else {
		goto L218
	}
L215:
	;
	if v893 != 0 {
		goto L214
	} else {
		goto L216
	}
L216:
	;
	v1241 = int32(0)
	v1242 = v915
	v1244 = v36
	v1247 = v899
	v1248 = v36
	v1254 = v906
	goto L11
L217:
	;
	v924 = v893
	goto L219
L218:
	;
	v924 = v922
	goto L219
L219:
	;
	v1241 = v924
	v1242 = v915
	v1244 = v896
	v1247 = v899
	v1248 = v36
	v1254 = v906
	goto L11
L220:
	;
	v928 = v926
	goto L222
L221:
	;
	v928 = int32(657717)
	goto L222
L222:
	;
	v929 = int32(2147483647)
	if base.Ui32(v929) <= base.Ui32(v509) {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v932 = v929
	goto L225
L224:
	;
	v932 = v509
	goto L225
L225:
	;
	v935 = F_memchr(m, v928, int32(0), v932)
	mBase = m.M
	if v935 != 0 {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	v938 = v937 + v928
	if int32(0) <= v509 {
		goto L230
	} else {
		goto L231
	}
L227:
	;
	v937 = v935 - v928
	goto L229
L228:
	;
	v937 = v932
	goto L229
L229:
	;
	goto L226
L230:
	;
	v1241 = v937
	v1242 = v592
	v1244 = v928
	v1247 = v596
	v1248 = v938
	v1254 = v597
	goto L11
L231:
	;
	goto L232
L232:
	;
	v941 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v938))))
	if v941 != 0 {
		goto L4
	} else {
		goto L233
	}
L233:
	;
	v1241 = v937
	v1242 = v592
	v1244 = v928
	v1247 = v596
	v1248 = v938
	v1254 = v597
	goto L11
L234:
	;
	v1227 = int32(0)
	goto L12
L235:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v30)+48))
	v961 = v509
	v963 = v946
	goto L136
L236:
	;
	goto L237
L237:
	;
	v947 = int32(0)
	F_pad(m, l0, int32(32), v404, v947, v595)
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L35
	} else {
		goto L238
	}
L238:
	;
	v1077 = v947
	goto L135
L239:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v977)))
	if v992 == int32(0) {
		v1010 = v972
		goto L241
	} else {
		goto L242
	}
L240:
	;
	if v1010 < int32(0) {
		v1364 = int32(61)
		goto L3
	} else {
		goto L250
	}
L241:
	;
	goto L240
L242:
	;
	v996 = v28 + int32(-60)
	if v996 == int32(0) {
		goto L244
	} else {
		goto L245
	}
L243:
	;
	if v1001 < int32(0) {
		goto L2
	} else {
		goto L247
	}
L244:
	;
	v1001 = int32(0)
	goto L243
L245:
	;
	goto L246
L246:
	;
	v1000 = F_wcrtomb(m, v996, v992)
	mBase = m.M
	v1001 = v1000
	goto L243
L247:
	;
	if base.Ui32(v961-v972) < base.Ui32(v1001) {
		v1010 = v972
		goto L241
	} else {
		goto L248
	}
L248:
	;
	v1008 = v972 + v1001
	if base.Ui32(v1008) < base.Ui32(v961) {
		v972 = v1008
		v977 = v977 + int32(4)
		goto L239
	} else {
		goto L249
	}
L249:
	;
	v1010 = v1008
	goto L241
L250:
	;
	F_pad(m, l0, int32(32), v404, v1010, v595)
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L35
	} else {
		goto L251
	}
L251:
	;
	if v1010 == int32(0) {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v1077 = int32(0)
	goto L135
L253:
	;
	goto L254
L254:
	;
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v30)+48))
	v1033 = int32(0)
	v1036 = v1023
	goto L255
L255:
	;
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v1036)))
	if v1051 == int32(0) {
		v1077 = v1010
		goto L135
	} else {
		goto L257
	}
L256:
	;
	v1077 = v1010
	goto L135
L257:
	;
	v1055 = v28 + int32(-60)
	if v1055 == int32(0) {
		goto L259
	} else {
		goto L260
	}
L258:
	;
	v1061 = v1060 + v1033
	if base.Ui32(v1010) < base.Ui32(v1061) {
		v1077 = v1010
		goto L135
	} else {
		goto L262
	}
L259:
	;
	v1060 = int32(0)
	goto L258
L260:
	;
	goto L261
L261:
	;
	v1059 = F_wcrtomb(m, v1055, v1051)
	mBase = m.M
	v1060 = v1059
	goto L258
L262:
	;
	F_out(m, l0, v28+int32(-60), v1060)
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L35
	} else {
		goto L263
	}
L263:
	;
	if base.Ui32(v1061) < base.Ui32(v1010) {
		v1033 = v1061
		v1036 = v1036 + int32(4)
		goto L255
	} else {
		goto L264
	}
L264:
	;
	goto L256
L265:
	;
	if v1077 < v404 {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	v1103 = v404
	goto L268
L267:
	;
	v1103 = v1077
	goto L268
L268:
	;
	v66 = v545
	v72 = v1103
	v79 = v95
	v85 = v405
	goto L7
L269:
	;
	v1108 = *(*float64)(unsafe.Add(mBase, uint32(v30)+48))
	v1109 = m.T0[l5].(func(*base.Module, int32, float64, int32, int32, int32, int32) int32)(m, l0, v1108, v404, v509, v595, v607)
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L35
	} else {
		goto L270
	}
L270:
	;
	if int32(0) <= v1109 {
		v66 = v545
		v72 = v1109
		v79 = v95
		v85 = v405
		goto L7
	} else {
		goto L271
	}
L271:
	;
	v1364 = int32(61)
	goto L3
L272:
	;
	if v85 == int32(0) {
		goto L10
	} else {
		goto L273
	}
L273:
	;
	v1126 = int32(1)
	goto L274
L274:
	;
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(l4+v1126<<(uint(int32(2))%32))))
	if v1149 != 0 {
		goto L276
	} else {
		goto L277
	}
L275:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v1126) {
		goto L281
	} else {
		goto L282
	}
L276:
	;
	F_pop_arg(m, l3+v1126<<(uint(int32(3))%32), v1149, l2, l6)
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L35
	} else {
		goto L279
	}
L277:
	;
	goto L278
L278:
	;
	goto L275
L279:
	;
	v1155 = int32(1)
	v1157 = v1126 + v1155
	if v1157 != int32(10) {
		v1126 = v1157
		goto L274
	} else {
		goto L280
	}
L280:
	;
	v1418 = v1155
	goto L1
L281:
	;
	v1418 = int32(1)
	goto L1
L282:
	;
	goto L283
L283:
	;
	v1170 = v1126
	goto L284
L284:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(l4+v1170<<(uint(int32(2))%32))))
	if v1193 != 0 {
		goto L13
	} else {
		goto L286
	}
L285:
	;
	v1418 = v1194
	goto L1
L286:
	;
	v1194 = int32(1)
	v1196 = v1170 + v1194
	if v1196 != int32(10) {
		v1170 = v1196
		goto L284
	} else {
		goto L287
	}
L287:
	;
	goto L285
L288:
	;
	v1260 = v1241
	goto L290
L289:
	;
	v1260 = v1258
	goto L290
L290:
	;
	if v1247^int32(2147483647) < v1260 {
		goto L4
	} else {
		goto L291
	}
L291:
	;
	v1265 = v1247 + v1260
	if v1265 < v404 {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1267 = v404
	goto L294
L293:
	;
	v1267 = v1265
	goto L294
L294:
	;
	if v196 < v1267 {
		v1364 = int32(61)
		goto L3
	} else {
		goto L295
	}
L295:
	;
	F_pad(m, l0, int32(32), v1267, v1265, v1242)
	mBase = m.M
	v1271 = m.ExcPending
	if v1271 != 0 {
		goto L35
	} else {
		goto L296
	}
L296:
	;
	F_out(m, l0, v1254, v1247)
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L35
	} else {
		goto L297
	}
L297:
	;
	F_pad(m, l0, int32(48), v1267, v1265, v1242^int32(65536))
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
		goto L35
	} else {
		goto L298
	}
L298:
	;
	F_pad(m, l0, int32(48), v1260, v1258, int32(0))
	mBase = m.M
	v1282 = m.ExcPending
	if v1282 != 0 {
		goto L35
	} else {
		goto L299
	}
L299:
	;
	F_out(m, l0, v1244, v1258)
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		goto L35
	} else {
		goto L300
	}
L300:
	;
	F_pad(m, l0, int32(32), v1267, v1265, v1242^int32(8192))
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L35
	} else {
		goto L301
	}
L301:
	;
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(v30)+60))
	v66 = v1290
	v72 = v1267
	v79 = v95
	v85 = v405
	goto L7
}
func F_printtup_shutdown(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v3 != 0 {
		F_pfree(m, v3)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			v6 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v6
			*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v6
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			if v10 != 0 {
				F_pfree(m, v10)
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					if v15 != 0 {
						F_MemoryContextDelete(m, v15)
						mBase = m.M
						v17 = m.ExcPending
						if v17 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
						return
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				if v15 != 0 {
					F_MemoryContextDelete(m, v15)
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
					return
				}
			}
		}
	} else {
		v6 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v6
		*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v6
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		if v10 != 0 {
			F_pfree(m, v10)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				if v15 != 0 {
					F_MemoryContextDelete(m, v15)
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
					return
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			if v15 != 0 {
				F_MemoryContextDelete(m, v15)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
				return
			}
		}
	}
}
func F_process_owned_by(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v12 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L19
	} else {
		goto L65
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L19
	} else {
		goto L61
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L19
	} else {
		goto L57
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L19
	} else {
		goto L53
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L19
	} else {
		goto L48
	}
L6:
	;
	if l2 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L7:
	;
	v15 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v19 = int32(370343)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, _consts[398])))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v23 == v15 {
		v42 = v22
		v43 = v23
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L9
L9:
	;
	v69 = F_list_copy_head(m, l1, v12-int32(1))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L19
	} else {
		goto L25
	}
L10:
	;
	if v43-v42 == int32(0) {
		v119 = v15
		v120 = int32(0)
		goto L6
	} else {
		goto L18
	}
L11:
	;
	goto L10
L12:
	;
	if v22 != v23 {
		v42 = v22
		v43 = v23
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v27 = v18
	v28 = v19
	goto L14
L14:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)))
	if v32 == int32(0) {
		v42 = v31
		v43 = v32
		goto L11
	} else {
		goto L16
	}
L15:
	;
	v42 = v31
	v43 = v32
	goto L11
L16:
	;
	v35 = int32(1)
	if v31 == v32 {
		v27 = v27 + v35
		v28 = v28 + v35
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return
L20:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	F_errmsg(m, int32(245368), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	F_errhint(m, int32(640721), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(496352), int32(1611), int32(23883))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v71+v72<<(uint(int32(2))%32)-int32(4))))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v80 = F_makeRangeVarFromNameList(m, v69)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L19
	} else {
		goto L26
	}
L26:
	;
	v83 = F_relation_openrv(m, v80, int32(1))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L19
	} else {
		goto L27
	}
L27:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v83)+48))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+119)))
	v88 = v86 - int32(102)
	v97 = (v88<<(uint(int32(7))%32) | int32(base.Ui32(v88&int32(254))>>(uint(int32(1))%32))) & int32(255)
	if base.Ui32(int32(8)) < base.Ui32(v97) {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	if int32(1)<<(uint(v97)%32)&int32(353) == int32(0) {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+80))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v85)+80))
	if v107 != v108 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v106)+68))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v85)+68))
	if v110 != v111 {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v83)+56))
	v114 = F_get_attnum(m, v113, v79)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L19
	} else {
		goto L32
	}
L32:
	;
	if v114 == int32(0) {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v119 = v83
	v120 = v114
	goto L6
L34:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v130 = F_sequenceIsOwned(m, v124, int32(105), v8+int32(-12), v8+int32(-24))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L19
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v132 = int32(1259)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if l2 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	if v130 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v137 = int32(105)
	goto L41
L40:
	;
	v137 = int32(97)
	goto L41
L41:
	;
	v138 = F_deleteDependencyRecordsForClass(m, v132, v133, v132, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L19
	} else {
		goto L42
	}
L42:
	;
	if v119 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v140 = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v140
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v119)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v140
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v147
	F_recordDependencyOn(m, v8+int32(-24), v8+int32(-12), v137)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L19
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	m.G0 = v10 - int32(-64)
	return
L46:
	;
	F_relation_close(m, v119, int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L19
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L19
	} else {
		goto L49
	}
L49:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v83)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v172 + int32(4)
	F_errmsg(m, int32(685758), v8+int32(-48))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L19
	} else {
		goto L50
	}
L50:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v83)+48))
	v182 = int32(*(*int8)(unsafe.Add(mBase, uint32(v181)+119)))
	F_errdetail_relkind_not_supported(m, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L19
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(496352), int32(1638), int32(23883))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L19
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L19
	} else {
		goto L54
	}
L54:
	;
	F_errmsg(m, int32(238105), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L19
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(496352), int32(1644), int32(23883))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L19
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L19
	} else {
		goto L58
	}
L58:
	;
	F_errmsg(m, int32(238160), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L19
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(496352), int32(1648), int32(23883))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L19
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L19
	} else {
		goto L62
	}
L62:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v83)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v229 + int32(4)
	F_errmsg(m, int32(71286), v8+int32(-32))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L19
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(496352), int32(1656), int32(23883))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L19
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L19
	} else {
		goto L66
	}
L66:
	;
	F_errmsg(m, int32(413386), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L19
	} else {
		goto L67
	}
L67:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
	v257 = F_get_rel_name(m, v256)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L19
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v255 + int32(4)
	F_errdetail(m, int32(648755), v10)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L19
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(496352), int32(1673), int32(23883))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L19
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pt_contained_circle(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_point_dt(m, v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*float64)(unsafe.Add(mBase, uint32(v3)+16))
		return base.F64_le(v5, v9)
	}
}
func F_pull_varnos(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v3 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v3
	v17 = F_query_or_expression_tree_walker_impl(m, l1, int32(896), v6+int32(4), v3)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
		m.G0 = v6 + int32(16)
		return v21
	}
}
func F_pull_vars_of_level(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v3 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v3
	v15 = F_query_or_expression_tree_walker_impl(m, l0, int32(898), v6+int32(8), v3)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		m.G0 = v6 + int32(16)
		return v19
	}
}
func F_pull_vars_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	v3 = int32(0)
	if l0 == v3 {
		v50 = v3
		return v50
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v7 != int32(319) {
			if v7 == int32(67) {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v34 + int32(1)
				v40 = F_query_tree_walker_impl(m, l0, int32(898), l1, int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v42 - int32(1)
					return v40
				}
			} else {
				if v7 != int32(6) {
					v48 = F_expression_tree_walker_impl(m, l0, int32(898), l1)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						v50 = v48
						return v50
					}
				} else {
					v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					if v14 != v15 {
						v50 = v3
						return v50
					} else {
						v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v18 = F_lappend(m, v17, l0)
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v18
							return int32(0)
						}
					}
				}
			}
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v25 != v26 {
				v50 = v3
				return v50
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v29 = F_lappend(m, v28, l0)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v29
					return int32(0)
				}
			}
		}
	}
}
func F_pullf_read_max(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
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
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v15 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = l0
	goto L4
L2:
	;
	v36 = l0
	v39 = v15
	goto L3
L3:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	if l1 < v43 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v29 == int32(0) {
		v22 = v27
		goto L4
	} else {
		goto L6
	}
L5:
	;
	v36 = v27
	v39 = v29
	goto L3
L6:
	;
	goto L5
L7:
	;
	m.G0 = v12 + int32(16)
	return v128
L8:
	;
	v45 = l1
	goto L10
L9:
	;
	v45 = v43
	goto L10
L10:
	;
	if v43 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v46 = v45
	goto L13
L12:
	;
	v46 = l1
	goto L13
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v48 = m.T0[v39].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v41, v42, v46, l2, v47, v43)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	if v48 <= int32(0) {
		v128 = v48
		goto L7
	} else {
		goto L16
	}
L16:
	;
	if l1 == v48 {
		v128 = v48
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v48 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v57
	v59 = l1 - v48
	if v59 <= int32(0) {
		v128 = v48
		goto L7
	} else {
		goto L22
	}
L19:
	;
	v56 = F__emscripten_memcpy_bulkmem(m, l3, v55, v48)
	mBase = m.M
	v57 = v56
	goto L21
L20:
	;
	v57 = l3
	goto L21
L21:
	;
	goto L18
L22:
	;
	v63 = v59
	v67 = v48
	goto L23
L23:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v72 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v128 = v119
	goto L7
L25:
	;
	v79 = l0
	goto L28
L26:
	;
	v93 = l0
	v96 = v72
	goto L27
L27:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
	if v63 < v100 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if v86 == int32(0) {
		v79 = v84
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v93 = v84
	v96 = v86
	goto L27
L30:
	;
	goto L29
L31:
	;
	v102 = v63
	goto L33
L32:
	;
	v102 = v100
	goto L33
L33:
	;
	if v100 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v103 = v102
	goto L36
L35:
	;
	v103 = v63
	goto L36
L36:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v107 = m.T0[v96].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v98, v99, v103, v12+int32(12), v106, v100)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L14
	} else {
		goto L37
	}
L37:
	;
	if v107 < int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v112 = F___memset(m, v57, int32(0), v67)
	mBase = m.M
	goto L41
L39:
	;
	goto L40
L40:
	;
	if v107 == int32(0) {
		v128 = v67
		goto L7
	} else {
		goto L42
	}
L41:
	;
	v128 = v107
	goto L7
L42:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v107 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v119 = v107 + v67
	v120 = v63 - v107
	if int32(0) < v120 {
		v63 = v120
		v67 = v119
		goto L23
	} else {
		goto L47
	}
L44:
	;
	v117 = F__emscripten_memcpy_bulkmem(m, v57+v67, v116, v107)
	mBase = m.M
	goto L46
L45:
	;
	goto L46
L46:
	;
	goto L43
L47:
	;
	goto L24
}
func F_push_child_plan(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v5 = F__emscripten_memcpy_bulkmem(m, l2, l0, int32(80))
	mBase = m.M
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v9 = F_lcons(m, v7, v8)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v9
		F_set_deparse_plan(m, l0, l1)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			return
		}
	}
}
func F_pushf_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
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
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+12))
	if v4 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		m.T0[v4].(func(*base.Module, int32))(m, v5)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v8 != 0 {
				v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v11 = F___memset(m, v8, int32(0), v10)
				mBase = m.M
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				F_pfree(m, v12)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					v17 = F___memset(m, l0, int32(0), int32(24))
					mBase = m.M
					F_pfree(m, l0)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				v17 = F___memset(m, l0, int32(0), int32(24))
				mBase = m.M
				F_pfree(m, l0)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v8 != 0 {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v11 = F___memset(m, v8, int32(0), v10)
			mBase = m.M
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			F_pfree(m, v12)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				v17 = F___memset(m, l0, int32(0), int32(24))
				mBase = m.M
				F_pfree(m, l0)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			v17 = F___memset(m, l0, int32(0), int32(24))
			mBase = m.M
			F_pfree(m, l0)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				return
			}
		}
	}
}
