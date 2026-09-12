package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_seg_lt(m *base.Module, l0 int32) int32 {
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_DirectFunctionCall2Coll(m, int32(6673), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return int32(base.Ui32(v6) >> (uint(int32(31)) % 32))
	}
}
func F_seg_over_left(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float32
	_ = v3
	var v4 int32
	_ = v4
	var v5 float32
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*float32)(unsafe.Add(mBase, uint32(v2)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*float32)(unsafe.Add(mBase, uint32(v4)+4))
	return base.F32_le(v3, v5)
}
func F_seg_yyensure_buffer_stack(m *base.Module, l0 int32) {
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
				F_yy_fatal_error_7(m, int32(717294))
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
					F_yy_fatal_error_7(m, int32(717294))
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
func F_seg_yyerror(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	if l1 == int32(0) {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+80))
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
		v19 = F_errsave_start(m, l1)
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
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						F_errmsg(m, int32(272289), int32(0))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l3
							F_errdetail(m, int32(71020), v9)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								v63 = int32(82)
								F_errsave_finish(m, l1, int32(330407), v63, int32(221714))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return
								} else {
									m.G0 = v9 + int32(32)
									return
								}
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
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						F_errmsg(m, int32(272289), int32(0))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+80))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v51
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l3
							F_errdetail(m, int32(737604), v9+int32(16))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								v63 = int32(90)
								F_errsave_finish(m, l1, int32(330407), v63, int32(221714))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
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
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		if v13 != int32(447) {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+80))
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
			v19 = F_errsave_start(m, l1)
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
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							F_errmsg(m, int32(272289), int32(0))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = l3
								F_errdetail(m, int32(71020), v9)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									v63 = int32(82)
									F_errsave_finish(m, l1, int32(330407), v63, int32(221714))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return
									} else {
										m.G0 = v9 + int32(32)
										return
									}
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
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							F_errmsg(m, int32(272289), int32(0))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+80))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v51
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l3
								F_errdetail(m, int32(737604), v9+int32(16))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									v63 = int32(90)
									F_errsave_finish(m, l1, int32(330407), v63, int32(221714))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
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
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
			if v16 != 0 {
				m.G0 = v9 + int32(32)
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+80))
				v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
				v19 = F_errsave_start(m, l1)
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
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								F_errmsg(m, int32(272289), int32(0))
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = l3
									F_errdetail(m, int32(71020), v9)
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return
									} else {
										v63 = int32(82)
										F_errsave_finish(m, l1, int32(330407), v63, int32(221714))
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return
										} else {
											m.G0 = v9 + int32(32)
											return
										}
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
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								F_errmsg(m, int32(272289), int32(0))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return
								} else {
									v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+80))
									*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v51
									*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l3
									F_errdetail(m, int32(737604), v9+int32(16))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return
									} else {
										v63 = int32(90)
										F_errsave_finish(m, l1, int32(330407), v63, int32(221714))
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
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
}
func F_seg_yyset_lineno(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v4 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v4+v5<<(uint(int32(2))%32))))
		if v9 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l0
			return
		} else {
			F_yy_fatal_error_7(m, int32(237497))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	} else {
		F_yy_fatal_error_7(m, int32(237497))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
