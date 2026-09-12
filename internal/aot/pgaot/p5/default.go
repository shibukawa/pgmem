package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetDefaultTablespace(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	if l0 == int32(116) {
		F_PrepareTempTablespaces(m)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _consts[310]))
			if int32(0) < v12 {
				v15 = int32(4405472)
				v17 = *(*int32)(unsafe.Add(mBase, _consts[311]))
				v19 = v17 + int32(1)
				if v19 < v12 {
					v22 = v19
				} else {
					v22 = int32(0)
				}
				*(*int32)(unsafe.Add(mBase, _consts[311])) = v22
				v25 = *(*int32)(unsafe.Add(mBase, _consts[312]))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v25+v22<<(uint(int32(2))%32))))
				v32 = v29
			} else {
				v32 = int32(0)
			}
			v68 = v32
			return v68
		}
	} else {
		v33 = int32(0)
		v35 = *(*int32)(unsafe.Add(mBase, _consts[313]))
		if v35 == v33 {
			v68 = v33
			return v68
		} else {
			v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
			if v38 == int32(0) {
				v68 = v33
				return v68
			} else {
				v42 = F_get_tablespace_oid(m, v35, int32(1))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, _consts[108]))
					if v42 != v46 {
						v48 = v42
					} else {
						v48 = int32(0)
					}
					if l1 == int32(0) {
						v68 = v48
						return v68
					} else {
						if v42 != v46 {
							v68 = v48
							return v68
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(1088))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(143152), int32(0))
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(498930), int32(1178), int32(418767))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
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
func F_findDefaultOnlyColumns(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v107 int32
	_ = v107
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v7 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v107
L2:
	;
	v107 = int32(0)
	goto L1
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v10 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v14 = v2
	v18 = v2
	goto L5
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19+v18<<(uint(int32(2))%32))))
	if v14 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v107 = v90
	goto L1
L7:
	;
	v96 = v18 + int32(1)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v96 < v97 {
		v14 = v90
		v18 = v96
		goto L5
	} else {
		goto L32
	}
L8:
	;
	if v82 == int32(0) {
		goto L2
	} else {
		goto L31
	}
L9:
	;
	if v23 == int32(0) {
		v90 = v14
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	if v23 == int32(0) {
		goto L2
	} else {
		goto L22
	}
L12:
	;
	v26 = int32(0)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v27 <= v26 {
		v82 = v14
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v31 = v14
	v32 = v26
	goto L14
L14:
	;
	v37 = v32 + int32(1)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v32<<(uint(int32(2))%32))))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v43 != int32(57) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v82 = v50
	goto L8
L16:
	;
	v46 = F_bms_del_member(m, v31, v37)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v50 = v31
	goto L18
L18:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v37 < v51 {
		v31 = v50
		v32 = v37
		goto L14
	} else {
		goto L21
	}
L19:
	;
	return int32(0)
L20:
	;
	v50 = v46
	goto L18
L21:
	;
	goto L15
L22:
	;
	v55 = int32(0)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v57 <= v55 {
		v82 = v55
		goto L8
	} else {
		goto L23
	}
L23:
	;
	v61 = v55
	v62 = v55
	goto L24
L24:
	;
	v67 = v62 + int32(1)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68+v62<<(uint(int32(2))%32))))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	if v73 == int32(57) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v82 = v78
	goto L8
L26:
	;
	v76 = F_bms_add_member(m, v61, v67)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L19
	} else {
		goto L29
	}
L27:
	;
	v78 = v61
	goto L28
L28:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v67 < v79 {
		v61 = v78
		v62 = v67
		goto L24
	} else {
		goto L30
	}
L29:
	;
	v78 = v76
	goto L28
L30:
	;
	goto L25
L31:
	;
	v90 = v82
	goto L7
L32:
	;
	goto L6
}
