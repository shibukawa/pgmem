package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ParameterAclLookup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = F_convert_GUC_name_for_parameter_acl(m, l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_cstring_to_text(m, v10)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = int32(0)
			v19 = F_GetSysCacheOid(m, int32(43), v14, v16, v16, v16)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if l1|v19 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67137668))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
							F_errmsg(m, int32(_a_F_ParameterAclLookup_0), v7)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_ParameterAclLookup_1), int32(50), int32(_a_F_ParameterAclLookup_2))
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					F_pfree(m, v10)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(16)
						return v19
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
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	v1 = int32(0)
	v3 = int32(_a_F_PostgresMainLongJmp_0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[0])) = v5 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[1])) = v1
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[2])) = v1
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[3])) = v1
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[4])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[5])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[6])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[7])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[8])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[9])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[10])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[11])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[12])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[13])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[14])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[15])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[16])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[17])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[18])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[19])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[20])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[21])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[22])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[23])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[24])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[25])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[26])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[27])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[28])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[29])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[30])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[31])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[32])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[33])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[34])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[35])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[36])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[37])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[38])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[39])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[40])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[41])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[42])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[43])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[44])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[45])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[46])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[47])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[48])) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[49])) = uint8(v1)
	goto L1
L1:
	;
	v157 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[50])) = v157
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[51])) = uint8(v157)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[52])) = uint8(v157)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[53])) = uint8(v157)
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[54]))
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
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[55])) = int32(0)
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
	v181 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[56])))
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
	v250 = m.G0
	v252 = v250 - int32(32)
	m.G0 = v252
	v255 = v252 + int32(12)
	v257 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[57]))
	F_hash_seq_init(m, v255, v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
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
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[58]))
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
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[59]))
	if v195 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[60]))
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
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[61])) = v213
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[62]))
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
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[63]))
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
	v248 = m.ExcPending
	if v248 != 0 {
		goto L2
	} else {
		goto L37
	}
L28:
	;
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[64]))
	if v228 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v230 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[65]))
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
	F_s_lock(m, v230+int32(76), int32(_a_F_PostgresMainLongJmp_1), int32(3869), int32(_a_F_PostgresMainLongJmp_2))
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
	v260 = F_hash_seq_search(m, v255)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	if v260 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v263 = v260
	goto L43
L41:
	;
	goto L42
L42:
	;
	m.G0 = v252 + int32(32)
	v283 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[60]))
	if v283 != 0 {
		goto L51
	} else {
		goto L52
	}
L43:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)+64))
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264)+85)))
	if v265 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L42
L45:
	;
	v268 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v264)+84)) = uint8(v268)
	F_PortalDrop(m, v264, v268)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L2
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v275 = F_hash_seq_search(m, v252+int32(12))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L2
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	if v275 != 0 {
		v263 = v275
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
	v285 = m.ExcPending
	if v285 != 0 {
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
	v288 = m.ExcPending
	if v288 != 0 {
		goto L2
	} else {
		goto L55
	}
L54:
	;
	goto L53
L55:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[66])))
	if v290 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v292 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[67]))
	m.T0[v292].(func(*base.Module))(m)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L2
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v297 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[68]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[69])) = v297
	F_FlushErrorState(m)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L2
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[70])))
	if v302 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v304 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[71])) = uint8(v304)
	goto L63
L62:
	;
	goto L63
L63:
	;
	v307 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[72])) = uint8(v307)
	v310 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[73])))
	if v310 == v307 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v313 = int32(_a_F_PostgresMainLongJmp_0)
	v315 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[0])) = v315 - int32(1)
	v320 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[71])))
	if v320 == int32(0) {
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
	v329 = m.ExcPending
	if v329 != 0 {
		goto L2
	} else {
		goto L70
	}
L67:
	;
	v324 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLongJmp[74])) = uint8(v324)
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
	v332 = m.ExcPending
	if v332 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLongJmp_3), int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLongJmp_4), int32(_a_F_PostgresMainLongJmp_5), int32(_a_F_PostgresMainLongJmp_6))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
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
			F_errmsg_internal(m, int32(_a_F_PrepareSortSupportFromIndexRel_0), v9)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_PrepareSortSupportFromIndexRel_1), int32(170), int32(_a_F_PrepareSortSupportFromIndexRel_2))
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
				F_errmsg_internal(m, int32(_a_F_PrepareSortSupportFromOrderingOp_0), v6)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_PrepareSortSupportFromOrderingOp_1), int32(146), int32(_a_F_PrepareSortSupportFromOrderingOp_2))
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
	var v57 int32
	_ = v57
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v65 int64
	_ = v65
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
	var v112 int32
	_ = v112
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
	*(*int64)(unsafe.Add(mBase, uint32(v40)+24)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v40)+16)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v40)+8)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v40))) = v42
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
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+16)) = v57
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v53)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v55)+8)) = v59
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v53)))
	*(*int64)(unsafe.Add(mBase, uint32(v55))) = v61
	goto L11
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55)+16)) = int32(0)
	v65 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v55)+8)) = v65
	*(*int64)(unsafe.Add(mBase, uint32(v55))) = v65
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
	v112 = v111
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
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v112)+24))
	F_pfree(m, v112)
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
		v112 = v117
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
	var v26 int32
	_ = v26
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
	var v54 int32
	_ = v54
	var __phi54 int32
	_ = __phi54
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
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	v8 = int32(16)
	v12 = (int32(base.Ui32(l1)>>(uint(v8)%32)) ^ l1) * int32(-2048144789)
	v17 = (int32(base.Ui32(v12)>>(uint(int32(13))%32)) ^ v12) * int32(-1028477387)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v26 = int32(base.Ui32(v17)>>(uint(v8)%32)) ^ v17
	goto L1
L1:
	;
	v30 = v26 & v22
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
	v26 = v30 + int32(1)
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
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+4)) = uint8(v104)
	v112 = v38
	goto L4
L8:
	;
	v99 = v33
	goto L7
L9:
	;
	goto L10
L10:
	;
	__phi52 = v44
	__phi54 = v33
	__phi55 = v47
	__phi56 = v22
	v52 = __phi52
	v54 = __phi54
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
	v99 = v55
	goto L7
L13:
	;
	v99 = v54
	goto L7
L14:
	;
	goto L15
L15:
	;
	v74 = *(*int64)(unsafe.Add(mBase, uint32(v55)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+40)) = v74
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v55)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+32)) = v76
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v55)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+24)) = v78
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v55)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v80
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v55)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+8)) = v82
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
	*(*int64)(unsafe.Add(mBase, uint32(v54))) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v88 = int32(1)
	v90 = v87 & (v52 + v88)
	v93 = v86 + v90*int32(48)
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+4)))
	if v94 == v88 {
		__phi52 = v90
		__phi54 = v55
		__phi55 = v93
		__phi56 = v87
		v52 = __phi52
		v54 = __phi54
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
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v47 int64
	_ = v47
	var v50 int32
	_ = v50
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v68 int64
	_ = v68
	var v73 int32
	_ = v73
	var v74 int64
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
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
	var v102 int32
	_ = v102
	var v104 int64
	_ = v104
	var v114 int64
	_ = v114
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int64
	_ = v234
	var v236 int64
	_ = v236
	var v238 int64
	_ = v238
	var v240 int64
	_ = v240
	var v242 int64
	_ = v242
	var v244 int64
	_ = v244
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v372 int64
	_ = v372
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int64
	_ = v420
	var v422 int64
	_ = v422
	var v424 int64
	_ = v424
	var v426 int64
	_ = v426
	var v428 int64
	_ = v428
	var v430 int64
	_ = v430
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v453 int64
	_ = v453
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v528 int32
	_ = v528
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	v15 = int32(16)
	v19 = (int32(base.Ui32(l1)>>(uint(v15)%32)) ^ l1) * int32(-2048144789)
	v24 = (int32(base.Ui32(v19)>>(uint(int32(13))%32)) ^ v19) * int32(-1028477387)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v34 = base.B2i32(base.Ui32(v28) < base.Ui32(v29))
	goto L1
L1:
	;
	if v34 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L26
	} else {
		goto L100
	}
L3:
	;
	goto L2
L4:
	;
	v528 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v528
	v34 = v528
	goto L1
L5:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v506)
	return v505
L6:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v491 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v490 + v491
	*(*uint8)(unsafe.Add(mBase, uint32(v483)+4)) = uint8(v491)
	*(*int32)(unsafe.Add(mBase, uint32(v483))) = l1
	v505 = v483
	v506 = int32(0)
	goto L5
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L26
	} else {
		goto L97
	}
L8:
	;
	v47 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v47 == int64(4294967296) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v299 = (int32(base.Ui32(v24)>>(uint(v15)%32)) ^ v24) & v298
	v302 = v297 + v299*int32(48)
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302)+4)))
	if v303 == int32(0) {
		v483 = v302
		goto L6
	} else {
		goto L66
	}
L11:
	;
	v50 = int32(0)
	v52 = int64(2)
	v54 = v47 << (uint(int64(1)) % 64)
	if base.Ui64(v54) <= base.Ui64(v52) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v34 = int32(1)
	goto L1
L13:
	;
	v57 = v52
	goto L15
L14:
	;
	v57 = v54
	goto L15
L15:
	;
	v58 = int64(1)
	if v57&(v57-v58) == int64(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v68 = v57
	goto L18
L17:
	;
	v68 = v58 << (uint(int64(64)-base.I64_clz(v57)) % 64)
	goto L18
L18:
	;
	if base.Ui64(v68*int64(48)) < base.Ui64(int64(2147483647)) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v77 = base.I32_wrap_i64(v68) * int32(48)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+112))
	if v79 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	goto L21
L21:
	;
	goto L3
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v102
	v104 = int64(1)
	if v68&(v68-v104) == int64(0) {
		goto L30
	} else {
		goto L31
	}
L23:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v84 = F_MemoryContextAllocExtended(m, v82, v77, int32(5))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v78)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+100)) = v88
	v93 = F_dsa_allocate_extended(m, v79, v77|int32(4), int32(5))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L26
	} else {
		goto L28
	}
L26:
	;
	return int32(0)
L27:
	;
	v102 = v84
	goto L22
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+96)) = v93
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v78)+112))
	v97 = F_dsa_get_address(m, v96, v93)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v102 = v97 + int32(4)
	goto L22
L30:
	;
	v114 = v68
	goto L32
L31:
	;
	v114 = v104 << (uint(int64(64)-base.I64_clz(v68)) % 64)
	goto L32
L32:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v114*int64(48)) {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v114
	v122 = base.I32_wrap_i64(v114) - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v122
	if v114 == int64(4294967296) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v131 = int32(-85899346)
	goto L36
L35:
	;
	v131 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i64_u(v114), float64(0.9)))
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v131
	if v74 != int64(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v141 = v50
	goto L41
L38:
	;
	goto L39
L39:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)+112))
	if v285 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L40:
	;
	v184 = v177
	v188 = v50
	goto L46
L41:
	;
	v151 = v73 + v141*int32(48)
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+4)))
	if v152 != int32(1) {
		v177 = v141
		goto L40
	} else {
		goto L43
	}
L42:
	;
	v177 = int32(0)
	goto L40
L43:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	v156 = int32(16)
	v160 = (int32(base.Ui32(v155)>>(uint(v156)%32)) ^ v155) * int32(-2048144789)
	v165 = (int32(base.Ui32(v160)>>(uint(int32(13))%32)) ^ v160) * int32(-1028477387)
	if (int32(base.Ui32(v165)>>(uint(v156)%32))^v165)&v122 == v141 {
		v177 = v141
		goto L40
	} else {
		goto L44
	}
L44:
	;
	v172 = v141 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v172)) < base.Ui64(v74) {
		v141 = v172
		goto L41
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	v194 = v73 + v184*int32(48)
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+4)))
	if v195 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L39
L48:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	v199 = int32(16)
	v203 = (int32(base.Ui32(v198)>>(uint(v199)%32)) ^ v198) * int32(-2048144789)
	v208 = (int32(base.Ui32(v203)>>(uint(int32(13))%32)) ^ v203) * int32(-1028477387)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v218 = int32(base.Ui32(v208)>>(uint(v199)%32)) ^ v208
	goto L51
L49:
	;
	goto L50
L50:
	;
	v261 = v184 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v261)) < base.Ui64(v74) {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v227 = v212 & v218
	v232 = v102 + v227*int32(48)
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232)+4)))
	if v233 != 0 {
		v218 = v227 + int32(1)
		goto L51
	} else {
		goto L53
	}
L52:
	;
	v234 = *(*int64)(unsafe.Add(mBase, uint32(v194)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v232)+40)) = v234
	v236 = *(*int64)(unsafe.Add(mBase, uint32(v194)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v232)+32)) = v236
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v194)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v232)+24)) = v238
	v240 = *(*int64)(unsafe.Add(mBase, uint32(v194)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v232)+16)) = v240
	v242 = *(*int64)(unsafe.Add(mBase, uint32(v194)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v232)+8)) = v242
	v244 = *(*int64)(unsafe.Add(mBase, uint32(v194)))
	*(*int64)(unsafe.Add(mBase, uint32(v232))) = v244
	goto L50
L53:
	;
	goto L52
L54:
	;
	v265 = v261
	goto L56
L55:
	;
	v265 = int32(0)
	goto L56
L56:
	;
	v267 = v188 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v267)) < base.Ui64(v74) {
		v184 = v265
		v188 = v267
		goto L46
	} else {
		goto L57
	}
L57:
	;
	goto L47
L58:
	;
	F_pfree(m, v73)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L26
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v284)+100))
	if v290 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	goto L12
L62:
	;
	F_dsa_free(m, v285, v290)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L26
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	goto L12
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v284)+100)) = int32(0)
	goto L64
L66:
	;
	v311 = v299
	v314 = int32(0)
	v315 = v302
	goto L67
L67:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v315)))
	if v322 == l1 {
		v505 = v315
		v506 = int32(1)
		goto L5
	} else {
		goto L69
	}
L68:
	;
	v483 = v461
	goto L6
L69:
	;
	v325 = v311 + int32(1)
	v326 = int32(16)
	v330 = (int32(base.Ui32(v322)>>(uint(v326)%32)) ^ v322) * int32(-2048144789)
	v335 = (int32(base.Ui32(v330)>>(uint(int32(13))%32)) ^ v330) * int32(-1028477387)
	v339 = (int32(base.Ui32(v335)>>(uint(v326)%32)) ^ v335) & v298
	if base.Ui32(v311) < base.Ui32(v339) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v343 = v311 + v341
	goto L72
L71:
	;
	v343 = v311
	goto L72
L72:
	;
	if base.Ui32(v343-v339) < base.Ui32(v314) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v346 = v325 & v298
	v349 = v297 + v346*int32(48)
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+4)))
	if v350 != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	v448 = v314 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v448) {
		goto L92
	} else {
		goto L93
	}
L76:
	;
	v358 = int32(0)
	v360 = v346
	goto L79
L77:
	;
	v389 = v349
	v392 = v346
	goto L78
L78:
	;
	if v311 != v392 {
		goto L86
	} else {
		goto L87
	}
L79:
	;
	v367 = v358 + int32(1)
	if int32(151) <= v367 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v389 = v382
	v392 = v379
	goto L78
L81:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v372 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v370), base.F64_convert_i64_u(v372)), float64(0.1)) != 0 {
		goto L4
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v379 = (v360 + int32(1)) & v298
	v382 = v297 + v379*int32(48)
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382)+4)))
	if v383 != 0 {
		v358 = v367
		v360 = v379
		goto L79
	} else {
		goto L85
	}
L84:
	;
	goto L83
L85:
	;
	goto L80
L86:
	;
	v404 = v389
	v407 = v392
	goto L89
L87:
	;
	goto L88
L88:
	;
	v483 = v315
	goto L6
L89:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v416 = v413 & (v407 - int32(1))
	v419 = v297 + v416*int32(48)
	v420 = *(*int64)(unsafe.Add(mBase, uint32(v419)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v404)+40)) = v420
	v422 = *(*int64)(unsafe.Add(mBase, uint32(v419)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v404)+32)) = v422
	v424 = *(*int64)(unsafe.Add(mBase, uint32(v419)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v404)+24)) = v424
	v426 = *(*int64)(unsafe.Add(mBase, uint32(v419)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v404)+16)) = v426
	v428 = *(*int64)(unsafe.Add(mBase, uint32(v419)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v404)+8)) = v428
	v430 = *(*int64)(unsafe.Add(mBase, uint32(v419)))
	*(*int64)(unsafe.Add(mBase, uint32(v404))) = v430
	if v311 != v416 {
		v404 = v419
		v407 = v416
		goto L89
	} else {
		goto L91
	}
L90:
	;
	goto L88
L91:
	;
	goto L90
L92:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v453 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v451), base.F64_convert_i64_u(v453)), float64(0.1)) != 0 {
		goto L4
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v458 = v325 & v298
	v461 = v297 + v458*int32(48)
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461)+4)))
	if v462 != 0 {
		v311 = v458
		v314 = v448
		v315 = v461
		goto L67
	} else {
		goto L96
	}
L95:
	;
	goto L94
L96:
	;
	goto L68
L97:
	;
	F_errmsg_internal(m, int32(_a_F_pagetable_insert_0), int32(0))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L26
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_pagetable_insert_1), int32(630), int32(_a_F_pagetable_insert_2))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L26
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	F_errmsg_internal(m, int32(_a_F_pagetable_insert_3), int32(0))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L26
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_pagetable_insert_1), int32(327), int32(_a_F_pagetable_insert_4))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L26
	} else {
		goto L102
	}
L102:
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
	v42 = *(*float64)(unsafe.Add(mBase, uint32(l0+v23+v39)))
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_palloc_extended[0]))
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+4)) = uint8(v3)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v14 = m.T0[v13].(func(*base.Module, int32, int32, int32) int32)(m, v5, l0, l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if base.B2i32(l1&int32(4) == v3)|base.B2i32(v14 == int32(0)) != 0 {
			return v14
		} else {
			if l0&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(l0)) == int32(0) {
				if l0 == int32(0) {
					return v14
				} else {
					v32 = l0 + v14
					v34 = v14 + int32(4)
					if base.Ui32(v34) < base.Ui32(v32) {
						v36 = v32
					} else {
						v36 = v34
					}
					v41 = (v14^int32(-1)+v36)&int32(-4) + int32(4)
					if v41 == int32(0) {
						return v14
					} else {
						base.MemoryFill(m, v14, int32(0), v41)
						return v14
					}
				}
			} else {
				if l0 == int32(0) {
				} else {
					base.MemoryFill(m, v14, int32(0), l0)
				}
				return v14
			}
		}
	}
}
func F_parseXidFromText(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn13954(m, l0, l1, l2, int32(_a_F_parseXidFromText_0), int32(1349), int32(1344), int32(1339), int32(_a_F_parseXidFromText_1))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_parse_format(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v414 int32
	_ = v414
	var v425 int32
	_ = v425
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v599 int32
	_ = v599
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
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v667 int32
	_ = v667
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v690 int32
	_ = v690
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v815 int32
	_ = v815
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v831 int32
	_ = v831
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v863 int32
	_ = v863
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v900 int32
	_ = v900
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v927 int32
	_ = v927
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v943 int32
	_ = v943
	var v947 int32
	_ = v947
	var v952 int32
	_ = v952
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v967 int32
	_ = v967
	var v972 int32
	_ = v972
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v19 == int32(0) {
		v720 = l0
		goto L15
	} else {
		goto L16
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L78
	} else {
		goto L265
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L78
	} else {
		goto L260
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L78
	} else {
		goto L256
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L78
	} else {
		goto L252
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L78
	} else {
		goto L248
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L78
	} else {
		goto L244
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L78
	} else {
		goto L240
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L78
	} else {
		goto L236
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L78
	} else {
		goto L232
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L78
	} else {
		goto L228
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L78
	} else {
		goto L224
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L78
	} else {
		goto L220
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L78
	} else {
		goto L216
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L78
	} else {
		goto L212
	}
L15:
	;
	v734 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v720)+6)) = uint8(v734)
	v736 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v720))) = uint8(v736)
	m.G0 = v17 + int32(16)
	return
L16:
	;
	v25 = l5 & int32(1)
	v28 = l0
	v29 = l1
	v36 = v19
	goto L17
L17:
	;
	if v25 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v720 = v705
	goto L15
L19:
	;
	v719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706))))
	if v719 != 0 {
		v28 = v705
		v29 = v706
		v36 = v719
		goto L17
	} else {
		goto L211
	}
L20:
	;
	if base.Ui32((v139-int32(126))&int32(255)) < base.Ui32(int32(163)) {
		goto L54
	} else {
		goto L55
	}
L21:
	;
	v132 = v29
	v139 = v36
	v140 = int32(0)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v45 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v132 = v29
	v139 = v36
	v140 = int32(0)
	goto L20
L25:
	;
	goto L26
L26:
	;
	v54 = l3
	v56 = v45
	goto L29
L27:
	;
	if v125&int32(255) == int32(0) {
		v705 = v28
		v706 = v122
		goto L19
	} else {
		goto L49
	}
L28:
	;
	v119 = v29 + v66
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v122 = v119
	v125 = v120
	v126 = v121
	goto L27
L29:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	if v63 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v122 = v29
	v125 = v36
	v126 = int32(0)
	goto L27
L31:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v66 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	goto L33
L33:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	if v115 != 0 {
		v54 = v54 + int32(16)
		v56 = v115
		goto L29
	} else {
		goto L48
	}
L34:
	;
	if v111 == int32(0) {
		goto L28
	} else {
		goto L47
	}
L35:
	;
	v111 = int32(0)
	goto L34
L36:
	;
	goto L37
L37:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v72 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v73 = v29
	v74 = v56
	v75 = v66
	v76 = v72
	goto L42
L39:
	;
	v99 = v56
	v103 = int32(0)
	goto L40
L40:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	v111 = v103 - v104
	goto L34
L41:
	;
	v99 = v94
	v103 = v96
	goto L40
L42:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if base.B2i32(v76 != v78)|base.B2i32(v78 == int32(0)) != 0 {
		v94 = v74
		v96 = v76
		goto L41
	} else {
		goto L44
	}
L43:
	;
	v94 = v88
	v96 = int32(0)
	goto L41
L44:
	;
	v84 = v75 - int32(1)
	if v84 == int32(0) {
		v94 = v74
		v96 = v76
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v87 = int32(1)
	v88 = v74 + v87
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
	if v89 != 0 {
		v73 = v73 + v87
		v74 = v88
		v75 = v84
		v76 = v89
		goto L42
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	goto L33
L48:
	;
	goto L30
L49:
	;
	v132 = v122
	v139 = v125
	v140 = v126
	goto L20
L50:
	;
	v705 = v28 + int32(12)
	v706 = v690
	goto L19
L51:
	;
	if v250 != int32(92) {
		goto L179
	} else {
		goto L180
	}
L52:
	;
	if v250 == int32(32) {
		goto L175
	} else {
		goto L176
	}
L53:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+6)) = uint8(v140)
	v287 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v287)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = v170
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v170)+4))
	if l5&int32(2) == int32(0) {
		goto L85
	} else {
		goto L86
	}
L54:
	;
	v248 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = v248
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	if v250 == v248 {
		v720 = v28
		goto L15
	} else {
		goto L75
	}
L55:
	;
	v152 = v139 & int32(255)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l4+v152<<(uint(int32(2))%32)-int32(128))))
	if v158 < int32(0) {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v163 = l2 + v158*int32(20)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v170 = v163
	v172 = v164
	goto L57
L57:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v170)+4))
	if v179 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L54
L59:
	;
	if v224 == int32(0) {
		goto L53
	} else {
		goto L72
	}
L60:
	;
	v224 = int32(0)
	goto L59
L61:
	;
	goto L62
L62:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	if v185 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v186 = v132
	v187 = v172
	v188 = v179
	v189 = v185
	goto L67
L64:
	;
	v212 = v172
	v216 = int32(0)
	goto L65
L65:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
	v224 = v216 - v217
	goto L59
L66:
	;
	v212 = v207
	v216 = v209
	goto L65
L67:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	if base.B2i32(v189 != v191)|base.B2i32(v191 == int32(0)) != 0 {
		v207 = v187
		v209 = v189
		goto L66
	} else {
		goto L69
	}
L68:
	;
	v207 = v201
	v209 = int32(0)
	goto L66
L69:
	;
	v197 = v188 - int32(1)
	if v197 == int32(0) {
		v207 = v187
		v209 = v189
		goto L66
	} else {
		goto L70
	}
L70:
	;
	v200 = int32(1)
	v201 = v187 + v200
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186)+1)))
	if v202 != 0 {
		v186 = v186 + v200
		v187 = v201
		v188 = v197
		v189 = v202
		goto L67
	} else {
		goto L71
	}
L71:
	;
	goto L68
L72:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v170)+20))
	if v227 == int32(0) {
		goto L54
	} else {
		goto L73
	}
L73:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227))))
	if v152 == v232 {
		v170 = v170 + int32(20)
		v172 = v227
		goto L57
	} else {
		goto L74
	}
L74:
	;
	goto L58
L75:
	;
	if base.B2i32(base.Ui32(l5) < base.Ui32(int32(4)))|base.B2i32(v250 == int32(34)) != 0 {
		goto L51
	} else {
		goto L76
	}
L76:
	;
	if base.B2i32(base.Ui32(v250) <= base.Ui32(int32(63)))&base.B2i32(int64(1)<<(uint(base.I64_extend_i32_u(v250))%64)&int64(864955565296582657) != int64(0)) != 0 {
		goto L52
	} else {
		goto L77
	}
L77:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	return
L79:
	;
	F_errcode(m, int32(117440642))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v273 = F_pg_mblen_cstr(m, v132)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L78
	} else {
		goto L81
	}
L81:
	;
	v275 = F_pnstrdup(m, v132, v273)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L78
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v275
	F_errmsg(m, int32(_a_F_parse_format_0), v17)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L78
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1446), int32(_a_F_parse_format_2))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L78
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	v480 = v132 + v290
	if v25 == int32(0) {
		v690 = v480
		goto L50
	} else {
		goto L151
	}
L86:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v170)+8))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	if v294&int32(_a_F_parse_format_3) != 0 {
		goto L91
	} else {
		goto L92
	}
L87:
	;
	if v469&int32(1024) == int32(0) {
		goto L85
	} else {
		goto L149
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6)+12)) = v466
	v469 = v466
	goto L87
