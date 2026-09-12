package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_extract_timetz(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_timetz_part_common(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_timetz_in(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v77 int64
	_ = v77
	var v82 int32
	_ = v82
	var v85 int64
	_ = v85
	var v88 int64
	_ = v88
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v94 int64
	_ = v94
	var v95 int64
	_ = v95
	var v98 int64
	_ = v98
	var v100 int32
	_ = v100
	v10 = m.G0
	v12 = v10 - int32(432)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v26 = F_ParseDateTime(m, v16, v12+int32(240), int32(129), v12+int32(128), v12+int32(16), v12+int32(376))
	mBase = m.M
	if v26 == int32(0) {
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v12)+376))
		v44 = F_DecodeTimeOnly(m, v12+int32(128), v12+int32(16), v33, v12+int32(124), v12+int32(384), v12+int32(428), v12+int32(380), v12+int32(8))
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
			return int32(0)
		} else {
			if v44 == int32(0) {
				v60 = F_palloc(m, int32(16))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					v62 = int64(*(*int32)(unsafe.Add(mBase, uint32(v12)+428)))
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)+384))
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v12)+388))
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v12)+392))
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v12)+380))
					*(*int32)(unsafe.Add(mBase, uint32(v60)+8)) = v66
					v68 = int32(60)
					v77 = v62 + base.I64_extend_i32_s(v63+(v64+v65*v68)*v68)*int64(1000000)
					*(*int64)(unsafe.Add(mBase, uint32(v60))) = v77
					if base.Ui32(int32(6)) < base.Ui32(v14) {
						v100 = v60
					} else {
						v82 = v14 << (uint(int32(3)) % 32)
						v85 = *(*int64)(unsafe.Add(mBase, uint32(v82)+uint32(_consts[1028])))
						v88 = *(*int64)(unsafe.Add(mBase, uint32(v82)+uint32(_consts[1029])))
						if int64(0) <= v77 {
							v91 = v77 + v88
							v92 = base.I64_rem_s(v91, v85)
							v98 = v91 - v92
						} else {
							v94 = v88 - v77
							v95 = base.I64_rem_s(v94, v85)
							v98 = v95 - v94
						}
						*(*int64)(unsafe.Add(mBase, uint32(v60))) = v98
						v100 = v60
					}
					m.G0 = v12 + int32(432)
					return v100
				}
			} else {
				v50 = v44
				F_DateTimeParseError(m, v50, v12+int32(8), v16, int32(390975), v15)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					v56 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v56)
					v100 = int32(0)
					m.G0 = v12 + int32(432)
					return v100
				}
			}
		}
	} else {
		v50 = v26
		F_DateTimeParseError(m, v50, v12+int32(8), v16, int32(390975), v15)
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return int32(0)
		} else {
			v56 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v56)
			v100 = int32(0)
			m.G0 = v12 + int32(432)
			return v100
		}
	}
}
func F_timetz_izone(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v24 int32
	_ = v24
	var v27 int64
	_ = v27
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v83 int32
	_ = v83
	var v85 int64
	_ = v85
	var v86 int32
	_ = v86
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v93 int64
	_ = v93
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v99 int64
	_ = v99
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v108 int64
	_ = v108
	var v110 int64
	_ = v110
	var v113 int64
	_ = v113
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v13 != 0 {
		if v13 != int32(2147483647) {
			if v13 != int32(-2147483648) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						v66 = F_DirectFunctionCall1Coll(m, int32(1292), int32(0), v12)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v66
							F_errmsg(m, int32(121225), v10)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(523621), int32(3206), int32(390606))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
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
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
				if v18 != int32(-2147483648) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v66 = F_DirectFunctionCall1Coll(m, int32(1292), int32(0), v12)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v66
								F_errmsg(m, int32(121225), v10)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(523621), int32(3206), int32(390606))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
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
				} else {
					v21 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
					if v21 == int64(-9223372036854775807-1) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								v41 = F_DirectFunctionCall1Coll(m, int32(1292), int32(0), v12)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v41
									F_errmsg(m, int32(368121), v10+int32(16))
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(523621), int32(3199), int32(390606))
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
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
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								v66 = F_DirectFunctionCall1Coll(m, int32(1292), int32(0), v12)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = v66
									F_errmsg(m, int32(121225), v10)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(523621), int32(3206), int32(390606))
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
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
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
			if v24 != int32(2147483647) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						v66 = F_DirectFunctionCall1Coll(m, int32(1292), int32(0), v12)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v66
							F_errmsg(m, int32(121225), v10)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(523621), int32(3206), int32(390606))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
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
			} else {
				v27 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
				if v27 != int64(9223372036854775807) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v66 = F_DirectFunctionCall1Coll(m, int32(1292), int32(0), v12)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v66
								F_errmsg(m, int32(121225), v10)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(523621), int32(3206), int32(390606))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
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
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v41 = F_DirectFunctionCall1Coll(m, int32(1292), int32(0), v12)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v41
								F_errmsg(m, int32(368121), v10+int32(16))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(523621), int32(3199), int32(390606))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
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
	} else {
		v54 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
		if v54 == int32(0) {
			v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v78 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
			v80 = F_palloc(m, int32(16))
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return int32(0)
			} else {
				v82 = *(*int64)(unsafe.Add(mBase, uint32(v77)))
				v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
				v85 = base.I64_div_s(v78, int64(-1000000))
				v86 = base.I32_wrap_i64(v85)
				*(*int32)(unsafe.Add(mBase, uint32(v80)+8)) = v86
				v91 = base.I64_extend_i32_s(v83-v86) * int64(1000000)
				v92 = v82 + v91
				v93 = int64(0)
				if v93 < v92 {
					v96 = v92
				} else {
					v96 = v93
				}
				v97 = v96 - v82
				v99 = base.I64_extend_i32_u(base.B2i32(v91 != v97))
				v102 = int64(86400000000)
				v103 = base.I64_div_u_s(v97-(v91|v99), v102)
				v108 = v82 + (v103+v99)*v102 + v91
				v110 = base.I64_rem_u_s(v108, v102)
				if base.Ui64(int64(86399999999)) < base.Ui64(v108) {
					v113 = v110
				} else {
					v113 = v108
				}
				*(*int64)(unsafe.Add(mBase, uint32(v80))) = v113
				m.G0 = v10 + int32(32)
				return v80
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int32(0)
				} else {
					v66 = F_DirectFunctionCall1Coll(m, int32(1292), int32(0), v12)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v66
						F_errmsg(m, int32(121225), v10)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(523621), int32(3206), int32(390606))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
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
func F_timetz_larger(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v10 = int64(1000000)
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	v13 = base.I64_extend_i32_s(v8)*v10 + v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v19 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v20 = base.I64_extend_i32_s(v15)*v10 + v19
	if v13 <= v20 {
		if v13 < v20 {
			return v14
		} else {
			if v15 < v8 {
				v25 = v7
			} else {
				v25 = v14
			}
			v26 = v25
			return v26
		}
	} else {
		v26 = v7
		return v26
	}
}
func F_timetz_part(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_timetz_part_common(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
