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
	v32 = l0 + v16<<(uint(int32(4))%32) + v13*int32(100)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+int32(20))+91)))
	if v35 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v39 = F_strncpy(m, v32+int32(24), v23, int32(64))
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
									F_errfinish(m, int32(_a_F_findTypeAnalyzeFunction_2), int32(2267), int32(_a_F_findTypeAnalyzeFunction_3))
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
							F_errfinish(m, int32(_a_F_findTypeAnalyzeFunction_2), int32(2261), int32(_a_F_findTypeAnalyzeFunction_3))
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
									F_errfinish(m, int32(_a_F_findTypeSendFunction_2), int32(2165), int32(_a_F_findTypeSendFunction_3))
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
												F_errfinish(m, int32(_a_F_findTypeSendFunction_2), int32(2172), int32(_a_F_findTypeSendFunction_3))
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
							F_errfinish(m, int32(_a_F_findTypeSendFunction_2), int32(2159), int32(_a_F_findTypeSendFunction_3))
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
	v11 = Fn13899(m, l0, int32(_a_F_findTypeTypmodinFunction_0), int32(2199), int32(_a_F_findTypeTypmodinFunction_1), int32(_a_F_findTypeTypmodinFunction_2), int32(2193), int32(2206), int32(_a_F_findTypeTypmodinFunction_3), int32(23), int32(1263))
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
	var v13 int32
	_ = v13
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
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v13 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v13 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+22)))
			v17 = v15 + v16
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+82)))
			if v18 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						v57 = F_format_type_be(m, l0)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v57
							F_errmsg(m, int32(_a_F_getTypeOutputInfo_0), v10+int32(32))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_getTypeOutputInfo_1), int32(3061), int32(_a_F_getTypeOutputInfo_2))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
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
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+104))
				if v21 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return
					} else {
						F_errcode(m, int32(52461700))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return
						} else {
							v77 = F_format_type_be(m, l0)
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v77
								F_errmsg(m, int32(_a_F_getTypeOutputInfo_3), v10+int32(16))
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_getTypeOutputInfo_1), int32(3066), int32(_a_F_getTypeOutputInfo_2))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
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
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v21
					v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+78)))
					if v25 != 0 {
						v30 = int32(0)
					} else {
						v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+76)))
						v30 = base.B2i32(v27 == int32(_a_F_getTypeOutputInfo_4))
					}
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v30)
					F_ReleaseCatCache(m, v13)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
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
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
				F_errmsg_internal(m, int32(_a_F_getTypeOutputInfo_5), v10)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_getTypeOutputInfo_1), int32(3054), int32(_a_F_getTypeOutputInfo_2))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
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
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_get_type_io_data[0]))
	if v16 == int32(0) {
		v20 = v13 + int32(12)
		F_boot_get_type_io_data(m, l0, l2, l3, l4, l5, l6, v20, v13+int32(8))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			switch l1 {
			case 0:
				v40 = v20
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
				*(*int32)(unsafe.Add(mBase, uint32(l7))) = v41
				m.G0 = v13 + int32(16)
				return
			case 1:
				v40 = v13 + int32(8)
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
				*(*int32)(unsafe.Add(mBase, uint32(l7))) = v41
				m.G0 = v13 + int32(16)
				return
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_get_type_io_data_0), int32(0))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_get_type_io_data_1), int32(2503), int32(_a_F_get_type_io_data_2))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
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
		v44 = F_SearchSysCache1(m, int32(82), l0)
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return
		} else {
			if v44 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
					F_errmsg_internal(m, int32(_a_F_get_type_io_data_3), v13)
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_get_type_io_data_1), int32(2511), int32(_a_F_get_type_io_data_2))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
				v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+22)))
				v50 = v48 + v49
				v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+76)))
				*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v51)
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+78)))
				*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v53)
				v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+128)))
				*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v55)
				v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+83)))
				*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v57)
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
				v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+22)))
				v61 = v59 + v60
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+92))
				if v62 != 0 {
					v64 = v62
				} else {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
					v64 = v63
				}
				*(*int32)(unsafe.Add(mBase, uint32(l6))) = v64
				if base.Ui32(l1) <= base.Ui32(int32(3)) {
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v50+l1<<(uint(int32(2))%32))+100))
					*(*int32)(unsafe.Add(mBase, uint32(l7))) = v71
				} else {
				}
				F_ReleaseCatCache(m, v44)
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return
				} else {
					m.G0 = v13 + int32(16)
					return
				}
			}
		}
	}
}
func F_has_type_privilege_id(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13935(m, l0, int32(_a_F_has_type_privilege_id_0), int32(1247))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_has_type_privilege_id_id(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13936(m, l0, int32(_a_F_has_type_privilege_id_id_0), int32(1247))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_has_type_privilege_id_name(m *base.Module, l0 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn13938(m, l0, int32(_a_F_has_type_privilege_id_name_0), int32(1247), int32(_a_F_has_type_privilege_id_name_1), int32(_a_F_has_type_privilege_id_name_2), int32(_a_F_has_type_privilege_id_name_3), int32(67137668), int32(1238))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_has_type_privilege_name_id(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13937(m, l0, int32(_a_F_has_type_privilege_name_id_0), int32(1247))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_has_type_privilege_name_name(m *base.Module, l0 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn13940(m, l0, int32(_a_F_has_type_privilege_name_name_0), int32(1247), int32(_a_F_has_type_privilege_name_name_1), int32(_a_F_has_type_privilege_name_name_2), int32(_a_F_has_type_privilege_name_name_3), int32(67137668), int32(1238))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
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
			v14 = F_SearchSysCache1(m, int32(57), v9)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				if v14 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
						F_errmsg_internal(m, int32(_a_F_typeIsOfTypedTable_0), v7)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_typeIsOfTypedTable_1), int32(3392), int32(_a_F_typeIsOfTypedTable_2))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
					v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+22)))
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v18+v19)+76))
					F_ReleaseCatCache(m, v14)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						v28 = base.B2i32(v21 == l1)
						m.G0 = v7 + int32(16)
						return v28
					}
				}
			}
		} else {
			v28 = int32(0)
			m.G0 = v7 + int32(16)
			return v28
		}
	}
}
func F_type_cache_syshash(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = F_GetSysCacheHashValue(m, int32(82), v4, int32(0))
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
	v3 = Fn14033(m, l0, int32(101))
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
	v3 = Fn14033(m, l0, int32(114))
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	v4 = int32(-1)
	if l1 < int32(0) {
		v54 = v4
		return v54
	} else {
		if base.Ui32(int32(2)) <= base.Ui32(l0-int32(1042)) {
			switch l0 - int32(1560) {
			case 0, 2:
				v50 = int32(8)
				v51 = base.I32_div_s(l1+int32(7), v50)
				v54 = v51 + v50
				return v54
			case 1:
				v54 = v4
				return v54
			default:
				if l0 != int32(1700) {
					v54 = v4
					return v54
				} else {
					v32 = int32(4)
					if l1 < v32 {
						v46 = int32(-1)
					} else {
						v46 = int32(base.Ui32(int32(base.Ui32(l1-v32)>>(uint(int32(16))%32))+int32(6))>>(uint(int32(1))%32))&int32(_a_F_type_maximum_size_0) + int32(8)
					}
					return v46
				}
			}
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_type_maximum_size[0]))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
			if base.Ui32(v15) <= base.Ui32(int32(41)) {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v15*int32(28))+uint32(_c_F_type_maximum_size[1])))
				v22 = v20
			} else {
				v22 = int32(1)
			}
			v23 = int32(4)
			return v22*(l1-v23) + v23
		}
	}
}