L89:
	;
	if v294&int32(4080) != 0 {
		goto L2
	} else {
		goto L148
	}
L90:
	;
	F_errmsg(m, int32(_a_F_parse_format_4), int32(0))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L78
	} else {
		goto L146
	}
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L78
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	switch v293 - int32(1) {
	case 0:
		v379 = v294
		goto L108
	case 1:
		goto L111
	case 2:
		goto L112
	case 3:
		goto L110
	default:
		v469 = v294
		goto L87
	case 5:
		goto L109
	case 6:
		goto L89
	case 7:
		goto L107
	case 8, 9:
		goto L100
	case 10:
		goto L105
	case 11:
		goto L104
	case 12:
		goto L102
	case 13, 29:
		goto L101
	case 14:
		goto L103
	case 16:
		goto L106
	case 18:
		goto L99
	}
L94:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L78
	} else {
		goto L95
	}
L95:
	;
	if v293 == int32(7) {
		goto L90
	} else {
		goto L96
	}
L96:
	;
	F_errmsg(m, int32(_a_F_parse_format_5), int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L78
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1197), int32(_a_F_parse_format_6))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L78
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L99:
	;
	if v294&int32(2) != 0 {
		goto L3
	} else {
		goto L145
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6)+32)) = int32(1)
	v469 = v294
	goto L87
L101:
	;
	if v294&int32(1024) != 0 {
		goto L4
	} else {
		goto L144
	}
L102:
	;
	if v294&int32(832) != 0 {
		goto L5
	} else {
		goto L143
	}
L103:
	;
	if v294&int32(64) != 0 {
		goto L6
	} else {
		goto L142
	}
L104:
	;
	if v294&int32(64) != 0 {
		goto L7
	} else {
		goto L140
	}
L105:
	;
	if v294&int32(64) != 0 {
		goto L8
	} else {
		goto L138
	}
L106:
	;
	if v294&int32(64) != 0 {
		goto L10
	} else {
		goto L132
	}
L107:
	;
	v466 = v294 | int32(32)
	goto L88
L108:
	;
	if v379&int32(2) != 0 {
		goto L12
	} else {
		goto L130
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6)+32)) = int32(1)
	v377 = v294 | int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l6)+12)) = v377
	v379 = v377
	goto L108
L110:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v367 != 0 {
		v469 = v294
		goto L87
	} else {
		goto L128
	}
L111:
	;
	if v294&int32(128) != 0 {
		goto L13
	} else {
		goto L120
	}
L112:
	;
	if v294&int32(128) != 0 {
		goto L14
	} else {
		goto L113
	}
L113:
	;
	if v294&int32(2048) != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l6)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l6)+20)) = v321 + int32(1)
	v469 = v294
	goto L87
L115:
	;
	goto L116
L116:
	;
	if v294&int32(2) != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l6)+4)) = v327 + int32(1)
	v469 = v294
	goto L87
L118:
	;
	goto L119
L119:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v331 + int32(1)
	v469 = v294
	goto L87
L120:
	;
	if v294&int32(10) == int32(0) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v342 = v294 | int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l6)+12)) = v342
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	*(*int32)(unsafe.Add(mBase, uint32(l6)+24)) = v344 + int32(1)
	v348 = v342
	goto L123
L122:
	;
	v348 = v294
	goto L123
L123:
	;
	if v348&int32(2) == int32(0) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6)+28)) = v363 + v364
	v469 = v348
	goto L87
L125:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v355 = v353 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v355
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v363 = v355
	v364 = v357
	goto L124
L126:
	;
	goto L127
L127:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v360 = v358 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l6)+4)) = v360
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v363 = v362
	v364 = v360
	goto L124
L128:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v368|v294&int32(8) != 0 {
		v469 = v294
		goto L87
	} else {
		goto L129
	}
L129:
	;
	v466 = v294 | int32(16)
	goto L88
L130:
	;
	if v379&int32(2048) != 0 {
		goto L11
	} else {
		goto L131
	}
L131:
	;
	v466 = v379 | int32(2)
	goto L88
L132:
	;
	if v294&int32(896) != 0 {
		goto L9
	} else {
		goto L133
	}
L133:
	;
	if v294&int32(2) == int32(0) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6)+8)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l6)+32)) = int32(1)
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	*(*int32)(unsafe.Add(mBase, uint32(l6)+16)) = v400
	v466 = v294 | int32(64)
	goto L88
L135:
	;
	goto L136
L136:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	if v404 != 0 {
		v469 = v294
		goto L87
	} else {
		goto L137
	}
L137:
	;
	v405 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l6)+32)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(l6)+8)) = v405
	v466 = v294 | int32(64)
	goto L88
L138:
	;
	v414 = v294 | int32(256)
	*(*int32)(unsafe.Add(mBase, uint32(l6)+12)) = v414
	if v294&int32(2) == int32(0) {
		v469 = v414
		goto L87
	} else {
		goto L139
	}
L139:
	;
	v466 = v294 | int32(_a_F_parse_format_7)
	goto L88
L140:
	;
	v425 = v294 | int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(l6)+12)) = v425
	if v294&int32(2) == int32(0) {
		v469 = v425
		goto L87
	} else {
		goto L141
	}
L141:
	;
	v466 = v294 | int32(_a_F_parse_format_8)
	goto L88
L142:
	;
	v466 = v294 | int32(768)
	goto L88
L143:
	;
	v466 = v294 | int32(128)
	goto L88
L144:
	;
	v466 = v294 | int32(1024)
	goto L88
L145:
	;
	v466 = v294 | int32(2048)
	goto L88
L146:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1347), int32(_a_F_parse_format_6))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L78
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	v466 = v294 | int32(_a_F_parse_format_3)
	goto L88
L149:
	;
	if v469&int32(-1057) != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	goto L85
L151:
	;
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480))))
	if v483 == int32(0) {
		v690 = v480
		goto L50
	} else {
		goto L152
	}
L152:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v486 == int32(0) {
		v690 = v480
		goto L50
	} else {
		goto L153
	}
L153:
	;
	v494 = l3
	v496 = v486
	goto L154
L154:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v494)+12))
	if v503 == int32(2) {
		goto L157
	} else {
		goto L158
	}
L155:
	;
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+6)))
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494)+8)))
	v559 = v557 | v558
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+6)) = uint8(v559)
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v494)+4))
	v690 = v480 + v561
	goto L50
L156:
	;
	goto L155
L157:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v494)+4))
	if v506 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L158:
	;
	goto L159
L159:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v494)+16))
	if v554 != 0 {
		v494 = v494 + int32(16)
		v496 = v554
		goto L154
	} else {
		goto L174
	}
L160:
	;
	if v551 == int32(0) {
		goto L156
	} else {
		goto L173
	}
L161:
	;
	v551 = int32(0)
	goto L160
L162:
	;
	goto L163
L163:
	;
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480))))
	if v512 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v513 = v480
	v514 = v496
	v515 = v506
	v516 = v512
	goto L168
L165:
	;
	v539 = v496
	v543 = int32(0)
	goto L166
L166:
	;
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v539))))
	v551 = v543 - v544
	goto L160
L167:
	;
	v539 = v534
	v543 = v536
	goto L166
L168:
	;
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514))))
	if base.B2i32(v516 != v518)|base.B2i32(v518 == int32(0)) != 0 {
		v534 = v514
		v536 = v516
		goto L167
	} else {
		goto L170
	}
L169:
	;
	v534 = v528
	v536 = int32(0)
	goto L167
L170:
	;
	v524 = v515 - int32(1)
	if v524 == int32(0) {
		v534 = v514
		v536 = v516
		goto L167
	} else {
		goto L171
	}
L171:
	;
	v527 = int32(1)
	v528 = v514 + v527
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513)+1)))
	if v529 != 0 {
		v513 = v513 + v527
		v514 = v528
		v515 = v524
		v516 = v529
		goto L168
	} else {
		goto L172
	}
L172:
	;
	goto L169
L173:
	;
	goto L159
L174:
	;
	v690 = v480
	goto L50
L175:
	;
	v567 = int32(5)
	goto L177
L176:
	;
	v567 = int32(4)
	goto L177
L177:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v567)
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	v570 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = v570
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+2)) = uint8(v570)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)) = uint8(v569)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+6)) = uint8(v570)
	v690 = v132 + int32(1)
	goto L50
L178:
	;
	v637 = F_pg_mblen_cstr(m, v636)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L78
	} else {
		goto L201
	}
L179:
	;
	if v250 != int32(34) {
		v636 = v132
		goto L178
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+1)))
	if v632 == int32(34) {
		goto L198
	} else {
		goto L199
	}
L182:
	;
	v585 = v28
	v590 = v132 + int32(1)
	goto L183
L183:
	;
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v590))))
	if v599 != int32(92) {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	v613 = F_pg_mblen_cstr(m, v612)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L78
	} else {
		goto L194
	}
L186:
	;
	if v599 == int32(0) {
		v720 = v585
		goto L15
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v590)+1)))
	if v610 != 0 {
		goto L191
	} else {
		goto L192
	}
L189:
	;
	if v599 != int32(34) {
		v612 = v590
		goto L185
	} else {
		goto L190
	}
L190:
	;
	v705 = v585
	v706 = v590 + int32(1)
	goto L19
L191:
	;
	v611 = v590 + int32(1)
	goto L193
L192:
	;
	v611 = v590
	goto L193
L193:
	;
	v612 = v611
	goto L185
L194:
	;
	v615 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v585))) = uint8(v615)
	v618 = v585 + int32(1)
	if v613 != 0 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	base.MemoryCopy(m, v618, v612, v613)
	goto L197
L196:
	;
	goto L197
L197:
	;
	v621 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v613+v618))) = uint8(v621)
	*(*uint8)(unsafe.Add(mBase, uint32(v585)+6)) = uint8(v621)
	*(*int32)(unsafe.Add(mBase, uint32(v585)+8)) = v621
	v585 = v585 + int32(12)
	v590 = v613 + v612
	goto L183
L198:
	;
	v635 = v132 + int32(1)
	goto L200
L199:
	;
	v635 = v132
	goto L200
L200:
	;
	v636 = v635
	goto L178
L201:
	;
	v639 = int32(0)
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v636))))
	v644 = int32(255)
	if base.B2i32(v25 == v639)|base.B2i32(base.Ui32(int32(93)) < base.Ui32((v641-int32(33))&v644))|base.B2i32(base.Ui32(int32(229)) < base.Ui32((v641&int32(223)-int32(91))&v644)) == v639 {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v676)
	v679 = v28 + int32(1)
	if v637 != 0 {
		goto L208
	} else {
		goto L209
	}
L203:
	;
	if base.Ui32((v641-int32(58))&int32(255)) < base.Ui32(int32(246)) {
		v676 = int32(4)
		goto L202
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	v667 = int32(5)
	if base.B2i32(v641 == int32(32))|base.B2i32(base.Ui32(v641-int32(9)) < base.Ui32(v667)) != 0 {
		v676 = v667
		goto L202
	} else {
		goto L207
	}
L206:
	;
	goto L205
L207:
	;
	v676 = int32(3)
	goto L202
L208:
	;
	base.MemoryCopy(m, v679, v636, v637)
	goto L210
L209:
	;
	goto L210
L210:
	;
	v682 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v637+v679))) = uint8(v682)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+6)) = uint8(v682)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = v682
	v690 = v636 + v637
	goto L50
L211:
	;
	goto L18
L212:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L78
	} else {
		goto L213
	}
L213:
	;
	F_errmsg(m, int32(_a_F_parse_format_9), int32(0))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L78
	} else {
		goto L214
	}
L214:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1205), int32(_a_F_parse_format_6))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L78
	} else {
		goto L215
	}
L215:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L216:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L78
	} else {
		goto L217
	}
L217:
	;
	F_errmsg(m, int32(_a_F_parse_format_10), int32(0))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L78
	} else {
		goto L218
	}
L218:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1221), int32(_a_F_parse_format_6))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L78
	} else {
		goto L219
	}
L219:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L220:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L78
	} else {
		goto L221
	}
L221:
	;
	F_errmsg(m, int32(_a_F_parse_format_11), int32(0))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L78
	} else {
		goto L222
	}
L222:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1248), int32(_a_F_parse_format_6))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L78
	} else {
		goto L223
	}
L223:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L224:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L78
	} else {
		goto L225
	}
L225:
	;
	F_errmsg(m, int32(_a_F_parse_format_12), int32(0))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L78
	} else {
		goto L226
	}
L226:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1252), int32(_a_F_parse_format_6))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L78
	} else {
		goto L227
	}
L227:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L228:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L78
	} else {
		goto L229
	}
L229:
	;
	F_errmsg(m, int32(_a_F_parse_format_13), int32(0))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L78
	} else {
		goto L230
	}
L230:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1264), int32(_a_F_parse_format_6))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L78
	} else {
		goto L231
	}
L231:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L232:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L78
	} else {
		goto L233
	}
L233:
	;
	F_errmsg(m, int32(_a_F_parse_format_14), int32(0))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L78
	} else {
		goto L234
	}
L234:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1268), int32(_a_F_parse_format_6))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L78
	} else {
		goto L235
	}
L235:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L236:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L78
	} else {
		goto L237
	}
L237:
	;
	F_errmsg(m, int32(_a_F_parse_format_15), int32(0))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L78
	} else {
		goto L238
	}
L238:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1288), int32(_a_F_parse_format_6))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L78
	} else {
		goto L239
	}
L239:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L240:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L78
	} else {
		goto L241
	}
L241:
	;
	F_errmsg(m, int32(_a_F_parse_format_16), int32(0))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L78
	} else {
		goto L242
	}
L242:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1298), int32(_a_F_parse_format_6))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L78
	} else {
		goto L243
	}
L243:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L244:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L78
	} else {
		goto L245
	}
L245:
	;
	F_errmsg(m, int32(_a_F_parse_format_17), int32(0))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L78
	} else {
		goto L246
	}
L246:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1308), int32(_a_F_parse_format_6))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L78
	} else {
		goto L247
	}
L247:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L248:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L78
	} else {
		goto L249
	}
L249:
	;
	F_errmsg(m, int32(_a_F_parse_format_18), int32(0))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L78
	} else {
		goto L250
	}
L250:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1317), int32(_a_F_parse_format_6))
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L78
	} else {
		goto L251
	}
L251:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L252:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L78
	} else {
		goto L253
	}
L253:
	;
	F_errmsg(m, int32(_a_F_parse_format_19), int32(0))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L78
	} else {
		goto L254
	}
L254:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1326), int32(_a_F_parse_format_6))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L78
	} else {
		goto L255
	}
L255:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L256:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L78
	} else {
		goto L257
	}
L257:
	;
	F_errmsg(m, int32(_a_F_parse_format_12), int32(0))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L78
	} else {
		goto L258
	}
L258:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1339), int32(_a_F_parse_format_6))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L78
	} else {
		goto L259
	}
L259:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L260:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L78
	} else {
		goto L261
	}
L261:
	;
	F_errmsg(m, int32(_a_F_parse_format_20), int32(0))
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L78
	} else {
		goto L262
	}
L262:
	;
	F_errdetail(m, int32(_a_F_parse_format_21), int32(0))
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L78
	} else {
		goto L263
	}
L263:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1354), int32(_a_F_parse_format_6))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L78
	} else {
		goto L264
	}
L264:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L265:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L78
	} else {
		goto L266
	}
L266:
	;
	F_errmsg(m, int32(_a_F_parse_format_22), int32(0))
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L78
	} else {
		goto L267
	}
L267:
	;
	F_errdetail(m, int32(_a_F_parse_format_23), int32(0))
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L78
	} else {
		goto L268
	}
L268:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1364), int32(_a_F_parse_format_6))
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L78
	} else {
		goto L269
	}
L269:
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
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
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
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
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
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
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v130 int64
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v218 int32
	_ = v218
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v301 int32
	_ = v301
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v417 int32
	_ = v417
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v500 int32
	_ = v500
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v614 int32
	_ = v614
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v697 int32
	_ = v697
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v737 int32
	_ = v737
	var v746 int32
	_ = v746
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
	v27 = F_strcspn(m, v26, int32(_a_F_parse_scram_secret_0))
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
	return v746
L11:
	;
	v737 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v737
	v746 = v737
	goto L10
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v43 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v55 == int32(0) {
		goto L11
	} else {
		goto L20
	}
L14:
	;
	v44 = F_strcspn(m, v43, int32(_a_F_parse_scram_secret_1))
	mBase = m.M
	v45 = v44 + v43
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v46 != 0 {
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
	v47 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v45))) = uint8(v47)
	v52 = v45 + int32(1)
	goto L19
L18:
	;
	v52 = int32(0)
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v52
	goto L16
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v60 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v72 == int32(0) {
		goto L11
	} else {
		goto L28
	}
L22:
	;
	v61 = F_strcspn(m, v60, int32(_a_F_parse_scram_secret_0))
	mBase = m.M
	v62 = v61 + v60
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v63 != 0 {
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
	v64 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v64)
	v69 = v62 + int32(1)
	goto L27
L26:
	;
	v69 = int32(0)
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v69
	goto L24
L28:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v77 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v89 == int32(0) {
		goto L11
	} else {
		goto L36
	}
L30:
	;
	v78 = F_strcspn(m, v77, int32(_a_F_parse_scram_secret_1))
	mBase = m.M
	v79 = v78 + v77
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v80 != 0 {
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
	v81 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v79))) = uint8(v81)
	v86 = v79 + int32(1)
	goto L35
L34:
	;
	v86 = int32(0)
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v86
	goto L32
L36:
	;
	v92 = int32(_a_F_parse_scram_secret_2)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_scram_secret[0])))
	if base.B2i32(v95 == int32(0))|base.B2i32(v95 != v98) != 0 {
		v116 = v95
		v117 = v98
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v116-v117 != 0 {
		goto L11
	} else {
		goto L44
	}
L38:
	;
	goto L37
L39:
	;
	v101 = v26
	v102 = v92
	goto L40
L40:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+1)))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	if v106 == int32(0) {
		v116 = v106
		v117 = v105
		goto L38
	} else {
		goto L42
	}
L41:
	;
	v116 = v106
	v117 = v105
	goto L38
L42:
	;
	v109 = int32(1)
	if v106 == v105 {
		v101 = v101 + v109
		v102 = v102 + v109
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(32)
	*(*int32)(unsafe.Add(mBase, _c_F_parse_scram_secret[1])) = int32(0)
	v130 = F_strtox_2(m, v43, v15+int32(8), int32(10), int64(2147483648))
	mBase = m.M
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = base.I32_wrap_i64(v130)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	if v134 != 0 {
		goto L11
	} else {
		goto L46
	}
L46:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_parse_scram_secret[1]))
	if v136 != 0 {
		goto L11
	} else {
		goto L47
	}
L47:
	;
	v137 = F_strlen(m, v60)
	mBase = m.M
	v141 = v137 * int32(3) >> (uint(int32(2)) % 32)
	goto L48
L48:
	;
	v142 = F_palloc(m, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v144 = F_strlen(m, v60)
	mBase = m.M
	v145 = int32(0)
	if v145 < v144 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	if v330 < int32(0) {
		goto L11
	} else {
		goto L96
	}
L51:
	;
	if v141 != 0 {
		goto L93
	} else {
		goto L94
	}
L52:
	;
	v154 = v60 + v144
	v155 = v60
	v159 = v145
	v160 = v142
	v162 = v145
	v163 = v145
	goto L55
L53:
	;
	v301 = v142
	goto L54
L54:
	;
	v330 = v301 - v142
	goto L50
L55:
	;
	v167 = v155 + int32(1)
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	if v168 != int32(61) {
		goto L63
	} else {
		goto L64
	}
L56:
	;
	if v289 != 0 {
		goto L51
	} else {
		goto L92
	}
L57:
	;
	if base.Ui32(v287) < base.Ui32(v154) {
		v155 = v287
		v159 = v289
		v160 = v290
		v162 = v292
		v163 = v293
		goto L55
	} else {
		goto L91
	}
L58:
	;
	if v141 < v160-v142+int32(1) {
		goto L51
	} else {
		goto L78
	}
L59:
	;
	v245 = v167
	v246 = int32(2)
	v249 = v163 << (uint(int32(6)) % 32)
	goto L58
L60:
	;
	v234 = v231 + v230<<(uint(int32(6))%32)
	v236 = v227 + int32(1)
	if v236 == int32(4) {
		v245 = v228
		v246 = v229
		v249 = v234
		goto L58
	} else {
		goto L77
	}
L61:
	;
	v227 = int32(3)
	v228 = v155 + int32(2)
	v229 = int32(1)
	v230 = v187
	v231 = v182
	goto L60
L62:
	;
	if base.Ui32(int32(125)) < base.Ui32((v206-int32(1))&int32(255)) {
		goto L51
	} else {
		goto L75
	}
L63:
	;
	v172 = v168 - int32(9)
	if base.B2i32(base.Ui32(int32(23)) < base.Ui32(v172))|base.B2i32(int32(1)<<(uint(v172)%32)&int32(_a_F_parse_scram_secret_3) == int32(0)) != 0 {
		v206 = v168
		v207 = v159
		v208 = v167
		v209 = v162
		v210 = v163
		goto L62
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v182 = int32(0)
	if v162 != 0 {
		v227 = v159
		v228 = v167
		v229 = v162
		v230 = v163
		v231 = v182
		goto L60
	} else {
		goto L67
	}
L66:
	;
	goto L51
L67:
	;
	switch v159 - int32(2) {
	case 0:
		goto L68
	case 1:
		goto L59
	default:
		goto L51
	}
L68:
	;
	if base.Ui32(v154) <= base.Ui32(v167) {
		goto L51
	} else {
		goto L69
	}
L69:
	;
	v187 = v163 << (uint(int32(6)) % 32)
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	if v188 == int32(61) {
		goto L61
	} else {
		goto L70
	}
L70:
	;
	v192 = v188 - int32(9)
	if int32(1)<<(uint(v192)%32)&int32(_a_F_parse_scram_secret_3) != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v200 = base.B2i32(base.Ui32(v192) <= base.Ui32(int32(23)))
	goto L73
L72:
	;
	v200 = int32(0)
	goto L73
L73:
	;
	if v200 != 0 {
		goto L51
	} else {
		goto L74
	}
L74:
	;
	v206 = v188
	v207 = int32(3)
	v208 = v155 + int32(2)
	v209 = int32(1)
	v210 = v187
	goto L62
L75:
	;
	v218 = int32(*(*int8)(unsafe.Add(mBase, uint32(v206)+uint32(_c_F_parse_scram_secret[2]))))
	if v218 < int32(0) {
		goto L51
	} else {
		goto L76
	}
L76:
	;
	v227 = v207
	v228 = v208
	v229 = v209
	v230 = v210
	v231 = v218
	goto L60
L77:
	;
	v287 = v228
	v289 = v236
	v290 = v160
	v292 = v229
	v293 = v234
	goto L57
L78:
	;
	v255 = int32(base.Ui32(v249) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v160))) = uint8(v255)
	v258 = v160 + int32(1)
	if base.Ui32(v246) < base.Ui32(int32(2)) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v262 = v246
	goto L81
L80:
	;
	v262 = int32(0)
	goto L81
L81:
	;
	if v262 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	if v141 < v258-v142+int32(1) {
		goto L51
	} else {
		goto L85
	}
L83:
	;
	v274 = v258
	goto L84
L84:
	;
	if v246 != 0 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	v270 = int32(base.Ui32(v249) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v160)+1)) = uint8(v270)
	v274 = v160 + int32(2)
	goto L84
L86:
	;
	v287 = v245
	v289 = int32(0)
	v290 = v284
	v292 = v285
	v293 = int32(0)
	goto L57
L87:
	;
	v284 = v274
	v285 = v246
	goto L86
L88:
	;
	goto L89
L89:
	;
	if v141 < v274-v142+int32(1) {
		goto L51
	} else {
		goto L90
	}
L90:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v274))) = uint8(v249)
	v284 = v274 + int32(1)
	v285 = int32(0)
	goto L86
L91:
	;
	goto L56
L92:
	;
	v301 = v290
	goto L54
L93:
	;
	base.MemoryFill(m, v142, int32(0), v141)
	goto L95
L94:
	;
	goto L95
L95:
	;
	v330 = int32(-1)
	goto L50
L96:
	;
	v333 = F_pstrdup(m, v60)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v333
	v336 = F_strlen(m, v77)
	mBase = m.M
	v340 = v336 * int32(3) >> (uint(int32(2)) % 32)
	goto L98
L98:
	;
	v341 = F_palloc(m, v340)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v343 = F_strlen(m, v77)
	mBase = m.M
	v344 = int32(0)
	if v344 < v343 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v529 != v530 {
		goto L11
	} else {
		goto L146
	}
L101:
	;
	if v340 != 0 {
		goto L143
	} else {
		goto L144
	}
L102:
	;
	v353 = v77 + v343
	v354 = v77
	v358 = v344
	v359 = v341
	v361 = v344
	v362 = v344
	goto L105
L103:
	;
	v500 = v341
	goto L104
L104:
	;
	v529 = v500 - v341
	goto L100
L105:
	;
	v366 = v354 + int32(1)
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354))))
	if v367 != int32(61) {
		goto L113
	} else {
		goto L114
	}
L106:
	;
	if v488 != 0 {
		goto L101
	} else {
		goto L142
	}
L107:
	;
	if base.Ui32(v486) < base.Ui32(v353) {
		v354 = v486
		v358 = v488
		v359 = v489
		v361 = v491
		v362 = v492
		goto L105
	} else {
		goto L141
	}
L108:
	;
	if v340 < v359-v341+int32(1) {
		goto L101
	} else {
		goto L128
	}
L109:
	;
	v444 = v366
	v445 = int32(2)
	v448 = v362 << (uint(int32(6)) % 32)
	goto L108
L110:
	;
	v433 = v430 + v429<<(uint(int32(6))%32)
	v435 = v426 + int32(1)
	if v435 == int32(4) {
		v444 = v427
		v445 = v428
		v448 = v433
		goto L108
	} else {
		goto L127
	}
L111:
	;
	v426 = int32(3)
	v427 = v354 + int32(2)
	v428 = int32(1)
	v429 = v386
	v430 = v381
	goto L110
L112:
	;
	if base.Ui32(int32(125)) < base.Ui32((v405-int32(1))&int32(255)) {
		goto L101
	} else {
		goto L125
	}
