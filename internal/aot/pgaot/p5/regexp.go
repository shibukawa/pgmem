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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
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
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v32
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(79220)
									F_errmsg(m, int32(469396), v10)
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(476522), int32(1155), int32(83880))
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
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
							F_parse_re_flags(m, v10+int32(24), v28)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)))
								if v40 == int32(1) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(645736)
											F_errmsg(m, int32(236351), v10+int32(16))
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(476522), int32(1166), int32(83880))
												mBase = m.M
												v96 = m.ExcPending
												if v96 != 0 {
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
									v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v50 = int32(0)
									v53 = F_setup_regexp_matches(m, v13, v18, v10+int32(24), v35-v43, v49, v50, v50, v50)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
										m.G0 = v10 + int32(32)
										return v55
									}
								}
							}
						}
					} else {
						v35 = int32(1)
						F_parse_re_flags(m, v10+int32(24), v28)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)))
							if v40 == int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(645736)
										F_errmsg(m, int32(236351), v10+int32(16))
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(476522), int32(1166), int32(83880))
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
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
								v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v50 = int32(0)
								v53 = F_setup_regexp_matches(m, v13, v18, v10+int32(24), v35-v43, v49, v50, v50, v50)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
									m.G0 = v10 + int32(32)
									return v55
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
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v32
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(79220)
								F_errmsg(m, int32(469396), v10)
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(476522), int32(1155), int32(83880))
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
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
						F_parse_re_flags(m, v10+int32(24), v28)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)))
							if v40 == int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(645736)
										F_errmsg(m, int32(236351), v10+int32(16))
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(476522), int32(1166), int32(83880))
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
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
								v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v50 = int32(0)
								v53 = F_setup_regexp_matches(m, v13, v18, v10+int32(24), v35-v43, v49, v50, v50, v50)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
									m.G0 = v10 + int32(32)
									return v55
								}
							}
						}
					}
				} else {
					v35 = int32(1)
					F_parse_re_flags(m, v10+int32(24), v28)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)))
						if v40 == int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(645736)
									F_errmsg(m, int32(236351), v10+int32(16))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(476522), int32(1166), int32(83880))
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
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
							v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v50 = int32(0)
							v53 = F_setup_regexp_matches(m, v13, v18, v10+int32(24), v35-v43, v49, v50, v50, v50)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
								m.G0 = v10 + int32(32)
								return v55
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
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum_packed(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
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
					F_parse_re_flags(m, v9+int32(8), v27)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+12)))
						if v30 != int32(1) {
							v35 = int32(0)
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v40 = F_setup_regexp_matches(m, v12, v17, v9+int32(8), v35, v36, int32(1), v35, v35)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
								if v42 == int32(0) {
									v45 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
									v59 = v2
									m.G0 = v9 + int32(16)
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
												m.G0 = v9 + int32(16)
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
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(646673)
									F_errmsg(m, int32(236351), v9)
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										F_errhint(m, int32(614971), int32(0))
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(476522), int32(1384), int32(311563))
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
				v27 = v2
				F_parse_re_flags(m, v9+int32(8), v27)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+12)))
					if v30 != int32(1) {
						v35 = int32(0)
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v40 = F_setup_regexp_matches(m, v12, v17, v9+int32(8), v35, v36, int32(1), v35, v35)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
							if v42 == int32(0) {
								v45 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
								v59 = v2
								m.G0 = v9 + int32(16)
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
											m.G0 = v9 + int32(16)
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
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(646673)
								F_errmsg(m, int32(236351), v9)
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int32(0)
								} else {
									F_errhint(m, int32(614971), int32(0))
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(476522), int32(1384), int32(311563))
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
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
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v12 == v2 {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = F_pg_detoast_datum_packed(m, v15)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
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
					v26 = v24
					v27 = F_init_MultiFuncCall(m, l0)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = int32(4449520)
						v30 = *(*int32)(unsafe.Add(mBase, _consts[9]))
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
						*(*int32)(unsafe.Add(mBase, _consts[9])) = v32
						F_parse_re_flags(m, v9+int32(8), v26)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v39 = F_pg_detoast_datum_copy(m, v38)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								v43 = int32(0)
								v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v48 = F_setup_regexp_matches(m, v39, v16, v9+int32(8), v43, v44, int32(1), v43, v43)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
									v53 = F_palloc(m, v50<<(uint(int32(2))%32))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = v53
										v56 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
										v57 = F_palloc(m, v56)
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v57
											*(*int32)(unsafe.Add(mBase, _consts[9])) = v30
											*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v48
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
													m.G0 = v9 + int32(16)
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
													m.G0 = v9 + int32(16)
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
				v26 = v2
				v27 = F_init_MultiFuncCall(m, l0)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = int32(4449520)
					v30 = *(*int32)(unsafe.Add(mBase, _consts[9]))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
					*(*int32)(unsafe.Add(mBase, _consts[9])) = v32
					F_parse_re_flags(m, v9+int32(8), v26)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v39 = F_pg_detoast_datum_copy(m, v38)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							v43 = int32(0)
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v48 = F_setup_regexp_matches(m, v39, v16, v9+int32(8), v43, v44, int32(1), v43, v43)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
								v53 = F_palloc(m, v50<<(uint(int32(2))%32))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = v53
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
									v57 = F_palloc(m, v56)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v57
										*(*int32)(unsafe.Add(mBase, _consts[9])) = v30
										*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v48
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
												m.G0 = v9 + int32(16)
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
												m.G0 = v9 + int32(16)
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
				m.G0 = v9 + int32(16)
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
				m.G0 = v9 + int32(16)
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
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
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v20 = F_pg_detoast_datum_packed(m, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(5) <= v22 {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				v26 = F_pg_detoast_datum_packed(m, v25)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
					v29 = v28
					v30 = v26
					v31 = int32(1)
					if base.I32_extend16_s(v29) < int32(3) {
						v53 = v2
						v54 = v31
						v55 = v31
						F_parse_re_flags(m, v10+int32(-8), v30)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+60)))
							if v60 == int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v184 = m.ExcPending
								if v184 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v187 = m.ExcPending
									if v187 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(646207)
										F_errmsg(m, int32(236351), v10+int32(-48))
										mBase = m.M
										v194 = m.ExcPending
										if v194 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(476522), int32(1955), int32(197686))
											mBase = m.M
											v199 = m.ExcPending
											if v199 != 0 {
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
								*(*uint8)(unsafe.Add(mBase, uint32(v12)+60)) = uint8(v63)
								v65 = int32(0)
								v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v75 = F_setup_regexp_matches(m, v15, v20, v10+int32(-8), v55-v63, v70, base.B2i32(v53 != v65), v65, v65)
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
									if v77 < v54 {
										v79 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v79)
										v117 = v65
										m.G0 = v12 - int32(-64)
										return v117
									} else {
										v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
										if v81 < v53 {
											v83 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v83)
											v117 = v65
											m.G0 = v12 - int32(-64)
											return v117
										} else {
											v85 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
											v86 = int32(1)
											v90 = v53 - v86
											if base.Ui32(v90) <= base.Ui32(v53) {
												v93 = v90
											} else {
												v93 = int32(0)
											}
											v97 = v85 + (v81*(v54-v86)+v93)<<(uint(int32(3))%32)
											v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
											if int32(0) <= v98 {
												v101 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
												if int32(0) <= v101 {
													v109 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
													v113 = F_DirectFunctionCall3Coll(m, int32(1492), int32(0), v109, v98+int32(1), v101-v98)
													mBase = m.M
													v114 = m.ExcPending
													if v114 != 0 {
														return int32(0)
													} else {
														v117 = v113
														m.G0 = v12 - int32(-64)
														return v117
													}
												} else {
													v105 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
													v117 = v65
													m.G0 = v12 - int32(-64)
													return v117
												}
											} else {
												v105 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
												v117 = v65
												m.G0 = v12 - int32(-64)
												return v117
											}
										}
									}
								}
							}
						}
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						if v36 <= int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v129 = m.ExcPending
								if v129 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v36
									*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(79220)
									F_errmsg(m, int32(469396), v12)
									mBase = m.M
									v135 = m.ExcPending
									if v135 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(476522), int32(1926), int32(197686))
										mBase = m.M
										v140 = m.ExcPending
										if v140 != 0 {
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
							if v29&int32(65535) == int32(3) {
								v53 = v2
								v54 = v31
								v55 = v36
								F_parse_re_flags(m, v10+int32(-8), v30)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+60)))
									if v60 == int32(1) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v184 = m.ExcPending
										if v184 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v187 = m.ExcPending
											if v187 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(646207)
												F_errmsg(m, int32(236351), v10+int32(-48))
												mBase = m.M
												v194 = m.ExcPending
												if v194 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(476522), int32(1955), int32(197686))
													mBase = m.M
													v199 = m.ExcPending
													if v199 != 0 {
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
										*(*uint8)(unsafe.Add(mBase, uint32(v12)+60)) = uint8(v63)
										v65 = int32(0)
										v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v75 = F_setup_regexp_matches(m, v15, v20, v10+int32(-8), v55-v63, v70, base.B2i32(v53 != v65), v65, v65)
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int32(0)
										} else {
											v77 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
											if v77 < v54 {
												v79 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v79)
												v117 = v65
												m.G0 = v12 - int32(-64)
												return v117
											} else {
												v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
												if v81 < v53 {
													v83 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v83)
													v117 = v65
													m.G0 = v12 - int32(-64)
													return v117
												} else {
													v85 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
													v86 = int32(1)
													v90 = v53 - v86
													if base.Ui32(v90) <= base.Ui32(v53) {
														v93 = v90
													} else {
														v93 = int32(0)
													}
													v97 = v85 + (v81*(v54-v86)+v93)<<(uint(int32(3))%32)
													v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
													if int32(0) <= v98 {
														v101 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
														if int32(0) <= v101 {
															v109 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
															v113 = F_DirectFunctionCall3Coll(m, int32(1492), int32(0), v109, v98+int32(1), v101-v98)
															mBase = m.M
															v114 = m.ExcPending
															if v114 != 0 {
																return int32(0)
															} else {
																v117 = v113
																m.G0 = v12 - int32(-64)
																return v117
															}
														} else {
															v105 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
															v117 = v65
															m.G0 = v12 - int32(-64)
															return v117
														}
													} else {
														v105 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
														v117 = v65
														m.G0 = v12 - int32(-64)
														return v117
													}
												}
											}
										}
									}
								}
							} else {
								v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
								if v43 <= int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v144 = m.ExcPending
									if v144 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v147 = m.ExcPending
										if v147 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v43
											*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(273821)
											F_errmsg(m, int32(469396), v10+int32(-32))
											mBase = m.M
											v155 = m.ExcPending
											if v155 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(476522), int32(1935), int32(197686))
												mBase = m.M
												v160 = m.ExcPending
												if v160 != 0 {
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
									if base.Ui32(v29&int32(65535)) < base.Ui32(int32(6)) {
										v53 = v2
										v54 = v43
										v55 = v36
										F_parse_re_flags(m, v10+int32(-8), v30)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int32(0)
										} else {
											v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+60)))
											if v60 == int32(1) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v184 = m.ExcPending
												if v184 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v187 = m.ExcPending
													if v187 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(646207)
														F_errmsg(m, int32(236351), v10+int32(-48))
														mBase = m.M
														v194 = m.ExcPending
														if v194 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(476522), int32(1955), int32(197686))
															mBase = m.M
															v199 = m.ExcPending
															if v199 != 0 {
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
												*(*uint8)(unsafe.Add(mBase, uint32(v12)+60)) = uint8(v63)
												v65 = int32(0)
												v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v75 = F_setup_regexp_matches(m, v15, v20, v10+int32(-8), v55-v63, v70, base.B2i32(v53 != v65), v65, v65)
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													v77 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
													if v77 < v54 {
														v79 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v79)
														v117 = v65
														m.G0 = v12 - int32(-64)
														return v117
													} else {
														v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
														if v81 < v53 {
															v83 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v83)
															v117 = v65
															m.G0 = v12 - int32(-64)
															return v117
														} else {
															v85 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
															v86 = int32(1)
															v90 = v53 - v86
															if base.Ui32(v90) <= base.Ui32(v53) {
																v93 = v90
															} else {
																v93 = int32(0)
															}
															v97 = v85 + (v81*(v54-v86)+v93)<<(uint(int32(3))%32)
															v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
															if int32(0) <= v98 {
																v101 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
																if int32(0) <= v101 {
																	v109 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
																	v113 = F_DirectFunctionCall3Coll(m, int32(1492), int32(0), v109, v98+int32(1), v101-v98)
																	mBase = m.M
																	v114 = m.ExcPending
																	if v114 != 0 {
																		return int32(0)
																	} else {
																		v117 = v113
																		m.G0 = v12 - int32(-64)
																		return v117
																	}
																} else {
																	v105 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
																	v117 = v65
																	m.G0 = v12 - int32(-64)
																	return v117
																}
															} else {
																v105 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
																v117 = v65
																m.G0 = v12 - int32(-64)
																return v117
															}
														}
													}
												}
											}
										}
									} else {
										v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
										if v50 < int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v167 = m.ExcPending
												if v167 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v50
													*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = int32(198141)
													F_errmsg(m, int32(469396), v10+int32(-16))
													mBase = m.M
													v175 = m.ExcPending
													if v175 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(476522), int32(1944), int32(197686))
														mBase = m.M
														v180 = m.ExcPending
														if v180 != 0 {
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
											v53 = v50
											v54 = v43
											v55 = v36
											F_parse_re_flags(m, v10+int32(-8), v30)
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return int32(0)
											} else {
												v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+60)))
												if v60 == int32(1) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v184 = m.ExcPending
													if v184 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v187 = m.ExcPending
														if v187 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(646207)
															F_errmsg(m, int32(236351), v10+int32(-48))
															mBase = m.M
															v194 = m.ExcPending
															if v194 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(476522), int32(1955), int32(197686))
																mBase = m.M
																v199 = m.ExcPending
																if v199 != 0 {
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
													*(*uint8)(unsafe.Add(mBase, uint32(v12)+60)) = uint8(v63)
													v65 = int32(0)
													v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v75 = F_setup_regexp_matches(m, v15, v20, v10+int32(-8), v55-v63, v70, base.B2i32(v53 != v65), v65, v65)
													mBase = m.M
													v76 = m.ExcPending
													if v76 != 0 {
														return int32(0)
													} else {
														v77 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
														if v77 < v54 {
															v79 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v79)
															v117 = v65
															m.G0 = v12 - int32(-64)
															return v117
														} else {
															v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
															if v81 < v53 {
																v83 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v83)
																v117 = v65
																m.G0 = v12 - int32(-64)
																return v117
															} else {
																v85 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
																v86 = int32(1)
																v90 = v53 - v86
																if base.Ui32(v90) <= base.Ui32(v53) {
																	v93 = v90
																} else {
																	v93 = int32(0)
																}
																v97 = v85 + (v81*(v54-v86)+v93)<<(uint(int32(3))%32)
																v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
																if int32(0) <= v98 {
																	v101 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
																	if int32(0) <= v101 {
																		v109 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
																		v113 = F_DirectFunctionCall3Coll(m, int32(1492), int32(0), v109, v98+int32(1), v101-v98)
																		mBase = m.M
																		v114 = m.ExcPending
																		if v114 != 0 {
																			return int32(0)
																		} else {
																			v117 = v113
																			m.G0 = v12 - int32(-64)
																			return v117
																		}
																	} else {
																		v105 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
																		v117 = v65
																		m.G0 = v12 - int32(-64)
																		return v117
																	}
																} else {
																	v105 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
																	v117 = v65
																	m.G0 = v12 - int32(-64)
																	return v117
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
				v29 = v22
				v30 = v2
				v31 = int32(1)
				if base.I32_extend16_s(v29) < int32(3) {
					v53 = v2
					v54 = v31
					v55 = v31
					F_parse_re_flags(m, v10+int32(-8), v30)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+60)))
						if v60 == int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v184 = m.ExcPending
							if v184 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v187 = m.ExcPending
								if v187 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(646207)
									F_errmsg(m, int32(236351), v10+int32(-48))
									mBase = m.M
									v194 = m.ExcPending
									if v194 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(476522), int32(1955), int32(197686))
										mBase = m.M
										v199 = m.ExcPending
										if v199 != 0 {
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
							*(*uint8)(unsafe.Add(mBase, uint32(v12)+60)) = uint8(v63)
							v65 = int32(0)
							v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v75 = F_setup_regexp_matches(m, v15, v20, v10+int32(-8), v55-v63, v70, base.B2i32(v53 != v65), v65, v65)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
								if v77 < v54 {
									v79 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v79)
									v117 = v65
									m.G0 = v12 - int32(-64)
									return v117
								} else {
									v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
									if v81 < v53 {
										v83 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v83)
										v117 = v65
										m.G0 = v12 - int32(-64)
										return v117
									} else {
										v85 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
										v86 = int32(1)
										v90 = v53 - v86
										if base.Ui32(v90) <= base.Ui32(v53) {
											v93 = v90
										} else {
											v93 = int32(0)
										}
										v97 = v85 + (v81*(v54-v86)+v93)<<(uint(int32(3))%32)
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
										if int32(0) <= v98 {
											v101 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
											if int32(0) <= v101 {
												v109 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
												v113 = F_DirectFunctionCall3Coll(m, int32(1492), int32(0), v109, v98+int32(1), v101-v98)
												mBase = m.M
												v114 = m.ExcPending
												if v114 != 0 {
													return int32(0)
												} else {
													v117 = v113
													m.G0 = v12 - int32(-64)
													return v117
												}
											} else {
												v105 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
												v117 = v65
												m.G0 = v12 - int32(-64)
												return v117
											}
										} else {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
											v117 = v65
											m.G0 = v12 - int32(-64)
											return v117
										}
									}
								}
							}
						}
					}
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					if v36 <= int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v36
								*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(79220)
								F_errmsg(m, int32(469396), v12)
								mBase = m.M
								v135 = m.ExcPending
								if v135 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(476522), int32(1926), int32(197686))
									mBase = m.M
									v140 = m.ExcPending
									if v140 != 0 {
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
						if v29&int32(65535) == int32(3) {
							v53 = v2
							v54 = v31
							v55 = v36
							F_parse_re_flags(m, v10+int32(-8), v30)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+60)))
								if v60 == int32(1) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v184 = m.ExcPending
									if v184 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v187 = m.ExcPending
										if v187 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(646207)
											F_errmsg(m, int32(236351), v10+int32(-48))
											mBase = m.M
											v194 = m.ExcPending
											if v194 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(476522), int32(1955), int32(197686))
												mBase = m.M
												v199 = m.ExcPending
												if v199 != 0 {
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
									*(*uint8)(unsafe.Add(mBase, uint32(v12)+60)) = uint8(v63)
									v65 = int32(0)
									v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v75 = F_setup_regexp_matches(m, v15, v20, v10+int32(-8), v55-v63, v70, base.B2i32(v53 != v65), v65, v65)
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
										if v77 < v54 {
											v79 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v79)
											v117 = v65
											m.G0 = v12 - int32(-64)
											return v117
										} else {
											v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
											if v81 < v53 {
												v83 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v83)
												v117 = v65
												m.G0 = v12 - int32(-64)
												return v117
											} else {
												v85 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
												v86 = int32(1)
												v90 = v53 - v86
												if base.Ui32(v90) <= base.Ui32(v53) {
													v93 = v90
												} else {
													v93 = int32(0)
												}
												v97 = v85 + (v81*(v54-v86)+v93)<<(uint(int32(3))%32)
												v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
												if int32(0) <= v98 {
													v101 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
													if int32(0) <= v101 {
														v109 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
														v113 = F_DirectFunctionCall3Coll(m, int32(1492), int32(0), v109, v98+int32(1), v101-v98)
														mBase = m.M
														v114 = m.ExcPending
														if v114 != 0 {
															return int32(0)
														} else {
															v117 = v113
															m.G0 = v12 - int32(-64)
															return v117
														}
													} else {
														v105 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
														v117 = v65
														m.G0 = v12 - int32(-64)
														return v117
													}
												} else {
													v105 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
													v117 = v65
													m.G0 = v12 - int32(-64)
													return v117
												}
											}
										}
									}
								}
							}
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							if v43 <= int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v144 = m.ExcPending
								if v144 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v147 = m.ExcPending
									if v147 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v43
										*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(273821)
										F_errmsg(m, int32(469396), v10+int32(-32))
										mBase = m.M
										v155 = m.ExcPending
										if v155 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(476522), int32(1935), int32(197686))
											mBase = m.M
											v160 = m.ExcPending
											if v160 != 0 {
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
								if base.Ui32(v29&int32(65535)) < base.Ui32(int32(6)) {
									v53 = v2
									v54 = v43
									v55 = v36
									F_parse_re_flags(m, v10+int32(-8), v30)
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int32(0)
									} else {
										v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+60)))
										if v60 == int32(1) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v184 = m.ExcPending
											if v184 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v187 = m.ExcPending
												if v187 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(646207)
													F_errmsg(m, int32(236351), v10+int32(-48))
													mBase = m.M
													v194 = m.ExcPending
													if v194 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(476522), int32(1955), int32(197686))
														mBase = m.M
														v199 = m.ExcPending
														if v199 != 0 {
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
											*(*uint8)(unsafe.Add(mBase, uint32(v12)+60)) = uint8(v63)
											v65 = int32(0)
											v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v75 = F_setup_regexp_matches(m, v15, v20, v10+int32(-8), v55-v63, v70, base.B2i32(v53 != v65), v65, v65)
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												v77 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
												if v77 < v54 {
													v79 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v79)
													v117 = v65
													m.G0 = v12 - int32(-64)
													return v117
												} else {
													v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
													if v81 < v53 {
														v83 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v83)
														v117 = v65
														m.G0 = v12 - int32(-64)
														return v117
													} else {
														v85 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
														v86 = int32(1)
														v90 = v53 - v86
														if base.Ui32(v90) <= base.Ui32(v53) {
															v93 = v90
														} else {
															v93 = int32(0)
														}
														v97 = v85 + (v81*(v54-v86)+v93)<<(uint(int32(3))%32)
														v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
														if int32(0) <= v98 {
															v101 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
															if int32(0) <= v101 {
																v109 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
																v113 = F_DirectFunctionCall3Coll(m, int32(1492), int32(0), v109, v98+int32(1), v101-v98)
																mBase = m.M
																v114 = m.ExcPending
																if v114 != 0 {
																	return int32(0)
																} else {
																	v117 = v113
																	m.G0 = v12 - int32(-64)
																	return v117
																}
															} else {
																v105 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
																v117 = v65
																m.G0 = v12 - int32(-64)
																return v117
															}
														} else {
															v105 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
															v117 = v65
															m.G0 = v12 - int32(-64)
															return v117
														}
													}
												}
											}
										}
									}
								} else {
									v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
									if v50 < int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v167 = m.ExcPending
											if v167 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v50
												*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = int32(198141)
												F_errmsg(m, int32(469396), v10+int32(-16))
												mBase = m.M
												v175 = m.ExcPending
												if v175 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(476522), int32(1944), int32(197686))
													mBase = m.M
													v180 = m.ExcPending
													if v180 != 0 {
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
										v53 = v50
										v54 = v43
										v55 = v36
										F_parse_re_flags(m, v10+int32(-8), v30)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int32(0)
										} else {
											v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+60)))
											if v60 == int32(1) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v184 = m.ExcPending
												if v184 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v187 = m.ExcPending
													if v187 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(646207)
														F_errmsg(m, int32(236351), v10+int32(-48))
														mBase = m.M
														v194 = m.ExcPending
														if v194 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(476522), int32(1955), int32(197686))
															mBase = m.M
															v199 = m.ExcPending
															if v199 != 0 {
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
												*(*uint8)(unsafe.Add(mBase, uint32(v12)+60)) = uint8(v63)
												v65 = int32(0)
												v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v75 = F_setup_regexp_matches(m, v15, v20, v10+int32(-8), v55-v63, v70, base.B2i32(v53 != v65), v65, v65)
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													v77 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
													if v77 < v54 {
														v79 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v79)
														v117 = v65
														m.G0 = v12 - int32(-64)
														return v117
													} else {
														v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
														if v81 < v53 {
															v83 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v83)
															v117 = v65
															m.G0 = v12 - int32(-64)
															return v117
														} else {
															v85 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
															v86 = int32(1)
															v90 = v53 - v86
															if base.Ui32(v90) <= base.Ui32(v53) {
																v93 = v90
															} else {
																v93 = int32(0)
															}
															v97 = v85 + (v81*(v54-v86)+v93)<<(uint(int32(3))%32)
															v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
															if int32(0) <= v98 {
																v101 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
																if int32(0) <= v101 {
																	v109 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
																	v113 = F_DirectFunctionCall3Coll(m, int32(1492), int32(0), v109, v98+int32(1), v101-v98)
																	mBase = m.M
																	v114 = m.ExcPending
																	if v114 != 0 {
																		return int32(0)
																	} else {
																		v117 = v113
																		m.G0 = v12 - int32(-64)
																		return v117
																	}
																} else {
																	v105 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
																	v117 = v65
																	m.G0 = v12 - int32(-64)
																	return v117
																}
															} else {
																v105 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
																v117 = v65
																m.G0 = v12 - int32(-64)
																return v117
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
