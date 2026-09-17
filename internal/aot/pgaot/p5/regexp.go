package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_regexp_count(m *base.Module, l0 int32) int32 {
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
	var v23 int32
	_ = v23
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = F_pg_detoast_datum_packed(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(4) <= v20 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				v24 = F_pg_detoast_datum_packed(m, v23)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
					v27 = v26
					v28 = v24
					if int32(3) <= v27 {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						if v32 <= int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v32
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_regexp_count_0)
									F_errmsg(m, int32(_a_F_regexp_count_1), v10)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_regexp_count_2), int32(1155), int32(_a_F_regexp_count_3))
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
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
							v35 = v32
							v37 = v10 + int32(24)
							F_parse_re_flags(m, v37, v28)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)))
								if v40 == int32(1) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_regexp_count_4)
											F_errmsg(m, int32(_a_F_regexp_count_5), v10+int32(16))
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_regexp_count_2), int32(1166), int32(_a_F_regexp_count_3))
												mBase = m.M
												v94 = m.ExcPending
												if v94 != 0 {
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
									v43 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)) = uint8(v43)
									v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v48 = int32(0)
									v51 = F_setup_regexp_matches(m, v13, v18, v37, v35-v43, v47, v48, v48, v48)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										v53 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
										m.G0 = v10 + int32(32)
										return v53
									}
								}
							}
						}
					} else {
						v35 = int32(1)
						v37 = v10 + int32(24)
						F_parse_re_flags(m, v37, v28)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)))
							if v40 == int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_regexp_count_4)
										F_errmsg(m, int32(_a_F_regexp_count_5), v10+int32(16))
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_regexp_count_2), int32(1166), int32(_a_F_regexp_count_3))
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
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
								v43 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)) = uint8(v43)
								v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v48 = int32(0)
								v51 = F_setup_regexp_matches(m, v13, v18, v37, v35-v43, v47, v48, v48, v48)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									v53 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
									m.G0 = v10 + int32(32)
									return v53
								}
							}
						}
					}
				}
			} else {
				v27 = v20
				v28 = int32(0)
				if int32(3) <= v27 {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					if v32 <= int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v32
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_regexp_count_0)
								F_errmsg(m, int32(_a_F_regexp_count_1), v10)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_regexp_count_2), int32(1155), int32(_a_F_regexp_count_3))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
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
						v35 = v32
						v37 = v10 + int32(24)
						F_parse_re_flags(m, v37, v28)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)))
							if v40 == int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_regexp_count_4)
										F_errmsg(m, int32(_a_F_regexp_count_5), v10+int32(16))
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_regexp_count_2), int32(1166), int32(_a_F_regexp_count_3))
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
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
								v43 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)) = uint8(v43)
								v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v48 = int32(0)
								v51 = F_setup_regexp_matches(m, v13, v18, v37, v35-v43, v47, v48, v48, v48)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									v53 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
									m.G0 = v10 + int32(32)
									return v53
								}
							}
						}
					}
				} else {
					v35 = int32(1)
					v37 = v10 + int32(24)
					F_parse_re_flags(m, v37, v28)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)))
						if v40 == int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_regexp_count_4)
									F_errmsg(m, int32(_a_F_regexp_count_5), v10+int32(16))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_regexp_count_2), int32(1166), int32(_a_F_regexp_count_3))
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
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
							v43 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)) = uint8(v43)
							v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v48 = int32(0)
							v51 = F_setup_regexp_matches(m, v13, v18, v37, v35-v43, v47, v48, v48, v48)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
								m.G0 = v10 + int32(32)
								return v53
							}
						}
					}
				}
			}
		}
	}
}
func F_regexp_match(m *base.Module, l0 int32) int32 {
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
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = F_pg_detoast_datum_packed(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(3) <= v20 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v24 = F_pg_detoast_datum_packed(m, v23)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v27 = v24
					F_parse_re_flags(m, v8+int32(8), v27)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+12)))
						if v30 != int32(1) {
							v35 = int32(0)
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v40 = F_setup_regexp_matches(m, v11, v16, v8+int32(8), v35, v36, int32(1), v35, v35)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
								if v42 == int32(0) {
									v45 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
									v59 = int32(0)
									m.G0 = v8 + int32(16)
									return v59
								} else {
									v47 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
									v50 = F_palloc(m, v47<<(uint(int32(2))%32))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v50
										v53 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
										v54 = F_palloc(m, v53)
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v40)+24)) = v54
											v57 = F_build_regexp_match_result(m, v40)
											mBase = m.M
											v58 = m.ExcPending
											if v58 != 0 {
												return int32(0)
											} else {
												v59 = v57
												m.G0 = v8 + int32(16)
												return v59
											}
										}
									}
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_regexp_match_0)
									F_errmsg(m, int32(_a_F_regexp_match_1), v8)
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										F_errhint(m, int32(_a_F_regexp_match_2), int32(0))
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_regexp_match_3), int32(1384), int32(_a_F_regexp_match_4))
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
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
				v27 = int32(0)
				F_parse_re_flags(m, v8+int32(8), v27)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+12)))
					if v30 != int32(1) {
						v35 = int32(0)
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v40 = F_setup_regexp_matches(m, v11, v16, v8+int32(8), v35, v36, int32(1), v35, v35)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
							if v42 == int32(0) {
								v45 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
								v59 = int32(0)
								m.G0 = v8 + int32(16)
								return v59
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
								v50 = F_palloc(m, v47<<(uint(int32(2))%32))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v50
									v53 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
									v54 = F_palloc(m, v53)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v40)+24)) = v54
										v57 = F_build_regexp_match_result(m, v40)
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return int32(0)
										} else {
											v59 = v57
											m.G0 = v8 + int32(16)
											return v59
										}
									}
								}
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_regexp_match_0)
								F_errmsg(m, int32(_a_F_regexp_match_1), v8)
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int32(0)
								} else {
									F_errhint(m, int32(_a_F_regexp_match_2), int32(0))
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_regexp_match_3), int32(1384), int32(_a_F_regexp_match_4))
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
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
func F_regexp_matches(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int64
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v13 == v2 {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum_packed(m, v16)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(3) <= v21 {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v25 = F_pg_detoast_datum_packed(m, v24)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = v25
					v28 = F_init_MultiFuncCall(m, l0)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v30 = int32(_a_F_regexp_matches_0)
						v31 = *(*int32)(unsafe.Add(mBase, _c_F_regexp_matches[0]))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
						*(*int32)(unsafe.Add(mBase, _c_F_regexp_matches[0])) = v33
						v36 = v10 + int32(8)
						F_parse_re_flags(m, v36, v27)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v40 = F_pg_detoast_datum_copy(m, v39)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								v42 = int32(0)
								v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v47 = F_setup_regexp_matches(m, v40, v17, v36, v42, v43, int32(1), v42, v42)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
									v52 = F_palloc(m, v49<<(uint(int32(2))%32))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v47)+20)) = v52
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
										v56 = F_palloc(m, v55)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v47)+24)) = v56
											*(*int32)(unsafe.Add(mBase, _c_F_regexp_matches[0])) = v31
											*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v47
											v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
											v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
											v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
											v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
											if v70 < v71 {
												v73 = F_build_regexp_match_result(m, v69)
												mBase = m.M
												v74 = m.ExcPending
												if v74 != 0 {
													return int32(0)
												} else {
													v75 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
													v76 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(v69)+16)) = v75 + v76
													v79 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
													*(*int64)(unsafe.Add(mBase, uint32(v68))) = v79 + int64(1)
													v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v83)+20)) = v76
													v94 = v73
													m.G0 = v10 + int32(16)
													return v94
												}
											} else {
												F_end_MultiFuncCall(m, l0)
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return int32(0)
												} else {
													v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v88)+20)) = int32(2)
													v91 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v91)
													v94 = int32(0)
													m.G0 = v10 + int32(16)
													return v94
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
				v27 = v2
				v28 = F_init_MultiFuncCall(m, l0)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = int32(_a_F_regexp_matches_0)
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_regexp_matches[0]))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
					*(*int32)(unsafe.Add(mBase, _c_F_regexp_matches[0])) = v33
					v36 = v10 + int32(8)
					F_parse_re_flags(m, v36, v27)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v40 = F_pg_detoast_datum_copy(m, v39)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							v42 = int32(0)
							v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v47 = F_setup_regexp_matches(m, v40, v17, v36, v42, v43, int32(1), v42, v42)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
								v52 = F_palloc(m, v49<<(uint(int32(2))%32))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v47)+20)) = v52
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
									v56 = F_palloc(m, v55)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v47)+24)) = v56
										*(*int32)(unsafe.Add(mBase, _c_F_regexp_matches[0])) = v31
										*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v47
										v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
										v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
										v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
										if v70 < v71 {
											v73 = F_build_regexp_match_result(m, v69)
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return int32(0)
											} else {
												v75 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
												v76 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(v69)+16)) = v75 + v76
												v79 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
												*(*int64)(unsafe.Add(mBase, uint32(v68))) = v79 + int64(1)
												v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v83)+20)) = v76
												v94 = v73
												m.G0 = v10 + int32(16)
												return v94
											}
										} else {
											F_end_MultiFuncCall(m, l0)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v88)+20)) = int32(2)
												v91 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v91)
												v94 = int32(0)
												m.G0 = v10 + int32(16)
												return v94
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
		v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
		v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
		v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
		v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
		if v70 < v71 {
			v73 = F_build_regexp_match_result(m, v69)
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return int32(0)
			} else {
				v75 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
				v76 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v69)+16)) = v75 + v76
				v79 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
				*(*int64)(unsafe.Add(mBase, uint32(v68))) = v79 + int64(1)
				v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v83)+20)) = v76
				v94 = v73
				m.G0 = v10 + int32(16)
				return v94
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v87 = m.ExcPending
			if v87 != 0 {
				return int32(0)
			} else {
				v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v88)+20)) = int32(2)
				v91 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v91)
				v94 = int32(0)
				m.G0 = v10 + int32(16)
				return v94
			}
		}
	}
}
func F_regexp_substr(m *base.Module, l0 int32) int32 {
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
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v21 = F_pg_detoast_datum_packed(m, v20)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(5) <= v23 {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				v27 = F_pg_detoast_datum_packed(m, v26)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
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
							return int32(0)
						} else {
							v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)))
							if v60 == int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v180 = m.ExcPending
								if v180 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v183 = m.ExcPending
									if v183 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(_a_F_regexp_substr_0)
										F_errmsg(m, int32(_a_F_regexp_substr_1), v11+int32(-48))
										mBase = m.M
										v190 = m.ExcPending
										if v190 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_regexp_substr_2), int32(1955), int32(_a_F_regexp_substr_3))
											mBase = m.M
											v195 = m.ExcPending
											if v195 != 0 {
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
								v63 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)) = uint8(v63)
								v65 = int32(0)
								v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v73 = F_setup_regexp_matches(m, v16, v21, v57, v55-v63, v68, base.B2i32(v54 != v65), v65, v65)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
									if v75 < v53 {
										v77 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v77)
										v111 = v65
										m.G0 = v13 - int32(-64)
										return v111
									} else {
										v79 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
										if v79 < v54 {
											v81 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
											v111 = v65
											m.G0 = v13 - int32(-64)
											return v111
										} else {
											v83 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
											v84 = int32(0)
											v93 = v83 + (v54-base.B2i32(v54 != v84)+v79*(v53-int32(1)))<<(uint(int32(3))%32)
											v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
											if v84 <= v94 {
												v97 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
												if int32(0) <= v97 {
													v105 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
													v109 = F_DirectFunctionCall3Coll(m, int32(1477), int32(0), v105, v94+int32(1), v97-v94)
													mBase = m.M
													v110 = m.ExcPending
													if v110 != 0 {
														return int32(0)
													} else {
														v111 = v109
														m.G0 = v13 - int32(-64)
														return v111
													}
												} else {
													v101 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v101)
													v111 = v65
													m.G0 = v13 - int32(-64)
													return v111
												}
											} else {
												v101 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v101)
												v111 = v65
												m.G0 = v13 - int32(-64)
												return v111
											}
										}
									}
								}
							}
						}
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						if v37 <= int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v122 = m.ExcPending
							if v122 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v37
									*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(_a_F_regexp_substr_4)
									F_errmsg(m, int32(_a_F_regexp_substr_5), v13)
									mBase = m.M
									v131 = m.ExcPending
									if v131 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_regexp_substr_2), int32(1926), int32(_a_F_regexp_substr_3))
										mBase = m.M
										v136 = m.ExcPending
										if v136 != 0 {
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
									return int32(0)
								} else {
									v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)))
									if v60 == int32(1) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v180 = m.ExcPending
										if v180 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v183 = m.ExcPending
											if v183 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(_a_F_regexp_substr_0)
												F_errmsg(m, int32(_a_F_regexp_substr_1), v11+int32(-48))
												mBase = m.M
												v190 = m.ExcPending
												if v190 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_regexp_substr_2), int32(1955), int32(_a_F_regexp_substr_3))
													mBase = m.M
													v195 = m.ExcPending
													if v195 != 0 {
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
										v63 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)) = uint8(v63)
										v65 = int32(0)
										v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v73 = F_setup_regexp_matches(m, v16, v21, v57, v55-v63, v68, base.B2i32(v54 != v65), v65, v65)
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return int32(0)
										} else {
											v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
											if v75 < v53 {
												v77 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v77)
												v111 = v65
												m.G0 = v13 - int32(-64)
												return v111
											} else {
												v79 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
												if v79 < v54 {
													v81 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
													v111 = v65
													m.G0 = v13 - int32(-64)
													return v111
												} else {
													v83 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
													v84 = int32(0)
													v93 = v83 + (v54-base.B2i32(v54 != v84)+v79*(v53-int32(1)))<<(uint(int32(3))%32)
													v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
													if v84 <= v94 {
														v97 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
														if int32(0) <= v97 {
															v105 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
															v109 = F_DirectFunctionCall3Coll(m, int32(1477), int32(0), v105, v94+int32(1), v97-v94)
															mBase = m.M
															v110 = m.ExcPending
															if v110 != 0 {
																return int32(0)
															} else {
																v111 = v109
																m.G0 = v13 - int32(-64)
																return v111
															}
														} else {
															v101 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v101)
															v111 = v65
															m.G0 = v13 - int32(-64)
															return v111
														}
													} else {
														v101 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v101)
														v111 = v65
														m.G0 = v13 - int32(-64)
														return v111
													}
												}
											}
										}
									}
								}
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
								if v44 <= int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v140 = m.ExcPending
									if v140 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v143 = m.ExcPending
										if v143 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v44
											*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = int32(_a_F_regexp_substr_7)
											F_errmsg(m, int32(_a_F_regexp_substr_5), v11+int32(-32))
											mBase = m.M
											v151 = m.ExcPending
											if v151 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_regexp_substr_2), int32(1935), int32(_a_F_regexp_substr_3))
												mBase = m.M
												v156 = m.ExcPending
												if v156 != 0 {
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
									if base.Ui32(v41) < base.Ui32(int32(6)) {
										v53 = v44
										v54 = v2
										v55 = v37
										v57 = v11 + int32(-8)
										F_parse_re_flags(m, v57, v31)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int32(0)
										} else {
											v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)))
											if v60 == int32(1) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v180 = m.ExcPending
												if v180 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v183 = m.ExcPending
													if v183 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(_a_F_regexp_substr_0)
														F_errmsg(m, int32(_a_F_regexp_substr_1), v11+int32(-48))
														mBase = m.M
														v190 = m.ExcPending
														if v190 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_regexp_substr_2), int32(1955), int32(_a_F_regexp_substr_3))
															mBase = m.M
															v195 = m.ExcPending
															if v195 != 0 {
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
												v63 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)) = uint8(v63)
												v65 = int32(0)
												v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v73 = F_setup_regexp_matches(m, v16, v21, v57, v55-v63, v68, base.B2i32(v54 != v65), v65, v65)
												mBase = m.M
												v74 = m.ExcPending
												if v74 != 0 {
													return int32(0)
												} else {
													v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
													if v75 < v53 {
														v77 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v77)
														v111 = v65
														m.G0 = v13 - int32(-64)
														return v111
													} else {
														v79 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
														if v79 < v54 {
															v81 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
															v111 = v65
															m.G0 = v13 - int32(-64)
															return v111
														} else {
															v83 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
															v84 = int32(0)
															v93 = v83 + (v54-base.B2i32(v54 != v84)+v79*(v53-int32(1)))<<(uint(int32(3))%32)
															v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
															if v84 <= v94 {
																v97 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
																if int32(0) <= v97 {
																	v105 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
																	v109 = F_DirectFunctionCall3Coll(m, int32(1477), int32(0), v105, v94+int32(1), v97-v94)
																	mBase = m.M
																	v110 = m.ExcPending
																	if v110 != 0 {
																		return int32(0)
																	} else {
																		v111 = v109
																		m.G0 = v13 - int32(-64)
																		return v111
																	}
																} else {
																	v101 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v101)
																	v111 = v65
																	m.G0 = v13 - int32(-64)
																	return v111
																}
															} else {
																v101 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v101)
																v111 = v65
																m.G0 = v13 - int32(-64)
																return v111
															}
														}
													}
												}
											}
										}
									} else {
										v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
										if v49 < int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v160 = m.ExcPending
											if v160 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v163 = m.ExcPending
												if v163 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v49
													*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = int32(_a_F_regexp_substr_8)
													F_errmsg(m, int32(_a_F_regexp_substr_5), v11+int32(-16))
													mBase = m.M
													v171 = m.ExcPending
													if v171 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_regexp_substr_2), int32(1944), int32(_a_F_regexp_substr_3))
														mBase = m.M
														v176 = m.ExcPending
														if v176 != 0 {
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
											v53 = v44
											v54 = v49
											v55 = v37
											v57 = v11 + int32(-8)
											F_parse_re_flags(m, v57, v31)
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return int32(0)
											} else {
												v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)))
												if v60 == int32(1) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v180 = m.ExcPending
													if v180 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v183 = m.ExcPending
														if v183 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(_a_F_regexp_substr_0)
															F_errmsg(m, int32(_a_F_regexp_substr_1), v11+int32(-48))
															mBase = m.M
															v190 = m.ExcPending
															if v190 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_regexp_substr_2), int32(1955), int32(_a_F_regexp_substr_3))
																mBase = m.M
																v195 = m.ExcPending
																if v195 != 0 {
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
													v63 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)) = uint8(v63)
													v65 = int32(0)
													v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v73 = F_setup_regexp_matches(m, v16, v21, v57, v55-v63, v68, base.B2i32(v54 != v65), v65, v65)
													mBase = m.M
													v74 = m.ExcPending
													if v74 != 0 {
														return int32(0)
													} else {
														v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
														if v75 < v53 {
															v77 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v77)
															v111 = v65
															m.G0 = v13 - int32(-64)
															return v111
														} else {
															v79 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
															if v79 < v54 {
																v81 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
																v111 = v65
																m.G0 = v13 - int32(-64)
																return v111
															} else {
																v83 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
																v84 = int32(0)
																v93 = v83 + (v54-base.B2i32(v54 != v84)+v79*(v53-int32(1)))<<(uint(int32(3))%32)
																v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
																if v84 <= v94 {
																	v97 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
																	if int32(0) <= v97 {
																		v105 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
																		v109 = F_DirectFunctionCall3Coll(m, int32(1477), int32(0), v105, v94+int32(1), v97-v94)
																		mBase = m.M
																		v110 = m.ExcPending
																		if v110 != 0 {
																			return int32(0)
																		} else {
																			v111 = v109
																			m.G0 = v13 - int32(-64)
																			return v111
																		}
																	} else {
																		v101 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v101)
																		v111 = v65
																		m.G0 = v13 - int32(-64)
																		return v111
																	}
																} else {
																	v101 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v101)
																	v111 = v65
																	m.G0 = v13 - int32(-64)
																	return v111
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
						return int32(0)
					} else {
						v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)))
						if v60 == int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v180 = m.ExcPending
							if v180 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v183 = m.ExcPending
								if v183 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(_a_F_regexp_substr_0)
									F_errmsg(m, int32(_a_F_regexp_substr_1), v11+int32(-48))
									mBase = m.M
									v190 = m.ExcPending
									if v190 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_regexp_substr_2), int32(1955), int32(_a_F_regexp_substr_3))
										mBase = m.M
										v195 = m.ExcPending
										if v195 != 0 {
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
							v63 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)) = uint8(v63)
							v65 = int32(0)
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v73 = F_setup_regexp_matches(m, v16, v21, v57, v55-v63, v68, base.B2i32(v54 != v65), v65, v65)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
								if v75 < v53 {
									v77 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v77)
									v111 = v65
									m.G0 = v13 - int32(-64)
									return v111
								} else {
									v79 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
									if v79 < v54 {
										v81 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
										v111 = v65
										m.G0 = v13 - int32(-64)
										return v111
									} else {
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
										v84 = int32(0)
										v93 = v83 + (v54-base.B2i32(v54 != v84)+v79*(v53-int32(1)))<<(uint(int32(3))%32)
										v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
										if v84 <= v94 {
											v97 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
											if int32(0) <= v97 {
												v105 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
												v109 = F_DirectFunctionCall3Coll(m, int32(1477), int32(0), v105, v94+int32(1), v97-v94)
												mBase = m.M
												v110 = m.ExcPending
												if v110 != 0 {
													return int32(0)
												} else {
													v111 = v109
													m.G0 = v13 - int32(-64)
													return v111
												}
											} else {
												v101 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v101)
												v111 = v65
												m.G0 = v13 - int32(-64)
												return v111
											}
										} else {
											v101 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v101)
											v111 = v65
											m.G0 = v13 - int32(-64)
											return v111
										}
									}
								}
							}
						}
					}
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					if v37 <= int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v122 = m.ExcPending
						if v122 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v37
								*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(_a_F_regexp_substr_4)
								F_errmsg(m, int32(_a_F_regexp_substr_5), v13)
								mBase = m.M
								v131 = m.ExcPending
								if v131 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_regexp_substr_2), int32(1926), int32(_a_F_regexp_substr_3))
									mBase = m.M
									v136 = m.ExcPending
									if v136 != 0 {
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
								return int32(0)
							} else {
								v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)))
								if v60 == int32(1) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v180 = m.ExcPending
									if v180 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v183 = m.ExcPending
										if v183 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(_a_F_regexp_substr_0)
											F_errmsg(m, int32(_a_F_regexp_substr_1), v11+int32(-48))
											mBase = m.M
											v190 = m.ExcPending
											if v190 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_regexp_substr_2), int32(1955), int32(_a_F_regexp_substr_3))
												mBase = m.M
												v195 = m.ExcPending
												if v195 != 0 {
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
									v63 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)) = uint8(v63)
									v65 = int32(0)
									v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v73 = F_setup_regexp_matches(m, v16, v21, v57, v55-v63, v68, base.B2i32(v54 != v65), v65, v65)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
										if v75 < v53 {
											v77 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v77)
											v111 = v65
											m.G0 = v13 - int32(-64)
											return v111
										} else {
											v79 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
											if v79 < v54 {
												v81 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
												v111 = v65
												m.G0 = v13 - int32(-64)
												return v111
											} else {
												v83 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
												v84 = int32(0)
												v93 = v83 + (v54-base.B2i32(v54 != v84)+v79*(v53-int32(1)))<<(uint(int32(3))%32)
												v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
												if v84 <= v94 {
													v97 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
													if int32(0) <= v97 {
														v105 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
														v109 = F_DirectFunctionCall3Coll(m, int32(1477), int32(0), v105, v94+int32(1), v97-v94)
														mBase = m.M
														v110 = m.ExcPending
														if v110 != 0 {
															return int32(0)
														} else {
															v111 = v109
															m.G0 = v13 - int32(-64)
															return v111
														}
													} else {
														v101 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v101)
														v111 = v65
														m.G0 = v13 - int32(-64)
														return v111
													}
												} else {
													v101 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v101)
													v111 = v65
													m.G0 = v13 - int32(-64)
													return v111
												}
											}
										}
									}
								}
							}
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							if v44 <= int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v143 = m.ExcPending
									if v143 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v44
										*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = int32(_a_F_regexp_substr_7)
										F_errmsg(m, int32(_a_F_regexp_substr_5), v11+int32(-32))
										mBase = m.M
										v151 = m.ExcPending
										if v151 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_regexp_substr_2), int32(1935), int32(_a_F_regexp_substr_3))
											mBase = m.M
											v156 = m.ExcPending
											if v156 != 0 {
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
								if base.Ui32(v41) < base.Ui32(int32(6)) {
									v53 = v44
									v54 = v2
									v55 = v37
									v57 = v11 + int32(-8)
									F_parse_re_flags(m, v57, v31)
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int32(0)
									} else {
										v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)))
										if v60 == int32(1) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v180 = m.ExcPending
											if v180 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v183 = m.ExcPending
												if v183 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(_a_F_regexp_substr_0)
													F_errmsg(m, int32(_a_F_regexp_substr_1), v11+int32(-48))
													mBase = m.M
													v190 = m.ExcPending
													if v190 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_regexp_substr_2), int32(1955), int32(_a_F_regexp_substr_3))
														mBase = m.M
														v195 = m.ExcPending
														if v195 != 0 {
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
											v63 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)) = uint8(v63)
											v65 = int32(0)
											v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v73 = F_setup_regexp_matches(m, v16, v21, v57, v55-v63, v68, base.B2i32(v54 != v65), v65, v65)
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return int32(0)
											} else {
												v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
												if v75 < v53 {
													v77 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v77)
													v111 = v65
													m.G0 = v13 - int32(-64)
													return v111
												} else {
													v79 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
													if v79 < v54 {
														v81 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
														v111 = v65
														m.G0 = v13 - int32(-64)
														return v111
													} else {
														v83 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
														v84 = int32(0)
														v93 = v83 + (v54-base.B2i32(v54 != v84)+v79*(v53-int32(1)))<<(uint(int32(3))%32)
														v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
														if v84 <= v94 {
															v97 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
															if int32(0) <= v97 {
																v105 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
																v109 = F_DirectFunctionCall3Coll(m, int32(1477), int32(0), v105, v94+int32(1), v97-v94)
																mBase = m.M
																v110 = m.ExcPending
																if v110 != 0 {
																	return int32(0)
																} else {
																	v111 = v109
																	m.G0 = v13 - int32(-64)
																	return v111
																}
															} else {
																v101 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v101)
																v111 = v65
																m.G0 = v13 - int32(-64)
																return v111
															}
														} else {
															v101 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v101)
															v111 = v65
															m.G0 = v13 - int32(-64)
															return v111
														}
													}
												}
											}
										}
									}
								} else {
									v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
									if v49 < int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v160 = m.ExcPending
										if v160 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v163 = m.ExcPending
											if v163 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v49
												*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = int32(_a_F_regexp_substr_8)
												F_errmsg(m, int32(_a_F_regexp_substr_5), v11+int32(-16))
												mBase = m.M
												v171 = m.ExcPending
												if v171 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_regexp_substr_2), int32(1944), int32(_a_F_regexp_substr_3))
													mBase = m.M
													v176 = m.ExcPending
													if v176 != 0 {
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
										v53 = v44
										v54 = v49
										v55 = v37
										v57 = v11 + int32(-8)
										F_parse_re_flags(m, v57, v31)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int32(0)
										} else {
											v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)))
											if v60 == int32(1) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v180 = m.ExcPending
												if v180 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v183 = m.ExcPending
													if v183 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(_a_F_regexp_substr_0)
														F_errmsg(m, int32(_a_F_regexp_substr_1), v11+int32(-48))
														mBase = m.M
														v190 = m.ExcPending
														if v190 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_regexp_substr_2), int32(1955), int32(_a_F_regexp_substr_3))
															mBase = m.M
															v195 = m.ExcPending
															if v195 != 0 {
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
												v63 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)) = uint8(v63)
												v65 = int32(0)
												v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v73 = F_setup_regexp_matches(m, v16, v21, v57, v55-v63, v68, base.B2i32(v54 != v65), v65, v65)
												mBase = m.M
												v74 = m.ExcPending
												if v74 != 0 {
													return int32(0)
												} else {
													v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
													if v75 < v53 {
														v77 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v77)
														v111 = v65
														m.G0 = v13 - int32(-64)
														return v111
													} else {
														v79 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
														if v79 < v54 {
															v81 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
															v111 = v65
															m.G0 = v13 - int32(-64)
															return v111
														} else {
															v83 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
															v84 = int32(0)
															v93 = v83 + (v54-base.B2i32(v54 != v84)+v79*(v53-int32(1)))<<(uint(int32(3))%32)
															v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
															if v84 <= v94 {
																v97 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
																if int32(0) <= v97 {
																	v105 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
																	v109 = F_DirectFunctionCall3Coll(m, int32(1477), int32(0), v105, v94+int32(1), v97-v94)
																	mBase = m.M
																	v110 = m.ExcPending
																	if v110 != 0 {
																		return int32(0)
																	} else {
																		v111 = v109
																		m.G0 = v13 - int32(-64)
																		return v111
																	}
																} else {
																	v101 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v101)
																	v111 = v65
																	m.G0 = v13 - int32(-64)
																	return v111
																}
															} else {
																v101 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v101)
																v111 = v65
																m.G0 = v13 - int32(-64)
																return v111
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