L113:
	;
	v371 = v367 - int32(9)
	if base.B2i32(base.Ui32(int32(23)) < base.Ui32(v371))|base.B2i32(int32(1)<<(uint(v371)%32)&int32(_a_F_parse_scram_secret_3) == int32(0)) != 0 {
		v405 = v367
		v406 = v358
		v407 = v366
		v408 = v361
		v409 = v362
		goto L112
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v381 = int32(0)
	if v361 != 0 {
		v426 = v358
		v427 = v366
		v428 = v361
		v429 = v362
		v430 = v381
		goto L110
	} else {
		goto L117
	}
L116:
	;
	goto L101
L117:
	;
	switch v358 - int32(2) {
	case 0:
		goto L118
	case 1:
		goto L109
	default:
		goto L101
	}
L118:
	;
	if base.Ui32(v353) <= base.Ui32(v366) {
		goto L101
	} else {
		goto L119
	}
L119:
	;
	v386 = v362 << (uint(int32(6)) % 32)
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366))))
	if v387 == int32(61) {
		goto L111
	} else {
		goto L120
	}
L120:
	;
	v391 = v387 - int32(9)
	if int32(1)<<(uint(v391)%32)&int32(_a_F_parse_scram_secret_3) != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v399 = base.B2i32(base.Ui32(v391) <= base.Ui32(int32(23)))
	goto L123
L122:
	;
	v399 = int32(0)
	goto L123
L123:
	;
	if v399 != 0 {
		goto L101
	} else {
		goto L124
	}
L124:
	;
	v405 = v387
	v406 = int32(3)
	v407 = v354 + int32(2)
	v408 = int32(1)
	v409 = v386
	goto L112
L125:
	;
	v417 = int32(*(*int8)(unsafe.Add(mBase, uint32(v405)+uint32(_c_F_parse_scram_secret[2]))))
	if v417 < int32(0) {
		goto L101
	} else {
		goto L126
	}
L126:
	;
	v426 = v406
	v427 = v407
	v428 = v408
	v429 = v409
	v430 = v417
	goto L110
L127:
	;
	v486 = v427
	v488 = v435
	v489 = v359
	v491 = v428
	v492 = v433
	goto L107
L128:
	;
	v454 = int32(base.Ui32(v448) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v359))) = uint8(v454)
	v457 = v359 + int32(1)
	if base.Ui32(v445) < base.Ui32(int32(2)) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v461 = v445
	goto L131
L130:
	;
	v461 = int32(0)
	goto L131
L131:
	;
	if v461 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	if v340 < v457-v341+int32(1) {
		goto L101
	} else {
		goto L135
	}
L133:
	;
	v473 = v457
	goto L134
L134:
	;
	if v445 != 0 {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	v469 = int32(base.Ui32(v448) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v359)+1)) = uint8(v469)
	v473 = v359 + int32(2)
	goto L134
L136:
	;
	v486 = v444
	v488 = int32(0)
	v489 = v483
	v491 = v484
	v492 = int32(0)
	goto L107
L137:
	;
	v483 = v473
	v484 = v445
	goto L136
L138:
	;
	goto L139
L139:
	;
	if v340 < v473-v341+int32(1) {
		goto L101
	} else {
		goto L140
	}
L140:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v473))) = uint8(v448)
	v483 = v473 + int32(1)
	v484 = int32(0)
	goto L136
L141:
	;
	goto L106
L142:
	;
	v500 = v489
	goto L104
L143:
	;
	base.MemoryFill(m, v341, int32(0), v340)
	goto L145
L144:
	;
	goto L145
L145:
	;
	v529 = int32(-1)
	goto L100
L146:
	;
	if v529 != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	base.MemoryCopy(m, l5, v341, v529)
	goto L149
L148:
	;
	goto L149
L149:
	;
	v533 = F_strlen(m, v89)
	mBase = m.M
	v537 = v533 * int32(3) >> (uint(int32(2)) % 32)
	goto L150
L150:
	;
	v538 = F_palloc(m, v537)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	v540 = F_strlen(m, v89)
	mBase = m.M
	v541 = int32(0)
	if v541 < v540 {
		goto L154
	} else {
		goto L155
	}
L152:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v726 != v727 {
		goto L11
	} else {
		goto L198
	}
L153:
	;
	if v537 != 0 {
		goto L195
	} else {
		goto L196
	}
L154:
	;
	v550 = v89 + v540
	v551 = v89
	v555 = v541
	v556 = v538
	v558 = v541
	v559 = v541
	goto L157
L155:
	;
	v697 = v538
	goto L156
L156:
	;
	v726 = v697 - v538
	goto L152
L157:
	;
	v563 = v551 + int32(1)
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551))))
	if v564 != int32(61) {
		goto L165
	} else {
		goto L166
	}
L158:
	;
	if v685 != 0 {
		goto L153
	} else {
		goto L194
	}
L159:
	;
	if base.Ui32(v683) < base.Ui32(v550) {
		v551 = v683
		v555 = v685
		v556 = v686
		v558 = v688
		v559 = v689
		goto L157
	} else {
		goto L193
	}
L160:
	;
	if v537 < v556-v538+int32(1) {
		goto L153
	} else {
		goto L180
	}
L161:
	;
	v641 = v563
	v642 = int32(2)
	v645 = v559 << (uint(int32(6)) % 32)
	goto L160
L162:
	;
	v630 = v627 + v626<<(uint(int32(6))%32)
	v632 = v623 + int32(1)
	if v632 == int32(4) {
		v641 = v624
		v642 = v625
		v645 = v630
		goto L160
	} else {
		goto L179
	}
L163:
	;
	v623 = int32(3)
	v624 = v551 + int32(2)
	v625 = int32(1)
	v626 = v583
	v627 = v578
	goto L162
L164:
	;
	if base.Ui32(int32(125)) < base.Ui32((v602-int32(1))&int32(255)) {
		goto L153
	} else {
		goto L177
	}
L165:
	;
	v568 = v564 - int32(9)
	if base.B2i32(base.Ui32(int32(23)) < base.Ui32(v568))|base.B2i32(int32(1)<<(uint(v568)%32)&int32(_a_F_parse_scram_secret_3) == int32(0)) != 0 {
		v602 = v564
		v603 = v555
		v604 = v563
		v605 = v558
		v606 = v559
		goto L164
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	v578 = int32(0)
	if v558 != 0 {
		v623 = v555
		v624 = v563
		v625 = v558
		v626 = v559
		v627 = v578
		goto L162
	} else {
		goto L169
	}
L168:
	;
	goto L153
L169:
	;
	switch v555 - int32(2) {
	case 0:
		goto L170
	case 1:
		goto L161
	default:
		goto L153
	}
L170:
	;
	if base.Ui32(v550) <= base.Ui32(v563) {
		goto L153
	} else {
		goto L171
	}
L171:
	;
	v583 = v559 << (uint(int32(6)) % 32)
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v563))))
	if v584 == int32(61) {
		goto L163
	} else {
		goto L172
	}
L172:
	;
	v588 = v584 - int32(9)
	if int32(1)<<(uint(v588)%32)&int32(_a_F_parse_scram_secret_3) != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v596 = base.B2i32(base.Ui32(v588) <= base.Ui32(int32(23)))
	goto L175
L174:
	;
	v596 = int32(0)
	goto L175
L175:
	;
	if v596 != 0 {
		goto L153
	} else {
		goto L176
	}
L176:
	;
	v602 = v584
	v603 = int32(3)
	v604 = v551 + int32(2)
	v605 = int32(1)
	v606 = v583
	goto L164
L177:
	;
	v614 = int32(*(*int8)(unsafe.Add(mBase, uint32(v602)+uint32(_c_F_parse_scram_secret[2]))))
	if v614 < int32(0) {
		goto L153
	} else {
		goto L178
	}
L178:
	;
	v623 = v603
	v624 = v604
	v625 = v605
	v626 = v606
	v627 = v614
	goto L162
L179:
	;
	v683 = v624
	v685 = v632
	v686 = v556
	v688 = v625
	v689 = v630
	goto L159
L180:
	;
	v651 = int32(base.Ui32(v645) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v556))) = uint8(v651)
	v654 = v556 + int32(1)
	if base.Ui32(v642) < base.Ui32(int32(2)) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v658 = v642
	goto L183
L182:
	;
	v658 = int32(0)
	goto L183
L183:
	;
	if v658 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	if v537 < v654-v538+int32(1) {
		goto L153
	} else {
		goto L187
	}
L185:
	;
	v670 = v654
	goto L186
L186:
	;
	if v642 != 0 {
		goto L189
	} else {
		goto L190
	}
L187:
	;
	v666 = int32(base.Ui32(v645) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v556)+1)) = uint8(v666)
	v670 = v556 + int32(2)
	goto L186
L188:
	;
	v683 = v641
	v685 = int32(0)
	v686 = v680
	v688 = v681
	v689 = int32(0)
	goto L159
L189:
	;
	v680 = v670
	v681 = v642
	goto L188
L190:
	;
	goto L191
L191:
	;
	if v537 < v670-v538+int32(1) {
		goto L153
	} else {
		goto L192
	}
L192:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v670))) = uint8(v645)
	v680 = v670 + int32(1)
	v681 = int32(0)
	goto L188
L193:
	;
	goto L158
L194:
	;
	v697 = v686
	goto L156
L195:
	;
	base.MemoryFill(m, v538, int32(0), v537)
	goto L197
L196:
	;
	goto L197
L197:
	;
	v726 = int32(-1)
	goto L152
L198:
	;
	if v726 != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	base.MemoryCopy(m, l6, v538, v726)
	goto L201
L200:
	;
	goto L201
L201:
	;
	v746 = int32(1)
	goto L10
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
	var v43 int32
	_ = v43
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
	if v36 == int32(0) {
		v49 = v33
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v12<<(uint(int32(2))%32)+v36)))
	if v43 == v45 {
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
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	if v5 != v6 {
		if v5 < v6 {
			v11 = int32(-1)
		} else {
			v11 = int32(1)
		}
		return v11
	} else {
		v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		v14 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
		if v13 != v14 {
			if v13 < v14 {
				v19 = int32(-1)
			} else {
				v19 = int32(1)
			}
			v21 = v19
		} else {
			v21 = int32(0)
		}
		return v21
	}
}
func F_pglz_decompress(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v189 int32
	_ = v189
	v6 = int32(0)
	v13 = l2 + l3
	v14 = l0 + l1
	if base.B2i32(l1 <= v6)|base.B2i32(l3 <= v6) == v6 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v189
L2:
	;
	if l4 != 0 {
		goto L42
	} else {
		goto L43
	}
L3:
	;
	v22 = l0
	v25 = l2
	goto L6
L4:
	;
	goto L5
L5:
	;
	v164 = l0
	v167 = l2
	goto L2
L6:
	;
	v35 = v22 + int32(1)
	if base.Ui32(v14) <= base.Ui32(v35) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v164 = v150
	v167 = v153
	goto L2
L8:
	;
	if base.Ui32(v14) <= base.Ui32(v150) {
		v164 = v150
		v167 = v153
		goto L2
	} else {
		goto L40
	}
L9:
	;
	v150 = v35
	v153 = v25
	goto L8
L10:
	;
	goto L11
L11:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v40 = v35
	v42 = v25
	v48 = v37
	v49 = int32(0)
	goto L12
L12:
	;
	if v48&int32(1) != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v150 = v126
	v153 = v138
	goto L8
L14:
	;
	if base.B2i32(base.Ui32(int32(6)) < base.Ui32(v49))|base.B2i32(base.Ui32(v14) <= base.Ui32(v126)) != 0 {
		v150 = v126
		v153 = v138
		goto L8
	} else {
		goto L38
	}
L15:
	;
	v53 = int32(-1)
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	v58 = v54&int32(15) + int32(3)
	if v58 != int32(18) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	*(*uint8)(unsafe.Add(mBase, uint32(v42))) = uint8(v120)
	v122 = int32(1)
	v126 = v40 + v122
	v138 = v42 + v122
	goto L14
L18:
	;
	v68 = v58
	v69 = v40 + int32(2)
	goto L20
L19:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+2)))
	v68 = v63 + int32(18)
	v69 = v40 + int32(3)
	goto L20
L20:
	;
	if base.Ui32(v14) < base.Ui32(v69) {
		v189 = v53
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)))
	v76 = v71 | v54<<(uint(int32(4))%32)&int32(3840)
	if base.B2i32(v76 == int32(0))|base.B2i32(v42-l2 < v76) != 0 {
		v189 = v53
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v82 = v13 - v42
	if v68 < v82 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v84 = v68
	goto L25
L24:
	;
	v84 = v82
	goto L25
L25:
	;
	if v76 < v84 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v87 = v76
	v89 = v42
	v91 = v84
	goto L29
L27:
	;
	v106 = v76
	v108 = v42
	v110 = v84
	goto L28
L28:
	;
	if v110 != 0 {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	if v87 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v106 = v103
	v108 = v100
	v110 = v101
	goto L28
L31:
	;
	base.MemoryCopy(m, v89, v89-v87, v87)
	goto L33
L32:
	;
	goto L33
L33:
	;
	v100 = v87 + v89
	v101 = v91 - v87
	v103 = v87 << (uint(int32(1)) % 32)
	if v103 < v101 {
		v87 = v103
		v89 = v100
		v91 = v101
		goto L29
	} else {
		goto L34
	}
L34:
	;
	goto L30
L35:
	;
	base.MemoryCopy(m, v108, v108-v106, v110)
	goto L37
L36:
	;
	goto L37
L37:
	;
	v126 = v69
	v138 = v108 + v110
	goto L14
L38:
	;
	v143 = int32(1)
	if base.Ui32(v138) < base.Ui32(v13) {
		v40 = v126
		v42 = v138
		v48 = int32(base.Ui32(v48&int32(254)) >> (uint(v143) % 32))
		v49 = v49 + v143
		goto L12
	} else {
		goto L39
	}
L39:
	;
	goto L13
L40:
	;
	if base.Ui32(v153) < base.Ui32(v13) {
		v22 = v150
		v25 = v153
		goto L6
	} else {
		goto L41
	}
L41:
	;
	goto L7
L42:
	;
	if base.B2i32(v164 != v14)|base.B2i32(v167 != v13) != 0 {
		v189 = int32(-1)
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v189 = v167 - l2
	goto L1
L45:
	;
	goto L44
}
func F_pgstatginindex_internal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
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
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+30)) = uint8(v3)
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+28)) = uint16(v3)
	v17 = F_relation_open(m, l0, int32(1))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+119)))
		if v22 != int32(105) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v99 = m.ExcPending
			if v99 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(151027844))
				mBase = m.M
				v102 = m.ExcPending
				if v102 != 0 {
					return int32(0)
				} else {
					v103 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v103 + int32(4)
					F_errmsg(m, int32(_a_F_pgstatginindex_internal_0), v10+int32(16))
					mBase = m.M
					v111 = m.ExcPending
					if v111 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pgstatginindex_internal_1), int32(523), int32(_a_F_pgstatginindex_internal_2))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+84))
			if v25 != int32(2742) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v99 = m.ExcPending
				if v99 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(151027844))
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return int32(0)
					} else {
						v103 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v103 + int32(4)
						F_errmsg(m, int32(_a_F_pgstatginindex_internal_0), v10+int32(16))
						mBase = m.M
						v111 = m.ExcPending
						if v111 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pgstatginindex_internal_1), int32(523), int32(_a_F_pgstatginindex_internal_2))
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+118)))
				if v28 == int32(116) {
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+24)))
					if v31 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v120 = m.ExcPending
						if v120 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v123 = m.ExcPending
							if v123 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_pgstatginindex_internal_3), int32(0))
								mBase = m.M
								v127 = m.ExcPending
								if v127 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_pgstatginindex_internal_1), int32(533), int32(_a_F_pgstatginindex_internal_2))
									mBase = m.M
									v132 = m.ExcPending
									if v132 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v17)+192))
						v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+18)))
						if v35 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v136 = m.ExcPending
							if v136 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(325))
								mBase = m.M
								v139 = m.ExcPending
								if v139 != 0 {
									return int32(0)
								} else {
									v140 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = v140 + int32(4)
									F_errmsg(m, int32(_a_F_pgstatginindex_internal_4), v10)
									mBase = m.M
									v146 = m.ExcPending
									if v146 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_pgstatginindex_internal_1), int32(540), int32(_a_F_pgstatginindex_internal_2))
										mBase = m.M
										v151 = m.ExcPending
										if v151 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v39 = F_ReadBuffer(m, v17, int32(0))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								F_LockBuffer(m, v39, int32(1))
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									if v39 < int32(0) {
										v47 = *(*int32)(unsafe.Add(mBase, _c_F_pgstatginindex_internal[0]))
										v53 = *(*int32)(unsafe.Add(mBase, uint32(v47+(v39^int32(-1))<<(uint(int32(2))%32))))
										v61 = v53
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, _c_F_pgstatginindex_internal[1]))
										v61 = v55 + v39<<(uint(int32(13))%32) + int32(-8192)
									}
									v62 = *(*int64)(unsafe.Add(mBase, uint32(v61)+40))
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)+36))
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+72))
									F_UnlockReleaseBuffer(m, v39)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return int32(0)
									} else {
										F_relation_close(m, v17, int32(1))
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											v73 = F_get_call_result_type(m, l1, int32(0), v10+int32(44))
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return int32(0)
											} else {
												if v73 != int32(1) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v155 = m.ExcPending
													if v155 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(_a_F_pgstatginindex_internal_5), int32(0))
														mBase = m.M
														v159 = m.ExcPending
														if v159 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_pgstatginindex_internal_1), int32(561), int32(_a_F_pgstatginindex_internal_2))
															mBase = m.M
															v164 = m.ExcPending
															if v164 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v63
													*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v64
													v79 = F_Int64GetDatum(m, v62)
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v79
														v82 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
														v87 = F_heap_form_tuple(m, v82, v10+int32(32), v10+int32(28))
														mBase = m.M
														v88 = m.ExcPending
														if v88 != 0 {
															return int32(0)
														} else {
															v89 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
															v90 = F_HeapTupleHeaderGetDatum(m, v89)
															mBase = m.M
															v91 = m.ExcPending
															if v91 != 0 {
																return int32(0)
															} else {
																m.G0 = v10 + int32(48)
																return v90
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
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v17)+192))
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+18)))
					if v35 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v136 = m.ExcPending
						if v136 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(325))
							mBase = m.M
							v139 = m.ExcPending
							if v139 != 0 {
								return int32(0)
							} else {
								v140 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v140 + int32(4)
								F_errmsg(m, int32(_a_F_pgstatginindex_internal_4), v10)
								mBase = m.M
								v146 = m.ExcPending
								if v146 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_pgstatginindex_internal_1), int32(540), int32(_a_F_pgstatginindex_internal_2))
									mBase = m.M
									v151 = m.ExcPending
									if v151 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v39 = F_ReadBuffer(m, v17, int32(0))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							F_LockBuffer(m, v39, int32(1))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								if v39 < int32(0) {
									v47 = *(*int32)(unsafe.Add(mBase, _c_F_pgstatginindex_internal[0]))
									v53 = *(*int32)(unsafe.Add(mBase, uint32(v47+(v39^int32(-1))<<(uint(int32(2))%32))))
									v61 = v53
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, _c_F_pgstatginindex_internal[1]))
									v61 = v55 + v39<<(uint(int32(13))%32) + int32(-8192)
								}
								v62 = *(*int64)(unsafe.Add(mBase, uint32(v61)+40))
								v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)+36))
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+72))
								F_UnlockReleaseBuffer(m, v39)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									F_relation_close(m, v17, int32(1))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										v73 = F_get_call_result_type(m, l1, int32(0), v10+int32(44))
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return int32(0)
										} else {
											if v73 != int32(1) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v155 = m.ExcPending
												if v155 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(_a_F_pgstatginindex_internal_5), int32(0))
													mBase = m.M
													v159 = m.ExcPending
													if v159 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_pgstatginindex_internal_1), int32(561), int32(_a_F_pgstatginindex_internal_2))
														mBase = m.M
														v164 = m.ExcPending
														if v164 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v63
												*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v64
												v79 = F_Int64GetDatum(m, v62)
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v79
													v82 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
													v87 = F_heap_form_tuple(m, v82, v10+int32(32), v10+int32(28))
													mBase = m.M
													v88 = m.ExcPending
													if v88 != 0 {
														return int32(0)
													} else {
														v89 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
														v90 = F_HeapTupleHeaderGetDatum(m, v89)
														mBase = m.M
														v91 = m.ExcPending
														if v91 != 0 {
															return int32(0)
														} else {
															m.G0 = v10 + int32(48)
															return v90
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
}
func F_pgstatindex_v1_5(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = F_textToQualifiedNameList(m, v3)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			v9 = F_makeRangeVarFromNameList(m, v7)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return int32(0)
			} else {
				v12 = F_relation_openrv(m, v9, int32(1))
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return int32(0)
				} else {
					v14 = F_pgstatindex_impl(m, v12, l0)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return int32(0)
					} else {
						return v14
					}
				}
			}
		}
	}
}
func F_pgstatindexbyid_v1_5(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_relation_open(m, v2, int32(1))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_pgstatindex_impl(m, v4, l0)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_pgstattuplebyid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_superuser(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_pgstattuplebyid_0), int32(0))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pgstattuplebyid_1), int32(218), int32(_a_F_pgstattuplebyid_2))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v27 = F_relation_open(m, v3, int32(1))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v29 = F_pgstat_relation(m, v27, l0)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					return v29
				}
			}
		}
	}
}
func F_phraseto_tsquery_byid(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14006(m, l0, int32(1), int32(4))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
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
	F_appendStringInfo(m, l0, int32(_a_F_plan_elem_desc_0), v10)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return
	} else {
		F_appendStringInfoString(m, l0, int32(_a_F_plan_elem_desc_1))
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
				F_appendStringInfoString(m, l0, int32(_a_F_plan_elem_desc_2))
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
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v225 int32
	_ = v225
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v282 int32
	_ = v282
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v339 int32
	_ = v339
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v467 int32
	_ = v467
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v623 int32
	_ = v623
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v726 int32
	_ = v726
	var v737 int32
	_ = v737
	var v743 int32
	_ = v743
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v769 int32
	_ = v769
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
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
	var v905 int32
	_ = v905
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v916 int32
	_ = v916
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v977 int32
	_ = v977
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1044 int32
	_ = v1044
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1060 int32
	_ = v1060
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1092 int32
	_ = v1092
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1156 int32
	_ = v1156
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1194 int32
	_ = v1194
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1203 int32
	_ = v1203
	var v1207 int32
	_ = v1207
	var v1220 int32
	_ = v1220
	var v1223 int32
	_ = v1223
	var v1226 int32
	_ = v1226
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1263 int32
	_ = v1263
	var v1265 int32
	_ = v1265
	var v1271 int32
	_ = v1271
	var v1273 int32
	_ = v1273
	var v1278 int32
	_ = v1278
	var v1281 int32
	_ = v1281
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1298 int32
	_ = v1298
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v9 == v11 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v1298
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
	v36 = v9
	goto L8
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v9))))
	if v15 != int32(121) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v18 = int32(1)
	v19 = v9 + v18
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v19
	v24 = F_slice_from_s(m, l0, v18, int32(_a_F_porter_ISO_8859_1_stem_0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	if v24 < int32(0) {
		v1298 = v24
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = int32(1)
	goto L2
L8:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v49 < v48 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v123)+4)) = v93
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v123))) = v125
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v136 < v127 {
		goto L35
	} else {
		goto L36
	}
L10:
	;
	goto L9
L11:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v92 != 0 {
		goto L26
	} else {
		goto L27
	}
L12:
	;
	v51 = v48
	goto L14
L13:
	;
	v51 = v49
	goto L14
L14:
	;
	goto L16
L15:
	;
	v92 = v87
	goto L11
L16:
	;
	if v48 == v51 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v87 = int32(0)
	goto L15
L18:
	;
	v92 = int32(-1)
	goto L11
L19:
	;
	goto L20
L20:
	;
	v63 = int32(1)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+v48))))
	if int32(121) < v66 {
		v87 = v63
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v68 = v66 - int32(97)
	if v68 < int32(0) {
		v87 = v63
		goto L15
	} else {
		goto L22
	}
L22:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v68)>>(uint(int32(3))%32)))+uint32(_c_F_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v74)>>(uint(v68&int32(7))%32))&int32(1) == int32(0) {
		v87 = v63
		goto L15
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v48 + int32(1)
	goto L24
L24:
	;
	goto L17
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v36
	v109 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v94 + v109
	v114 = F_slice_from_s(m, l0, v109, int32(_a_F_porter_ISO_8859_1_stem_1))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L5
	} else {
		goto L31
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v36
	if v93 <= v36 {
		goto L10
	} else {
		goto L30
	}
L27:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v94
	if v93 == v94 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97+v94))))
	if v99 == int32(121) {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	goto L26
L30:
	;
	v106 = v36 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v106
	v36 = v106
	goto L8
L31:
	;
	if v114 < int32(0) {
		v1298 = v114
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v118)+8)) = int32(1)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v36 = v121
	goto L8
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v127
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v357
	if v357 <= v127 {
		goto L96
	} else {
		goto L97
	}
L34:
	;
	if v176 < int32(0) {
		goto L33
	} else {
		goto L49
	}
L35:
	;
	v138 = v127
	goto L37
L36:
	;
	v138 = v136
	goto L37
L37:
	;
	v145 = v127
	goto L39
L38:
	;
	v176 = v156
	goto L34
L39:
	;
	if v145 == v138 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v176 = int32(-1)
	goto L34
L42:
	;
	goto L43
L43:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149+v145))))
	if int32(121) < v151 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v168 = v145 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v168
	v145 = v168
	goto L39
L45:
	;
	v153 = v151 - int32(97)
	if v153 < int32(0) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v156 = int32(1)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v153)>>(uint(int32(3))%32)))+uint32(_c_F_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v160)>>(uint(v153&int32(7))%32))&v156 != 0 {
		goto L38
	} else {
		goto L47
	}
L47:
	;
	goto L44
L49:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v180 = v179 + v176
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v180
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v191 < v180 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if v234 < int32(0) {
		goto L33
	} else {
		goto L64
	}
L51:
	;
	v193 = v180
	goto L53
