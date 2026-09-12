package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pq_beginmessage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	F_initStringInfo(m, l0)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
		return
	}
}
func F_pq_begintypsend(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
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
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	F_initStringInfo(m, l0)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v5 <= v6+int32(1) {
			F_appendStringInfoChar(m, l0, int32(0))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v26 <= v27+int32(1) {
					F_appendStringInfoChar(m, l0, int32(0))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if v47 <= v48+int32(1) {
							F_appendStringInfoChar(m, l0, int32(0))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								if v68 <= v69+int32(1) {
									F_appendStringInfoChar(m, l0, int32(0))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return
									} else {
										return
									}
								} else {
									v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v78 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v76+v69))) = uint8(v78)
									v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v82 = v80 + int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v82
									v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*uint8)(unsafe.Add(mBase, uint32(v84+v82))) = uint8(v78)
									return
								}
							}
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v57 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v55+v48))) = uint8(v57)
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v61 = v59 + int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v61
							v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*uint8)(unsafe.Add(mBase, uint32(v63+v61))) = uint8(v57)
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v68 <= v69+int32(1) {
								F_appendStringInfoChar(m, l0, int32(0))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return
								} else {
									return
								}
							} else {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v78 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v76+v69))) = uint8(v78)
								v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v82 = v80 + int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v82
								v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*uint8)(unsafe.Add(mBase, uint32(v84+v82))) = uint8(v78)
								return
							}
						}
					}
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v36 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v34+v27))) = uint8(v36)
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v40 = v38 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v40
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*uint8)(unsafe.Add(mBase, uint32(v42+v40))) = uint8(v36)
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v47 <= v48+int32(1) {
						F_appendStringInfoChar(m, l0, int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v68 <= v69+int32(1) {
								F_appendStringInfoChar(m, l0, int32(0))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return
								} else {
									return
								}
							} else {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v78 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v76+v69))) = uint8(v78)
								v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v82 = v80 + int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v82
								v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*uint8)(unsafe.Add(mBase, uint32(v84+v82))) = uint8(v78)
								return
							}
						}
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v57 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v55+v48))) = uint8(v57)
						v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v61 = v59 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v61
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*uint8)(unsafe.Add(mBase, uint32(v63+v61))) = uint8(v57)
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if v68 <= v69+int32(1) {
							F_appendStringInfoChar(m, l0, int32(0))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return
							} else {
								return
							}
						} else {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v78 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v76+v69))) = uint8(v78)
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v82 = v80 + int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v82
							v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*uint8)(unsafe.Add(mBase, uint32(v84+v82))) = uint8(v78)
							return
						}
					}
				}
			}
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v15 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v13+v6))) = uint8(v15)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v19 = v17 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v19
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*uint8)(unsafe.Add(mBase, uint32(v21+v19))) = uint8(v15)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v26 <= v27+int32(1) {
				F_appendStringInfoChar(m, l0, int32(0))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v47 <= v48+int32(1) {
						F_appendStringInfoChar(m, l0, int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v68 <= v69+int32(1) {
								F_appendStringInfoChar(m, l0, int32(0))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return
								} else {
									return
								}
							} else {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v78 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v76+v69))) = uint8(v78)
								v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v82 = v80 + int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v82
								v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*uint8)(unsafe.Add(mBase, uint32(v84+v82))) = uint8(v78)
								return
							}
						}
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v57 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v55+v48))) = uint8(v57)
						v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v61 = v59 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v61
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*uint8)(unsafe.Add(mBase, uint32(v63+v61))) = uint8(v57)
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if v68 <= v69+int32(1) {
							F_appendStringInfoChar(m, l0, int32(0))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return
							} else {
								return
							}
						} else {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v78 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v76+v69))) = uint8(v78)
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v82 = v80 + int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v82
							v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*uint8)(unsafe.Add(mBase, uint32(v84+v82))) = uint8(v78)
							return
						}
					}
				}
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v36 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v34+v27))) = uint8(v36)
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v40 = v38 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v40
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*uint8)(unsafe.Add(mBase, uint32(v42+v40))) = uint8(v36)
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v47 <= v48+int32(1) {
					F_appendStringInfoChar(m, l0, int32(0))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if v68 <= v69+int32(1) {
							F_appendStringInfoChar(m, l0, int32(0))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return
							} else {
								return
							}
						} else {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v78 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v76+v69))) = uint8(v78)
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v82 = v80 + int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v82
							v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*uint8)(unsafe.Add(mBase, uint32(v84+v82))) = uint8(v78)
							return
						}
					}
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v57 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v55+v48))) = uint8(v57)
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v61 = v59 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v61
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*uint8)(unsafe.Add(mBase, uint32(v63+v61))) = uint8(v57)
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v68 <= v69+int32(1) {
						F_appendStringInfoChar(m, l0, int32(0))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return
						} else {
							return
						}
					} else {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v78 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v76+v69))) = uint8(v78)
						v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v82 = v80 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v82
						v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*uint8)(unsafe.Add(mBase, uint32(v84+v82))) = uint8(v78)
						return
					}
				}
			}
		}
	}
}
func F_pq_buffer_remaining_data(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, _consts[372]))
	v4 = *(*int32)(unsafe.Add(mBase, _consts[370]))
	return v2 - v4
}
func F_pq_endmessage(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
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
	v2 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+12)))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v8 = m.T0[v7].(func(*base.Module, int32, int32, int32) int32)(m, v2, v3, v4)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		F_pfree(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
			return
		}
	}
}
func F_pq_endtypsend(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3))) = v4 << (uint(int32(2)) % 32)
	return v3
}
func F_pq_getbyte(m *base.Module) int32 {
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
	var v22 int32
	_ = v22
	goto L2
L1:
	;
	*(*int32)(unsafe.Add(mBase, _consts[370])) = v4 + int32(1)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+uint32(_consts[371]))))
	return v22
