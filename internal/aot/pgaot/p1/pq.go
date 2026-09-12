package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pq_cleanup_redirect_to_shm_mq(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[146])) = v4
	*(*int32)(unsafe.Add(mBase, _consts[461])) = v4
	return
}
func F_pq_getmsgbyte(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v4 <= v3 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(16908800))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(424599), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(518207), int32(404), int32(366825))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
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
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v3 + int32(1)
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v3))))
		return v29
	}
}
func F_pq_gettcpusertimeout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l0 == v2 {
		v22 = v2
	} else {
		v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
		if v10 == int32(1) {
			v22 = v2
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+412))
			if v13 != 0 {
				v22 = v13
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+396))
				if v14 != 0 {
					v22 = v14
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(4)
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(396))))
					v22 = v20
				}
			}
		}
	}
	m.G0 = v6 + int32(16)
	return v22
}
func F_pq_recvbuf(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	v4 = *(*int32)(unsafe.Add(mBase, _consts[453]))
	if int32(0) < v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = *(*int32)(unsafe.Add(mBase, _consts[454]))
	if v4 < v8 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	v172 = *(*int32)(unsafe.Add(mBase, _consts[455]))
	if v172 != 0 {
		goto L54
	} else {
		goto L55
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v166))) = int32(0)
	goto L3
L5:
	;
	v10 = int32(4461008)
	v12 = v4 + v10
	v13 = v8 - v4
	if v10 == v12 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[453])) = int32(0)
	v166 = int32(4460976)
	goto L4
L8:
	;
	*(*int32)(unsafe.Add(mBase, _consts[454])) = v13
	v166 = int32(4460980)
	goto L4
L9:
	;
	goto L8
L10:
	;
	v17 = v10 + v13
	if base.Ui32(v12-v17) <= base.Ui32(int32(0)-v13<<(uint(int32(1))%32)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v24 = F___memcpy(m, v10, v12, v13)
	mBase = m.M
	goto L8
L12:
	;
	goto L13
L13:
	;
	v27 = (v10 ^ v12) & int32(3)
	if base.Ui32(v10) < base.Ui32(v12) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	if v129 == int32(0) {
		goto L9
	} else {
		goto L50
	}
L15:
	;
	if base.Ui32(v13) <= base.Ui32(int32(3)) {
		v128 = v12
		v129 = v13
		v130 = v10
		goto L14
	} else {
		goto L46
	}
L16:
	;
	if v27 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	if v27 != 0 {
		v89 = v13
		goto L29
	} else {
		goto L30
	}
L19:
	;
	v128 = v12
	v129 = v13
	v130 = v10
	goto L14
L20:
	;
	goto L21
L21:
	;
	goto L22
L22:
	;
	goto L15
L29:
	;
	if v89 == int32(0) {
		goto L9
	} else {
		goto L42
	}
L30:
	;
	if v17&int32(3) != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v54 = v13
	goto L34
L32:
	;
	v69 = v13
	goto L33
L33:
	;
	if base.Ui32(v69) <= base.Ui32(int32(3)) {
		v89 = v69
		goto L29
	} else {
		goto L38
	}
L34:
	;
	if v54 == int32(0) {
		goto L9
	} else {
		goto L36
	}
L35:
	;
	v69 = v60
	goto L33
L36:
	;
	v60 = v54 - int32(1)
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v60))))
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+uint32(_consts[456]))) = uint8(v63)
	if (v10+v60)&int32(3) != 0 {
		v54 = v60
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v76 = v69
	goto L39
L39:
	;
	v80 = v76 - int32(4)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v12+v80)))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+uint32(_consts[456]))) = v83
	if base.Ui32(int32(3)) < base.Ui32(v80) {
		v76 = v80
		goto L39
	} else {
		goto L41
	}
L40:
	;
	v89 = v80
	goto L29
L41:
	;
	goto L40
L42:
	;
	v96 = v89
	goto L43
L43:
	;
	v100 = v96 - int32(1)
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v100))))
	*(*uint8)(unsafe.Add(mBase, uint32(v100)+uint32(_consts[456]))) = uint8(v103)
	if v100 != 0 {
		v96 = v100
		goto L43
	} else {
		goto L45
	}
L44:
	;
	goto L9
L45:
	;
	goto L44
L46:
	;
	v113 = v12
	v114 = v13
	v115 = v10
	goto L47
L47:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	*(*int32)(unsafe.Add(mBase, uint32(v115))) = v117
	v119 = int32(4)
	v120 = v113 + v119
	v122 = v115 + v119
	v124 = v114 - v119
	if base.Ui32(int32(3)) < base.Ui32(v124) {
		v113 = v120
		v114 = v124
		v115 = v122
		goto L47
	} else {
		goto L49
	}
L48:
	;
	v128 = v120
	v129 = v124
	v130 = v122
	goto L14
L49:
	;
	goto L48
L50:
	;
	v135 = v128
	v136 = v129
	v137 = v130
	goto L51