L52:
	;
	v193 = v191
	goto L53
L53:
	;
	v199 = v180
	goto L55
L54:
	;
	v234 = int32(1)
	goto L50
L55:
	;
	if v199 == v193 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v234 = int32(-1)
	goto L50
L58:
	;
	goto L59
L59:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206+v199))))
	if int32(121) < v208 {
		goto L54
	} else {
		goto L60
	}
L60:
	;
	v210 = v208 - int32(97)
	if v210 < int32(0) {
		goto L54
	} else {
		goto L61
	}
L61:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v210)>>(uint(int32(3))%32)))+uint32(_c_F_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v216)>>(uint(v210&int32(7))%32))&int32(1) == int32(0) {
		goto L54
	} else {
		goto L62
	}
L62:
	;
	v225 = v199 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v225
	v199 = v225
	goto L55
L64:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v238 = v237 + v234
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v238
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v240)+4)) = v238
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v250 < v249 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	if v290 < int32(0) {
		goto L33
	} else {
		goto L80
	}
L66:
	;
	v252 = v249
	goto L68
L67:
	;
	v252 = v250
	goto L68
L68:
	;
	v259 = v249
	goto L70
L69:
	;
	v290 = v270
	goto L65
L70:
	;
	if v259 == v252 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v290 = int32(-1)
	goto L65
L73:
	;
	goto L74
L74:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263+v259))))
	if int32(121) < v265 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v282 = v259 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v282
	v259 = v282
	goto L70
L76:
	;
	v267 = v265 - int32(97)
	if v267 < int32(0) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v270 = int32(1)
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v267)>>(uint(int32(3))%32)))+uint32(_c_F_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v274)>>(uint(v267&int32(7))%32))&v270 != 0 {
		goto L69
	} else {
		goto L78
	}
L78:
	;
	goto L75
L80:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v294 = v293 + v290
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v294
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v305 < v294 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	if v348 < int32(0) {
		goto L33
	} else {
		goto L95
	}
L82:
	;
	v307 = v294
	goto L84
L83:
	;
	v307 = v305
	goto L84
L84:
	;
	v313 = v294
	goto L86
L85:
	;
	v348 = int32(1)
	goto L81
L86:
	;
	if v313 == v307 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v348 = int32(-1)
	goto L81
L89:
	;
	goto L90
L90:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320+v313))))
	if int32(121) < v322 {
		goto L85
	} else {
		goto L91
	}
L91:
	;
	v324 = v322 - int32(97)
	if v324 < int32(0) {
		goto L85
	} else {
		goto L92
	}
L92:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v324)>>(uint(int32(3))%32)))+uint32(_c_F_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v330)>>(uint(v324&int32(7))%32))&int32(1) == int32(0) {
		goto L85
	} else {
		goto L93
	}
L93:
	;
	v339 = v313 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v339
	v313 = v339
	goto L86
L95:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v351))) = v352 + v348
	goto L33
L96:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v395
	v399 = v395 - int32(1)
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v399 <= v400 {
		goto L110
	} else {
		goto L111
	}
L97:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361+v357-int32(1)))))
	if v365 != int32(115) {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v370 = F_find_among_b(m, l0, int32(_a_F_porter_ISO_8859_1_stem_2), int32(4))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L5
	} else {
		goto L99
	}
L99:
	;
	if v370 == int32(0) {
		goto L96
	} else {
		goto L100
	}
L100:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v374
	switch v370 - int32(1) {
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
	v390 = F_slice_del(m, l0)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L5
	} else {
		goto L108
	}
L102:
	;
	v386 = F_slice_from_s(m, l0, int32(1), int32(_a_F_porter_ISO_8859_1_stem_3))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L5
	} else {
		goto L106
	}
L103:
	;
	v380 = F_slice_from_s(m, l0, int32(2), int32(_a_F_porter_ISO_8859_1_stem_4))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L5
	} else {
		goto L104
	}
L104:
	;
	if int32(0) <= v380 {
		goto L96
	} else {
		goto L105
	}
L105:
	;
	v1298 = v380
	goto L1
L106:
	;
	if int32(0) <= v386 {
		goto L96
	} else {
		goto L107
	}
L107:
	;
	v1298 = v386
	goto L1
L108:
	;
	if v390 < int32(0) {
		v1298 = v390
		goto L1
	} else {
		goto L109
	}
L109:
	;
	goto L96
L110:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v709
	v711 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v709
	v714 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v709 <= v714 {
		v794 = v711
		goto L188
	} else {
		goto L189
	}
L111:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402+v399))))
	switch v404 - int32(100) {
	case 0, 3:
		goto L112
	default:
		goto L110
	}
L112:
	;
	v409 = F_find_among_b(m, l0, int32(_a_F_porter_ISO_8859_1_stem_5), int32(3))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L5
	} else {
		goto L113
	}
L113:
	;
	if v409 == int32(0) {
		goto L110
	} else {
		goto L114
	}
L114:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v413
	switch v409 - int32(1) {
	case 0:
		goto L116
	case 1:
		goto L115
	default:
		goto L110
	}
L115:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v441 = v434
	goto L122
L116:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v417)+4))
	if v413 < v418 {
		goto L110
	} else {
		goto L117
	}
L117:
	;
	v422 = F_slice_from_s(m, l0, int32(2), int32(_a_F_porter_ISO_8859_1_stem_6))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L5
	} else {
		goto L118
	}
L118:
	;
	if int32(0) <= v422 {
		goto L110
	} else {
		goto L119
	}
L119:
	;
	v1298 = v422
	goto L1
L120:
	;
	if v475 < int32(0) {
		goto L110
	} else {
		goto L132
	}
L121:
	;
	v475 = v455
	goto L120
L122:
	;
	if v441 <= v435 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v475 = int32(-1)
	goto L120
L125:
	;
	goto L126
L126:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446+v441-int32(1)))))
	if int32(121) < v450 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v467 = v441 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v467
	v441 = v467
	goto L122
L128:
	;
	v452 = v450 - int32(97)
	if v452 < int32(0) {
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v455 = int32(1)
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v452)>>(uint(int32(3))%32)))+uint32(_c_F_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v459)>>(uint(v452&int32(7))%32))&v455 != 0 {
		goto L121
	} else {
		goto L130
	}
L130:
	;
	goto L127
L132:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v478 + (v413 - v426)
	v482 = F_slice_del(m, l0)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L5
	} else {
		goto L133
	}
L133:
	;
	if v482 < int32(0) {
		v1298 = v482
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v488 = v486 - v487
	v490 = v486 - int32(1)
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v490 <= v491 {
		v535 = v486
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v536)+4))
	if v535 != v537 {
		goto L110
	} else {
		goto L146
	}
L136:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493+v490))))
	if base.B2i32(v495&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v495)%32)&int32(68514004) == int32(0)) != 0 {
		v535 = v486
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v509 = F_find_among_b(m, l0, int32(_a_F_porter_ISO_8859_1_stem_7), int32(13))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L5
	} else {
		goto L138
	}
L138:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v512 = v511 + v488
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v512
	switch v509 - int32(1) {
	case 0:
		goto L140
	case 1:
		goto L139
	case 2:
		v535 = v512
		goto L135
	default:
		goto L110
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v512
	v524 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v512 <= v524 {
		goto L110
	} else {
		goto L143
	}
L140:
	;
	v518 = F_insert_s(m, l0, v512, v512, int32(1), int32(_a_F_porter_ISO_8859_1_stem_8))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L5
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v512
	if int32(0) <= v518 {
		goto L110
	} else {
		goto L142
	}
L142:
	;
	v1298 = v518
	goto L1
L143:
	;
	v527 = v512 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v527
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v527
	v530 = F_slice_del(m, l0)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L5
	} else {
		goto L144
	}
L144:
	;
	if int32(0) <= v530 {
		goto L110
	} else {
		goto L145
	}
L145:
	;
	v1298 = v530
	goto L1
L146:
	;
	v539 = int32(0)
	v547 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L150
L147:
	;
	if v693 == int32(0) {
		goto L110
	} else {
		goto L185
	}
L148:
	;
	if v588 != 0 {
		v693 = v539
		goto L147
	} else {
		goto L160
	}
L149:
	;
	v588 = v585
	goto L148
L150:
	;
	if v547 <= v548 {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	v585 = int32(0)
	goto L149
L152:
	;
	v588 = int32(-1)
	goto L148
L153:
	;
	goto L154
L154:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v559+v547-int32(1)))))
	if int32(121) < v563 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v547 - int32(1)
	goto L159
L156:
	;
	v565 = v563 - int32(89)
	if v565 < int32(0) {
		goto L155
	} else {
		goto L157
	}
L157:
	;
	v568 = int32(1)
	v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v565)>>(uint(int32(3))%32)))+uint32(_c_F_porter_ISO_8859_1_stem[1]))))
	if int32(base.Ui32(v572)>>(uint(v565&int32(7))%32))&v568 != 0 {
		v585 = v568
		goto L149
	} else {
		goto L158
	}
L158:
	;
	goto L155
L159:
	;
	goto L151
L160:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L163
L161:
	;
	if v641 != 0 {
		v693 = v539
		goto L147
	} else {
		goto L172
	}
L162:
	;
	v641 = v637
	goto L161
L163:
	;
	if v597 <= v598 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	v637 = int32(0)
	goto L162
L165:
	;
	v641 = int32(-1)
	goto L161
L166:
	;
	goto L167
L167:
	;
	v610 = int32(1)
	v611 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611+v597-v610))))
	if int32(121) < v615 {
		v637 = v610
		goto L162
	} else {
		goto L168
	}
L168:
	;
	v617 = v615 - int32(97)
	if v617 < int32(0) {
		v637 = v610
		goto L162
	} else {
		goto L169
	}
L169:
	;
	v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v617)>>(uint(int32(3))%32)))+uint32(_c_F_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v623)>>(uint(v617&int32(7))%32))&int32(1) == int32(0) {
		v637 = v610
		goto L162
	} else {
		goto L170
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v597 - int32(1)
	goto L171
L171:
	;
	goto L164
L172:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v650 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L175
L173:
	;
	v693 = base.B2i32(v690 == int32(0))
	goto L147
L174:
	;
	v690 = v687
	goto L173
L175:
	;
	if v649 <= v650 {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	v687 = int32(0)
	goto L174
L177:
	;
	v690 = int32(-1)
	goto L173
L178:
	;
	goto L179
L179:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v661+v649-int32(1)))))
	if int32(121) < v665 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v649 - int32(1)
	goto L184
L181:
	;
	v667 = v665 - int32(97)
	if v667 < int32(0) {
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v670 = int32(1)
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v667)>>(uint(int32(3))%32)))+uint32(_c_F_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v674)>>(uint(v667&int32(7))%32))&v670 != 0 {
		v687 = v670
		goto L174
	} else {
		goto L183
	}
L183:
	;
	goto L180
L184:
	;
	goto L176
L185:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v697 = v696 + v488
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v697
	v701 = F_insert_s(m, l0, v697, v697, int32(1), int32(_a_F_porter_ISO_8859_1_stem_9))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L5
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v697
	if v701 < int32(0) {
		v1298 = v701
		goto L1
	} else {
		goto L187
	}
L187:
	;
	goto L110
L188:
	;
	if v794 < int32(0) {
		v1298 = v794
		goto L1
	} else {
		goto L208
	}
L189:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v716+v709-int32(1)))))
	if v720|int32(32) != int32(121) {
		v794 = v711
		goto L188
	} else {
		goto L190
	}
L190:
	;
	v726 = v709 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v726
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v726
	v737 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v743 = v726
	goto L193
L191:
	;
	if v777 < int32(0) {
		v794 = v711
		goto L188
	} else {
		goto L203
	}
L192:
	;
	v777 = v757
	goto L191
L193:
	;
	if v743 <= v737 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v777 = int32(-1)
	goto L191
L196:
	;
	goto L197
L197:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v748+v743-int32(1)))))
	if int32(121) < v752 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v769 = v743 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v769
	v743 = v769
	goto L193
L199:
	;
	v754 = v752 - int32(97)
	if v754 < int32(0) {
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v757 = int32(1)
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v754)>>(uint(int32(3))%32)))+uint32(_c_F_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v761)>>(uint(v754&int32(7))%32))&v757 != 0 {
		goto L192
	} else {
		goto L201
	}
L201:
	;
	goto L198
L203:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v780 - v777
	v783 = int32(1)
	v786 = F_slice_from_s(m, l0, v783, int32(_a_F_porter_ISO_8859_1_stem_10))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L5
	} else {
		goto L204
	}
L204:
	;
	if int32(0) <= v786 {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v793 = v783
	goto L207
L206:
	;
	v793 = v786 >> (uint(int32(31)) % 32) & v786
	goto L207
L207:
	;
	v794 = v793
	goto L188
L208:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v798
	v800 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v798
	v803 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v798-int32(2) <= v803 {
		v916 = v800
		goto L209
	} else {
		goto L210
	}
L209:
	;
	if v916 < int32(0) {
		v1298 = v916
		goto L1
	} else {
		goto L255
	}
L210:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v809 = int32(1)
	v811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v807+v798-v809))))
	if base.B2i32(v811&int32(224) != int32(96))|base.B2i32(v809<<(uint(v811)%32)&int32(_a_F_porter_ISO_8859_1_stem_11) == int32(0)) != 0 {
		v916 = v800
		goto L209
	} else {
		goto L211
	}
L211:
	;
	v825 = F_find_among_b(m, l0, int32(_a_F_porter_ISO_8859_1_stem_12), int32(20))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L5
	} else {
		goto L212
	}
L212:
	;
	if v825 == int32(0) {
		v916 = v800
		goto L209
	} else {
		goto L213
	}
L213:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v829
	v831 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v831)+4))
	if v829 < v832 {
		v916 = v800
		goto L209
	} else {
		goto L214
	}
L214:
	;
	switch v825 - int32(1) {
	case 0:
		goto L228
	case 1:
		goto L227
	case 2:
		goto L226
	case 3:
		goto L225
	case 4:
		goto L224
	case 5:
		goto L223
	case 6:
		goto L222
	case 7:
		goto L221
	case 8:
		goto L220
	case 9:
		goto L219
	case 10:
		goto L218
	case 11:
		goto L217
	case 12:
		goto L216
	default:
		goto L215
	}
L215:
	;
	v916 = int32(1)
	goto L209
L216:
	;
	v910 = F_slice_from_s(m, l0, int32(3), int32(_a_F_porter_ISO_8859_1_stem_13))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L5
	} else {
		goto L253
	}
L217:
	;
	v904 = F_slice_from_s(m, l0, int32(3), int32(_a_F_porter_ISO_8859_1_stem_14))
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L5
	} else {
		goto L251
	}
L218:
	;
	v898 = F_slice_from_s(m, l0, int32(3), int32(_a_F_porter_ISO_8859_1_stem_15))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L5
	} else {
		goto L249
	}
L219:
	;
	v892 = F_slice_from_s(m, l0, int32(3), int32(_a_F_porter_ISO_8859_1_stem_16))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L5
	} else {
		goto L247
	}
L220:
	;
	v886 = F_slice_from_s(m, l0, int32(2), int32(_a_F_porter_ISO_8859_1_stem_17))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L5
	} else {
		goto L245
	}
L221:
	;
	v880 = F_slice_from_s(m, l0, int32(3), int32(_a_F_porter_ISO_8859_1_stem_18))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L5
	} else {
		goto L243
	}
L222:
	;
	v874 = F_slice_from_s(m, l0, int32(3), int32(_a_F_porter_ISO_8859_1_stem_19))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L5
	} else {
		goto L241
	}
L223:
	;
	v868 = F_slice_from_s(m, l0, int32(1), int32(_a_F_porter_ISO_8859_1_stem_20))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L5
	} else {
		goto L239
	}
L224:
	;
	v862 = F_slice_from_s(m, l0, int32(3), int32(_a_F_porter_ISO_8859_1_stem_21))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L5
	} else {
		goto L237
	}
L225:
	;
	v856 = F_slice_from_s(m, l0, int32(4), int32(_a_F_porter_ISO_8859_1_stem_22))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L5
	} else {
		goto L235
	}
L226:
	;
	v850 = F_slice_from_s(m, l0, int32(4), int32(_a_F_porter_ISO_8859_1_stem_23))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L5
	} else {
		goto L233
	}
L227:
	;
	v844 = F_slice_from_s(m, l0, int32(4), int32(_a_F_porter_ISO_8859_1_stem_24))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L5
	} else {
		goto L231
	}
L228:
	;
	v838 = F_slice_from_s(m, l0, int32(4), int32(_a_F_porter_ISO_8859_1_stem_25))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L5
	} else {
		goto L229
	}
L229:
	;
	if int32(0) <= v838 {
		goto L215
	} else {
		goto L230
	}
L230:
	;
	v916 = v838
	goto L209
L231:
	;
	if int32(0) <= v844 {
		goto L215
	} else {
		goto L232
	}
L232:
	;
	v916 = v844
	goto L209
L233:
	;
	if int32(0) <= v850 {
		goto L215
	} else {
		goto L234
	}
L234:
	;
	v916 = v850
	goto L209
L235:
	;
	if int32(0) <= v856 {
		goto L215
	} else {
		goto L236
	}
L236:
	;
	v916 = v856
	goto L209
L237:
	;
	if int32(0) <= v862 {
		goto L215
	} else {
		goto L238
	}
L238:
	;
	v916 = v862
	goto L209
L239:
	;
	if int32(0) <= v868 {
		goto L215
	} else {
		goto L240
	}
L240:
	;
	v916 = v868
	goto L209
L241:
	;
	if int32(0) <= v874 {
		goto L215
	} else {
		goto L242
	}
L242:
	;
	v916 = v874
	goto L209
L243:
	;
	if int32(0) <= v880 {
		goto L215
	} else {
		goto L244
	}
L244:
	;
	v916 = v880
	goto L209
L245:
	;
	if int32(0) <= v886 {
		goto L215
	} else {
		goto L246
	}
L246:
	;
	v916 = v886
	goto L209
L247:
	;
	if int32(0) <= v892 {
		goto L215
	} else {
		goto L248
	}
L248:
	;
	v916 = v892
	goto L209
L249:
	;
	if int32(0) <= v898 {
		goto L215
	} else {
		goto L250
	}
L250:
	;
	v916 = v898
	goto L209
L251:
	;
	if int32(0) <= v904 {
		goto L215
	} else {
		goto L252
	}
L252:
	;
	v916 = v904
	goto L209
L253:
	;
	if v910 < int32(0) {
		v916 = v910
		goto L209
	} else {
		goto L254
	}
L254:
	;
	goto L215
L255:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v921
	v923 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v921
	v926 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v921-int32(2) <= v926 {
		v977 = v923
		goto L256
	} else {
		goto L257
	}
L256:
	;
	if v977 < int32(0) {
		v1298 = v977
		goto L1
	} else {
		goto L272
	}
L257:
	;
	v930 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v932 = int32(1)
	v934 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v930+v921-v932))))
	if base.B2i32(v934&int32(224) != int32(96))|base.B2i32(v932<<(uint(v934)%32)&int32(_a_F_porter_ISO_8859_1_stem_26) == int32(0)) != 0 {
		v977 = v923
		goto L256
	} else {
		goto L258
	}
L258:
	;
	v948 = F_find_among_b(m, l0, int32(_a_F_porter_ISO_8859_1_stem_27), int32(7))
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L5
	} else {
		goto L259
	}
L259:
	;
	if v948 == int32(0) {
		v977 = v923
		goto L256
	} else {
		goto L260
	}
L260:
	;
	v952 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v952
	v954 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v954)+4))
	if v952 < v955 {
		v977 = v923
		goto L256
	} else {
		goto L261
	}
L261:
	;
	switch v948 - int32(1) {
	case 0:
		goto L265
	case 1:
		goto L264
	case 2:
		goto L263
	default:
		goto L262
	}
L262:
	;
	v977 = int32(1)
	goto L256
L263:
	;
	v971 = F_slice_del(m, l0)
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L5
	} else {
		goto L270
	}
L264:
	;
	v967 = F_slice_from_s(m, l0, int32(2), int32(_a_F_porter_ISO_8859_1_stem_28))
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L5
	} else {
		goto L268
	}
L265:
	;
	v961 = F_slice_from_s(m, l0, int32(2), int32(_a_F_porter_ISO_8859_1_stem_29))
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L5
	} else {
		goto L266
	}
L266:
	;
	if int32(0) <= v961 {
		goto L262
	} else {
		goto L267
	}
L267:
	;
	v977 = v961
	goto L256
L268:
	;
	if int32(0) <= v967 {
		goto L262
	} else {
		goto L269
	}
L269:
	;
	v977 = v967
	goto L256
L270:
	;
	if v971 < int32(0) {
		v977 = v971
		goto L256
	} else {
		goto L271
	}
L271:
	;
	goto L262
L272:
	;
	v982 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v982
	v984 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v982
	v988 = v982 - int32(1)
	v989 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v988 <= v989 {
		v1044 = v984
		goto L273
	} else {
		goto L274
	}
L273:
	;
	if v1044 < int32(0) {
		v1298 = v1044
		goto L1
	} else {
		goto L288
	}
L274:
	;
	v991 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v993 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v991+v988))))
	if base.B2i32(v993&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v993)%32)&int32(_a_F_porter_ISO_8859_1_stem_30) == int32(0)) != 0 {
		v1044 = v984
		goto L273
	} else {
		goto L275
	}
L275:
	;
	v1007 = F_find_among_b(m, l0, int32(_a_F_porter_ISO_8859_1_stem_31), int32(19))
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L5
	} else {
		goto L276
	}
L276:
	;
	if v1007 == int32(0) {
		v1044 = v984
		goto L273
	} else {
		goto L277
	}
L277:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1011
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v1013)))
	if v1011 < v1014 {
		v1044 = v984
		goto L273
	} else {
		goto L278
	}
L278:
	;
	switch v1007 - int32(1) {
	case 0:
		goto L281
	case 1:
		goto L280
	default:
		goto L279
	}
L279:
	;
	v1044 = int32(1)
	goto L273
L280:
	;
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1011 <= v1022 {
		v1044 = v984
		goto L273
	} else {
		goto L284
	}
L281:
	;
	v1018 = F_slice_del(m, l0)
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L5
	} else {
		goto L282
	}
L282:
	;
	if int32(0) <= v1018 {
		goto L279
	} else {
		goto L283
	}
L283:
	;
	v1044 = v1018
	goto L273
L284:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1026 = int32(1)
	v1028 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1024+v1011-v1026))))
	if base.Ui32(v1026) < base.Ui32((v1028-int32(115))&int32(255)) {
		v1044 = v984
		goto L273
	} else {
		goto L285
	}
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1011 - int32(1)
	v1038 = F_slice_del(m, l0)
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L5
	} else {
		goto L286
	}
L286:
	;
	if v1038 < int32(0) {
		v1044 = v1038
		goto L273
	} else {
		goto L287
	}
L287:
	;
	goto L279
L288:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1049
	v1051 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1049
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1049 <= v1054 {
		v1237 = v1051
		goto L289
	} else {
		goto L290
	}
L289:
	;
	if v1237 < int32(0) {
		v1298 = v1237
		goto L1
	} else {
		goto L339
	}
L290:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1056+v1049-int32(1)))))
	if v1060 != int32(101) {
		v1237 = v1051
		goto L289
	} else {
		goto L291
	}
L291:
	;
	v1064 = v1049 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1064
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1064
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v1067)))
	if v1049 <= v1068 {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1067)+4))
	if v1049 <= v1070 {
		v1237 = v1051
		goto L289
	} else {
		goto L295
	}
L293:
	;
	goto L294
L294:
	;
	v1232 = F_slice_del(m, l0)
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L5
	} else {
		goto L335
	}
L295:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L299
L296:
	;
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1226 + (v1064 - v1072)
	goto L294
L297:
	;
	if v1121 != 0 {
		goto L296
	} else {
		goto L309
	}
L298:
	;
	v1121 = v1118
	goto L297
L299:
	;
	if v1080 <= v1081 {
		goto L301
	} else {
		goto L302
	}
L300:
	;
	v1118 = int32(0)
	goto L298
L301:
	;
	v1121 = int32(-1)
	goto L297
L302:
	;
	goto L303
L303:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1096 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1092+v1080-int32(1)))))
	if int32(121) < v1096 {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1080 - int32(1)
	goto L308
L305:
	;
	v1098 = v1096 - int32(89)
	if v1098 < int32(0) {
		goto L304
	} else {
		goto L306
	}
L306:
	;
	v1101 = int32(1)
	v1105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1098)>>(uint(int32(3))%32)))+uint32(_c_F_porter_ISO_8859_1_stem[1]))))
	if int32(base.Ui32(v1105)>>(uint(v1098&int32(7))%32))&v1101 != 0 {
		v1118 = v1101
		goto L298
	} else {
		goto L307
	}
L307:
	;
	goto L304
L308:
	;
	goto L300
L309:
	;
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L312
L310:
	;
	if v1174 != 0 {
		goto L296
	} else {
		goto L321
	}
L311:
	;
	v1174 = v1170
	goto L310
L312:
	;
	if v1130 <= v1131 {
		goto L314
	} else {
		goto L315
	}
L313:
	;
	v1170 = int32(0)
	goto L311
L314:
	;
	v1174 = int32(-1)
	goto L310
L315:
	;
	goto L316
L316:
	;
	v1143 = int32(1)
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1144+v1130-v1143))))
	if int32(121) < v1148 {
		v1170 = v1143
		goto L311
	} else {
		goto L317
	}
L317:
	;
	v1150 = v1148 - int32(97)
	if v1150 < int32(0) {
		v1170 = v1143
		goto L311
	} else {
		goto L318
	}
L318:
	;
	v1156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1150)>>(uint(int32(3))%32)))+uint32(_c_F_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v1156)>>(uint(v1150&int32(7))%32))&int32(1) == int32(0) {
		v1170 = v1143
		goto L311
	} else {
		goto L319
	}
L319:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1130 - int32(1)
	goto L320
L320:
	;
	goto L313
L321:
	;
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L324
L322:
	;
	if v1223 == int32(0) {
		v1237 = v1051
		goto L289
	} else {
		goto L334
	}
L323:
	;
	v1223 = v1220
	goto L322
L324:
	;
	if v1182 <= v1183 {
		goto L326
	} else {
		goto L327
	}
