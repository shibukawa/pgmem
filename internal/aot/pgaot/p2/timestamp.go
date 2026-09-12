package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_timestamp_cmp_date(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v18 int64
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v7 == int32(-2147483648) {
		v18 = int64(-9223372036854775807 - 1)
		v27 = base.B2i32(v4 < v18) - base.B2i32(v18 < v4)
	} else {
		if v7 == int32(2147483647) {
			v18 = int64(9223372036854775807)
			v27 = base.B2i32(v4 < v18) - base.B2i32(v18 < v4)
		} else {
			if int32(106751982) < v7 {
				if v4 == int64(9223372036854775807) {
					v26 = int32(-1)
				} else {
					v26 = int32(1)
				}
				v27 = v26
			} else {
				v18 = base.I64_extend_i32_s(v7) * int64(86400000000)
				v27 = base.B2i32(v4 < v18) - base.B2i32(v18 < v4)
			}
		}
	}
	return int32(0) - v27
}
func F_timestamp_cmp_timestamptz(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v15 = F_timestamp2timestamptz_opt_overflow(m, v12, v7+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if v10 == int64(9223372036854775807) {
			v23 = int32(-1)
		} else {
			v23 = int32(1)
		}
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		if int32(0) < v24 {
			v37 = v23
		} else {
			if v10 == int64(-9223372036854775807-1) {
				v31 = int32(1)
			} else {
				v31 = int32(-1)
			}
			if v24 < int32(0) {
				v37 = v31
			} else {
				v37 = base.B2i32(v10 < v15) - base.B2i32(v15 < v10)
			}
		}
		m.G0 = v7 + int32(16)
		return v37
	}
}
func F_timestamp_eq_timestamptz(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v15 = F_timestamp2timestamptz_opt_overflow(m, v12, v7+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		m.G0 = v7 + int32(16)
		return base.B2i32(v19 == int32(0)) & base.B2i32(v10 == v15)
	}
}
func F_timestamp_finite(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	return base.B2i32(base.Ui64(v3+int64(9223372036854775807)) < base.Ui64(int64(-2)))
}
func F_timestamp_ge_timestamptz(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v15 = F_timestamp2timestamptz_opt_overflow(m, v12, v7+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		if int32(0) < v21 {
			v29 = base.B2i32(v10 != int64(9223372036854775807))
		} else {
			if v21 < int32(0) {
				v29 = base.B2i32(v10 == int64(-9223372036854775807-1))
			} else {
				v29 = base.B2i32(v10 <= v15)
			}
		}
		m.G0 = v7 + int32(16)
		return v29
	}
}
func F_timestamp_izone(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v64 int64
	_ = v64
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	if base.Ui64(int64(2)) <= base.Ui64(v10-int64(9223372036854775807)) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
		if v16 != 0 {
			if v16 != int32(2147483647) {
				if v16 != int32(-2147483648) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return int32(0)
						} else {
							v103 = F_DirectFunctionCall1Coll(m, int32(1288), int32(0), v15)
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v103
								F_errmsg(m, int32(111400), v7)
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(486977), int32(6413), int32(365674))
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
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
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
					if v21 != int32(-2147483648) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int32(0)
							} else {
								v103 = F_DirectFunctionCall1Coll(m, int32(1288), int32(0), v15)
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v103
									F_errmsg(m, int32(111400), v7)
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(486977), int32(6413), int32(365674))
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
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
						v24 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
						if v24 == int64(-9223372036854775807-1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									v44 = F_DirectFunctionCall1Coll(m, int32(1288), int32(0), v15)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v44
										F_errmsg(m, int32(344297), v7+int32(16))
										mBase = m.M
										v51 = m.ExcPending
										if v51 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(486977), int32(6406), int32(365674))
											mBase = m.M
											v56 = m.ExcPending
											if v56 != 0 {
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
							v97 = m.ExcPending
							if v97 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return int32(0)
								} else {
									v103 = F_DirectFunctionCall1Coll(m, int32(1288), int32(0), v15)
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = v103
										F_errmsg(m, int32(111400), v7)
										mBase = m.M
										v108 = m.ExcPending
										if v108 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(486977), int32(6413), int32(365674))
											mBase = m.M
											v113 = m.ExcPending
											if v113 != 0 {
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
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				if v27 != int32(2147483647) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return int32(0)
						} else {
							v103 = F_DirectFunctionCall1Coll(m, int32(1288), int32(0), v15)
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v103
								F_errmsg(m, int32(111400), v7)
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(486977), int32(6413), int32(365674))
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
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
					v30 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
					if v30 != int64(9223372036854775807) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int32(0)
							} else {
								v103 = F_DirectFunctionCall1Coll(m, int32(1288), int32(0), v15)
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v103
									F_errmsg(m, int32(111400), v7)
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(486977), int32(6413), int32(365674))
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
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
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								v44 = F_DirectFunctionCall1Coll(m, int32(1288), int32(0), v15)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v44
									F_errmsg(m, int32(344297), v7+int32(16))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(486977), int32(6406), int32(365674))
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
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
			v57 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
			if v57 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v97 = m.ExcPending
				if v97 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return int32(0)
					} else {
						v103 = F_DirectFunctionCall1Coll(m, int32(1288), int32(0), v15)
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v103
							F_errmsg(m, int32(111400), v7)
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(486977), int32(6413), int32(365674))
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
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
				v58 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
				v60 = base.I64_div_s(v58, int64(1000000))
				v64 = base.I64_extend32_s(v60)*int64(-1000000) + v10
				if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v64+int64(211813488000000000)) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(395441), int32(0))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(486977), int32(6422), int32(365674))
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
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
					v71 = v64
					v72 = F_Int64GetDatum(m, v71)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(32)
						return v72
					}
				}
			}
		}
	} else {
		v71 = v10
		v72 = F_Int64GetDatum(m, v71)
		mBase = m.M
		v73 = m.ExcPending
		if v73 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(32)
			return v72
		}
	}
}
func F_timestamp_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(1509)
	return int32(0)
}
func F_timestamp_time(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	if base.Ui64(v9-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
		v14 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
		v42 = int32(0)
		m.G0 = v6 + int32(48)
		return v42
	} else {
		v17 = int32(0)
		v22 = F_timestamp2tm(m, v9, v17, v6+int32(4), v6, v17, v17)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			if v22 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(395441), int32(0))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(489677), int32(1984), int32(368252))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
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
				v26 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6))))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
				v30 = int32(60)
				v40 = F_Int64GetDatum(m, v26+base.I64_extend_i32_s(v27+(v28+v29*v30)*v30)*int64(1000000))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v42 = v40
					m.G0 = v6 + int32(48)
					return v42
				}
			}
		}
	}
}
