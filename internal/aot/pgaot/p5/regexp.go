package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_regexp_count(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v19 = F_pg_detoast_datum_packed(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(4) <= v21 {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
				v25 = F_pg_detoast_datum_packed(m, v24)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					v27 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
					v28 = v27
					v29 = v25
					if int32(3) <= v28 {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
						if v33 <= int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v33
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_regexp_count_0)
									F_errmsg(m, int32(_a_F_regexp_count_1), v11)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_regexp_count_2), int32(1155), int32(_a_F_regexp_count_3))
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v36 = v33
							v38 = v11 + int32(24)
							F_parse_re_flags(m, v38, v29)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int64(0)
							} else {
								v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)))
								if v41 == int32(1) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(_a_F_regexp_count_4)
											F_errmsg(m, int32(_a_F_regexp_count_5), v11+int32(16))
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_regexp_count_2), int32(1166), int32(_a_F_regexp_count_3))
												mBase = m.M
												v95 = m.ExcPending
												if v95 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v44 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v44)
									v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v49 = int32(0)
									v52 = F_setup_regexp_matches(m, v14, v19, v38, v36-v44, v48, v49, v49, v49)
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int64(0)
									} else {
										v54 = int64(*(*int32)(unsafe.Add(mBase, uint32(v52)+4)))
										m.G0 = v11 + int32(32)
										return v54
									}
								}
							}
						}
					} else {
						v36 = int32(1)
						v38 = v11 + int32(24)
						F_parse_re_flags(m, v38, v29)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int64(0)
						} else {
							v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)))
							if v41 == int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(_a_F_regexp_count_4)
										F_errmsg(m, int32(_a_F_regexp_count_5), v11+int32(16))
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_regexp_count_2), int32(1166), int32(_a_F_regexp_count_3))
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v44 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v44)
								v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v49 = int32(0)
								v52 = F_setup_regexp_matches(m, v14, v19, v38, v36-v44, v48, v49, v49, v49)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int64(0)
								} else {
									v54 = int64(*(*int32)(unsafe.Add(mBase, uint32(v52)+4)))
									m.G0 = v11 + int32(32)
									return v54
								}
							}
						}
					}
				}
			} else {
				v28 = v21
				v29 = int32(0)
				if int32(3) <= v28 {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					if v33 <= int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v33
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_regexp_count_0)
								F_errmsg(m, int32(_a_F_regexp_count_1), v11)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_regexp_count_2), int32(1155), int32(_a_F_regexp_count_3))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v36 = v33
						v38 = v11 + int32(24)
						F_parse_re_flags(m, v38, v29)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int64(0)
						} else {
							v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)))
							if v41 == int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(_a_F_regexp_count_4)
										F_errmsg(m, int32(_a_F_regexp_count_5), v11+int32(16))
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_regexp_count_2), int32(1166), int32(_a_F_regexp_count_3))
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v44 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v44)
								v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v49 = int32(0)
								v52 = F_setup_regexp_matches(m, v14, v19, v38, v36-v44, v48, v49, v49, v49)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int64(0)
								} else {
									v54 = int64(*(*int32)(unsafe.Add(mBase, uint32(v52)+4)))
									m.G0 = v11 + int32(32)
									return v54
								}
							}
						}
					}
				} else {
					v36 = int32(1)
					v38 = v11 + int32(24)
					F_parse_re_flags(m, v38, v29)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int64(0)
					} else {
						v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)))
						if v41 == int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(_a_F_regexp_count_4)
									F_errmsg(m, int32(_a_F_regexp_count_5), v11+int32(16))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_regexp_count_2), int32(1166), int32(_a_F_regexp_count_3))
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v44 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v44)
							v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v49 = int32(0)
							v52 = F_setup_regexp_matches(m, v14, v19, v38, v36-v44, v48, v49, v49, v49)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int64(0)
							} else {
								v54 = int64(*(*int32)(unsafe.Add(mBase, uint32(v52)+4)))
								m.G0 = v11 + int32(32)
								return v54
							}
						}
					}
				}
			}
		}
	}
}
func F_regexp_match(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int64
	_ = v61
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v16 = F_pg_detoast_datum_packed(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(3) <= v20 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				v24 = F_pg_detoast_datum_packed(m, v23)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					v27 = v24
					F_parse_re_flags(m, v8+int32(8), v27)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int64(0)
					} else {
						v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+12)))
						if v30 != int32(1) {
							v35 = int32(0)
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v40 = F_setup_regexp_matches(m, v11, v16, v8+int32(8), v35, v36, int32(1), v35, v35)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int64(0)
							} else {
								v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
								if v42 == int32(0) {
									v45 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
									v61 = int64(0)
									m.G0 = v8 + int32(16)
									return v61
								} else {
									v49 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
									v50 = F_palloc_mul(m, int32(8), v49)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v50
										v54 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
										v55 = F_palloc_mul(m, int32(1), v54)
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v40)+24)) = v55
											v58 = F_build_regexp_match_result(m, v40)
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return int64(0)
											} else {
												v61 = base.I64_extend_i32_u(v58)
												m.G0 = v8 + int32(16)
												return v61
											}
										}
									}
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_regexp_match_0)
									F_errmsg(m, int32(_a_F_regexp_match_1), v8)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return int64(0)
									} else {
										F_errhint(m, int32(_a_F_regexp_match_2), int32(0))
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_regexp_match_3), int32(1384), int32(_a_F_regexp_match_4))
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
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
				v27 = int32(0)
				F_parse_re_flags(m, v8+int32(8), v27)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int64(0)
				} else {
					v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+12)))
					if v30 != int32(1) {
						v35 = int32(0)
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v40 = F_setup_regexp_matches(m, v11, v16, v8+int32(8), v35, v36, int32(1), v35, v35)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int64(0)
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
							if v42 == int32(0) {
								v45 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
								v61 = int64(0)
								m.G0 = v8 + int32(16)
								return v61
							} else {
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
								v50 = F_palloc_mul(m, int32(8), v49)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v50
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
									v55 = F_palloc_mul(m, int32(1), v54)
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v40)+24)) = v55
										v58 = F_build_regexp_match_result(m, v40)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int64(0)
										} else {
											v61 = base.I64_extend_i32_u(v58)
											m.G0 = v8 + int32(16)
											return v61
										}
									}
								}
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_regexp_match_0)
								F_errmsg(m, int32(_a_F_regexp_match_1), v8)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int64(0)
								} else {
									F_errhint(m, int32(_a_F_regexp_match_2), int32(0))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_regexp_match_3), int32(1384), int32(_a_F_regexp_match_4))
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
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
		}
	}
}
func F_regexp_matches(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int64
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int64
	_ = v96
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v14 == v2 {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v18 = F_pg_detoast_datum_packed(m, v17)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			v22 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(3) <= v22 {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				v26 = F_pg_detoast_datum_packed(m, v25)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int64(0)
				} else {
					v28 = v26
					v29 = F_init_MultiFuncCall(m, l0)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int64(0)
					} else {
						v31 = int32(_a_F_regexp_matches_0)
						v32 = *(*int32)(unsafe.Add(mBase, _c_F_regexp_matches[0]))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
						*(*int32)(unsafe.Add(mBase, _c_F_regexp_matches[0])) = v34
						v37 = v11 + int32(8)
						F_parse_re_flags(m, v37, v28)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int64(0)
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v41 = F_pg_detoast_datum_copy(m, v40)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int64(0)
							} else {
								v43 = int32(0)
								v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v48 = F_setup_regexp_matches(m, v41, v18, v37, v43, v44, int32(1), v43, v43)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int64(0)
								} else {
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
									v52 = F_palloc_mul(m, int32(8), v51)
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = v52
										v56 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
										v57 = F_palloc_mul(m, int32(1), v56)
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v57
											*(*int32)(unsafe.Add(mBase, _c_F_regexp_matches[0])) = v32
											*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v48
											v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
											v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
											v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
											v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
											if v71 < v72 {
												v74 = F_build_regexp_match_result(m, v70)
												mBase = m.M
												v75 = m.ExcPending
												if v75 != 0 {
													return int64(0)
												} else {
													v76 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
													v77 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(v70)+16)) = v76 + v77
													v80 = *(*int64)(unsafe.Add(mBase, uint32(v69)))
													*(*int64)(unsafe.Add(mBase, uint32(v69))) = v80 + int64(1)
													v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v84)+20)) = v77
													v96 = base.I64_extend_i32_u(v74)
													m.G0 = v11 + int32(16)
													return v96
												}
											} else {
												F_end_MultiFuncCall(m, l0)
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return int64(0)
												} else {
													v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v90)+20)) = int32(2)
													v93 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v93)
													v96 = int64(0)
													m.G0 = v11 + int32(16)
													return v96
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
				v28 = v2
				v29 = F_init_MultiFuncCall(m, l0)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int64(0)
				} else {
					v31 = int32(_a_F_regexp_matches_0)
					v32 = *(*int32)(unsafe.Add(mBase, _c_F_regexp_matches[0]))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
					*(*int32)(unsafe.Add(mBase, _c_F_regexp_matches[0])) = v34
					v37 = v11 + int32(8)
					F_parse_re_flags(m, v37, v28)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int64(0)
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v41 = F_pg_detoast_datum_copy(m, v40)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int64(0)
						} else {
							v43 = int32(0)
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v48 = F_setup_regexp_matches(m, v41, v18, v37, v43, v44, int32(1), v43, v43)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int64(0)
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
								v52 = F_palloc_mul(m, int32(8), v51)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = v52
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
									v57 = F_palloc_mul(m, int32(1), v56)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v57
										*(*int32)(unsafe.Add(mBase, _c_F_regexp_matches[0])) = v32
										*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v48
										v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
										v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
										v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
										v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
										if v71 < v72 {
											v74 = F_build_regexp_match_result(m, v70)
											mBase = m.M
											v75 = m.ExcPending
											if v75 != 0 {
												return int64(0)
											} else {
												v76 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
												v77 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(v70)+16)) = v76 + v77
												v80 = *(*int64)(unsafe.Add(mBase, uint32(v69)))
												*(*int64)(unsafe.Add(mBase, uint32(v69))) = v80 + int64(1)
												v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v84)+20)) = v77
												v96 = base.I64_extend_i32_u(v74)
												m.G0 = v11 + int32(16)
												return v96
											}
										} else {
											F_end_MultiFuncCall(m, l0)
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int64(0)
											} else {
												v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v90)+20)) = int32(2)
												v93 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v93)
												v96 = int64(0)
												m.G0 = v11 + int32(16)
												return v96
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
	} else {
		v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
		v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
		v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
		v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
		if v71 < v72 {
			v74 = F_build_regexp_match_result(m, v70)
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return int64(0)
			} else {
				v76 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
				v77 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v70)+16)) = v76 + v77
				v80 = *(*int64)(unsafe.Add(mBase, uint32(v69)))
				*(*int64)(unsafe.Add(mBase, uint32(v69))) = v80 + int64(1)
				v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v84)+20)) = v77
				v96 = base.I64_extend_i32_u(v74)
				m.G0 = v11 + int32(16)
				return v96
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return int64(0)
			} else {
				v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v90)+20)) = int32(2)
				v93 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v93)
				v96 = int64(0)
				m.G0 = v11 + int32(16)
				return v96
			}
		}
	}
}
func F_regexp_substr(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int64
	_ = v107
	var v113 int64
	_ = v113
	var v114 int32
	_ = v114
	var v118 int64
	_ = v118
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int64(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v21 = F_pg_detoast_datum_packed(m, v20)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int64(0)
		} else {
			v23 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(5) <= v23 {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
				v27 = F_pg_detoast_datum_packed(m, v26)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int64(0)
				} else {
					v29 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
					v30 = v29
					v31 = v27
					v32 = int32(1)
					if base.I32_extend16_s(v30) < int32(3) {
						v53 = v32
						v54 = v2
						v55 = v32
						v57 = v11 + int32(-8)
						F_parse_re_flags(m, v57, v31)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int64(0)
						} else {
							v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)))
							if v60 == int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v184 = m.ExcPending
								if v184 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v187 = m.ExcPending
									if v187 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(_a_F_regexp_substr_0)
										F_errmsg(m, int32(_a_F_regexp_substr_1), v11+int32(-48))
										mBase = m.M
										v194 = m.ExcPending
										if v194 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_regexp_substr_2), int32(1956), int32(_a_F_regexp_substr_3))
											mBase = m.M
											v199 = m.ExcPending
											if v199 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v63 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)) = uint8(v63)
								v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v68 = int32(0)
								v72 = F_setup_regexp_matches(m, v16, v21, v57, v55-v63, v67, base.B2i32(v54 != v68), v68, v68)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int64(0)
								} else {
									v74 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
									if v74 < v53 {
										v76 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
										v118 = int64(0)
										m.G0 = v13 - int32(-64)
										return v118
									} else {
										v79 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
										if v79 < v54 {
											v81 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
											v118 = int64(0)
											m.G0 = v13 - int32(-64)
											return v118
										} else {
											v84 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
											v85 = int32(0)
											v94 = v84 + (v54-base.B2i32(v54 != v85)+v79*(v53-int32(1)))<<(uint(int32(3))%32)
											v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
											if v85 <= v95 {
												v98 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
												if int32(0) <= v98 {
													v107 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v72))))
													v113 = F_DirectFunctionCall3Coll(m, int32(1689), int32(0), v107, base.I64_extend_i32_s(v95+int32(1)), base.I64_extend_i32_s(v98-v95))
													mBase = m.M
													v114 = m.ExcPending
													if v114 != 0 {
														return int64(0)
													} else {
														v118 = v113
														m.G0 = v13 - int32(-64)
														return v118
													}
												} else {
													v102 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v102)
													v118 = int64(0)
													m.G0 = v13 - int32(-64)
													return v118
												}
											} else {
												v102 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v102)
												v118 = int64(0)
												m.G0 = v13 - int32(-64)
												return v118
											}
										}
									}
								}
							}
						}
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
						if v37 <= int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v129 = m.ExcPending
								if v129 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v37
									*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(_a_F_regexp_substr_4)
									F_errmsg(m, int32(_a_F_regexp_substr_5), v13)
									mBase = m.M
									v135 = m.ExcPending
									if v135 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_regexp_substr_2), int32(1927), int32(_a_F_regexp_substr_3))
										mBase = m.M
										v140 = m.ExcPending
										if v140 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v41 = v30 & int32(_a_F_regexp_substr_6)
							if v41 == int32(3) {
								v53 = v32
								v54 = v2
								v55 = v37
								v57 = v11 + int32(-8)
								F_parse_re_flags(m, v57, v31)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int64(0)
								} else {
									v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)))
									if v60 == int32(1) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v184 = m.ExcPending
										if v184 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v187 = m.ExcPending
											if v187 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(_a_F_regexp_substr_0)
												F_errmsg(m, int32(_a_F_regexp_substr_1), v11+int32(-48))
												mBase = m.M
												v194 = m.ExcPending
												if v194 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_regexp_substr_2), int32(1956), int32(_a_F_regexp_substr_3))
													mBase = m.M
													v199 = m.ExcPending
													if v199 != 0 {
														return int64(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v63 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)) = uint8(v63)
										v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v68 = int32(0)
										v72 = F_setup_regexp_matches(m, v16, v21, v57, v55-v63, v67, base.B2i32(v54 != v68), v68, v68)
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int64(0)
										} else {
											v74 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
											if v74 < v53 {
												v76 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
												v118 = int64(0)
												m.G0 = v13 - int32(-64)
												return v118
											} else {
												v79 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
												if v79 < v54 {
													v81 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
													v118 = int64(0)
													m.G0 = v13 - int32(-64)
													return v118
												} else {
													v84 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
													v85 = int32(0)
													v94 = v84 + (v54-base.B2i32(v54 != v85)+v79*(v53-int32(1)))<<(uint(int32(3))%32)
													v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
													if v85 <= v95 {
														v98 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
														if int32(0) <= v98 {
															v107 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v72))))
															v113 = F_DirectFunctionCall3Coll(m, int32(1689), int32(0), v107, base.I64_extend_i32_s(v95+int32(1)), base.I64_extend_i32_s(v98-v95))
															mBase = m.M
															v114 = m.ExcPending
															if v114 != 0 {
																return int64(0)
															} else {
																v118 = v113
																m.G0 = v13 - int32(-64)
																return v118
															}
														} else {
															v102 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v102)
															v118 = int64(0)
															m.G0 = v13 - int32(-64)
															return v118
														}
													} else {
														v102 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v102)
														v118 = int64(0)
														m.G0 = v13 - int32(-64)
														return v118
													}
												}
											}
										}
									}
								}
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
								if v44 <= int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v144 = m.ExcPending
									if v144 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v147 = m.ExcPending
										if v147 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v44
											*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = int32(_a_F_regexp_substr_7)
											F_errmsg(m, int32(_a_F_regexp_substr_5), v11+int32(-32))
											mBase = m.M
											v155 = m.ExcPending
											if v155 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_regexp_substr_2), int32(1936), int32(_a_F_regexp_substr_3))
												mBase = m.M
												v160 = m.ExcPending
												if v160 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									if base.Ui32(v41) < base.Ui32(int32(6)) {
										v53 = v44
										v54 = v2
										v55 = v37
										v57 = v11 + int32(-8)
										F_parse_re_flags(m, v57, v31)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int64(0)
										} else {
											v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)))
											if v60 == int32(1) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v184 = m.ExcPending
												if v184 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v187 = m.ExcPending
													if v187 != 0 {
														return int64(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(_a_F_regexp_substr_0)
														F_errmsg(m, int32(_a_F_regexp_substr_1), v11+int32(-48))
														mBase = m.M
														v194 = m.ExcPending
														if v194 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_regexp_substr_2), int32(1956), int32(_a_F_regexp_substr_3))
															mBase = m.M
															v199 = m.ExcPending
															if v199 != 0 {
																return int64(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v63 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)) = uint8(v63)
												v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v68 = int32(0)
												v72 = F_setup_regexp_matches(m, v16, v21, v57, v55-v63, v67, base.B2i32(v54 != v68), v68, v68)
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return int64(0)
												} else {
													v74 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
													if v74 < v53 {
														v76 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
														v118 = int64(0)
														m.G0 = v13 - int32(-64)
														return v118
													} else {
														v79 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
														if v79 < v54 {
															v81 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
															v118 = int64(0)
															m.G0 = v13 - int32(-64)
															return v118
														} else {
															v84 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
															v85 = int32(0)
															v94 = v84 + (v54-base.B2i32(v54 != v85)+v79*(v53-int32(1)))<<(uint(int32(3))%32)
															v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
															if v85 <= v95 {
																v98 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
																if int32(0) <= v98 {
																	v107 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v72))))
																	v113 = F_DirectFunctionCall3Coll(m, int32(1689), int32(0), v107, base.I64_extend_i32_s(v95+int32(1)), base.I64_extend_i32_s(v98-v95))
																	mBase = m.M
																	v114 = m.ExcPending
																	if v114 != 0 {
																		return int64(0)
																	} else {
																		v118 = v113
																		m.G0 = v13 - int32(-64)
																		return v118
																	}
																} else {
																	v102 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v102)
																	v118 = int64(0)
																	m.G0 = v13 - int32(-64)
																	return v118
																}
															} else {
																v102 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v102)
																v118 = int64(0)
																m.G0 = v13 - int32(-64)
																return v118
															}
														}
													}
												}
											}
										}
									} else {
										v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
										if v49 < int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v167 = m.ExcPending
												if v167 != 0 {
													return int64(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v49
													*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = int32(_a_F_regexp_substr_8)
													F_errmsg(m, int32(_a_F_regexp_substr_5), v11+int32(-16))
													mBase = m.M
													v175 = m.ExcPending
													if v175 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_regexp_substr_2), int32(1945), int32(_a_F_regexp_substr_3))
														mBase = m.M
														v180 = m.ExcPending
														if v180 != 0 {
															return int64(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											v53 = v44
											v54 = v49
											v55 = v37
											v57 = v11 + int32(-8)
											F_parse_re_flags(m, v57, v31)
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return int64(0)
											} else {
												v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)))
												if v60 == int32(1) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v184 = m.ExcPending
													if v184 != 0 {
														return int64(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v187 = m.ExcPending
														if v187 != 0 {
															return int64(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(_a_F_regexp_substr_0)
															F_errmsg(m, int32(_a_F_regexp_substr_1), v11+int32(-48))
															mBase = m.M
															v194 = m.ExcPending
															if v194 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_regexp_substr_2), int32(1956), int32(_a_F_regexp_substr_3))
																mBase = m.M
																v199 = m.ExcPending
																if v199 != 0 {
																	return int64(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													v63 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)) = uint8(v63)
													v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v68 = int32(0)
													v72 = F_setup_regexp_matches(m, v16, v21, v57, v55-v63, v67, base.B2i32(v54 != v68), v68, v68)
													mBase = m.M
													v73 = m.ExcPending
													if v73 != 0 {
														return int64(0)
													} else {
														v74 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
														if v74 < v53 {
															v76 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
															v118 = int64(0)
															m.G0 = v13 - int32(-64)
															return v118
														} else {
															v79 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
															if v79 < v54 {
																v81 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
																v118 = int64(0)
																m.G0 = v13 - int32(-64)
																return v118
															} else {
																v84 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
																v85 = int32(0)
																v94 = v84 + (v54-base.B2i32(v54 != v85)+v79*(v53-int32(1)))<<(uint(int32(3))%32)
																v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
																if v85 <= v95 {
																	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
																	if int32(0) <= v98 {
																		v107 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v72))))
																		v113 = F_DirectFunctionCall3Coll(m, int32(1689), int32(0), v107, base.I64_extend_i32_s(v95+int32(1)), base.I64_extend_i32_s(v98-v95))
																		mBase = m.M
																		v114 = m.ExcPending
																		if v114 != 0 {
																			return int64(0)
																		} else {
																			v118 = v113
																			m.G0 = v13 - int32(-64)
																			return v118
																		}
																	} else {
																		v102 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v102)
																		v118 = int64(0)
																		m.G0 = v13 - int32(-64)
																		return v118
																	}
																} else {
																	v102 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v102)
																	v118 = int64(0)
																	m.G0 = v13 - int32(-64)
																	return v118
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
				}
			} else {
				v30 = v23
				v31 = v2
				v32 = int32(1)
				if base.I32_extend16_s(v30) < int32(3) {
					v53 = v32
					v54 = v2
					v55 = v32
					v57 = v11 + int32(-8)
					F_parse_re_flags(m, v57, v31)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int64(0)
					} else {
						v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)))
						if v60 == int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v184 = m.ExcPending
							if v184 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v187 = m.ExcPending
								if v187 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(_a_F_regexp_substr_0)
									F_errmsg(m, int32(_a_F_regexp_substr_1), v11+int32(-48))
									mBase = m.M
									v194 = m.ExcPending
									if v194 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_regexp_substr_2), int32(1956), int32(_a_F_regexp_substr_3))
										mBase = m.M
										v199 = m.ExcPending
										if v199 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v63 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)) = uint8(v63)
							v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v68 = int32(0)
							v72 = F_setup_regexp_matches(m, v16, v21, v57, v55-v63, v67, base.B2i32(v54 != v68), v68, v68)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int64(0)
							} else {
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
								if v74 < v53 {
									v76 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
									v118 = int64(0)
									m.G0 = v13 - int32(-64)
									return v118
								} else {
									v79 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
									if v79 < v54 {
										v81 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
										v118 = int64(0)
										m.G0 = v13 - int32(-64)
										return v118
									} else {
										v84 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
										v85 = int32(0)
										v94 = v84 + (v54-base.B2i32(v54 != v85)+v79*(v53-int32(1)))<<(uint(int32(3))%32)
										v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
										if v85 <= v95 {
											v98 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
											if int32(0) <= v98 {
												v107 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v72))))
												v113 = F_DirectFunctionCall3Coll(m, int32(1689), int32(0), v107, base.I64_extend_i32_s(v95+int32(1)), base.I64_extend_i32_s(v98-v95))
												mBase = m.M
												v114 = m.ExcPending
												if v114 != 0 {
													return int64(0)
												} else {
													v118 = v113
													m.G0 = v13 - int32(-64)
													return v118
												}
											} else {
												v102 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v102)
												v118 = int64(0)
												m.G0 = v13 - int32(-64)
												return v118
											}
										} else {
											v102 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v102)
											v118 = int64(0)
											m.G0 = v13 - int32(-64)
											return v118
										}
									}
								}
							}
						}
					}
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					if v37 <= int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v37
								*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(_a_F_regexp_substr_4)
								F_errmsg(m, int32(_a_F_regexp_substr_5), v13)
								mBase = m.M
								v135 = m.ExcPending
								if v135 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_regexp_substr_2), int32(1927), int32(_a_F_regexp_substr_3))
									mBase = m.M
									v140 = m.ExcPending
									if v140 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v41 = v30 & int32(_a_F_regexp_substr_6)
						if v41 == int32(3) {
							v53 = v32
							v54 = v2
							v55 = v37
							v57 = v11 + int32(-8)
							F_parse_re_flags(m, v57, v31)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int64(0)
							} else {
								v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)))
								if v60 == int32(1) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v184 = m.ExcPending
									if v184 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v187 = m.ExcPending
										if v187 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(_a_F_regexp_substr_0)
											F_errmsg(m, int32(_a_F_regexp_substr_1), v11+int32(-48))
											mBase = m.M
											v194 = m.ExcPending
											if v194 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_regexp_substr_2), int32(1956), int32(_a_F_regexp_substr_3))
												mBase = m.M
												v199 = m.ExcPending
												if v199 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v63 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)) = uint8(v63)
									v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v68 = int32(0)
									v72 = F_setup_regexp_matches(m, v16, v21, v57, v55-v63, v67, base.B2i32(v54 != v68), v68, v68)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int64(0)
									} else {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
										if v74 < v53 {
											v76 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
											v118 = int64(0)
											m.G0 = v13 - int32(-64)
											return v118
										} else {
											v79 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
											if v79 < v54 {
												v81 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
												v118 = int64(0)
												m.G0 = v13 - int32(-64)
												return v118
											} else {
												v84 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
												v85 = int32(0)
												v94 = v84 + (v54-base.B2i32(v54 != v85)+v79*(v53-int32(1)))<<(uint(int32(3))%32)
												v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
												if v85 <= v95 {
													v98 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
													if int32(0) <= v98 {
														v107 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v72))))
														v113 = F_DirectFunctionCall3Coll(m, int32(1689), int32(0), v107, base.I64_extend_i32_s(v95+int32(1)), base.I64_extend_i32_s(v98-v95))
														mBase = m.M
														v114 = m.ExcPending
														if v114 != 0 {
															return int64(0)
														} else {
															v118 = v113
															m.G0 = v13 - int32(-64)
															return v118
														}
													} else {
														v102 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v102)
														v118 = int64(0)
														m.G0 = v13 - int32(-64)
														return v118
													}
												} else {
													v102 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v102)
													v118 = int64(0)
													m.G0 = v13 - int32(-64)
													return v118
												}
											}
										}
									}
								}
							}
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
							if v44 <= int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v144 = m.ExcPending
								if v144 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v147 = m.ExcPending
									if v147 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v44
										*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = int32(_a_F_regexp_substr_7)
										F_errmsg(m, int32(_a_F_regexp_substr_5), v11+int32(-32))
										mBase = m.M
										v155 = m.ExcPending
										if v155 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_regexp_substr_2), int32(1936), int32(_a_F_regexp_substr_3))
											mBase = m.M
											v160 = m.ExcPending
											if v160 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								if base.Ui32(v41) < base.Ui32(int32(6)) {
									v53 = v44
									v54 = v2
									v55 = v37
									v57 = v11 + int32(-8)
									F_parse_re_flags(m, v57, v31)
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int64(0)
									} else {
										v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)))
										if v60 == int32(1) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v184 = m.ExcPending
											if v184 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v187 = m.ExcPending
												if v187 != 0 {
													return int64(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(_a_F_regexp_substr_0)
													F_errmsg(m, int32(_a_F_regexp_substr_1), v11+int32(-48))
													mBase = m.M
													v194 = m.ExcPending
													if v194 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_regexp_substr_2), int32(1956), int32(_a_F_regexp_substr_3))
														mBase = m.M
														v199 = m.ExcPending
														if v199 != 0 {
															return int64(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											v63 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)) = uint8(v63)
											v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v68 = int32(0)
											v72 = F_setup_regexp_matches(m, v16, v21, v57, v55-v63, v67, base.B2i32(v54 != v68), v68, v68)
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return int64(0)
											} else {
												v74 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
												if v74 < v53 {
													v76 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
													v118 = int64(0)
													m.G0 = v13 - int32(-64)
													return v118
												} else {
													v79 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
													if v79 < v54 {
														v81 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
														v118 = int64(0)
														m.G0 = v13 - int32(-64)
														return v118
													} else {
														v84 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
														v85 = int32(0)
														v94 = v84 + (v54-base.B2i32(v54 != v85)+v79*(v53-int32(1)))<<(uint(int32(3))%32)
														v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
														if v85 <= v95 {
															v98 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
															if int32(0) <= v98 {
																v107 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v72))))
																v113 = F_DirectFunctionCall3Coll(m, int32(1689), int32(0), v107, base.I64_extend_i32_s(v95+int32(1)), base.I64_extend_i32_s(v98-v95))
																mBase = m.M
																v114 = m.ExcPending
																if v114 != 0 {
																	return int64(0)
																} else {
																	v118 = v113
																	m.G0 = v13 - int32(-64)
																	return v118
																}
															} else {
																v102 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v102)
																v118 = int64(0)
																m.G0 = v13 - int32(-64)
																return v118
															}
														} else {
															v102 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v102)
															v118 = int64(0)
															m.G0 = v13 - int32(-64)
															return v118
														}
													}
												}
											}
										}
									}
								} else {
									v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
									if v49 < int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v167 = m.ExcPending
											if v167 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v49
												*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = int32(_a_F_regexp_substr_8)
												F_errmsg(m, int32(_a_F_regexp_substr_5), v11+int32(-16))
												mBase = m.M
												v175 = m.ExcPending
												if v175 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_regexp_substr_2), int32(1945), int32(_a_F_regexp_substr_3))
													mBase = m.M
													v180 = m.ExcPending
													if v180 != 0 {
														return int64(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v53 = v44
										v54 = v49
										v55 = v37
										v57 = v11 + int32(-8)
										F_parse_re_flags(m, v57, v31)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int64(0)
										} else {
											v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)))
											if v60 == int32(1) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v184 = m.ExcPending
												if v184 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v187 = m.ExcPending
													if v187 != 0 {
														return int64(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(_a_F_regexp_substr_0)
														F_errmsg(m, int32(_a_F_regexp_substr_1), v11+int32(-48))
														mBase = m.M
														v194 = m.ExcPending
														if v194 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_regexp_substr_2), int32(1956), int32(_a_F_regexp_substr_3))
															mBase = m.M
															v199 = m.ExcPending
															if v199 != 0 {
																return int64(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v63 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)) = uint8(v63)
												v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v68 = int32(0)
												v72 = F_setup_regexp_matches(m, v16, v21, v57, v55-v63, v67, base.B2i32(v54 != v68), v68, v68)
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return int64(0)
												} else {
													v74 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
													if v74 < v53 {
														v76 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
														v118 = int64(0)
														m.G0 = v13 - int32(-64)
														return v118
													} else {
														v79 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
														if v79 < v54 {
															v81 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
															v118 = int64(0)
															m.G0 = v13 - int32(-64)
															return v118
														} else {
															v84 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
															v85 = int32(0)
															v94 = v84 + (v54-base.B2i32(v54 != v85)+v79*(v53-int32(1)))<<(uint(int32(3))%32)
															v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
															if v85 <= v95 {
																v98 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
																if int32(0) <= v98 {
																	v107 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v72))))
																	v113 = F_DirectFunctionCall3Coll(m, int32(1689), int32(0), v107, base.I64_extend_i32_s(v95+int32(1)), base.I64_extend_i32_s(v98-v95))
																	mBase = m.M
																	v114 = m.ExcPending
																	if v114 != 0 {
																		return int64(0)
																	} else {
																		v118 = v113
																		m.G0 = v13 - int32(-64)
																		return v118
																	}
																} else {
																	v102 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v102)
																	v118 = int64(0)
																	m.G0 = v13 - int32(-64)
																	return v118
																}
															} else {
																v102 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v102)
																v118 = int64(0)
																m.G0 = v13 - int32(-64)
																return v118
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
			}
		}
	}
}