L325:
	;
	v1220 = int32(0)
	goto L323
L326:
	;
	v1223 = int32(-1)
	goto L322
L327:
	;
	goto L328
L328:
	;
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1194+v1182-int32(1)))))
	if int32(121) < v1198 {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1182 - int32(1)
	goto L333
L330:
	;
	v1200 = v1198 - int32(97)
	if v1200 < int32(0) {
		goto L329
	} else {
		goto L331
	}
L331:
	;
	v1203 = int32(1)
	v1207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1200)>>(uint(int32(3))%32)))+uint32(_c_F_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v1207)>>(uint(v1200&int32(7))%32))&v1203 != 0 {
		v1220 = v1203
		goto L323
	} else {
		goto L332
	}
L332:
	;
	goto L329
L333:
	;
	goto L325
L334:
	;
	goto L296
L335:
	;
	if int32(0) <= v1232 {
		goto L336
	} else {
		goto L337
	}
L336:
	;
	v1236 = int32(1)
	goto L338
L337:
	;
	v1236 = v1232
	goto L338
L338:
	;
	v1237 = v1236
	goto L289
L339:
	;
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1243
	v1245 = F_r_Step_5b(m, l0)
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L5
	} else {
		goto L340
	}
L340:
	;
	if v1245 < int32(0) {
		v1298 = v1245
		goto L1
	} else {
		goto L341
	}
L341:
	;
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1249
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v1251)+8))
	if v1252 == int32(0) {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1249
	v1298 = int32(1)
	goto L1
L343:
	;
	goto L344
L344:
	;
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1261 < v1260 {
		goto L346
	} else {
		goto L347
	}
L345:
	;
	v1298 = v1286
	goto L1
L346:
	;
	v1263 = v1260
	goto L348
L347:
	;
	v1263 = v1261
	goto L348
L348:
	;
	v1265 = v1260
	goto L349
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1265
	if v1265 != v1261 {
		goto L352
	} else {
		goto L353
	}
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1265
	v1281 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1265 + v1281
	v1286 = F_slice_from_s(m, l0, v1281, int32(_a_F_porter_ISO_8859_1_stem_32))
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L5
	} else {
		goto L357
	}
L351:
	;
	goto L350
L352:
	;
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1271+v1265))))
	if v1273 == int32(89) {
		goto L351
	} else {
		goto L355
	}
L353:
	;
	goto L354
L354:
	;
	if v1265 == v1263 {
		goto L342
	} else {
		goto L356
	}
L355:
	;
	goto L354
L356:
	;
	v1278 = v1265 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1278
	v1265 = v1278
	goto L349
L357:
	;
	if int32(0) <= v1286 {
		goto L344
	} else {
		goto L358
	}
L358:
	;
	goto L345
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
	var v173 int32
	_ = v173
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
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v838 int32
	_ = v838
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v857 int32
	_ = v857
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
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
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1016 int32
	_ = v1016
	var v1021 int32
	_ = v1021
	var v1025 int32
	_ = v1025
	var v1028 int32
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1046 int32
	_ = v1046
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1154 int32
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1166 int32
	_ = v1166
	var v1183 int32
	_ = v1183
	var v1190 int32
	_ = v1190
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1237 int32
	_ = v1237
	var v1239 int32
	_ = v1239
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1284 int32
	_ = v1284
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1296 int32
	_ = v1296
	var v1312 int32
	_ = v1312
	var v1319 int32
	_ = v1319
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1353 int32
	_ = v1353
	var v1355 int32
	_ = v1355
	var v1362 int32
	_ = v1362
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1385 int32
	_ = v1385
	var v1403 int32
	_ = v1403
	var v1405 int32
	_ = v1405
	var v1413 int32
	_ = v1413
	var v1417 int32
	_ = v1417
	var v1419 int32
	_ = v1419
	var v1425 int32
	_ = v1425
	var v1442 int32
	_ = v1442
	var v1449 int32
	_ = v1449
	var v1452 int32
	_ = v1452
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1468 int32
	_ = v1468
	var v1470 int32
	_ = v1470
	var v1473 int32
	_ = v1473
	var v1475 int32
	_ = v1475
	var v1479 int32
	_ = v1479
	var v1485 int32
	_ = v1485
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1511 int32
	_ = v1511
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1521 int32
	_ = v1521
	var v1523 int32
	_ = v1523
	var v1530 int32
	_ = v1530
	var v1532 int32
	_ = v1532
	var v1534 int32
	_ = v1534
	var v1536 int32
	_ = v1536
	var v1549 int32
	_ = v1549
	var v1551 int32
	_ = v1551
	var v1553 int32
	_ = v1553
	var v1571 int32
	_ = v1571
	var v1573 int32
	_ = v1573
	var v1581 int32
	_ = v1581
	var v1585 int32
	_ = v1585
	var v1587 int32
	_ = v1587
	var v1593 int32
	_ = v1593
	var v1602 int32
	_ = v1602
	var v1617 int32
	_ = v1617
	var v1620 int32
	_ = v1620
	var v1623 int32
	_ = v1623
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1633 int32
	_ = v1633
	var v1635 int32
	_ = v1635
	var v1638 int32
	_ = v1638
	var v1640 int32
	_ = v1640
	var v1643 int32
	_ = v1643
	var v1647 int32
	_ = v1647
	var v1649 int32
	_ = v1649
	var v1651 int32
	_ = v1651
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1669 int32
	_ = v1669
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1756 int32
	_ = v1756
	var v1761 int32
	_ = v1761
	var v1763 int32
	_ = v1763
	var v1766 int32
	_ = v1766
	var v1770 int32
	_ = v1770
	var v1772 int32
	_ = v1772
	var v1774 int32
	_ = v1774
	var v1788 int32
	_ = v1788
	var v1789 int32
	_ = v1789
	var v1792 int32
	_ = v1792
	var v1794 int32
	_ = v1794
	var v1795 int32
	_ = v1795
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1811 int32
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1818 int32
	_ = v1818
	var v1822 int32
	_ = v1822
	var v1824 int32
	_ = v1824
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1831 int32
	_ = v1831
	var v1833 int32
	_ = v1833
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1851 int32
	_ = v1851
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1862 int32
	_ = v1862
	var v1864 int32
	_ = v1864
	var v1866 int32
	_ = v1866
	var v1868 int32
	_ = v1868
	var v1878 int32
	_ = v1878
	var v1879 int32
	_ = v1879
	var v1885 int32
	_ = v1885
	var v1889 int32
	_ = v1889
	var v1891 int32
	_ = v1891
	var v1894 int32
	_ = v1894
	var v1896 int32
	_ = v1896
	var v1900 int32
	_ = v1900
	var v1904 int32
	_ = v1904
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1910 int32
	_ = v1910
	var v1912 int32
	_ = v1912
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1943 int32
	_ = v1943
	var v1944 int32
	_ = v1944
	var v1946 int32
	_ = v1946
	var v1948 int32
	_ = v1948
	var v1955 int32
	_ = v1955
	var v1957 int32
	_ = v1957
	var v1959 int32
	_ = v1959
	var v1961 int32
	_ = v1961
	var v1974 int32
	_ = v1974
	var v1976 int32
	_ = v1976
	var v1978 int32
	_ = v1978
	var v1996 int32
	_ = v1996
	var v1998 int32
	_ = v1998
	var v2006 int32
	_ = v2006
	var v2010 int32
	_ = v2010
	var v2012 int32
	_ = v2012
	var v2018 int32
	_ = v2018
	var v2035 int32
	_ = v2035
	var v2042 int32
	_ = v2042
	var v2055 int32
	_ = v2055
	var v2056 int32
	_ = v2056
	var v2057 int32
	_ = v2057
	var v2073 int32
	_ = v2073
	var v2074 int32
	_ = v2074
	var v2076 int32
	_ = v2076
	var v2078 int32
	_ = v2078
	var v2085 int32
	_ = v2085
	var v2087 int32
	_ = v2087
	var v2089 int32
	_ = v2089
	var v2091 int32
	_ = v2091
	var v2104 int32
	_ = v2104
	var v2106 int32
	_ = v2106
	var v2108 int32
	_ = v2108
	var v2126 int32
	_ = v2126
	var v2128 int32
	_ = v2128
	var v2136 int32
	_ = v2136
	var v2140 int32
	_ = v2140
	var v2142 int32
	_ = v2142
	var v2148 int32
	_ = v2148
	var v2164 int32
	_ = v2164
	var v2171 int32
	_ = v2171
	var v2184 int32
	_ = v2184
	var v2185 int32
	_ = v2185
	var v2186 int32
	_ = v2186
	var v2202 int32
	_ = v2202
	var v2203 int32
	_ = v2203
	var v2205 int32
	_ = v2205
	var v2207 int32
	_ = v2207
	var v2214 int32
	_ = v2214
	var v2216 int32
	_ = v2216
	var v2218 int32
	_ = v2218
	var v2220 int32
	_ = v2220
	var v2233 int32
	_ = v2233
	var v2235 int32
	_ = v2235
	var v2237 int32
	_ = v2237
	var v2255 int32
	_ = v2255
	var v2257 int32
	_ = v2257
	var v2265 int32
	_ = v2265
	var v2269 int32
	_ = v2269
	var v2271 int32
	_ = v2271
	var v2277 int32
	_ = v2277
	var v2294 int32
	_ = v2294
	var v2301 int32
	_ = v2301
	var v2304 int32
	_ = v2304
	var v2310 int32
	_ = v2310
	var v2311 int32
	_ = v2311
	var v2314 int32
	_ = v2314
	var v2315 int32
	_ = v2315
	var v2321 int32
	_ = v2321
	var v2323 int32
	_ = v2323
	var v2324 int32
	_ = v2324
	var v2327 int32
	_ = v2327
	var v2329 int32
	_ = v2329
	var v2330 int32
	_ = v2330
	var v2339 int32
	_ = v2339
	var v2341 int32
	_ = v2341
	var v2347 int32
	_ = v2347
	var v2348 int32
	_ = v2348
	var v2351 int32
	_ = v2351
	var v2360 int32
	_ = v2360
	var v2362 int32
	_ = v2362
	var v2367 int32
	_ = v2367
	var v2369 int32
	_ = v2369
	var v2376 int32
	_ = v2376
	var v2379 int32
	_ = v2379
	var v2383 int32
	_ = v2383
	var v2390 int32
	_ = v2390
	var v2391 int32
	_ = v2391
	var v2405 int32
	_ = v2405
	var v2410 int32
	_ = v2410
	var v2415 int32
	_ = v2415
	var v2416 int32
	_ = v2416
	var v2428 int32
	_ = v2428
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
	return v2428
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
	v25 = F_slice_from_s(m, l0, v19, int32(_a_F_porter_UTF_8_stem_0))
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
		v2428 = v25
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
	v236 = F_slice_from_s(m, l0, v231, int32(_a_F_porter_UTF_8_stem_1))
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
	v129 = v114&int32(63) | (v72<<(uint(int32(18))%32)&int32(_a_F_porter_UTF_8_stem_2) | v81<<(uint(int32(12))%32) | v97<<(uint(int32(6))%32))
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
	v129 = v72<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_3) | v81<<(uint(int32(6))%32) | v97
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
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v134)>>(uint(int32(3))%32)))+uint32(_c_F_porter_UTF_8_stem[0]))))
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
	v173 = v161
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
	v173 = v164
	goto L12
L41:
	;
	goto L42
L42:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162+v164))))
	if v168 == int32(121) {
		goto L11
	} else {
		goto L43
	}
L43:
	;
	v171 = v165
	v173 = v164
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
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173+v181))))
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
	v200 = int32(*(*int8)(unsafe.Add(mBase, uint32(v173+v197))))
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
		v2428 = v236
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
	v337 = v322&int32(63) | (v280<<(uint(int32(18))%32)&int32(_a_F_porter_UTF_8_stem_2) | v289<<(uint(int32(12))%32) | v305<<(uint(int32(6))%32))
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
	v337 = v280<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_3) | v289<<(uint(int32(6))%32) | v305
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
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v342)>>(uint(int32(3))%32)))+uint32(_c_F_porter_UTF_8_stem[0]))))
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
	v459 = v444&int32(63) | (v402<<(uint(int32(18))%32)&int32(_a_F_porter_UTF_8_stem_2) | v411<<(uint(int32(12))%32) | v427<<(uint(int32(6))%32))
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
	v459 = v402<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_3) | v411<<(uint(int32(6))%32) | v427
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
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v464)>>(uint(int32(3))%32)))+uint32(_c_F_porter_UTF_8_stem[0]))))
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
	v584 = v569&int32(63) | (v527<<(uint(int32(18))%32)&int32(_a_F_porter_UTF_8_stem_2) | v536<<(uint(int32(12))%32) | v552<<(uint(int32(6))%32))
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
	v584 = v527<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_3) | v536<<(uint(int32(6))%32) | v552
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
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v589)>>(uint(int32(3))%32)))+uint32(_c_F_porter_UTF_8_stem[0]))))
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
	v706 = v691&int32(63) | (v649<<(uint(int32(18))%32)&int32(_a_F_porter_UTF_8_stem_2) | v658<<(uint(int32(12))%32) | v674<<(uint(int32(6))%32))
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
	v706 = v649<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_3) | v658<<(uint(int32(6))%32) | v674
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
	v717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v711)>>(uint(int32(3))%32)))+uint32(_c_F_porter_UTF_8_stem[0]))))
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
	v758 = F_find_among_b(m, l0, int32(_a_F_porter_UTF_8_stem_4), int32(4))
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
	v774 = F_slice_from_s(m, l0, int32(1), int32(_a_F_porter_UTF_8_stem_5))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L5
	} else {
		goto L180
	}
L177:
	;
	v768 = F_slice_from_s(m, l0, int32(2), int32(_a_F_porter_UTF_8_stem_6))
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
	v2428 = v768
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
	v2428 = v774
	goto L1
L182:
	;
	if v778 < int32(0) {
		v2428 = v778
		goto L1
	} else {
		goto L183
	}
L183:
	;
	goto L170
L184:
	;
	v1468 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1468
	v1470 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1468
	v1473 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1468 <= v1473 {
		v1635 = v1470
		goto L311
	} else {
		goto L312
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
	v797 = F_find_among_b(m, l0, int32(_a_F_porter_UTF_8_stem_7), int32(3))
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
	v827 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v829 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v838 = v827
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
	v810 = F_slice_from_s(m, l0, int32(2), int32(_a_F_porter_UTF_8_stem_8))
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
	v2428 = v810
	goto L1
L194:
	;
	if v944 < int32(0) {
		goto L184
	} else {
		goto L212
	}
L195:
	;
	v944 = int32(-1)
	goto L194
L196:
	;
	if v838 <= v828 {
		goto L195
	} else {
		goto L198
	}
L198:
	;
	v845 = int32(1)
	v846 = v838 - v845
	v848 = int32(*(*int8)(unsafe.Add(mBase, uint32(v829+v846))))
	v850 = v848 & int32(255)
	if base.B2i32(v846 == v828)|base.B2i32(int32(0) <= v848) != 0 {
		v908 = v850
		v912 = v845
		goto L199
	} else {
		goto L200
	}
L199:
	;
	if int32(121) < v908 {
		goto L207
	} else {
		goto L208
	}
L200:
	;
	v857 = v850 & int32(63)
	v859 = v838 - int32(2)
	v861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v829+v859))))
	v863 = v861 << (uint(int32(6)) % 32)
	if base.B2i32(v859 != v828)&base.B2i32(base.Ui32(v861) < base.Ui32(int32(192))) == int32(0) {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v908 = v863&int32(1984) | v857
	v912 = int32(2)
	goto L199
L202:
	;
	goto L203
L203:
	;
	v876 = v863&int32(4032) | v857
	v878 = v838 - int32(3)
	v880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v829+v878))))
	if base.B2i32(v878 != v828)&base.B2i32(base.Ui32(v880) < base.Ui32(int32(224))) == int32(0) {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v908 = v880<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_3) | v876
	v912 = int32(3)
	goto L199
L205:
	;
	goto L206
L206:
	;
	v898 = int32(4)
	v900 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v838+v829-v898))))
	v908 = v880<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_9) | v900&int32(7)<<(uint(int32(18))%32) | v876
	v912 = v898
	goto L199
L207:
	;
	v929 = v838 - v912
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v929
	v838 = v929
	goto L196
L208:
	;
	v914 = v908 - int32(97)
	if v914 < int32(0) {
		goto L207
	} else {
		goto L209
	}
L209:
	;
	v920 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v914)>>(uint(int32(3))%32)))+uint32(_c_F_porter_UTF_8_stem[0]))))
	if int32(base.Ui32(v920)>>(uint(v914&int32(7))%32))&int32(1) == int32(0) {
		goto L207
	} else {
		goto L210
	}
L210:
	;
	v944 = v912
	goto L194
L212:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v947 + (v801 - v814)
	v951 = F_slice_del(m, l0)
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L5
	} else {
		goto L213
	}
L213:
	;
	if v951 < int32(0) {
		v2428 = v951
		goto L1
	} else {
		goto L214
	}
L214:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v956 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v957 = v955 - v956
	v959 = v955 - int32(1)
	v960 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v959 <= v960 {
		v1056 = v955
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+4))
	if v1056 != v1058 {
		goto L184
	} else {
		goto L245
	}
L216:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v964 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v962+v959))))
	if base.B2i32(v964&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v964)%32)&int32(68514004) == int32(0)) != 0 {
		v1056 = v955
		goto L215
	} else {
		goto L217
	}
L217:
	;
	v978 = F_find_among_b(m, l0, int32(_a_F_porter_UTF_8_stem_10), int32(13))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L5
	} else {
		goto L218
	}
L218:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v981 = v980 + v957
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v981
	switch v978 - int32(1) {
	case 0:
		goto L220
	case 1:
		goto L219
	case 2:
		v1056 = v981
		goto L215
	default:
		goto L184
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v981
	v993 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v994 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L225
L220:
	;
	v987 = F_insert_s(m, l0, v981, v981, int32(1), int32(_a_F_porter_UTF_8_stem_11))
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L5
	} else {
		goto L221
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v981
	if int32(0) <= v987 {
		goto L184
	} else {
		goto L222
	}
L222:
	;
	v2428 = v987
	goto L1
L223:
	;
	if v1046 < int32(0) {
		goto L184
	} else {
		goto L242
	}
L225:
	;
	goto L226
L226:
	;
	goto L227
L227:
	;
	v1001 = v981
	v1003 = int32(1)
	goto L230
L229:
	;
	v1046 = v1028
	goto L223
L230:
	;
	if v1001 <= v994 {
		goto L232
	} else {
		goto L233
	}
L231:
	;
	goto L229
L232:
	;
	v1046 = int32(-1)
	goto L223
L233:
	;
	goto L234
L234:
	;
	v1008 = v1001 - int32(1)
	v1010 = int32(*(*int8)(unsafe.Add(mBase, uint32(v993+v1008))))
	if base.B2i32(int32(0) <= v1010)|base.B2i32(v1008 <= v994) != 0 {
		v1028 = v1008
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v1032 = int32(1)
	if v1032 < v1003 {
		v1001 = v1028
		v1003 = v1003 - v1032
		goto L230
	} else {
		goto L241
	}
L236:
	;
	v1016 = v1008
	goto L237
L237:
	;
	v1021 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v993+v1016))))
	if base.Ui32(int32(191)) < base.Ui32(v1021) {
		v1028 = v1016
		goto L235
	} else {
		goto L239
	}
L238:
	;
	v1028 = v994
	goto L235
L239:
	;
	v1025 = v1016 - int32(1)
	if v994 < v1025 {
		v1016 = v1025
		goto L237
	} else {
		goto L240
	}
L240:
	;
	goto L238
L241:
	;
	goto L231
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1046
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1046
	v1051 = F_slice_del(m, l0)
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L5
	} else {
		goto L243
	}
L243:
	;
	if int32(0) <= v1051 {
		goto L184
	} else {
		goto L244
	}
L244:
	;
	v2428 = v1051
	goto L1
L245:
	;
	v1060 = int32(0)
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L249
L246:
	;
	if v1452 == int32(0) {
		goto L184
	} else {
		goto L308
	}
L247:
	;
	if v1190 != 0 {
		v1452 = v1060
		goto L246
	} else {
		goto L265
	}
L248:
	;
	v1190 = v1183
	goto L247
L249:
	;
	if v1073 <= v1074 {
		v1183 = int32(-1)
		goto L248
	} else {
		goto L251
	}
L250:
	;
	v1183 = int32(0)
	goto L248
L251:
	;
	v1091 = int32(1)
	v1092 = v1073 - v1091
	v1094 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1075+v1092))))
	v1096 = v1094 & int32(255)
	if base.B2i32(v1092 == v1074)|base.B2i32(int32(0) <= v1094) != 0 {
		v1154 = v1096
		v1158 = v1091
		goto L252
	} else {
		goto L253
	}
L252:
	;
	if int32(121) < v1154 {
		goto L260
	} else {
		goto L261
	}
L253:
	;
	v1103 = v1096 & int32(63)
	v1105 = v1073 - int32(2)
	v1107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075+v1105))))
	v1109 = v1107 << (uint(int32(6)) % 32)
	if base.B2i32(v1105 != v1074)&base.B2i32(base.Ui32(v1107) < base.Ui32(int32(192))) == int32(0) {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v1154 = v1109&int32(1984) | v1103
	v1158 = int32(2)
	goto L252
L255:
	;
	goto L256
L256:
	;
	v1122 = v1109&int32(4032) | v1103
	v1124 = v1073 - int32(3)
	v1126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075+v1124))))
	if base.B2i32(v1124 != v1074)&base.B2i32(base.Ui32(v1126) < base.Ui32(int32(224))) == int32(0) {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v1154 = v1126<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_3) | v1122
	v1158 = int32(3)
	goto L252
L258:
	;
	goto L259
L259:
	;
	v1144 = int32(4)
	v1146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1073+v1075-v1144))))
	v1154 = v1126<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_9) | v1146&int32(7)<<(uint(int32(18))%32) | v1122
	v1158 = v1144
	goto L252
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1073 - v1158
	goto L264
L261:
	;
	v1160 = v1154 - int32(89)
	if v1160 < int32(0) {
		goto L260
	} else {
		goto L262
	}
L262:
	;
	v1166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1160)>>(uint(int32(3))%32)))+uint32(_c_F_porter_UTF_8_stem[1]))))
	if int32(base.Ui32(v1166)>>(uint(v1160&int32(7))%32))&int32(1) == int32(0) {
		goto L260
	} else {
		goto L263
	}
L263:
	;
	v1190 = v1158
	goto L247
L264:
	;
	goto L250
L265:
	;
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L268
L266:
	;
	if v1319 != 0 {
		v1452 = v1060
		goto L246
	} else {
		goto L289
	}
L267:
	;
	v1319 = v1312
	goto L266
L268:
	;
	if v1203 <= v1204 {
		v1312 = int32(-1)
		goto L267
	} else {
		goto L270
	}
L269:
	;
	v1312 = int32(0)
	goto L267
L270:
	;
	v1221 = int32(1)
	v1222 = v1203 - v1221
	v1224 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1205+v1222))))
	v1226 = v1224 & int32(255)
	if base.B2i32(v1222 == v1204)|base.B2i32(int32(0) <= v1224) != 0 {
		v1284 = v1226
		v1288 = v1221
		goto L271
	} else {
		goto L272
	}
L271:
	;
	if int32(121) < v1284 {
		goto L279
	} else {
		goto L280
	}
L272:
	;
	v1233 = v1226 & int32(63)
	v1235 = v1203 - int32(2)
	v1237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1205+v1235))))
	v1239 = v1237 << (uint(int32(6)) % 32)
	if base.B2i32(v1235 != v1204)&base.B2i32(base.Ui32(v1237) < base.Ui32(int32(192))) == int32(0) {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v1284 = v1239&int32(1984) | v1233
	v1288 = int32(2)
	goto L271
L274:
	;
	goto L275
L275:
	;
	v1252 = v1239&int32(4032) | v1233
	v1254 = v1203 - int32(3)
	v1256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1205+v1254))))
	if base.B2i32(v1254 != v1204)&base.B2i32(base.Ui32(v1256) < base.Ui32(int32(224))) == int32(0) {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	v1284 = v1256<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_3) | v1252
	v1288 = int32(3)
	goto L271
L277:
	;
	goto L278
L278:
	;
	v1274 = int32(4)
	v1276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1203+v1205-v1274))))
	v1284 = v1256<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_9) | v1276&int32(7)<<(uint(int32(18))%32) | v1252
	v1288 = v1274
	goto L271
L279:
	;
	v1319 = v1288
	goto L266
L280:
	;
	goto L281
L281:
	;
	v1290 = v1284 - int32(97)
	if v1290 < int32(0) {
		goto L282
	} else {
		goto L283
	}
L282:
	;
	v1319 = v1288
	goto L266
L283:
	;
	goto L284
L284:
	;
	v1296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1290)>>(uint(int32(3))%32)))+uint32(_c_F_porter_UTF_8_stem[0]))))
	if int32(base.Ui32(v1296)>>(uint(v1290&int32(7))%32))&int32(1) == int32(0) {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1319 = v1288
	goto L266
L286:
	;
	goto L287
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1203 - v1288
	goto L288
L288:
	;
	goto L269
L289:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L292
L290:
	;
	v1452 = base.B2i32(v1449 == int32(0))
	goto L246
L291:
	;
	v1449 = v1442
	goto L290
L292:
	;
	if v1332 <= v1333 {
		v1442 = int32(-1)
		goto L291
	} else {
		goto L294
	}
L293:
	;
	v1442 = int32(0)
	goto L291
L294:
	;
	v1350 = int32(1)
	v1351 = v1332 - v1350
	v1353 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1334+v1351))))
	v1355 = v1353 & int32(255)
	if base.B2i32(v1351 == v1333)|base.B2i32(int32(0) <= v1353) != 0 {
		v1413 = v1355
		v1417 = v1350
		goto L295
	} else {
		goto L296
	}
L295:
	;
	if int32(121) < v1413 {
		goto L303
	} else {
		goto L304
	}
