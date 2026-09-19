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
	var v94 int64
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v120 int64
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, _c_F_xid8in[0])) = v2
	v20 = F_strtox_2(m, v7, v11+int32(44), v2, int64(-1))
	mBase = m.M
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_xid8in[0]))
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
						v120 = v55
						m.G0 = v11 + int32(48)
						v124 = F_Int64GetDatum(m, v120)
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return int32(0)
						} else {
							return v124
						}
					} else {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(_a_F_xid8in_0)
							*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v7
							F_errmsg(m, int32(_a_F_xid8in_1), v11+int32(16))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								F_errsave_finish(m, v8, int32(_a_F_xid8in_2), int32(1010), int32(_a_F_xid8in_3))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int32(0)
								} else {
									v120 = v55
									m.G0 = v11 + int32(48)
									v124 = F_Int64GetDatum(m, v120)
									mBase = m.M
									v125 = m.ExcPending
									if v125 != 0 {
										return int32(0)
									} else {
										return v124
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
					if base.B2i32(base.Ui32(v82-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v82 == int32(32)) != 0 {
						v80 = v80 + int32(1)
						continue
					} else {
						break
					}
					break
				}
				if v82 == int32(0) {
					v120 = v20
					m.G0 = v11 + int32(48)
					v124 = F_Int64GetDatum(m, v120)
					mBase = m.M
					v125 = m.ExcPending
					if v125 != 0 {
						return int32(0)
					} else {
						return v124
					}
				} else {
					v94 = int64(0)
					v95 = F_errsave_start(m, v8)
					mBase = m.M
					v96 = m.ExcPending
					if v96 != 0 {
						return int32(0)
					} else {
						if v95 == int32(0) {
							v120 = v94
							m.G0 = v11 + int32(48)
							v124 = F_Int64GetDatum(m, v120)
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
								return int32(0)
							} else {
								return v124
							}
						} else {
							F_errcode(m, int32(33685634))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v7
								*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = int32(_a_F_xid8in_0)
								F_errmsg(m, int32(_a_F_xid8in_4), v11+int32(32))
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return int32(0)
								} else {
									F_errsave_finish(m, v8, int32(_a_F_xid8in_2), int32(1026), int32(_a_F_xid8in_3))
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return int32(0)
									} else {
										v120 = v94
										m.G0 = v11 + int32(48)
										v124 = F_Int64GetDatum(m, v120)
										mBase = m.M
										v125 = m.ExcPending
										if v125 != 0 {
											return int32(0)
										} else {
											return v124
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
					v120 = v32
					m.G0 = v11 + int32(48)
					v124 = F_Int64GetDatum(m, v120)
					mBase = m.M
					v125 = m.ExcPending
					if v125 != 0 {
						return int32(0)
					} else {
						return v124
					}
				} else {
					F_errcode(m, int32(33685634))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v7
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_xid8in_0)
						F_errmsg(m, int32(_a_F_xid8in_4), v11)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, v8, int32(_a_F_xid8in_2), int32(1004), int32(_a_F_xid8in_3))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								v120 = v32
								m.G0 = v11 + int32(48)
								v124 = F_Int64GetDatum(m, v120)
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return int32(0)
								} else {
									return v124
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
				v120 = v32
				m.G0 = v11 + int32(48)
				v124 = F_Int64GetDatum(m, v120)
				mBase = m.M
				v125 = m.ExcPending
				if v125 != 0 {
					return int32(0)
				} else {
					return v124
				}
			} else {
				F_errcode(m, int32(33685634))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v7
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_xid8in_0)
					F_errmsg(m, int32(_a_F_xid8in_4), v11)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						F_errsave_finish(m, v8, int32(_a_F_xid8in_2), int32(1004), int32(_a_F_xid8in_3))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							v120 = v32
							m.G0 = v11 + int32(48)
							v124 = F_Int64GetDatum(m, v120)
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
								return int32(0)
							} else {
								return v124
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
			F_errmsg(m, int32(_a_F_xmlparse_0), int32(0))
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				F_errdetail(m, int32(_a_F_xmlparse_1), int32(0))
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_xmlparse_2), int32(1004), int32(_a_F_xmlparse_3))
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
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13842(m, l0, int32(_a_F_xmltext_0), int32(542))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_xpath_exists(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13842(m, l0, int32(_a_F_xpath_exists_0), int32(_a_F_xpath_exists_1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
