package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_j2date(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	v8 = l0 + int32(32044)
	v9 = int32(146097)
	v10 = base.I32_div_u_s(v8, v9)
	v11 = int32(3)
	v17 = int32(2)
	v22 = base.I32_div_u_s((v10*int32(1073595727)+v8)<<(uint(v17)%32)|v11, v9)
	v25 = l0 + v10*v11 + v22 + int32(32104)
	v26 = int32(1461)
	v27 = base.I32_div_u_s(v25, v26)
	v30 = v27*int32(-1461) + v25
	v32 = v30 << (uint(v17) % 32)
	if base.Ui32(v26) <= base.Ui32(v32) {
		v38 = base.I32_rem_u_s(v30+int32(305), int32(365))
		v43 = v38
	} else {
		v42 = base.I32_rem_u_s(v30+int32(306), int32(366))
		v43 = v42
	}
	v45 = base.I32_div_u_s(v32, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v45 + v27<<(uint(int32(2))%32) - int32(4800)
	v53 = v43 + int32(123)
	v57 = int32(base.Ui32(v53*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v53 - int32(base.Ui32(v57*int32(7834))>>(uint(int32(8))%32))
	v67 = base.I32_rem_u_s(v57+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v67 + int32(1)
	return
}
func F_j2day(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	v4 = int32(7)
	v5 = base.I32_rem_s(l0+int32(1), v4)
	if v5 < int32(0) {
		v10 = v5 + v4
	} else {
		v10 = v5
	}
	return v10
}
func F_jsonpath_yyensure_buffer_stack(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v4 == int32(0) {
		v8 = F_palloc(m, int32(4))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v8
			if v8 == int32(0) {
				F_yy_fatal_error_5(m, int32(658225))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = int64(4294967296)
				return
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if base.Ui32(v18-int32(1)) <= base.Ui32(v17) {
			v23 = v18 + int32(8)
			v26 = F_repalloc(m, v4, v23<<(uint(int32(2))%32))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v26
				if v26 == int32(0) {
					F_yy_fatal_error_5(m, int32(658225))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v34 = v26 + v31<<(uint(int32(2))%32)
					v35 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v34))) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+24)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+16)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+8)) = v35
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v23
					return
				}
			}
		} else {
			return
		}
	}
}
func F_jsonpath_yyerror(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	if l0 == int32(0) {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
		v19 = F_errsave_start(m, l0)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			if v18 == int32(0) {
				if v19 == int32(0) {
					m.G0 = v9 + int32(32)
					return
				} else {
					F_errcode(m, int32(16801924))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
						F_errmsg(m, int32(64089), v9)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							v48 = int32(378)
							F_errsave_finish(m, l0, int32(309113), v48, int32(206955))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								m.G0 = v9 + int32(32)
								return
							}
						}
					}
				}
			} else {
				if v19 == int32(0) {
					m.G0 = v9 + int32(32)
					return
				} else {
					F_errcode(m, int32(16801924))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v39
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l2
						F_errmsg(m, int32(64117), v9+int32(16))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							v48 = int32(386)
							F_errsave_finish(m, l0, int32(309113), v48, int32(206955))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								m.G0 = v9 + int32(32)
								return
							}
						}
					}
				}
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v13 != int32(447) {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
			v19 = F_errsave_start(m, l0)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				if v18 == int32(0) {
					if v19 == int32(0) {
						m.G0 = v9 + int32(32)
						return
					} else {
						F_errcode(m, int32(16801924))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
							F_errmsg(m, int32(64089), v9)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								v48 = int32(378)
								F_errsave_finish(m, l0, int32(309113), v48, int32(206955))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return
								} else {
									m.G0 = v9 + int32(32)
									return
								}
							}
						}
					}
				} else {
					if v19 == int32(0) {
						m.G0 = v9 + int32(32)
						return
					} else {
						F_errcode(m, int32(16801924))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v39
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l2
							F_errmsg(m, int32(64117), v9+int32(16))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return
							} else {
								v48 = int32(386)
								F_errsave_finish(m, l0, int32(309113), v48, int32(206955))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return
								} else {
									m.G0 = v9 + int32(32)
									return
								}
							}
						}
					}
				}
			}
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
			if v16 != 0 {
				m.G0 = v9 + int32(32)
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
				v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
				v19 = F_errsave_start(m, l0)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					if v18 == int32(0) {
						if v19 == int32(0) {
							m.G0 = v9 + int32(32)
							return
						} else {
							F_errcode(m, int32(16801924))
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
								F_errmsg(m, int32(64089), v9)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return
								} else {
									v48 = int32(378)
									F_errsave_finish(m, l0, int32(309113), v48, int32(206955))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return
									} else {
										m.G0 = v9 + int32(32)
										return
									}
								}
							}
						}
					} else {
						if v19 == int32(0) {
							m.G0 = v9 + int32(32)
							return
						} else {
							F_errcode(m, int32(16801924))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v39
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l2
								F_errmsg(m, int32(64117), v9+int32(16))
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return
								} else {
									v48 = int32(386)
									F_errsave_finish(m, l0, int32(309113), v48, int32(206955))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return
									} else {
										m.G0 = v9 + int32(32)
										return
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