L296:
	;
	v1362 = v1355 & int32(63)
	v1364 = v1332 - int32(2)
	v1366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1334+v1364))))
	v1368 = v1366 << (uint(int32(6)) % 32)
	if base.B2i32(v1364 != v1333)&base.B2i32(base.Ui32(v1366) < base.Ui32(int32(192))) == int32(0) {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	v1413 = v1368&int32(1984) | v1362
	v1417 = int32(2)
	goto L295
L298:
	;
	goto L299
L299:
	;
	v1381 = v1368&int32(4032) | v1362
	v1383 = v1332 - int32(3)
	v1385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1334+v1383))))
	if base.B2i32(v1383 != v1333)&base.B2i32(base.Ui32(v1385) < base.Ui32(int32(224))) == int32(0) {
		goto L300
	} else {
		goto L301
	}
L300:
	;
	v1413 = v1385<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_3) | v1381
	v1417 = int32(3)
	goto L295
L301:
	;
	goto L302
L302:
	;
	v1403 = int32(4)
	v1405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1332+v1334-v1403))))
	v1413 = v1385<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_9) | v1405&int32(7)<<(uint(int32(18))%32) | v1381
	v1417 = v1403
	goto L295
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1332 - v1417
	goto L307
L304:
	;
	v1419 = v1413 - int32(97)
	if v1419 < int32(0) {
		goto L303
	} else {
		goto L305
	}
L305:
	;
	v1425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1419)>>(uint(int32(3))%32)))+uint32(_c_F_porter_UTF_8_stem[0]))))
	if int32(base.Ui32(v1425)>>(uint(v1419&int32(7))%32))&int32(1) == int32(0) {
		goto L303
	} else {
		goto L306
	}
L306:
	;
	v1449 = v1417
	goto L290
L307:
	;
	goto L293
L308:
	;
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1456 = v1455 + v957
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1456
	v1460 = F_insert_s(m, l0, v1456, v1456, int32(1), int32(_a_F_porter_UTF_8_stem_12))
	mBase = m.M
	v1461 = m.ExcPending
	if v1461 != 0 {
		goto L5
	} else {
		goto L309
	}
L309:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1456
	if v1460 < int32(0) {
		v2428 = v1460
		goto L1
	} else {
		goto L310
	}
L310:
	;
	goto L184
L311:
	;
	if v1635 < int32(0) {
		v2428 = v1635
		goto L1
	} else {
		goto L337
	}
L312:
	;
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1475+v1468-int32(1)))))
	if v1479|int32(32) != int32(121) {
		v1635 = v1470
		goto L311
	} else {
		goto L313
	}
L313:
	;
	v1485 = v1468 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1485
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1511 = v1485
	goto L316
L314:
	;
	if v1617 < int32(0) {
		v1635 = v1470
		goto L311
	} else {
		goto L332
	}
L315:
	;
	v1617 = int32(-1)
	goto L314
L316:
	;
	if v1511 <= v1501 {
		goto L315
	} else {
		goto L318
	}
L318:
	;
	v1518 = int32(1)
	v1519 = v1511 - v1518
	v1521 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1502+v1519))))
	v1523 = v1521 & int32(255)
	if base.B2i32(v1519 == v1501)|base.B2i32(int32(0) <= v1521) != 0 {
		v1581 = v1523
		v1585 = v1518
		goto L319
	} else {
		goto L320
	}
L319:
	;
	if int32(121) < v1581 {
		goto L327
	} else {
		goto L328
	}
L320:
	;
	v1530 = v1523 & int32(63)
	v1532 = v1511 - int32(2)
	v1534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1502+v1532))))
	v1536 = v1534 << (uint(int32(6)) % 32)
	if base.B2i32(v1532 != v1501)&base.B2i32(base.Ui32(v1534) < base.Ui32(int32(192))) == int32(0) {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v1581 = v1536&int32(1984) | v1530
	v1585 = int32(2)
	goto L319
L322:
	;
	goto L323
L323:
	;
	v1549 = v1536&int32(4032) | v1530
	v1551 = v1511 - int32(3)
	v1553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1502+v1551))))
	if base.B2i32(v1551 != v1501)&base.B2i32(base.Ui32(v1553) < base.Ui32(int32(224))) == int32(0) {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	v1581 = v1553<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_3) | v1549
	v1585 = int32(3)
	goto L319
L325:
	;
	goto L326
L326:
	;
	v1571 = int32(4)
	v1573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1511+v1502-v1571))))
	v1581 = v1553<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_9) | v1573&int32(7)<<(uint(int32(18))%32) | v1549
	v1585 = v1571
	goto L319
L327:
	;
	v1602 = v1511 - v1585
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1602
	v1511 = v1602
	goto L316
L328:
	;
	v1587 = v1581 - int32(97)
	if v1587 < int32(0) {
		goto L327
	} else {
		goto L329
	}
L329:
	;
	v1593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1587)>>(uint(int32(3))%32)))+uint32(_c_F_porter_UTF_8_stem[0]))))
	if int32(base.Ui32(v1593)>>(uint(v1587&int32(7))%32))&int32(1) == int32(0) {
		goto L327
	} else {
		goto L330
	}
L330:
	;
	v1617 = v1585
	goto L314
L332:
	;
	v1620 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1620 - v1617
	v1623 = int32(1)
	v1626 = F_slice_from_s(m, l0, v1623, int32(_a_F_porter_UTF_8_stem_13))
	mBase = m.M
	v1627 = m.ExcPending
	if v1627 != 0 {
		goto L5
	} else {
		goto L333
	}
L333:
	;
	if int32(0) <= v1626 {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	v1633 = v1623
	goto L336
L335:
	;
	v1633 = v1626 >> (uint(int32(31)) % 32) & v1626
	goto L336
L336:
	;
	v1635 = v1633
	goto L311
L337:
	;
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1638
	v1640 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1638
	v1643 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1638-int32(2) <= v1643 {
		v1756 = v1640
		goto L338
	} else {
		goto L339
	}
L338:
	;
	if v1756 < int32(0) {
		v2428 = v1756
		goto L1
	} else {
		goto L384
	}
L339:
	;
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1649 = int32(1)
	v1651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1647+v1638-v1649))))
	if base.B2i32(v1651&int32(224) != int32(96))|base.B2i32(v1649<<(uint(v1651)%32)&int32(_a_F_porter_UTF_8_stem_14) == int32(0)) != 0 {
		v1756 = v1640
		goto L338
	} else {
		goto L340
	}
L340:
	;
	v1665 = F_find_among_b(m, l0, int32(_a_F_porter_UTF_8_stem_15), int32(20))
	mBase = m.M
	v1666 = m.ExcPending
	if v1666 != 0 {
		goto L5
	} else {
		goto L341
	}
L341:
	;
	if v1665 == int32(0) {
		v1756 = v1640
		goto L338
	} else {
		goto L342
	}
L342:
	;
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1669
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(v1671)+4))
	if v1669 < v1672 {
		v1756 = v1640
		goto L338
	} else {
		goto L343
	}
L343:
	;
	switch v1665 - int32(1) {
	case 0:
		goto L357
	case 1:
		goto L356
	case 2:
		goto L355
	case 3:
		goto L354
	case 4:
		goto L353
	case 5:
		goto L352
	case 6:
		goto L351
	case 7:
		goto L350
	case 8:
		goto L349
	case 9:
		goto L348
	case 10:
		goto L347
	case 11:
		goto L346
	case 12:
		goto L345
	default:
		goto L344
	}
L344:
	;
	v1756 = int32(1)
	goto L338
L345:
	;
	v1750 = F_slice_from_s(m, l0, int32(3), int32(_a_F_porter_UTF_8_stem_16))
	mBase = m.M
	v1751 = m.ExcPending
	if v1751 != 0 {
		goto L5
	} else {
		goto L382
	}
L346:
	;
	v1744 = F_slice_from_s(m, l0, int32(3), int32(_a_F_porter_UTF_8_stem_17))
	mBase = m.M
	v1745 = m.ExcPending
	if v1745 != 0 {
		goto L5
	} else {
		goto L380
	}
L347:
	;
	v1738 = F_slice_from_s(m, l0, int32(3), int32(_a_F_porter_UTF_8_stem_18))
	mBase = m.M
	v1739 = m.ExcPending
	if v1739 != 0 {
		goto L5
	} else {
		goto L378
	}
L348:
	;
	v1732 = F_slice_from_s(m, l0, int32(3), int32(_a_F_porter_UTF_8_stem_19))
	mBase = m.M
	v1733 = m.ExcPending
	if v1733 != 0 {
		goto L5
	} else {
		goto L376
	}
L349:
	;
	v1726 = F_slice_from_s(m, l0, int32(2), int32(_a_F_porter_UTF_8_stem_20))
	mBase = m.M
	v1727 = m.ExcPending
	if v1727 != 0 {
		goto L5
	} else {
		goto L374
	}
L350:
	;
	v1720 = F_slice_from_s(m, l0, int32(3), int32(_a_F_porter_UTF_8_stem_21))
	mBase = m.M
	v1721 = m.ExcPending
	if v1721 != 0 {
		goto L5
	} else {
		goto L372
	}
L351:
	;
	v1714 = F_slice_from_s(m, l0, int32(3), int32(_a_F_porter_UTF_8_stem_22))
	mBase = m.M
	v1715 = m.ExcPending
	if v1715 != 0 {
		goto L5
	} else {
		goto L370
	}
L352:
	;
	v1708 = F_slice_from_s(m, l0, int32(1), int32(_a_F_porter_UTF_8_stem_23))
	mBase = m.M
	v1709 = m.ExcPending
	if v1709 != 0 {
		goto L5
	} else {
		goto L368
	}
L353:
	;
	v1702 = F_slice_from_s(m, l0, int32(3), int32(_a_F_porter_UTF_8_stem_24))
	mBase = m.M
	v1703 = m.ExcPending
	if v1703 != 0 {
		goto L5
	} else {
		goto L366
	}
L354:
	;
	v1696 = F_slice_from_s(m, l0, int32(4), int32(_a_F_porter_UTF_8_stem_25))
	mBase = m.M
	v1697 = m.ExcPending
	if v1697 != 0 {
		goto L5
	} else {
		goto L364
	}
L355:
	;
	v1690 = F_slice_from_s(m, l0, int32(4), int32(_a_F_porter_UTF_8_stem_26))
	mBase = m.M
	v1691 = m.ExcPending
	if v1691 != 0 {
		goto L5
	} else {
		goto L362
	}
L356:
	;
	v1684 = F_slice_from_s(m, l0, int32(4), int32(_a_F_porter_UTF_8_stem_27))
	mBase = m.M
	v1685 = m.ExcPending
	if v1685 != 0 {
		goto L5
	} else {
		goto L360
	}
L357:
	;
	v1678 = F_slice_from_s(m, l0, int32(4), int32(_a_F_porter_UTF_8_stem_28))
	mBase = m.M
	v1679 = m.ExcPending
	if v1679 != 0 {
		goto L5
	} else {
		goto L358
	}
L358:
	;
	if int32(0) <= v1678 {
		goto L344
	} else {
		goto L359
	}
L359:
	;
	v1756 = v1678
	goto L338
L360:
	;
	if int32(0) <= v1684 {
		goto L344
	} else {
		goto L361
	}
L361:
	;
	v1756 = v1684
	goto L338
L362:
	;
	if int32(0) <= v1690 {
		goto L344
	} else {
		goto L363
	}
L363:
	;
	v1756 = v1690
	goto L338
L364:
	;
	if int32(0) <= v1696 {
		goto L344
	} else {
		goto L365
	}
L365:
	;
	v1756 = v1696
	goto L338
L366:
	;
	if int32(0) <= v1702 {
		goto L344
	} else {
		goto L367
	}
L367:
	;
	v1756 = v1702
	goto L338
L368:
	;
	if int32(0) <= v1708 {
		goto L344
	} else {
		goto L369
	}
L369:
	;
	v1756 = v1708
	goto L338
L370:
	;
	if int32(0) <= v1714 {
		goto L344
	} else {
		goto L371
	}
L371:
	;
	v1756 = v1714
	goto L338
L372:
	;
	if int32(0) <= v1720 {
		goto L344
	} else {
		goto L373
	}
L373:
	;
	v1756 = v1720
	goto L338
L374:
	;
	if int32(0) <= v1726 {
		goto L344
	} else {
		goto L375
	}
L375:
	;
	v1756 = v1726
	goto L338
L376:
	;
	if int32(0) <= v1732 {
		goto L344
	} else {
		goto L377
	}
L377:
	;
	v1756 = v1732
	goto L338
L378:
	;
	if int32(0) <= v1738 {
		goto L344
	} else {
		goto L379
	}
L379:
	;
	v1756 = v1738
	goto L338
L380:
	;
	if int32(0) <= v1744 {
		goto L344
	} else {
		goto L381
	}
L381:
	;
	v1756 = v1744
	goto L338
L382:
	;
	if v1750 < int32(0) {
		v1756 = v1750
		goto L338
	} else {
		goto L383
	}
L383:
	;
	goto L344
L384:
	;
	v1761 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1761
	v1763 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1761
	v1766 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1761-int32(2) <= v1766 {
		v1818 = v1763
		goto L385
	} else {
		goto L386
	}
L385:
	;
	if v1818 < int32(0) {
		v2428 = v1818
		goto L1
	} else {
		goto L401
	}
L386:
	;
	v1770 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1772 = int32(1)
	v1774 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1770+v1761-v1772))))
	if base.B2i32(v1774&int32(224) != int32(96))|base.B2i32(v1772<<(uint(v1774)%32)&int32(_a_F_porter_UTF_8_stem_29) == int32(0)) != 0 {
		v1818 = v1763
		goto L385
	} else {
		goto L387
	}
L387:
	;
	v1788 = F_find_among_b(m, l0, int32(_a_F_porter_UTF_8_stem_30), int32(7))
	mBase = m.M
	v1789 = m.ExcPending
	if v1789 != 0 {
		goto L5
	} else {
		goto L388
	}
L388:
	;
	if v1788 == int32(0) {
		v1818 = v1763
		goto L385
	} else {
		goto L389
	}
L389:
	;
	v1792 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1792
	v1794 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1795 = *(*int32)(unsafe.Add(mBase, uint32(v1794)+4))
	if v1792 < v1795 {
		v1818 = v1763
		goto L385
	} else {
		goto L390
	}
L390:
	;
	switch v1788 - int32(1) {
	case 0:
		goto L394
	case 1:
		goto L393
	case 2:
		goto L392
	default:
		goto L391
	}
L391:
	;
	v1818 = int32(1)
	goto L385
L392:
	;
	v1811 = F_slice_del(m, l0)
	mBase = m.M
	v1812 = m.ExcPending
	if v1812 != 0 {
		goto L5
	} else {
		goto L399
	}
L393:
	;
	v1807 = F_slice_from_s(m, l0, int32(2), int32(_a_F_porter_UTF_8_stem_31))
	mBase = m.M
	v1808 = m.ExcPending
	if v1808 != 0 {
		goto L5
	} else {
		goto L397
	}
L394:
	;
	v1801 = F_slice_from_s(m, l0, int32(2), int32(_a_F_porter_UTF_8_stem_32))
	mBase = m.M
	v1802 = m.ExcPending
	if v1802 != 0 {
		goto L5
	} else {
		goto L395
	}
L395:
	;
	if int32(0) <= v1801 {
		goto L391
	} else {
		goto L396
	}
L396:
	;
	v1818 = v1801
	goto L385
L397:
	;
	if int32(0) <= v1807 {
		goto L391
	} else {
		goto L398
	}
L398:
	;
	v1818 = v1807
	goto L385
L399:
	;
	if v1811 < int32(0) {
		v1818 = v1811
		goto L385
	} else {
		goto L400
	}
L400:
	;
	goto L391
L401:
	;
	v1822 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1822
	v1824 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1822
	v1828 = v1822 - int32(1)
	v1829 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1828 <= v1829 {
		v1885 = v1824
		goto L402
	} else {
		goto L403
	}
L402:
	;
	if v1885 < int32(0) {
		v2428 = v1885
		goto L1
	} else {
		goto L417
	}
L403:
	;
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1833 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1831+v1828))))
	if base.B2i32(v1833&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1833)%32)&int32(_a_F_porter_UTF_8_stem_33) == int32(0)) != 0 {
		v1885 = v1824
		goto L402
	} else {
		goto L404
	}
L404:
	;
	v1847 = F_find_among_b(m, l0, int32(_a_F_porter_UTF_8_stem_34), int32(19))
	mBase = m.M
	v1848 = m.ExcPending
	if v1848 != 0 {
		goto L5
	} else {
		goto L405
	}
L405:
	;
	if v1847 == int32(0) {
		v1885 = v1824
		goto L402
	} else {
		goto L406
	}
L406:
	;
	v1851 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1851
	v1853 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1854 = *(*int32)(unsafe.Add(mBase, uint32(v1853)))
	if v1851 < v1854 {
		v1885 = v1824
		goto L402
	} else {
		goto L407
	}
L407:
	;
	switch v1847 - int32(1) {
	case 0:
		goto L410
	case 1:
		goto L409
	default:
		goto L408
	}
L408:
	;
	v1885 = int32(1)
	goto L402
L409:
	;
	v1862 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1851 <= v1862 {
		v1885 = v1824
		goto L402
	} else {
		goto L413
	}
L410:
	;
	v1858 = F_slice_del(m, l0)
	mBase = m.M
	v1859 = m.ExcPending
	if v1859 != 0 {
		goto L5
	} else {
		goto L411
	}
L411:
	;
	if int32(0) <= v1858 {
		goto L408
	} else {
		goto L412
	}
L412:
	;
	v1885 = v1858
	goto L402
L413:
	;
	v1864 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1866 = int32(1)
	v1868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1864+v1851-v1866))))
	if base.Ui32(v1866) < base.Ui32((v1868-int32(115))&int32(255)) {
		v1885 = v1824
		goto L402
	} else {
		goto L414
	}
L414:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1851 - int32(1)
	v1878 = F_slice_del(m, l0)
	mBase = m.M
	v1879 = m.ExcPending
	if v1879 != 0 {
		goto L5
	} else {
		goto L415
	}
L415:
	;
	if v1878 < int32(0) {
		v1885 = v1878
		goto L402
	} else {
		goto L416
	}
L416:
	;
	goto L408
L417:
	;
	v1889 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1889
	v1891 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1889
	v1894 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1889 <= v1894 {
		v2315 = v1891
		goto L418
	} else {
		goto L419
	}
L418:
	;
	if v2315 < int32(0) {
		v2428 = v2315
		goto L1
	} else {
		goto L492
	}
L419:
	;
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1900 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1896+v1889-int32(1)))))
	if v1900 != int32(101) {
		v2315 = v1891
		goto L418
	} else {
		goto L420
	}
L420:
	;
	v1904 = v1889 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1904
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1904
	v1907 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1908 = *(*int32)(unsafe.Add(mBase, uint32(v1907)))
	if v1889 <= v1908 {
		goto L421
	} else {
		goto L422
	}
L421:
	;
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(v1907)+4))
	if v1889 <= v1910 {
		v2315 = v1891
		goto L418
	} else {
		goto L424
	}
L422:
	;
	goto L423
L423:
	;
	v2310 = F_slice_del(m, l0)
	mBase = m.M
	v2311 = m.ExcPending
	if v2311 != 0 {
		goto L5
	} else {
		goto L488
	}
L424:
	;
	v1912 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1925 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1927 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L428
L425:
	;
	v2304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2304 + (v1904 - v1912)
	goto L423
L426:
	;
	if v2042 != 0 {
		goto L425
	} else {
		goto L444
	}
L427:
	;
	v2042 = v2035
	goto L426
L428:
	;
	if v1925 <= v1926 {
		v2035 = int32(-1)
		goto L427
	} else {
		goto L430
	}
L429:
	;
	v2035 = int32(0)
	goto L427
L430:
	;
	v1943 = int32(1)
	v1944 = v1925 - v1943
	v1946 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1927+v1944))))
	v1948 = v1946 & int32(255)
	if base.B2i32(v1944 == v1926)|base.B2i32(int32(0) <= v1946) != 0 {
		v2006 = v1948
		v2010 = v1943
		goto L431
	} else {
		goto L432
	}
L431:
	;
	if int32(121) < v2006 {
		goto L439
	} else {
		goto L440
	}
L432:
	;
	v1955 = v1948 & int32(63)
	v1957 = v1925 - int32(2)
	v1959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1927+v1957))))
	v1961 = v1959 << (uint(int32(6)) % 32)
	if base.B2i32(v1957 != v1926)&base.B2i32(base.Ui32(v1959) < base.Ui32(int32(192))) == int32(0) {
		goto L433
	} else {
		goto L434
	}
L433:
	;
	v2006 = v1961&int32(1984) | v1955
	v2010 = int32(2)
	goto L431
L434:
	;
	goto L435
L435:
	;
	v1974 = v1961&int32(4032) | v1955
	v1976 = v1925 - int32(3)
	v1978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1927+v1976))))
	if base.B2i32(v1976 != v1926)&base.B2i32(base.Ui32(v1978) < base.Ui32(int32(224))) == int32(0) {
		goto L436
	} else {
		goto L437
	}
L436:
	;
	v2006 = v1978<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_3) | v1974
	v2010 = int32(3)
	goto L431
L437:
	;
	goto L438
L438:
	;
	v1996 = int32(4)
	v1998 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1925+v1927-v1996))))
	v2006 = v1978<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_9) | v1998&int32(7)<<(uint(int32(18))%32) | v1974
	v2010 = v1996
	goto L431
L439:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1925 - v2010
	goto L443
L440:
	;
	v2012 = v2006 - int32(89)
	if v2012 < int32(0) {
		goto L439
	} else {
		goto L441
	}
L441:
	;
	v2018 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2012)>>(uint(int32(3))%32)))+uint32(_c_F_porter_UTF_8_stem[1]))))
	if int32(base.Ui32(v2018)>>(uint(v2012&int32(7))%32))&int32(1) == int32(0) {
		goto L439
	} else {
		goto L442
	}
L442:
	;
	v2042 = v2010
	goto L426
L443:
	;
	goto L429
L444:
	;
	v2055 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2056 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2057 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L447
L445:
	;
	if v2171 != 0 {
		goto L425
	} else {
		goto L468
	}
L446:
	;
	v2171 = v2164
	goto L445
L447:
	;
	if v2055 <= v2056 {
		v2164 = int32(-1)
		goto L446
	} else {
		goto L449
	}
L448:
	;
	v2164 = int32(0)
	goto L446
L449:
	;
	v2073 = int32(1)
	v2074 = v2055 - v2073
	v2076 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2057+v2074))))
	v2078 = v2076 & int32(255)
	if base.B2i32(v2074 == v2056)|base.B2i32(int32(0) <= v2076) != 0 {
		v2136 = v2078
		v2140 = v2073
		goto L450
	} else {
		goto L451
	}
L450:
	;
	if int32(121) < v2136 {
		goto L458
	} else {
		goto L459
	}
L451:
	;
	v2085 = v2078 & int32(63)
	v2087 = v2055 - int32(2)
	v2089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2057+v2087))))
	v2091 = v2089 << (uint(int32(6)) % 32)
	if base.B2i32(v2087 != v2056)&base.B2i32(base.Ui32(v2089) < base.Ui32(int32(192))) == int32(0) {
		goto L452
	} else {
		goto L453
	}
L452:
	;
	v2136 = v2091&int32(1984) | v2085
	v2140 = int32(2)
	goto L450
L453:
	;
	goto L454
L454:
	;
	v2104 = v2091&int32(4032) | v2085
	v2106 = v2055 - int32(3)
	v2108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2057+v2106))))
	if base.B2i32(v2106 != v2056)&base.B2i32(base.Ui32(v2108) < base.Ui32(int32(224))) == int32(0) {
		goto L455
	} else {
		goto L456
	}
L455:
	;
	v2136 = v2108<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_3) | v2104
	v2140 = int32(3)
	goto L450
L456:
	;
	goto L457
L457:
	;
	v2126 = int32(4)
	v2128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2055+v2057-v2126))))
	v2136 = v2108<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_9) | v2128&int32(7)<<(uint(int32(18))%32) | v2104
	v2140 = v2126
	goto L450
L458:
	;
	v2171 = v2140
	goto L445
L459:
	;
	goto L460
L460:
	;
	v2142 = v2136 - int32(97)
	if v2142 < int32(0) {
		goto L461
	} else {
		goto L462
	}
L461:
	;
	v2171 = v2140
	goto L445
L462:
	;
	goto L463
L463:
	;
	v2148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2142)>>(uint(int32(3))%32)))+uint32(_c_F_porter_UTF_8_stem[0]))))
	if int32(base.Ui32(v2148)>>(uint(v2142&int32(7))%32))&int32(1) == int32(0) {
		goto L464
	} else {
		goto L465
	}
L464:
	;
	v2171 = v2140
	goto L445
L465:
	;
	goto L466
L466:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2055 - v2140
	goto L467
L467:
	;
	goto L448
L468:
	;
	v2184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2186 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L471
L469:
	;
	if v2301 == int32(0) {
		v2315 = v1891
		goto L418
	} else {
		goto L487
	}
L470:
	;
	v2301 = v2294
	goto L469
L471:
	;
	if v2184 <= v2185 {
		v2294 = int32(-1)
		goto L470
	} else {
		goto L473
	}
L472:
	;
	v2294 = int32(0)
	goto L470
L473:
	;
	v2202 = int32(1)
	v2203 = v2184 - v2202
	v2205 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2186+v2203))))
	v2207 = v2205 & int32(255)
	if base.B2i32(v2203 == v2185)|base.B2i32(int32(0) <= v2205) != 0 {
		v2265 = v2207
		v2269 = v2202
		goto L474
	} else {
		goto L475
	}
L474:
	;
	if int32(121) < v2265 {
		goto L482
	} else {
		goto L483
	}
L475:
	;
	v2214 = v2207 & int32(63)
	v2216 = v2184 - int32(2)
	v2218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2186+v2216))))
	v2220 = v2218 << (uint(int32(6)) % 32)
	if base.B2i32(v2216 != v2185)&base.B2i32(base.Ui32(v2218) < base.Ui32(int32(192))) == int32(0) {
		goto L476
	} else {
		goto L477
	}
L476:
	;
	v2265 = v2220&int32(1984) | v2214
	v2269 = int32(2)
	goto L474
L477:
	;
	goto L478
L478:
	;
	v2233 = v2220&int32(4032) | v2214
	v2235 = v2184 - int32(3)
	v2237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2186+v2235))))
	if base.B2i32(v2235 != v2185)&base.B2i32(base.Ui32(v2237) < base.Ui32(int32(224))) == int32(0) {
		goto L479
	} else {
		goto L480
	}
