package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecTypeSetColNames(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v3 = int32(0)
	if l1 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v8 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = v3
	goto L4
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v16 <= v13 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	goto L1
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18+v13<<(uint(int32(2))%32))))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v24 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v44 = v13 + int32(1)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v44 < v45 {
		v13 = v44
		goto L4
	} else {
		goto L11
	}
L8:
	;
	v32 = l0 + v16<<(uint(int32(3))%32) + v13*int32(100)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+int32(28))+91)))
	if v35 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v39 = F_strncpy(m, v32+int32(32), v23, int32(64))
	mBase = m.M
	v40 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+63)) = uint8(v40)
	goto L10
L10:
	;
	goto L7
L11:
	;
	goto L5
}
func F_LookupTypeName(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_LookupTypeNameExtended(m, l0, l1, int32(0), int32(1), l2)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_LookupTypeNameOid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v13 = F_LookupTypeNameExtended(m, v3, l0, v3, int32(1), l1)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 == int32(0) {
			if l1 != 0 {
				v48 = v3
				m.G0 = v8 + int32(16)
				return v48
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = F_TypeNameToString(m, l0)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v26
							F_errmsg(m, int32(_a_F_LookupTypeNameOid_0), v8)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								F_parser_errposition(m, int32(0), v33)
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_LookupTypeNameOid_1), int32(245), int32(_a_F_LookupTypeNameOid_2))
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
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
		} else {
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+22)))
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v41+v42)))
			F_ReleaseCatCache(m, v13)
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				v48 = v44
				m.G0 = v8 + int32(16)
				return v48
			}
		}
	}
}
func F_findTypeAnalyzeFunction(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = int32(2281)
	v10 = int32(1)
	v14 = F_LookupFuncName(m, l0, v10, v6+int32(28), v10)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 != 0 {
			v18 = F_get_func_rettype(m, v14)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v18 != int32(16) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(117833860))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							v55 = F_NameListToString(m, l0)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(_a_F_findTypeAnalyzeFunction_0)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v55
								F_errmsg(m, int32(_a_F_findTypeAnalyzeFunction_1), v6+int32(16))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_findTypeAnalyzeFunction_2), int32(2311), int32(_a_F_findTypeAnalyzeFunction_3))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
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
					m.G0 = v6 + int32(32)
					return v14
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(52461700))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v37 = F_func_signature_string(m, l0, int32(1), int32(0), v6+int32(28))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v37
						F_errmsg(m, int32(_a_F_findTypeAnalyzeFunction_4), v6)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_findTypeAnalyzeFunction_2), int32(2305), int32(_a_F_findTypeAnalyzeFunction_3))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
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
func F_findTypeSendFunction(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = l1
	v9 = int32(1)
	v13 = F_LookupFuncName(m, l0, v9, v6+int32(44), v9)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 != 0 {
			v17 = F_get_func_rettype(m, v13)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				if v17 != int32(17) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(117833860))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							v80 = F_NameListToString(m, l0)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = int32(_a_F_findTypeSendFunction_0)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v80
								F_errmsg(m, int32(_a_F_findTypeSendFunction_1), v6+int32(32))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_findTypeSendFunction_2), int32(2209), int32(_a_F_findTypeSendFunction_3))
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
					}
				} else {
					v21 = F_func_volatile(m, v13)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						if v21 != int32(118) {
							m.G0 = v6 + int32(48)
							return v13
						} else {
							v27 = F_errstart(m, int32(19), int32(0))
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								if v27 == int32(0) {
									m.G0 = v6 + int32(48)
									return v13
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return int32(0)
									} else {
										v34 = F_NameListToString(m, l0)
										mBase = m.M
										v35 = m.ExcPending
										if v35 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v34
											F_errmsg(m, int32(_a_F_findTypeSendFunction_4), v6+int32(16))
											mBase = m.M
											v41 = m.ExcPending
											if v41 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_findTypeSendFunction_2), int32(2216), int32(_a_F_findTypeSendFunction_3))
												mBase = m.M
												v46 = m.ExcPending
												if v46 != 0 {
													return int32(0)
												} else {
													m.G0 = v6 + int32(48)
													return v13
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
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(52461700))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					v62 = F_func_signature_string(m, l0, int32(1), int32(0), v6+int32(44))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v62
						F_errmsg(m, int32(_a_F_findTypeSendFunction_5), v6)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_findTypeSendFunction_2), int32(2203), int32(_a_F_findTypeSendFunction_3))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
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
func F_findTypeTypmodinFunction(m *base.Module, l0 int32) int32 {
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v11 = Fn14264(m, l0, int32(_a_F_findTypeTypmodinFunction_0), int32(2243), int32(_a_F_findTypeTypmodinFunction_1), int32(_a_F_findTypeTypmodinFunction_2), int32(2237), int32(2250), int32(_a_F_findTypeTypmodinFunction_3), int32(23), int32(1263))
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		return v11
	}
}
func F_getTypeOutputInfo(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v14 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		if v14 != 0 {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+22)))
			v18 = v16 + v17
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+82)))
			if v19 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						v58 = F_format_type_be(m, l0)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v58
							F_errmsg(m, int32(_a_F_getTypeOutputInfo_0), v10+int32(32))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_getTypeOutputInfo_1), int32(3235), int32(_a_F_getTypeOutputInfo_2))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return
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
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
				if v22 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return
					} else {
						F_errcode(m, int32(52461700))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return
						} else {
							v78 = F_format_type_be(m, l0)
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v78
								F_errmsg(m, int32(_a_F_getTypeOutputInfo_3), v10+int32(16))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_getTypeOutputInfo_1), int32(3240), int32(_a_F_getTypeOutputInfo_2))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return
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
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v22
					v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+78)))
					if v26 != 0 {
						v31 = int32(0)
					} else {
						v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+76)))
						v31 = base.B2i32(v28 == int32(_a_F_getTypeOutputInfo_4))
					}
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v31)
					F_ReleaseCatCache(m, v14)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						m.G0 = v10 + int32(48)
						return
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
				F_errmsg_internal(m, int32(_a_F_getTypeOutputInfo_5), v10)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_getTypeOutputInfo_1), int32(3228), int32(_a_F_getTypeOutputInfo_2))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
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
func F_get_type_io_data(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
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
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_get_type_io_data[0]))
	if v16 == int32(0) {
		v20 = v13 + int32(12)
		F_boot_get_type_io_data(m, l0, l2, l3, l4, l5, l6, v20, v13+int32(8), v13+int32(4))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			switch l1 {
			case 0:
				v42 = v20
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
				*(*int32)(unsafe.Add(mBase, uint32(l7))) = v43
				m.G0 = v13 + int32(16)
				return
			case 1:
				v42 = v13 + int32(8)
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
				*(*int32)(unsafe.Add(mBase, uint32(l7))) = v43
				m.G0 = v13 + int32(16)
				return
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_get_type_io_data_0), int32(0))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_get_type_io_data_1), int32(2677), int32(_a_F_get_type_io_data_2))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
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
		v47 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l0))
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return
		} else {
			if v47 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
					F_errmsg_internal(m, int32(_a_F_get_type_io_data_3), v13)
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_get_type_io_data_1), int32(2685), int32(_a_F_get_type_io_data_2))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
				v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+22)))
				v53 = v51 + v52
				v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+76)))
				*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v54)
				v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+78)))
				*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v56)
				v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+128)))
				*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v58)
				v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+83)))
				*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v60)
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
				v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+22)))
				v64 = v62 + v63
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+92))
				if v65 != 0 {
					v67 = v65
				} else {
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
					v67 = v66
				}
				*(*int32)(unsafe.Add(mBase, uint32(l6))) = v67
				if base.Ui32(l1) <= base.Ui32(int32(3)) {
					v74 = *(*int32)(unsafe.Add(mBase, uint32(v53+l1<<(uint(int32(2))%32))+100))
					*(*int32)(unsafe.Add(mBase, uint32(l7))) = v74
				} else {
				}
				F_ReleaseCatCache(m, v47)
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return
				} else {
					m.G0 = v13 + int32(16)
					return
				}
			}
		}
	}
}
func F_has_type_privilege_id(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14302(m, l0, int32(_a_F_has_type_privilege_id_0), int32(1247))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_has_type_privilege_id_id(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14303(m, l0, int32(_a_F_has_type_privilege_id_id_0), int32(1247))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_has_type_privilege_id_name(m *base.Module, l0 int32) int64 {
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn14305(m, l0, int32(_a_F_has_type_privilege_id_name_0), int32(1247), int32(_a_F_has_type_privilege_id_name_1), int32(_a_F_has_type_privilege_id_name_2), int32(_a_F_has_type_privilege_id_name_3), int32(67137668), int32(1366))
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		return v9
	}
}
func F_has_type_privilege_name_id(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14304(m, l0, int32(_a_F_has_type_privilege_name_id_0), int32(1247))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_has_type_privilege_name_name(m *base.Module, l0 int32) int64 {
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn14307(m, l0, int32(_a_F_has_type_privilege_name_name_0), int32(1247), int32(_a_F_has_type_privilege_name_name_1), int32(_a_F_has_type_privilege_name_name_2), int32(_a_F_has_type_privilege_name_name_3), int32(67137668), int32(1366))
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		return v9
	}
}
func F_makeTypeNameFromOid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v4 = F_palloc0(m, int32(32))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = int32(-1)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+28)) = v8
		*(*int32)(unsafe.Add(mBase, uint32(v4)+20)) = v8
		*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(68)
		return v4
	}
}
func F_typeIsOfTypedTable(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = F_typeOrDomainTypeRelid(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 != 0 {
			v15 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(v9))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				if v15 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
						F_errmsg_internal(m, int32(_a_F_typeIsOfTypedTable_0), v7)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_typeIsOfTypedTable_1), int32(3404), int32(_a_F_typeIsOfTypedTable_2))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
					v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+22)))
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v19+v20)+76))
					F_ReleaseCatCache(m, v15)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v29 = base.B2i32(v22 == l1)
						m.G0 = v7 + int32(16)
						return v29
					}
				}
			}
		} else {
			v29 = int32(0)
			m.G0 = v7 + int32(16)
			return v29
		}
	}
}
func F_type_cache_syshash(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0))))
	v6 = F_GetSysCacheHashValue(m, int32(82), v4, int64(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_type_is_enum(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14396(m, l0, int32(101))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_type_is_range(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14396(m, l0, int32(114))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_type_maximum_size(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	v4 = int32(-1)
	if l1 < int32(0) {
		v59 = v4
		return v59
	} else {
		if base.Ui32(int32(2)) <= base.Ui32(l0-int32(1042)) {
			switch l0 - int32(1560) {
			case 0, 2:
				v55 = int32(8)
				v56 = base.I32_div_s(l1+int32(7), v55)
				v59 = v56 + v55
				return v59
			case 1:
				v59 = v4
				return v59
			default:
				if l0 != int32(1700) {
					v59 = v4
					return v59
				} else {
					v37 = int32(4)
					if l1 < v37 {
						v51 = int32(-1)
					} else {
						v51 = int32(base.Ui32(int32(base.Ui32(l1-v37)>>(uint(int32(16))%32))+int32(6))>>(uint(int32(1))%32))&int32(_a_F_type_maximum_size_0) + int32(8)
					}
					return v51
				}
			}
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_type_maximum_size[0]))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
			if base.B2i32(v15 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(v15)) != 0 {
				v27 = int32(1)
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v15*int32(28))+uint32(_c_F_type_maximum_size[1])))
				v27 = v26
			}
			v28 = int32(4)
			return v27*(l1-v28) + v28
		}
	}
}