L51:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	*(*uint8)(unsafe.Add(mBase, uint32(v137))) = uint8(v139)
	v141 = int32(1)
	v146 = v136 - v141
	if v146 != 0 {
		v135 = v135 + v141
		v136 = v146
		v137 = v137 + v141
		goto L51
	} else {
		goto L53
	}
L52:
	;
	goto L9
L53:
	;
	goto L52
L54:
	;
	v173 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v172)+4)) = uint8(v173)
	v175 = int32(-1)
	goto L59
L55:
	;
	goto L56
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L61
	} else {
		goto L72
	}
L57:
	;
	return v228
L58:
	;
	if v189 == int32(0) {
		v228 = v175
		goto L57
	} else {
		goto L71
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	v182 = *(*int32)(unsafe.Add(mBase, _consts[455]))
	v184 = *(*int32)(unsafe.Add(mBase, _consts[454]))
	v189 = F_secure_read(m, v182, v184+int32(4461008), int32(8192)-v184)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	if v196 == int32(0) {
		v228 = v175
		goto L57
	} else {
		goto L65
	}
L61:
	;
	return int32(0)
L62:
	;
	if int32(0) <= v189 {
		goto L58
	} else {
		goto L63
	}
L63:
	;
	v196 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v196 == int32(27) {
		goto L59
	} else {
		goto L64
	}
L64:
	;
	goto L60
L65:
	;
	v203 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L61
	} else {
		goto L66
	}
L66:
	;
	if v203 == int32(0) {
		v228 = v175
		goto L57
	} else {
		goto L67
	}
L67:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L61
	} else {
		goto L68
	}
L68:
	;
	F_errmsg(m, int32(307484), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L61
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(522971), int32(942), int32(355977))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L61
	} else {
		goto L70
	}
L70:
	;
	return int32(-1)
L71:
	;
	v223 = int32(4460976)
	v225 = *(*int32)(unsafe.Add(mBase, _consts[454]))
	*(*int32)(unsafe.Add(mBase, _consts[454])) = v225 + v189
	v228 = int32(0)
	goto L57
L72:
	;
	F_errcode(m, int32(50332160))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L61
	} else {
		goto L73
	}
L73:
	;
	F_errmsg(m, int32(268422), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L61
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(522971), int32(886), int32(353309))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L61
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pq_sendbytes(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v5 int32
	_ = v5
	F_appendBinaryStringInfo(m, l0, l1, l2)
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_pq_sendcountedtext(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v64 int32
	_ = v64
	v5 = F_pg_server_to_client(m, l1, l2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		if l1 != v5 {
			v8 = F_strlen(m, v5)
			mBase = m.M
			F_enlargeStringInfo(m, l0, int32(4))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v15 = int32(24)
				v17 = int32(65280)
				v19 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v12+v13))) = v8<<(uint(v15)%32) | v8&v17<<(uint(v19)%32) | (int32(base.Ui32(v8)>>(uint(v19)%32))&v17 | int32(base.Ui32(v8)>>(uint(v15)%32)))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12 + int32(4)
				F_appendBinaryStringInfoNT(m, l0, v5, v8)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					F_pfree(m, v5)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			F_enlargeStringInfo(m, l0, int32(4))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v44 = int32(24)
				v46 = int32(65280)
				v48 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v41+v42))) = l2<<(uint(v44)%32) | l2&v46<<(uint(v48)%32) | (int32(base.Ui32(l2)>>(uint(v48)%32))&v46 | int32(base.Ui32(l2)>>(uint(v44)%32)))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v41 + int32(4)
				F_appendBinaryStringInfoNT(m, l0, l1, l2)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_pq_sendtext(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v5 = F_pg_server_to_client(m, l1, l2)
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		if l1 != v5 {
			v8 = F_strlen(m, v5)
			F_appendBinaryStringInfo(m, l0, v5, v8)
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				F_pfree(m, v5)
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			F_appendBinaryStringInfo(m, l0, l1, l2)
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_pq_setkeepalivesidle(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l0
	if l1 == int32(0) {
	} else {
		v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)))
		if v11 == int32(1) {
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+400))
			if l0 == v14 {
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+384))
				if int32(0) < v16 {
					if l0 == int32(0) {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+384))
						*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v47
					} else {
					}
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+400)) = v50
				} else {
					v19 = int32(0)
					v21 = m.G0
					v23 = v21 - int32(16)
					m.G0 = v23
					if l1 == v19 {
						v39 = v19
					} else {
						v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)))
						if v27 == int32(1) {
							v39 = v19
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+400))
							if v30 != 0 {
								v39 = v30
							} else {
								v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+384))
								if v31 != 0 {
									v39 = v31
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = int32(4)
									v37 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(384))))
									v39 = v37
								}
							}
						}
					}
					m.G0 = v23 + int32(16)
					if int32(0) <= v39 {
						if l0 == int32(0) {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+384))
							*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v47
						} else {
						}
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+400)) = v50
					} else {
					}
				}
			}
		}
	}
	m.G0 = v6 + int32(16)
	return
}