L479:
	;
	v2265 = v2237<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_3) | v2233
	v2269 = int32(3)
	goto L474
L480:
	;
	goto L481
L481:
	;
	v2255 = int32(4)
	v2257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2184+v2186-v2255))))
	v2265 = v2237<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_9) | v2257&int32(7)<<(uint(int32(18))%32) | v2233
	v2269 = v2255
	goto L474
L482:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2184 - v2269
	goto L486
L483:
	;
	v2271 = v2265 - int32(97)
	if v2271 < int32(0) {
		goto L482
	} else {
		goto L484
	}
L484:
	;
	v2277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2271)>>(uint(int32(3))%32)))+uint32(_c_F_porter_UTF_8_stem[0]))))
	if int32(base.Ui32(v2277)>>(uint(v2271&int32(7))%32))&int32(1) == int32(0) {
		goto L482
	} else {
		goto L485
	}
L485:
	;
	v2301 = v2269
	goto L469
L486:
	;
	goto L472
L487:
	;
	goto L425
L488:
	;
	if int32(0) <= v2310 {
		goto L489
	} else {
		goto L490
	}
L489:
	;
	v2314 = int32(1)
	goto L491
L490:
	;
	v2314 = v2310
	goto L491
L491:
	;
	v2315 = v2314
	goto L418
L492:
	;
	v2321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2321
	v2323 = F_r_Step_5b(m, l0)
	mBase = m.M
	v2324 = m.ExcPending
	if v2324 != 0 {
		goto L5
	} else {
		goto L493
	}
L493:
	;
	if v2323 < int32(0) {
		v2428 = v2323
		goto L1
	} else {
		goto L494
	}
L494:
	;
	v2327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2327
	v2329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2330 = *(*int32)(unsafe.Add(mBase, uint32(v2329)+8))
	if v2330 == int32(0) {
		goto L495
	} else {
		goto L496
	}
L495:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2327
	v2428 = int32(1)
	goto L1
L496:
	;
	goto L497
L497:
	;
	v2339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2341 = v2339
	goto L499
L498:
	;
	v2428 = v2415
	goto L1
L499:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2341
	v2347 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2348 != v2341 {
		goto L502
	} else {
		goto L503
	}
L500:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2341
	v2410 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2341 + v2410
	v2415 = F_slice_from_s(m, l0, v2410, int32(_a_F_porter_UTF_8_stem_35))
	mBase = m.M
	v2416 = m.ExcPending
	if v2416 != 0 {
		goto L5
	} else {
		goto L527
	}
L501:
	;
	goto L500
L502:
	;
	v2351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2341+v2347))))
	if v2351 == int32(89) {
		goto L501
	} else {
		goto L505
	}
L503:
	;
	goto L504
L504:
	;
	goto L508
L505:
	;
	goto L504
L506:
	;
	if v2405 < int32(0) {
		goto L495
	} else {
		goto L526
	}
L508:
	;
	goto L509
L509:
	;
	goto L510
L510:
	;
	v2360 = v2341
	v2362 = int32(1)
	goto L513
L512:
	;
	v2405 = v2390
	goto L506
L513:
	;
	if v2348 <= v2360 {
		goto L515
	} else {
		goto L516
	}
L514:
	;
	goto L512
L515:
	;
	v2405 = int32(-1)
	goto L506
L516:
	;
	goto L517
L517:
	;
	v2367 = v2360 + int32(1)
	v2369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2347+v2360))))
	if base.Ui32(v2369) < base.Ui32(int32(192)) {
		v2390 = v2367
		goto L518
	} else {
		goto L519
	}
L518:
	;
	v2391 = int32(1)
	if v2391 < v2362 {
		v2360 = v2390
		v2362 = v2362 - v2391
		goto L513
	} else {
		goto L525
	}
L519:
	;
	if v2348 <= v2367 {
		v2390 = v2367
		goto L518
	} else {
		goto L520
	}
L520:
	;
	v2376 = v2367
	goto L521
L521:
	;
	v2379 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2347+v2376))))
	if int32(-65) < v2379 {
		v2390 = v2376
		goto L518
	} else {
		goto L523
	}
L522:
	;
	v2390 = v2348
	goto L518
L523:
	;
	v2383 = v2376 + int32(1)
	if v2383 != v2348 {
		v2376 = v2383
		goto L521
	} else {
		goto L524
	}
L524:
	;
	goto L522
L525:
	;
	goto L514
L526:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2405
	v2341 = v2405
	goto L499
L527:
	;
	if int32(0) <= v2415 {
		goto L497
	} else {
		goto L528
	}
L528:
	;
	goto L498
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
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_hash[0]))
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v14 = v11 - int32(2)
	if base.B2i32(base.Ui32(int32(8)) < base.Ui32(v14))|base.B2i32(int32(base.Ui32(int32(487))>>(uint(v14)%32))&int32(1) == int32(0)) != 0 {
		v31 = int32(0)
	} else {
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v14<<(uint(int32(2))%32))+uint32(_c_F_prefix_init[0])))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
		v31 = v30
	}
	if int32(32) < v31 {
		v76 = int32(-12)
		m.G0 = v8 + int32(48)
		return v76
	} else {
		v35 = v31 + int32(2)
		v38 = F_pullf_read_max(m, l2, v35, v8+int32(44), v8)
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return int32(0)
		} else {
			if v38 < int32(0) {
				v76 = v38
				m.G0 = v8 + int32(48)
				return v76
			} else {
				if v38 != v35 {
					F_px_debug(m, int32(_a_F_prefix_init_0), int32(0))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						v71 = int32(-100)
						base.MemoryFill(m, v8, int32(0), int32(34))
						v76 = v71
						m.G0 = v8 + int32(48)
						return v76
					}
				} else {
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
					v51 = v50 + v31
					v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51-int32(2)))))
					v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
					if v54 == v55 {
						v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51-int32(1)))))
						v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
						if v60 == v61 {
							v71 = int32(0)
							base.MemoryFill(m, v8, int32(0), int32(34))
							v76 = v71
							m.G0 = v8 + int32(48)
							return v76
						} else {
							v64 = int32(0)
							F_px_debug(m, int32(_a_F_prefix_init_1), v64)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = int32(1)
								v71 = v64
								base.MemoryFill(m, v8, int32(0), int32(34))
								v76 = v71
								m.G0 = v8 + int32(48)
								return v76
							}
						}
					} else {
						v64 = int32(0)
						F_px_debug(m, int32(_a_F_prefix_init_1), v64)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = int32(1)
							v71 = v64
							base.MemoryFill(m, v8, int32(0), int32(34))
							v76 = v71
							m.G0 = v8 + int32(48)
							return v76
						}
					}
				}
			}
		}
	}
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
	var v17 int32
	_ = v17
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
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
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
	var v92 int32
	_ = v92
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
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
	var v161 int32
	_ = v161
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
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	v3 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if l1 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v196
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v10 <= int32(0) {
		v196 = v3
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
	v17 = v3
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
	v196 = v32
	goto L1
L8:
	;
	return int32(0)
L9:
	;
	v32 = F_lappend(m, v17, v28)
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
		v17 = v32
		goto L6
	} else {
		goto L11
	}
L11:
	;
	goto L7
L12:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
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
	v183 = v9 + int32(100)
	goto L12
L16:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v9)+100))
	if v107 == int32(0) {
		v196 = v92
		goto L1
	} else {
		goto L35
	}
L17:
	;
	v45 = v3
	v47 = v3
	goto L20
L18:
	;
	goto L19
L19:
	;
	v183 = v9 + int32(100)
	goto L12
L20:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v9)+100))
	if v50 == int32(0) {
		v92 = v45
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v92 != 0 {
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
		v92 = v45
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v47<<(uint(int32(2))%32))))
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
	v83 = F_lappend(m, v45, v74)
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
	v92 = v45
	goto L22
L32:
	;
	v86 = v47 + int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v86 < v87 {
		v45 = v83
		v47 = v86
		goto L20
	} else {
		goto L33
	}
L33:
	;
	v92 = v83
	goto L22
L34:
	;
	goto L19
L35:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v110 <= int32(0) {
		v196 = v92
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v116 = int32(0)
	v119 = v92
	goto L37
L37:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v124+v116<<(uint(int32(2))%32))))
	v129 = int32(0)
	if v119 == v129 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v196 = v175
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
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	if v135 <= int32(0) {
		v161 = v129
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v167 = v161
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
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v119)+12))
	v144 = int32(0)
	goto L48
L48:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v142+v144<<(uint(int32(2))%32))))
	v153 = base.B2i32(v152 == v128)
	if v152 == v128 {
		v161 = v153
		goto L43
	} else {
		goto L50
	}
L49:
	;
	v161 = v153
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
		v183 = v9 + int32(100)
		goto L12
	} else {
		goto L55
	}
L53:
	;
	v175 = v119
	goto L54
L54:
	;
	v177 = v116 + int32(1)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v177 < v178 {
		v116 = v177
		v119 = v175
		goto L37
	} else {
		goto L57
	}
L55:
	;
	v173 = F_lappend(m, v119, v128)
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
	v196 = v191
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	v23 = v12
	v26 = int32(0)
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
	if v23 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	F_get_rule_expr(m, v33, l1, int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	v51 = int32(0)
	goto L13
L13:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v28+v26<<(uint(int32(2))%32))))
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
	v39 = m.ExcPending
	if v39 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v41 = v23 + int32(4)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if base.Ui32(v41) < base.Ui32(v44+v45<<(uint(int32(2))%32)) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v50 = v41
	goto L18
L17:
	;
	v50 = int32(0)
	goto L18
L18:
	;
	v51 = v50
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
	v64 = v26 + int32(1)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v64 < v65 {
		v23 = v51
		v26 = v64
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v131 int32
	_ = v131
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
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
	var v522 int32
	_ = v522
	var v528 int32
	_ = v528
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v562 int32
	_ = v562
	var v585 int64
	_ = v585
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v625 int64
	_ = v625
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int64
	_ = v654
	var v660 int32
	_ = v660
	var v686 int64
	_ = v686
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v696 int64
	_ = v696
	var v700 int32
	_ = v700
	var v739 int64
	_ = v739
	var v743 int32
	_ = v743
	var v769 int64
	_ = v769
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v779 int64
	_ = v779
	var v783 int32
	_ = v783
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v817 int64
	_ = v817
	var v821 int64
	_ = v821
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int64
	_ = v835
	var v836 int32
	_ = v836
	var v842 int64
	_ = v842
	var v843 int32
	_ = v843
	var v848 int32
	_ = v848
	var v849 int64
	_ = v849
	var v850 int64
	_ = v850
	var v856 int32
	_ = v856
	var v860 int64
	_ = v860
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v881 int32
	_ = v881
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v918 int32
	_ = v918
	var v922 int64
	_ = v922
	var v929 int32
	_ = v929
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v959 int64
	_ = v959
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v968 int32
	_ = v968
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v979 int32
	_ = v979
	var v988 int32
	_ = v988
	var v992 int32
	_ = v992
	var v1009 int32
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1035 int32
	_ = v1035
	var v1040 int32
	_ = v1040
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1069 int32
	_ = v1069
	var v1073 int32
	_ = v1073
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1093 int32
	_ = v1093
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1125 float64
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1130 int32
	_ = v1130
	var v1143 int32
	_ = v1143
	var v1167 int32
	_ = v1167
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1188 int32
	_ = v1188
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1247 int32
	_ = v1247
	var v1260 int32
	_ = v1260
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1273 int32
	_ = v1273
	var v1279 int32
	_ = v1279
	var v1281 int32
	_ = v1281
	var v1286 int32
	_ = v1286
	var v1288 int32
	_ = v1288
	var v1292 int32
	_ = v1292
	var v1294 int32
	_ = v1294
	var v1299 int32
	_ = v1299
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1387 int32
	_ = v1387
	var v1443 int32
	_ = v1443
	v8 = int32(0)
	v29 = m.G0
	v31 = v29 + int32(-64)
	m.G0 = v31
	*(*int32)(unsafe.Add(mBase, uint32(v31)+60)) = l1
	v39 = v29 + int32(-24)
	v41 = l1
	v54 = v8
	v59 = v8
	goto L5
L1:
	;
	m.G0 = v31 - int32(-64)
	return v1443
L2:
	;
	v1443 = int32(-1)
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_printf_core[0])) = v1387
	goto L2
L4:
	;
	v1387 = int32(61)
	goto L3
L5:
	;
	v70 = v41
	v76 = int32(0)
	v83 = v54
	v88 = v59
	goto L7
L6:
	;
	v1443 = int32(0)
	goto L1
L7:
	;
	if v83^int32(2147483647) < v76 {
		goto L4
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	v100 = v76 + v83
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v101 != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	goto L8
L11:
	;
	v1279 = v1268 - v1264
	if v1279 < v1260 {
		goto L289
	} else {
		goto L290
	}
L12:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+39)) = uint8(v1247)
	v1260 = int32(1)
	v1263 = v603
	v1264 = v29 + int32(-25)
	v1267 = v607
	v1268 = v39
	v1273 = v608
	goto L11
L13:
	;
	v1387 = int32(28)
	goto L3
L14:
	;
	v109 = v70
	v113 = v101
	goto L17
L15:
	;
	goto L16
L16:
	;
	if l0 != 0 {
		v1443 = v100
		goto L1
	} else {
		goto L273
	}
L17:
	;
	v131 = v113 & int32(255)
	if v131 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v1130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+1)))
	v109 = v109 + int32(1)
	v113 = v1130
	goto L17
L20:
	;
	v202 = v181 - v70
	v204 = v100 ^ int32(2147483647)
	if v204 < v202 {
		goto L4
	} else {
		goto L31
	}
L21:
	;
	v175 = v109
	v181 = v109
	goto L20
L22:
	;
	goto L23
L23:
	;
	if v131 != int32(37) {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v143 = v109
	v147 = v109
	goto L25
L25:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+1)))
	if v164 != int32(37) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v175 = v171
	v181 = v168
	goto L20
L27:
	;
	v175 = v147
	v181 = v143
	goto L20
L28:
	;
	goto L29
L29:
	;
	v168 = v143 + int32(1)
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+2)))
	v171 = v147 + int32(2)
	if v169 == int32(37) {
		v143 = v168
		v147 = v171
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
	F_out(m, l0, v70, v202)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	if v202 != 0 {
		v70 = v175
		v76 = v202
		v83 = v100
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
	*(*int32)(unsafe.Add(mBase, uint32(v31)+60)) = v175
	v212 = v175 + int32(1)
	v213 = int32(-1)
	v214 = int32(*(*int8)(unsafe.Add(mBase, uint32(v175)+1)))
	v216 = v214 - int32(48)
	if base.Ui32(int32(9)) < base.Ui32(v216) {
		v225 = v212
		v226 = v213
		v227 = v88
		goto L38
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+60)) = v225
	v229 = int32(0)
	v230 = int32(*(*int8)(unsafe.Add(mBase, uint32(v225))))
	v232 = v230 - int32(32)
	if base.Ui32(int32(31)) < base.Ui32(v232) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+2)))
	if v219 != int32(36) {
		v225 = v212
		v226 = v213
		v227 = v88
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v225 = v175 + int32(3)
	v226 = v216
	v227 = int32(1)
	goto L38
L41:
	;
	if v293 == int32(42) {
		goto L51
	} else {
		goto L52
	}
L42:
	;
	v292 = v225
	v293 = v230
	v294 = v229
	goto L41
L43:
	;
	goto L44
L44:
	;
	v236 = int32(1) << (uint(v232) % 32)
	if v236&int32(_a_F_printf_core_0) == int32(0) {
		v292 = v225
		v293 = v230
		v294 = v229
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v242 = v236
	v248 = v225
	v253 = v229
	goto L46
L46:
	;
	v270 = v248 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+60)) = v270
	v272 = v242 | v253
	v273 = int32(*(*int8)(unsafe.Add(mBase, uint32(v248)+1)))
	v274 = int32(32)
	v275 = v273 - v274
	if base.Ui32(v274) <= base.Ui32(v275) {
		v292 = v270
		v293 = v273
		v294 = v272
		goto L41
	} else {
		goto L48
	}
L47:
	;
	v292 = v270
	v293 = v273
	v294 = v272
	goto L41
L48:
	;
	v279 = int32(1) << (uint(v275) % 32)
	if v279&int32(_a_F_printf_core_0) != 0 {
		v242 = v279
		v248 = v270
		v253 = v272
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411))))
	if v418 != int32(46) {
		goto L82
	} else {
		goto L83
	}
L51:
	;
	v312 = int32(*(*int8)(unsafe.Add(mBase, uint32(v292)+1)))
	v314 = v312 - int32(48)
	if base.Ui32(int32(9)) < base.Ui32(v314) {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	goto L53
L53:
	;
	v361 = v29 + int32(-4)
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	v368 = int32(*(*int8)(unsafe.Add(mBase, uint32(v367))))
	v370 = v368 - int32(48)
	if base.Ui32(int32(9)) < base.Ui32(v370) {
		goto L68
	} else {
		goto L69
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+60)) = v349
	if int32(0) <= v351 {
		v411 = v349
		v413 = v294
		v414 = v351
		v415 = v352
		goto L50
	} else {
		goto L66
	}
L55:
	;
	if v227 != 0 {
		goto L13
	} else {
		goto L62
	}
L56:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292)+2)))
	if v317 != int32(36) {
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
	v349 = v292 + int32(3)
	v351 = v332
	v352 = int32(1)
	goto L54
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4+v314<<(uint(int32(2))%32)))) = int32(10)
	v332 = int32(0)
	goto L58
L60:
	;
	goto L61
L61:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l3+v314<<(uint(int32(3))%32))))
	v332 = v331
	goto L58
L62:
	;
	v337 = v292 + int32(1)
	if l0 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+60)) = v337
	v341 = int32(0)
	v411 = v337
	v413 = v294
	v414 = v341
	v415 = v341
	goto L50
L64:
	;
	goto L65
L65:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v343 + int32(4)
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
	v349 = v337
	v351 = v347
	v352 = int32(0)
	goto L54
L66:
	;
	v411 = v349
	v413 = v294 | int32(_a_F_printf_core_1)
	v414 = int32(0) - v351
	v415 = v352
	goto L50
L67:
	;
	if v407 < int32(0) {
		goto L4
	} else {
		goto L80
	}
L68:
	;
	v407 = int32(0)
	goto L67
L69:
	;
	goto L70
L70:
	;
	v375 = v370
	v376 = int32(0)
	v377 = v367
	goto L71
L71:
	;
	if base.Ui32(v376) <= base.Ui32(int32(214748364)) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v407 = v392
	goto L67
L73:
	;
	v385 = v376 * int32(10)
	if base.Ui32(v385^int32(2147483647)) < base.Ui32(v375) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v392 = int32(-1)
	goto L75
L75:
	;
	v394 = v377 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v361))) = v394
	v396 = int32(*(*int8)(unsafe.Add(mBase, uint32(v377)+1)))
	v398 = v396 - int32(48)
	if base.Ui32(v398) < base.Ui32(int32(10)) {
		v375 = v398
		v376 = v392
		v377 = v394
		goto L71
	} else {
		goto L79
	}
L76:
	;
	v390 = int32(-1)
	goto L78
L77:
	;
	v390 = v375 + v385
	goto L78
L78:
	;
	v392 = v390
	goto L75
L79:
	;
	goto L72
L80:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v31)+60))
	v411 = v410
	v413 = v294
	v414 = v407
	v415 = v227
	goto L50
L81:
	;
	v522 = v517
	v528 = int32(0)
	goto L111
L82:
	;
	v517 = v411
	v518 = int32(-1)
	v520 = int32(0)
	goto L81
L83:
	;
	goto L84
L84:
	;
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411)+1)))
	if v422 == int32(42) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v425 = int32(*(*int8)(unsafe.Add(mBase, uint32(v411)+2)))
	v427 = v425 - int32(48)
	if base.Ui32(int32(9)) < base.Ui32(v427) {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	goto L87
L87:
	;
	v464 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+60)) = v411 + v464
	v469 = v29 + int32(-4)
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v476 = int32(*(*int8)(unsafe.Add(mBase, uint32(v475))))
	v478 = v476 - int32(48)
	if base.Ui32(int32(9)) < base.Ui32(v478) {
		goto L99
	} else {
		goto L100
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+60)) = v458
	v517 = v458
	v518 = v460
	v520 = base.B2i32(int32(0) <= v460)
	goto L81
L89:
	;
	if v415 != 0 {
		goto L13
	} else {
		goto L96
	}
L90:
	;
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411)+3)))
	if v430 != int32(36) {
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
	v458 = v411 + int32(4)
	v460 = v447
	goto L88
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4+v427<<(uint(int32(2))%32)))) = int32(10)
	v447 = int32(0)
	goto L92
L94:
	;
	goto L95
L95:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l3+v427<<(uint(int32(3))%32))))
	v447 = v446
	goto L92
L96:
	;
	v449 = v411 + int32(2)
	v450 = int32(0)
	if l0 == v450 {
		v458 = v449
		v460 = v450
		goto L88
	} else {
		goto L97
	}
L97:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v453 + int32(4)
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v453)))
	v458 = v449
	v460 = v457
	goto L88
L98:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v31)+60))
	v517 = v516
	v518 = v515
	v520 = v464
	goto L81
L99:
	;
	v515 = int32(0)
	goto L98
L100:
	;
	goto L101
L101:
	;
	v483 = v478
	v484 = int32(0)
	v485 = v475
	goto L102
L102:
	;
	if base.Ui32(v484) <= base.Ui32(int32(214748364)) {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v515 = v500
	goto L98
L104:
	;
	v493 = v484 * int32(10)
	if base.Ui32(v493^int32(2147483647)) < base.Ui32(v483) {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	v500 = int32(-1)
	goto L106
L106:
	;
	v502 = v485 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v469))) = v502
	v504 = int32(*(*int8)(unsafe.Add(mBase, uint32(v485)+1)))
	v506 = v504 - int32(48)
	if base.Ui32(v506) < base.Ui32(int32(10)) {
		v483 = v506
		v484 = v500
		v485 = v502
		goto L102
	} else {
		goto L110
	}
L107:
	;
	v498 = int32(-1)
	goto L109
L108:
	;
	v498 = v483 + v493
	goto L109
L109:
	;
	v500 = v498
	goto L106
L110:
	;
	goto L103
L111:
	;
	v549 = int32(28)
	v550 = int32(*(*int8)(unsafe.Add(mBase, uint32(v522))))
	if base.Ui32(v550-int32(123)) < base.Ui32(int32(-58)) {
		v1387 = v549
		goto L3
	} else {
		goto L113
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+60)) = v556
	if v562 != int32(27) {
		goto L116
	} else {
		goto L117
	}
L113:
	;
	v555 = int32(1)
	v556 = v522 + v555
	v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528*int32(58)+v550)+uint32(_c_F_printf_core[1]))))
	if base.Ui32((v562-v555)&int32(255)) < base.Ui32(int32(8)) {
		v522 = v556
		v528 = v562
		goto L111
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v599&int32(32) != 0 {
		goto L2
	} else {
		goto L130
	}
L116:
	;
	if v562 == int32(0) {
		v1387 = v549
		goto L3
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	if int32(0) <= v226 {
		v1387 = v549
		goto L3
	} else {
		goto L128
	}
L119:
	;
	if int32(0) <= v226 {
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
	*(*int32)(unsafe.Add(mBase, uint32(l4+v226<<(uint(int32(2))%32)))) = v562
	v41 = v556
	v54 = v100
	v59 = v415
	goto L5
L124:
	;
	goto L125
L125:
	;
	v585 = *(*int64)(unsafe.Add(mBase, uint32(l3+v226<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+48)) = v585
	goto L115
L126:
	;
	F_pop_arg(m, v29+int32(-16), v562, l2, l6)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L35
	} else {
		goto L127
	}
L127:
	;
	goto L115
L128:
	;
	v595 = int32(0)
	if l0 == v595 {
		v70 = v556
		v76 = v595
		v83 = v100
		v88 = v415
		goto L7
	} else {
		goto L129
	}
L129:
	;
	goto L115
L130:
	;
	v603 = v413 & int32(-65537)
	if v413&int32(_a_F_printf_core_1) != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v606 = v603
	goto L133
L132:
	;
	v606 = v413
	goto L133
L133:
	;
	v607 = int32(0)
	v608 = int32(_a_F_printf_core_2)
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522))))
	v610 = base.I32_extend8_s(v609)
	if v609&int32(15) == int32(3) {
		goto L151
	} else {
		goto L152
	}
L134:
	;
	if v520&base.B2i32(v518 < int32(0)) != 0 {
		goto L4
	} else {
		goto L270
	}
L135:
	;
	F_pad(m, l0, int32(32), v414, v1093, v606^int32(_a_F_printf_core_1))
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L35
	} else {
		goto L266
	}
L136:
	;
	v988 = int32(0)
	v992 = v979
	goto L240
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+12)) = int32(0)
	*(*uint32)(unsafe.Add(mBase, uint32(v31)+8)) = uint32(v959)
	v973 = v29 + int32(-56)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+48)) = v973
	v976 = int32(-1)
	v979 = v973
	goto L136
L138:
	;
	if v518 != 0 {
		goto L236
	} else {
		goto L237
	}
L139:
	;
	v959 = *(*int64)(unsafe.Add(mBase, uint32(v31)+48))
	if v959 != int64(0) {
		goto L137
	} else {
		goto L235
	}
L140:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	if v943 != 0 {
		goto L221
	} else {
		goto L222
	}
L141:
	;
	v942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+48)))
	v1247 = v942
	goto L12
L142:
	;
	if v520&base.B2i32(v905 < int32(0)) != 0 {
		goto L4
	} else {
		goto L211
	}
L143:
	;
	if base.Ui64(int64(4294967296)) <= base.Ui64(v835) {
		goto L196
	} else {
		goto L197
	}
L144:
	;
	v817 = *(*int64)(unsafe.Add(mBase, uint32(v31)+48))
	if v817 < int64(0) {
		goto L186
	} else {
		goto L187
	}
L145:
	;
	v739 = *(*int64)(unsafe.Add(mBase, uint32(v31)+48))
	if v739 != int64(0) {
		goto L176
	} else {
		goto L177
	}
