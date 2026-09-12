package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_xid8in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v20 int64
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int64
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v91 int64
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v119 int64
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, _consts[140])) = v2
	v20 = F_strtox_2(m, v7, v11+int32(44), v2, int64(-1))
	mBase = m.M
	v22 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v22 != 0 {
		v26 = base.B2i32(v22 != int32(68))
	} else {
		v26 = int32(0)
	}
	if v26 == int32(0) {
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
		if v29 != v7 {
			if v22 == int32(68) {
				v55 = int64(0)
				v56 = F_errsave_start(m, v8)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					if v56 == int32(0) {
						v119 = v55
						m.G0 = v11 + int32(48)
						v123 = F_Int64GetDatum(m, v119)
						mBase = m.M
						v124 = m.ExcPending
						if v124 != 0 {
							return int32(0)
						} else {
							return v123
						}
					} else {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(543294)
							*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v7
							F_errmsg(m, int32(188171), v11+int32(16))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								F_errsave_finish(m, v8, int32(488479), int32(1010), int32(227058))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int32(0)
								} else {
									v119 = v55
									m.G0 = v11 + int32(48)
									v123 = F_Int64GetDatum(m, v119)
									mBase = m.M
									v124 = m.ExcPending
									if v124 != 0 {
										return int32(0)
									} else {
										return v123
									}
								}
							}
						}
					}
				}
			} else {
				v80 = v29
				for {
					v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
					if base.Ui32(v82-int32(9)) < base.Ui32(int32(5)) {
						v80 = v80 + int32(1)
						continue
					} else {
					}
					if v82 == int32(32) {
						v80 = v80 + int32(1)
						continue
					} else {
						break
					}
					break
				}
				if v82 == int32(0) {
					v119 = v20
					m.G0 = v11 + int32(48)
					v123 = F_Int64GetDatum(m, v119)
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return int32(0)
					} else {
						return v123
					}
				} else {
					v91 = int64(0)
					v92 = F_errsave_start(m, v8)
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return int32(0)
					} else {
						if v92 == int32(0) {
							v119 = v91
							m.G0 = v11 + int32(48)
							v123 = F_Int64GetDatum(m, v119)
							mBase = m.M
							v124 = m.ExcPending
							if v124 != 0 {
								return int32(0)
							} else {
								return v123
							}
						} else {
							F_errcode(m, int32(33685634))
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v7
								*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = int32(543294)
								F_errmsg(m, int32(703551), v11+int32(32))
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return int32(0)
								} else {
									F_errsave_finish(m, v8, int32(488479), int32(1026), int32(227058))
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
										return int32(0)
									} else {
										v119 = v91
										m.G0 = v11 + int32(48)
										v123 = F_Int64GetDatum(m, v119)
										mBase = m.M
										v124 = m.ExcPending
										if v124 != 0 {
											return int32(0)
										} else {
											return v123
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v32 = int64(0)
			v33 = F_errsave_start(m, v8)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				if v33 == int32(0) {
					v119 = v32
					m.G0 = v11 + int32(48)
					v123 = F_Int64GetDatum(m, v119)
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return int32(0)
					} else {
						return v123
					}
				} else {
					F_errcode(m, int32(33685634))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v7
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(543294)
						F_errmsg(m, int32(703551), v11)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, v8, int32(488479), int32(1004), int32(227058))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								v119 = v32
								m.G0 = v11 + int32(48)
								v123 = F_Int64GetDatum(m, v119)
								mBase = m.M
								v124 = m.ExcPending
								if v124 != 0 {
									return int32(0)
								} else {
									return v123
								}
							}
						}
					}
				}
			}
		}
	} else {
		v32 = int64(0)
		v33 = F_errsave_start(m, v8)
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			if v33 == int32(0) {
				v119 = v32
				m.G0 = v11 + int32(48)
				v123 = F_Int64GetDatum(m, v119)
				mBase = m.M
				v124 = m.ExcPending
				if v124 != 0 {
					return int32(0)
				} else {
					return v123
				}
			} else {
				F_errcode(m, int32(33685634))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v7
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(543294)
					F_errmsg(m, int32(703551), v11)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						F_errsave_finish(m, v8, int32(488479), int32(1004), int32(227058))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							v119 = v32
							m.G0 = v11 + int32(48)
							v123 = F_Int64GetDatum(m, v119)
							mBase = m.M
							v124 = m.ExcPending
							if v124 != 0 {
								return int32(0)
							} else {
								return v123
							}
						}
					}
				}
			}
		}
	}
}
func F_xmlparse(m *base.Module) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	F_errstart_cold(m, int32(21), int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			F_errmsg(m, int32(359098), int32(0))
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				F_errdetail(m, int32(557450), int32(0))
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(491910), int32(1004), int32(356873))
					v22 = m.ExcPending
					if v22 != 0 {
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
func F_xmltext(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	F_errstart_cold(m, int32(21), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_errmsg(m, int32(359098), int32(0))
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errdetail(m, int32(557450), int32(0))
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(491910), int32(542), int32(62233))
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
	}
}
func F_xpath_exists(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	F_errstart_cold(m, int32(21), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_errmsg(m, int32(359098), int32(0))
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errdetail(m, int32(557450), int32(0))
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(491910), int32(4579), int32(114440))
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
	}
}