L2:
	;
	v4 = *(*int32)(unsafe.Add(mBase, _consts[370]))
	v6 = *(*int32)(unsafe.Add(mBase, _consts[372]))
	if v4 < v6 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	return int32(-1)
L4:
	;
	v8 = F_pq_recvbuf(m)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	if v8 == int32(0) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	goto L3
}
func F_pq_getkeepalivescount(m *base.Module, l0 int32) int32 {
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
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+408))
			if v13 != 0 {
				v22 = v13
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
				if v14 != 0 {
					v22 = v14
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(4)
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(392))))
					v22 = v20
				}
			}
		}
	}
	m.G0 = v6 + int32(16)
	return v22
}
func F_pq_getmsgend(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v2 != v3 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			F_errcode(m, int32(16908800))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				F_errmsg(m, int32(106514), int32(0))
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					F_errfinish(m, int32(476777), int32(640), int32(412190))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		return
	}
}
func F_pq_getmsgrawstring(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = v4 + v5
	if v6&int32(3) == int32(0) {
		v30 = v6
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v64 = v63 + v4
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v65 <= v64 {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	v63 = v55 - v6
	goto L1
L3:
	;
	v34 = v30
	goto L12
L4:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v14 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v63 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v19 = v6
	goto L8
L8:
	;
	v23 = v19 + int32(1)
	if v23&int32(3) == int32(0) {
		v30 = v23
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v55 = v23
	goto L2
L10:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v28 != 0 {
		v19 = v23
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v43 = int32(-2139062144)
	if (int32(16843008)-v40|v40)&v43 == v43 {
		v34 = v34 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v49 = v34
	goto L15
L14:
	;
	goto L13
L15:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v53 != 0 {
		v49 = v49 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v55 = v49
	goto L2
L17:
	;
	goto L16
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v64 + int32(1)
	return v6
L21:
	;
	return int32(0)
L22:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	F_errmsg(m, int32(390440), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(476777), int32(624), int32(317627))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pq_sendfloat8(m *base.Module, l0 int32, l1 float64) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	F_enlargeStringInfo(m, l0, int32(8))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v11 = base.I64_reinterpret_f64(l1)
		v12 = int64(56)
		v14 = int64(65280)
		v16 = int64(40)
		v19 = int64(16711680)
		v21 = int64(24)
		v23 = int64(4278190080)
		v25 = int64(8)
		*(*int64)(unsafe.Add(mBase, uint32(v8+v9))) = v11<<(uint(v12)%64) | v11&v14<<(uint(v16)%64) | (v11&v19<<(uint(v21)%64) | v11&v23<<(uint(v25)%64)) | (int64(base.Ui64(v11)>>(uint(v25)%64))&v23 | int64(base.Ui64(v11)>>(uint(v21)%64))&v19 | (int64(base.Ui64(v11)>>(uint(v16)%64))&v14 | int64(base.Ui64(v11)>>(uint(v12)%64))))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v8 + int32(8)
		return
	}
}
func F_pq_set_parallel_leader(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _consts[92])) = l1
	*(*int32)(unsafe.Add(mBase, _consts[93])) = l0
	return
}
func F_pq_setkeepalivesinterval(m *base.Module, l0 int32, l1 int32) {
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
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+404))
			if l0 == v14 {
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+388))
				if int32(0) < v16 {
					if l0 == int32(0) {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+388))
						*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v47
					} else {
					}
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+404)) = v50
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
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+404))
							if v30 != 0 {
								v39 = v30
							} else {
								v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+388))
								if v31 != 0 {
									v39 = v31
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = int32(4)
									v37 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(388))))
									v39 = v37
								}
							}
						}
					}
					m.G0 = v23 + int32(16)
					if int32(0) <= v39 {
						if l0 == int32(0) {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+388))
							*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v47
						} else {
						}
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+404)) = v50
					} else {
					}
				}
			}
		}
	}
	m.G0 = v6 + int32(16)
	return
}