L146:
	;
	v654 = *(*int64)(unsafe.Add(mBase, uint32(v31)+48))
	if v654 != int64(0) {
		goto L169
	} else {
		goto L170
	}
L147:
	;
	v644 = int32(8)
	if base.Ui32(v518) <= base.Ui32(v644) {
		goto L166
	} else {
		goto L167
	}
L148:
	;
	v627 = int32(0)
	switch v528 {
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
		v70 = v556
		v76 = v627
		v83 = v100
		v88 = v415
		goto L7
	case 6:
		goto L160
	case 7:
		goto L159
	}
L149:
	;
	v625 = *(*int64)(unsafe.Add(mBase, uint32(v31)+48))
	v834 = v607
	v835 = v625
	v836 = int32(_a_F_printf_core_2)
	goto L143
L150:
	;
	switch v618 - int32(65) {
	case 0, 4, 5, 6:
		goto L134
	case 1, 3:
		v1260 = v518
		v1263 = v606
		v1264 = v70
		v1267 = v607
		v1268 = v39
		v1273 = v608
		goto L11
	case 2:
		goto L139
	default:
		goto L157
	}
L151:
	;
	v617 = v610 & int32(-45)
	goto L153
L152:
	;
	v617 = v610
	goto L153
L153:
	;
	if v528 != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v618 = v617
	goto L156
L155:
	;
	v618 = v610
	goto L156
L156:
	;
	switch v618 - int32(88) {
	case 0, 32:
		v651 = v618
		v652 = v518
		v653 = v606
		goto L146
	case 1, 2, 3, 4, 5, 6, 7, 8, 10, 16, 18, 19, 20, 21, 25, 26, 28, 30, 31:
		v1260 = v518
		v1263 = v606
		v1264 = v70
		v1267 = v607
		v1268 = v39
		v1273 = v608
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
	if v618 == int32(83) {
		goto L138
	} else {
		goto L158
	}
L158:
	;
	v1260 = v518
	v1263 = v606
	v1264 = v70
	v1267 = v607
	v1268 = v39
	v1273 = v608
	goto L11
L159:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v641))) = base.I64_extend_i32_s(v100)
	v70 = v556
	v76 = v627
	v83 = v100
	v88 = v415
	goto L7
L160:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v639))) = v100
	v70 = v556
	v76 = v627
	v83 = v100
	v88 = v415
	goto L7
L161:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v637))) = uint8(v100)
	v70 = v556
	v76 = v627
	v83 = v100
	v88 = v415
	goto L7
L162:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	*(*uint16)(unsafe.Add(mBase, uint32(v635))) = uint16(v100)
	v70 = v556
	v76 = v627
	v83 = v100
	v88 = v415
	goto L7
L163:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v632))) = base.I64_extend_i32_s(v100)
	v70 = v556
	v76 = v627
	v83 = v100
	v88 = v415
	goto L7
L164:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v630))) = v100
	v70 = v556
	v76 = v627
	v83 = v100
	v88 = v415
	goto L7
L165:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v628))) = v100
	v70 = v556
	v76 = v627
	v83 = v100
	v88 = v415
	goto L7
L166:
	;
	v647 = v644
	goto L168
L167:
	;
	v647 = v518
	goto L168
L168:
	;
	v651 = int32(120)
	v652 = v647
	v653 = v606 | int32(8)
	goto L146
L169:
	;
	v660 = v39
	v686 = v654
	goto L172
L170:
	;
	v700 = v39
	goto L171
L171:
	;
	if base.B2i32(v653&int32(8) == int32(0))|base.B2i32(v654 == int64(0)) != 0 {
		v905 = v652
		v908 = v653
		v909 = v700
		v912 = v607
		v918 = v608
		v922 = v654
		goto L142
	} else {
		goto L175
	}
L172:
	;
	v688 = v660 - int32(1)
	v692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v686)&int32(15))+uint32(_c_F_printf_core[2]))))
	v693 = v692 | v651&int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v688))) = uint8(v693)
	v696 = int64(base.Ui64(v686) >> (uint(int64(4)) % 64))
	if v696 != int64(0) {
		v660 = v688
		v686 = v696
		goto L172
	} else {
		goto L174
	}
L173:
	;
	v700 = v688
	goto L171
L174:
	;
	goto L173
L175:
	;
	v905 = v652
	v908 = v653
	v909 = v700
	v912 = int32(2)
	v918 = int32(base.Ui32(v651)>>(uint(int32(4))%32)) + int32(_a_F_printf_core_2)
	v922 = v654
	goto L142
L176:
	;
	v743 = v39
	v769 = v739
	goto L179
L177:
	;
	v783 = v39
	goto L178
L178:
	;
	if v606&int32(8) == int32(0) {
		v905 = v518
		v908 = v606
		v909 = v783
		v912 = v607
		v918 = v608
		v922 = v739
		goto L142
	} else {
		goto L182
	}
L179:
	;
	v771 = v743 - int32(1)
	v776 = base.I32_wrap_i64(v769)&int32(7) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v771))) = uint8(v776)
	v779 = int64(base.Ui64(v769) >> (uint(int64(3)) % 64))
	if v779 != int64(0) {
		v743 = v771
		v769 = v779
		goto L179
	} else {
		goto L181
	}
L180:
	;
	v783 = v771
	goto L178
L181:
	;
	goto L180
L182:
	;
	v814 = v29 + int32(-23) - v783
	if v814 < v518 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v816 = v518
	goto L185
L184:
	;
	v816 = v814
	goto L185
L185:
	;
	v905 = v816
	v908 = v606
	v909 = v783
	v912 = v607
	v918 = v608
	v922 = v739
	goto L142
L186:
	;
	v821 = int64(0) - v817
	*(*int64)(unsafe.Add(mBase, uint32(v31)+48)) = v821
	v834 = int32(1)
	v835 = v821
	v836 = int32(_a_F_printf_core_2)
	goto L143
L187:
	;
	goto L188
L188:
	;
	if v606&int32(2048) != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v834 = int32(1)
	v835 = v817
	v836 = int32(_a_F_printf_core_3)
	goto L143
L190:
	;
	goto L191
L191:
	;
	v832 = v606 & int32(1)
	if v832 != 0 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v833 = int32(_a_F_printf_core_4)
	goto L194
L193:
	;
	v833 = int32(_a_F_printf_core_2)
	goto L194
L194:
	;
	v834 = v832
	v835 = v817
	v836 = v833
	goto L143
L195:
	;
	v905 = v518
	v908 = v606
	v909 = v895
	v912 = v834
	v918 = v836
	v922 = v835
	goto L142
L196:
	;
	v842 = v835
	v843 = v39
	goto L199
L197:
	;
	v860 = v835
	v861 = v39
	goto L198
L198:
	;
	v865 = base.I32_wrap_i64(v860)
	if base.Ui64(int64(10)) <= base.Ui64(v860) {
		goto L202
	} else {
		goto L203
	}
L199:
	;
	v848 = v843 - int32(1)
	v849 = int64(10)
	v850 = base.I64_div_u_s(v842, v849)
	v856 = base.I32_wrap_i64(v842-v850*v849) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v848))) = uint8(v856)
	if base.Ui64(int64(42949672959)) < base.Ui64(v842) {
		v842 = v850
		v843 = v848
		goto L199
	} else {
		goto L201
	}
L200:
	;
	v860 = v850
	v861 = v848
	goto L198
L201:
	;
	goto L200
L202:
	;
	v869 = v861
	v870 = v865
	goto L205
L203:
	;
	v886 = v861
	v887 = v865
	goto L204
L204:
	;
	if v887 != 0 {
		goto L208
	} else {
		goto L209
	}
L205:
	;
	v874 = v869 - int32(1)
	v875 = int32(10)
	v876 = base.I32_div_u_s(v870, v875)
	v881 = v870 - v876*v875 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v874))) = uint8(v881)
	if base.Ui32(int32(99)) < base.Ui32(v870) {
		v869 = v874
		v870 = v876
		goto L205
	} else {
		goto L207
	}
L206:
	;
	v886 = v874
	v887 = v876
	goto L204
L207:
	;
	goto L206
L208:
	;
	v891 = v886 - int32(1)
	v893 = v887 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v891))) = uint8(v893)
	v895 = v891
	goto L210
L209:
	;
	v895 = v886
	goto L210
L210:
	;
	goto L195
L211:
	;
	if v520 != 0 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v929 = v908 & int32(-65537)
	goto L214
L213:
	;
	v929 = v908
	goto L214
L214:
	;
	if base.B2i32(v922 != int64(0))|v905 == int32(0) {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v1260 = int32(0)
	v1263 = v929
	v1264 = v39
	v1267 = v912
	v1268 = v39
	v1273 = v918
	goto L11
L216:
	;
	goto L217
L217:
	;
	v939 = base.B2i32(v922 == int64(0)) + (v39 - v909)
	if v939 < v905 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v941 = v905
	goto L220
L219:
	;
	v941 = v939
	goto L220
L220:
	;
	v1260 = v941
	v1263 = v929
	v1264 = v909
	v1267 = v912
	v1268 = v39
	v1273 = v918
	goto L11
L221:
	;
	v945 = v943
	goto L223
L222:
	;
	v945 = int32(_a_F_printf_core_5)
	goto L223
L223:
	;
	v946 = int32(2147483647)
	if base.Ui32(v946) <= base.Ui32(v518) {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v949 = v946
	goto L226
L225:
	;
	v949 = v518
	goto L226
L226:
	;
	v952 = F_memchr(m, v945, int32(0), v949)
	mBase = m.M
	if v952 != 0 {
		goto L228
	} else {
		goto L229
	}
L227:
	;
	v955 = v954 + v945
	if int32(0) <= v518 {
		goto L231
	} else {
		goto L232
	}
L228:
	;
	v954 = v952 - v945
	goto L230
L229:
	;
	v954 = v949
	goto L230
L230:
	;
	goto L227
L231:
	;
	v1260 = v954
	v1263 = v603
	v1264 = v945
	v1267 = v607
	v1268 = v955
	v1273 = v608
	goto L11
L232:
	;
	goto L233
L233:
	;
	v958 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v955))))
	if v958 != 0 {
		goto L4
	} else {
		goto L234
	}
L234:
	;
	v1260 = v954
	v1263 = v603
	v1264 = v945
	v1267 = v607
	v1268 = v955
	v1273 = v608
	goto L11
L235:
	;
	v1247 = int32(0)
	goto L12
L236:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	v976 = v518
	v979 = v963
	goto L136
L237:
	;
	goto L238
L238:
	;
	v964 = int32(0)
	F_pad(m, l0, int32(32), v414, v964, v606)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L35
	} else {
		goto L239
	}
L239:
	;
	v1093 = v964
	goto L135
L240:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v992)))
	if v1009 == int32(0) {
		v1027 = v988
		goto L242
	} else {
		goto L243
	}
L241:
	;
	if v1027 < int32(0) {
		v1387 = int32(61)
		goto L3
	} else {
		goto L251
	}
L242:
	;
	goto L241
L243:
	;
	v1013 = v29 + int32(-60)
	if v1013 == int32(0) {
		goto L245
	} else {
		goto L246
	}
L244:
	;
	if v1018 < int32(0) {
		goto L2
	} else {
		goto L248
	}
L245:
	;
	v1018 = int32(0)
	goto L244
L246:
	;
	goto L247
L247:
	;
	v1017 = F_wcrtomb(m, v1013, v1009)
	mBase = m.M
	v1018 = v1017
	goto L244
L248:
	;
	if base.Ui32(v976-v988) < base.Ui32(v1018) {
		v1027 = v988
		goto L242
	} else {
		goto L249
	}
L249:
	;
	v1025 = v988 + v1018
	if base.Ui32(v1025) < base.Ui32(v976) {
		v988 = v1025
		v992 = v992 + int32(4)
		goto L240
	} else {
		goto L250
	}
L250:
	;
	v1027 = v1025
	goto L242
L251:
	;
	F_pad(m, l0, int32(32), v414, v1027, v606)
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L35
	} else {
		goto L252
	}
L252:
	;
	if v1027 == int32(0) {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v1093 = int32(0)
	goto L135
L254:
	;
	goto L255
L255:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	v1051 = int32(0)
	v1052 = v1040
	goto L256
L256:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v1052)))
	if v1069 == int32(0) {
		v1093 = v1027
		goto L135
	} else {
		goto L258
	}
L257:
	;
	v1093 = v1027
	goto L135
L258:
	;
	v1073 = v29 + int32(-60)
	if v1073 == int32(0) {
		goto L260
	} else {
		goto L261
	}
L259:
	;
	v1079 = v1078 + v1051
	if base.Ui32(v1027) < base.Ui32(v1079) {
		v1093 = v1027
		goto L135
	} else {
		goto L263
	}
L260:
	;
	v1078 = int32(0)
	goto L259
L261:
	;
	goto L262
L262:
	;
	v1077 = F_wcrtomb(m, v1073, v1069)
	mBase = m.M
	v1078 = v1077
	goto L259
L263:
	;
	F_out(m, l0, v1073, v1078)
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L35
	} else {
		goto L264
	}
L264:
	;
	if base.Ui32(v1079) < base.Ui32(v1027) {
		v1051 = v1079
		v1052 = v1052 + int32(4)
		goto L256
	} else {
		goto L265
	}
L265:
	;
	goto L257
L266:
	;
	if v1093 < v414 {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v1120 = v414
	goto L269
L268:
	;
	v1120 = v1093
	goto L269
L269:
	;
	v70 = v556
	v76 = v1120
	v83 = v100
	v88 = v415
	goto L7
L270:
	;
	v1125 = *(*float64)(unsafe.Add(mBase, uint32(v31)+48))
	v1126 = m.T0[l5].(func(*base.Module, int32, float64, int32, int32, int32, int32, int32) int32)(m, l0, v1125, v414, v518, v606, v618, v528)
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L35
	} else {
		goto L271
	}
L271:
	;
	if int32(0) <= v1126 {
		v70 = v556
		v76 = v1126
		v83 = v100
		v88 = v415
		goto L7
	} else {
		goto L272
	}
L272:
	;
	v1387 = int32(61)
	goto L3
L273:
	;
	if v88 == int32(0) {
		goto L10
	} else {
		goto L274
	}
L274:
	;
	v1143 = int32(1)
	goto L275
L275:
	;
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(l4+v1143<<(uint(int32(2))%32))))
	if v1167 != 0 {
		goto L277
	} else {
		goto L278
	}
L276:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v1143) {
		goto L282
	} else {
		goto L283
	}
L277:
	;
	F_pop_arg(m, l3+v1143<<(uint(int32(3))%32), v1167, l2, l6)
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L35
	} else {
		goto L280
	}
L278:
	;
	goto L279
L279:
	;
	goto L276
L280:
	;
	v1173 = int32(1)
	v1175 = v1143 + v1173
	if v1175 != int32(10) {
		v1143 = v1175
		goto L275
	} else {
		goto L281
	}
L281:
	;
	v1443 = v1173
	goto L1
L282:
	;
	v1443 = int32(1)
	goto L1
L283:
	;
	goto L284
L284:
	;
	v1188 = v1143
	goto L285
L285:
	;
	v1212 = *(*int32)(unsafe.Add(mBase, uint32(l4+v1188<<(uint(int32(2))%32))))
	if v1212 != 0 {
		goto L13
	} else {
		goto L287
	}
L286:
	;
	v1443 = v1213
	goto L1
L287:
	;
	v1213 = int32(1)
	v1215 = v1188 + v1213
	if v1215 != int32(10) {
		v1188 = v1215
		goto L285
	} else {
		goto L288
	}
L288:
	;
	goto L286
L289:
	;
	v1281 = v1260
	goto L291
L290:
	;
	v1281 = v1279
	goto L291
L291:
	;
	if v1267^int32(2147483647) < v1281 {
		goto L4
	} else {
		goto L292
	}
L292:
	;
	v1286 = v1281 + v1267
	if v1286 < v414 {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	v1288 = v414
	goto L295
L294:
	;
	v1288 = v1286
	goto L295
L295:
	;
	if base.Ui32(v204) < base.Ui32(v1288) {
		v1387 = int32(61)
		goto L3
	} else {
		goto L296
	}
L296:
	;
	F_pad(m, l0, int32(32), v1288, v1286, v1263)
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L35
	} else {
		goto L297
	}
L297:
	;
	F_out(m, l0, v1273, v1267)
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L35
	} else {
		goto L298
	}
L298:
	;
	F_pad(m, l0, int32(48), v1288, v1286, v1263^int32(_a_F_printf_core_6))
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L35
	} else {
		goto L299
	}
L299:
	;
	F_pad(m, l0, int32(48), v1281, v1279, int32(0))
	mBase = m.M
	v1303 = m.ExcPending
	if v1303 != 0 {
		goto L35
	} else {
		goto L300
	}
L300:
	;
	F_out(m, l0, v1264, v1279)
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L35
	} else {
		goto L301
	}
L301:
	;
	F_pad(m, l0, int32(32), v1288, v1286, v1263^int32(_a_F_printf_core_1))
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
		goto L35
	} else {
		goto L302
	}
L302:
	;
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v31)+60))
	v70 = v1311
	v76 = v1288
	v83 = v100
	v88 = v415
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
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
	v249 = m.ExcPending
	if v249 != 0 {
		goto L18
	} else {
		goto L63
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L18
	} else {
		goto L59
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L18
	} else {
		goto L55
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L18
	} else {
		goto L51
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L18
	} else {
		goto L46
	}
L6:
	;
	if l2 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L7:
	;
	v15 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v19 = int32(_a_F_process_owned_by_0)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_process_owned_by[0])))
	if base.B2i32(v22 == v15)|base.B2i32(v22 != v25) != 0 {
		v43 = v22
		v44 = v25
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L9
L9:
	;
	v70 = F_list_copy_head(m, l1, v12-int32(1))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L18
	} else {
		goto L24
	}
L10:
	;
	if v43-v44 == int32(0) {
		v121 = v15
		v123 = int32(0)
		goto L6
	} else {
		goto L17
	}
L11:
	;
	goto L10
L12:
	;
	v28 = v18
	v29 = v19
	goto L13
L13:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	if v33 == int32(0) {
		v43 = v33
		v44 = v32
		goto L11
	} else {
		goto L15
	}
L14:
	;
	v43 = v33
	v44 = v32
	goto L11
L15:
	;
	v36 = int32(1)
	if v33 == v32 {
		v28 = v28 + v36
		v29 = v29 + v36
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return
L19:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	F_errmsg(m, int32(_a_F_process_owned_by_1), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	F_errhint(m, int32(_a_F_process_owned_by_2), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_process_owned_by_3), int32(1611), int32(_a_F_process_owned_by_4))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v72+v73<<(uint(int32(2))%32)-int32(4))))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v81 = F_makeRangeVarFromNameList(m, v70)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	v84 = F_relation_openrv(m, v81, int32(1))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L18
	} else {
		goto L26
	}
L26:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)+48))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+119)))
	v89 = v87 - int32(102)
	v94 = int32(1)
	v98 = (v89<<(uint(int32(7))%32) | int32(base.Ui32(v89&int32(254))>>(uint(v94)%32))) & int32(255)
	if base.B2i32(base.Ui32(int32(8)) < base.Ui32(v98))|base.B2i32(v94<<(uint(v98)%32)&int32(353) == int32(0)) != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+80))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v86)+80))
	if v109 != v110 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108)+68))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v86)+68))
	if v112 != v113 {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v84)+56))
	v116 = F_get_attnum(m, v115, v80)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L18
	} else {
		goto L30
	}
L30:
	;
	if v116 == int32(0) {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v121 = v84
	v123 = v116
	goto L6
L32:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v132 = F_sequenceIsOwned(m, v126, int32(105), v8+int32(-12), v8+int32(-24))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L18
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v134 = int32(1259)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if l2 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	if v132 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v139 = int32(105)
	goto L39
L38:
	;
	v139 = int32(97)
	goto L39
L39:
	;
	v140 = F_deleteDependencyRecordsForClass(m, v134, v135, v134, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L18
	} else {
		goto L40
	}
L40:
	;
	if v121 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v142 = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v142
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v121)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v142
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v149
	F_recordDependencyOn(m, v8+int32(-24), v8+int32(-12), v139)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L18
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	m.G0 = v10 - int32(-64)
	return
L44:
	;
	F_relation_close(m, v121, int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L18
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L18
	} else {
		goto L47
	}
L47:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v84)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v174 + int32(4)
	F_errmsg(m, int32(_a_F_process_owned_by_5), v8+int32(-48))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L18
	} else {
		goto L48
	}
L48:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v84)+48))
	v184 = int32(*(*int8)(unsafe.Add(mBase, uint32(v183)+119)))
	F_errdetail_relkind_not_supported(m, v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L18
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_process_owned_by_3), int32(1638), int32(_a_F_process_owned_by_4))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L18
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L18
	} else {
		goto L52
	}
L52:
	;
	F_errmsg(m, int32(_a_F_process_owned_by_6), int32(0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L18
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_process_owned_by_3), int32(1644), int32(_a_F_process_owned_by_4))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L18
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L18
	} else {
		goto L56
	}
L56:
	;
	F_errmsg(m, int32(_a_F_process_owned_by_7), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L18
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_process_owned_by_3), int32(1648), int32(_a_F_process_owned_by_4))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L18
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L18
	} else {
		goto L60
	}
L60:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v84)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v231 + int32(4)
	F_errmsg(m, int32(_a_F_process_owned_by_8), v8+int32(-32))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L18
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_process_owned_by_3), int32(1656), int32(_a_F_process_owned_by_4))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L18
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L18
	} else {
		goto L64
	}
L64:
	;
	F_errmsg(m, int32(_a_F_process_owned_by_9), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L18
	} else {
		goto L65
	}
L65:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
	v259 = F_get_rel_name(m, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L18
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v257 + int32(4)
	F_errdetail(m, int32(_a_F_process_owned_by_10), v10)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L18
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_process_owned_by_3), int32(1673), int32(_a_F_process_owned_by_4))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L18
	} else {
		goto L68
	}
L68:
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
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13977(m, l0, l1, int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_pull_vars_of_level(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13978(m, l0, l1, int32(898), int32(0))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
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
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	v3 = int32(0)
	if l0 == v3 {
		v39 = v3
		return v39
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v7 != int32(319) {
			if v7 == int32(67) {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v20 + int32(1)
				v26 = F_query_tree_walker_impl(m, l0, int32(898), l1, int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v30 - int32(1)
					return v26
				}
			} else {
				if v7 != int32(6) {
					v36 = F_expression_tree_walker_impl(m, l0, int32(898), l1)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v39 = v36
						return v39
					}
				} else {
					v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					if v14 != v15 {
						v39 = v3
						return v39
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v42 = F_lappend(m, v41, l0)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v42
							return int32(0)
						}
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v17 != v18 {
				v39 = v3
				return v39
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v42 = F_lappend(m, v41, l0)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v42
					return int32(0)
				}
			}
		}
	}
}
func F_pullf_read_max(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
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
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v14 == int32(0) {
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
	v35 = l0
	v36 = v14
	goto L3
L3:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	if l1 < v40 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v27 == int32(0) {
		v22 = v25
		goto L4
	} else {
		goto L6
	}
L5:
	;
	v35 = v25
	v36 = v27
	goto L3
L6:
	;
	goto L5
L7:
	;
	m.G0 = v11 + int32(16)
	return v120
L8:
	;
	v42 = l1
	goto L10
L9:
	;
	v42 = v40
	goto L10
L10:
	;
	if v40 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v43 = v42
	goto L13
L12:
	;
	v43 = l1
	goto L13
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v45 = m.T0[v36].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v38, v39, v43, l2, v44, v40)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	if base.B2i32(v45 <= int32(0))|base.B2i32(l1 == v45) != 0 {
		v120 = v45
		goto L7
	} else {
		goto L16
	}
L16:
	;
	if v45 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	base.MemoryCopy(m, l3, v53, v45)
	goto L19
L18:
	;
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = l3
	v56 = l1 - v45
	if v56 <= int32(0) {
		v120 = v45
		goto L7
	} else {
		goto L20
	}
L20:
	;
	v60 = v56
	v63 = v45
	goto L21
L21:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v68 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v120 = v112
	goto L7
L23:
	;
	v76 = l0
	goto L26
L24:
	;
	v89 = l0
	v90 = v68
	goto L25
L25:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	if v60 < v94 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if v81 == int32(0) {
		v76 = v79
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v89 = v79
	v90 = v81
	goto L25
L28:
	;
	goto L27
L29:
	;
	v96 = v60
	goto L31
L30:
	;
	v96 = v94
	goto L31
L31:
	;
	if v94 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v97 = v96
	goto L34
L33:
	;
	v97 = v60
	goto L34
L34:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v101 = m.T0[v90].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v92, v93, v97, v11+int32(12), v100, v94)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L14
	} else {
		goto L35
	}
L35:
	;
	if v101 < int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if v63 != 0 {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	goto L38
L38:
	;
	if v101 == int32(0) {
		v120 = v63
		goto L7
	} else {
		goto L43
	}
L39:
	;
	v120 = v101
	goto L7
L40:
	;
	base.MemoryFill(m, l3, int32(0), v63)
	goto L42
L41:
	;
	goto L42
L42:
	;
	goto L39
L43:
	;
	if v101 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	base.MemoryCopy(m, l3+v63, v110, v101)
	goto L46
L45:
	;
	goto L46
L46:
	;
	v112 = v101 + v63
	v113 = v60 - v101
	if int32(0) < v113 {
		v60 = v113
		v63 = v112
		goto L21
	} else {
		goto L47
	}
L47:
	;
	goto L22
}
func F_push_child_plan(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	base.MemoryCopy(m, l2, l0, int32(80))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v8 = F_lcons(m, v6, v7)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v8
		F_set_deparse_plan(m, l0, l1)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
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
				if v10 != 0 {
					base.MemoryFill(m, v8, int32(0), v10)
				} else {
				}
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				F_pfree(m, v12)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					base.MemoryFill(m, l0, int32(0), int32(24))
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
				base.MemoryFill(m, l0, int32(0), int32(24))
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
			if v10 != 0 {
				base.MemoryFill(m, v8, int32(0), v10)
			} else {
			}
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			F_pfree(m, v12)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				base.MemoryFill(m, l0, int32(0), int32(24))
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
			base.MemoryFill(m, l0, int32(0), int32(24))
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
