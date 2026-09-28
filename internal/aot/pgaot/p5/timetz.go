package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_extract_timetz(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_timetz_part_common(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_timetz_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v75 int64
	_ = v75
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v92 int64
	_ = v92
	var v107 int64
	_ = v107
	v12 = m.G0
	v14 = v12 - int32(432)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v23 = v14 + int32(128)
	v25 = v14 + int32(16)
	v28 = F_ParseDateTime(m, v18, v14+int32(240), int32(129), v23, v25, v14+int32(376))
	mBase = m.M
	if v28 == int32(0) {
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v14)+376))
		v42 = F_DecodeTimeOnly(m, v23, v25, v31, v14+int32(124), v14+int32(384), v14+int32(428), v14+int32(380), v14+int32(8))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int64(0)
		} else {
			if v42 == int32(0) {
				v58 = F_palloc(m, int32(16))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int64(0)
				} else {
					v60 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+428)))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v14)+384))
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v14)+388))
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v14)+392))
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v14)+380))
					*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v64
					v66 = int32(60)
					v75 = v60 + base.I64_extend_i32_s(v61+(v62+v63*v66)*v66)*int64(1000000)
					*(*int64)(unsafe.Add(mBase, uint32(v58))) = v75
					if base.Ui32(v17) <= base.Ui32(int32(6)) {
						v80 = v17 << (uint(int32(3)) % 32)
						v81 = *(*int64)(unsafe.Add(mBase, uint32(v80)+uint32(_c_F_timetz_in[0])))
						v82 = *(*int64)(unsafe.Add(mBase, uint32(v80)+uint32(_c_F_timetz_in[1])))
						if int64(0) <= v75 {
							v85 = v75 + v82
							v86 = base.I64_rem_s(v85, v81)
							v92 = v85 - v86
						} else {
							v88 = v82 - v75
							v89 = base.I64_rem_s(v88, v81)
							v92 = v89 - v88
						}
						*(*int64)(unsafe.Add(mBase, uint32(v58))) = v92
					} else {
					}
					v107 = base.I64_extend_i32_u(v58)
					m.G0 = v14 + int32(432)
					return v107
				}
			} else {
				v48 = v42
				F_DateTimeParseError(m, v48, v14+int32(8), v18, int32(_a_F_timetz_in_0), v16)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int64(0)
				} else {
					v54 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v54)
					v107 = int64(0)
					m.G0 = v14 + int32(432)
					return v107
				}
			}
		}
	} else {
		v48 = v28
		F_DateTimeParseError(m, v48, v14+int32(8), v18, int32(_a_F_timetz_in_0), v16)
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return int64(0)
		} else {
			v54 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v54)
			v107 = int64(0)
			m.G0 = v14 + int32(432)
			return v107
		}
	}
}
func F_timetz_izone(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int64
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v81 int64
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int64
	_ = v86
	var v87 int32
	_ = v87
	var v89 int64
	_ = v89
	var v90 int32
	_ = v90
	var v95 int64
	_ = v95
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v103 int64
	_ = v103
	var v106 int64
	_ = v106
	var v107 int64
	_ = v107
	var v112 int64
	_ = v112
	var v114 int64
	_ = v114
	var v117 int64
	_ = v117
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = v12 & int64(4294967295)
	v15 = base.I32_wrap_i64(v12)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v16 != 0 {
		if v16 != int32(2147483647) {
			if v16 != int32(-2147483648) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int64(0)
					} else {
						v69 = F_DirectFunctionCall1Coll(m, int32(1401), int32(0), v14)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int64(0)
						} else {
							*(*uint32)(unsafe.Add(mBase, uint32(v10))) = uint32(v69)
							F_errmsg(m, int32(_a_F_timetz_izone_0), v10)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_timetz_izone_1), int32(3270), int32(_a_F_timetz_izone_2))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int64(0)
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
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				if v21 != int32(-2147483648) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int64(0)
						} else {
							v69 = F_DirectFunctionCall1Coll(m, int32(1401), int32(0), v14)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int64(0)
							} else {
								*(*uint32)(unsafe.Add(mBase, uint32(v10))) = uint32(v69)
								F_errmsg(m, int32(_a_F_timetz_izone_0), v10)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_timetz_izone_1), int32(3270), int32(_a_F_timetz_izone_2))
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int64(0)
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
					v24 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
					if v24 == int64(-9223372036854775807-1) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int64(0)
							} else {
								v44 = F_DirectFunctionCall1Coll(m, int32(1401), int32(0), v14)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int64(0)
								} else {
									*(*uint32)(unsafe.Add(mBase, uint32(v10)+16)) = uint32(v44)
									F_errmsg(m, int32(_a_F_timetz_izone_3), v10+int32(16))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_timetz_izone_1), int32(3263), int32(_a_F_timetz_izone_2))
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return int64(0)
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
						v63 = m.ExcPending
						if v63 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int64(0)
							} else {
								v69 = F_DirectFunctionCall1Coll(m, int32(1401), int32(0), v14)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int64(0)
								} else {
									*(*uint32)(unsafe.Add(mBase, uint32(v10))) = uint32(v69)
									F_errmsg(m, int32(_a_F_timetz_izone_0), v10)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_timetz_izone_1), int32(3270), int32(_a_F_timetz_izone_2))
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int64(0)
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
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
			if v27 != int32(2147483647) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int64(0)
					} else {
						v69 = F_DirectFunctionCall1Coll(m, int32(1401), int32(0), v14)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int64(0)
						} else {
							*(*uint32)(unsafe.Add(mBase, uint32(v10))) = uint32(v69)
							F_errmsg(m, int32(_a_F_timetz_izone_0), v10)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_timetz_izone_1), int32(3270), int32(_a_F_timetz_izone_2))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int64(0)
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
				v30 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
				if v30 != int64(9223372036854775807) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int64(0)
						} else {
							v69 = F_DirectFunctionCall1Coll(m, int32(1401), int32(0), v14)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int64(0)
							} else {
								*(*uint32)(unsafe.Add(mBase, uint32(v10))) = uint32(v69)
								F_errmsg(m, int32(_a_F_timetz_izone_0), v10)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_timetz_izone_1), int32(3270), int32(_a_F_timetz_izone_2))
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int64(0)
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
					v38 = m.ExcPending
					if v38 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int64(0)
						} else {
							v44 = F_DirectFunctionCall1Coll(m, int32(1401), int32(0), v14)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int64(0)
							} else {
								*(*uint32)(unsafe.Add(mBase, uint32(v10)+16)) = uint32(v44)
								F_errmsg(m, int32(_a_F_timetz_izone_3), v10+int32(16))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_timetz_izone_1), int32(3263), int32(_a_F_timetz_izone_2))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int64(0)
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
		v57 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
		if v57 == int32(0) {
			v80 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
			v81 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
			v83 = F_palloc(m, int32(16))
			mBase = m.M
			v84 = m.ExcPending
			if v84 != 0 {
				return int64(0)
			} else {
				v85 = base.I32_wrap_i64(v80)
				v86 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
				v87 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
				v89 = base.I64_div_s(v81, int64(-1000000))
				v90 = base.I32_wrap_i64(v89)
				*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v90
				v95 = base.I64_extend_i32_s(v87-v90) * int64(1000000)
				v96 = v86 + v95
				v97 = int64(0)
				if v97 < v96 {
					v100 = v96
				} else {
					v100 = v97
				}
				v101 = v100 - v86
				v103 = base.I64_extend_i32_u(base.B2i32(v95 != v101))
				v106 = int64(86400000000)
				v107 = base.I64_div_u_s(v101-(v95|v103), v106)
				v112 = v86 + (v107+v103)*v106 + v95
				v114 = base.I64_rem_u_s(v112, v106)
				if base.Ui64(int64(86399999999)) < base.Ui64(v112) {
					v117 = v114
				} else {
					v117 = v112
				}
				*(*int64)(unsafe.Add(mBase, uint32(v83))) = v117
				m.G0 = v10 + int32(32)
				return base.I64_extend_i32_u(v83)
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int64(0)
				} else {
					v69 = F_DirectFunctionCall1Coll(m, int32(1401), int32(0), v14)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int64(0)
					} else {
						*(*uint32)(unsafe.Add(mBase, uint32(v10))) = uint32(v69)
						F_errmsg(m, int32(_a_F_timetz_izone_0), v10)
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_timetz_izone_1), int32(3270), int32(_a_F_timetz_izone_2))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return int64(0)
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
func F_timetz_larger(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = base.I32_wrap_i64(v8)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v12 = int64(1000000)
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	v15 = base.I64_extend_i32_s(v10)*v12 + v14
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v17 = base.I32_wrap_i64(v16)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v22 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	v23 = base.I64_extend_i32_s(v18)*v12 + v22
	if v23 < v15 {
		v28 = v8
	} else {
		if v15 < v23 {
			v28 = v16
		} else {
			if v18 < v10 {
				v27 = v8
			} else {
				v27 = v16
			}
			v28 = v27
		}
	}
	return v28 & int64(4294967295)
}
func F_timetz_part(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_timetz_part_common(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
