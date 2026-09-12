package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bool_accum_inv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v3 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(349639), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(492355), int32(370), int32(32441))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v4 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(349639), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(492355), int32(370), int32(32441))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v7 != 0 {
			} else {
				v8 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
				*(*int64)(unsafe.Add(mBase, uint32(v4))) = v8 - int64(1)
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v12 == int32(0) {
				} else {
					v15 = *(*int64)(unsafe.Add(mBase, uint32(v4)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v4)+8)) = v15 - int64(1)
				}
			}
			return v4
		}
	}
}
func F_bool_increment(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(base.B2i32(l1 != v4))
	return base.B2i32(l1 == v4)
}
func F_call_bool_check_hook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = int32(1)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v14 == int32(0) {
		v81 = v13
		m.G0 = v11 - int32(-64)
		return v81
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[249])) = int32(50856066)
		v21 = int32(0)
		*(*int32)(unsafe.Add(mBase, _consts[423])) = v21
		*(*int32)(unsafe.Add(mBase, _consts[251])) = v21
		*(*int32)(unsafe.Add(mBase, _consts[252])) = v21
		v29 = m.T0[v14].(func(*base.Module, int32, int32, int32) int32)(m, l1, l2, l3)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			if v29 != 0 {
				v81 = v13
				m.G0 = v11 - int32(-64)
				return v81
			} else {
				v34 = F_errstart(m, l4, int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					if v34 != 0 {
						v37 = *(*int32)(unsafe.Add(mBase, _consts[249]))
						F_errcode(m, v37)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, _consts[423]))
							if v41 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v41
								F_errmsg_internal(m, int32(204631), v9+int32(-16))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									v59 = *(*int32)(unsafe.Add(mBase, _consts[251]))
									if v59 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v59
										F_errdetail_internal(m, int32(204631), v9+int32(-48))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											v67 = *(*int32)(unsafe.Add(mBase, _consts[252]))
											if v67 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
												F_errhint(m, int32(204631), v11)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(494737), int32(6836), int32(312309))
													mBase = m.M
													v76 = m.ExcPending
													if v76 != 0 {
														return int32(0)
													} else {
														F_FlushErrorState(m)
														mBase = m.M
														v79 = m.ExcPending
														if v79 != 0 {
															return int32(0)
														} else {
															v81 = int32(0)
															m.G0 = v11 - int32(-64)
															return v81
														}
													}
												}
											} else {
												F_errfinish(m, int32(494737), int32(6836), int32(312309))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													F_FlushErrorState(m)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														v81 = int32(0)
														m.G0 = v11 - int32(-64)
														return v81
													}
												}
											}
										}
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, _consts[252]))
										if v67 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
											F_errhint(m, int32(204631), v11)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(494737), int32(6836), int32(312309))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													F_FlushErrorState(m)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														v81 = int32(0)
														m.G0 = v11 - int32(-64)
														return v81
													}
												}
											}
										} else {
											F_errfinish(m, int32(494737), int32(6836), int32(312309))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												F_FlushErrorState(m)
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return int32(0)
												} else {
													v81 = int32(0)
													m.G0 = v11 - int32(-64)
													return v81
												}
											}
										}
									}
								}
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v49
								*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v48
								F_errmsg(m, int32(483105), v9+int32(-32))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									v59 = *(*int32)(unsafe.Add(mBase, _consts[251]))
									if v59 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v59
										F_errdetail_internal(m, int32(204631), v9+int32(-48))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											v67 = *(*int32)(unsafe.Add(mBase, _consts[252]))
											if v67 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
												F_errhint(m, int32(204631), v11)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(494737), int32(6836), int32(312309))
													mBase = m.M
													v76 = m.ExcPending
													if v76 != 0 {
														return int32(0)
													} else {
														F_FlushErrorState(m)
														mBase = m.M
														v79 = m.ExcPending
														if v79 != 0 {
															return int32(0)
														} else {
															v81 = int32(0)
															m.G0 = v11 - int32(-64)
															return v81
														}
													}
												}
											} else {
												F_errfinish(m, int32(494737), int32(6836), int32(312309))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													F_FlushErrorState(m)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														v81 = int32(0)
														m.G0 = v11 - int32(-64)
														return v81
													}
												}
											}
										}
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, _consts[252]))
										if v67 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
											F_errhint(m, int32(204631), v11)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(494737), int32(6836), int32(312309))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													F_FlushErrorState(m)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														v81 = int32(0)
														m.G0 = v11 - int32(-64)
														return v81
													}
												}
											}
										} else {
											F_errfinish(m, int32(494737), int32(6836), int32(312309))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												F_FlushErrorState(m)
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return int32(0)
												} else {
													v81 = int32(0)
													m.G0 = v11 - int32(-64)
													return v81
												}
											}
										}
									}
								}
							}
						}
					} else {
						F_FlushErrorState(m)
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							v81 = int32(0)
							m.G0 = v11 - int32(-64)
							return v81
						}
					}
				}
			}
		}
	}
}
func F_parse_bool(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	v3 = F_strlen(m, l0)
	mBase = m.M
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v5 - int32(48) {
	case 0:
		v37 = int32(1)
		if v3 != v37 {
			v47 = int32(0)
			if l1 != 0 {
				v56 = v47
				v58 = v47
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v56)
				v60 = v58
			} else {
				v60 = v47
			}
			v66 = v60
		} else {
			if l1 != 0 {
				v42 = v37
				v56 = int32(0)
				v58 = v42
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v56)
				v60 = v58
				v66 = v60
			} else {
				v66 = int32(1)
			}
		}
	case 1:
		v34 = int32(1)
		if v3 != v34 {
			v47 = int32(0)
			if l1 != 0 {
				v56 = v47
				v58 = v47
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v56)
				v60 = v58
			} else {
				v60 = v47
			}
		} else {
			v51 = v34
			if l1 == int32(0) {
				v60 = v51
			} else {
				v56 = v51
				v58 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v56)
				v60 = v58
			}
		}
		v66 = v60
	default:
		v47 = int32(0)
		if l1 != 0 {
			v56 = v47
			v58 = v47
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v56)
			v60 = v58
		} else {
			v60 = v47
		}
		v66 = v60
	case 22, 54:
		v11 = F_pg_strncasecmp(m, l0, int32(357768), v3)
		mBase = m.M
		if v11 != 0 {
			v47 = int32(0)
			if l1 != 0 {
				v56 = v47
				v58 = v47
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v56)
				v60 = v58
			} else {
				v60 = v47
			}
			v66 = v60
		} else {
			if l1 != 0 {
				v42 = int32(1)
				v56 = int32(0)
				v58 = v42
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v56)
				v60 = v58
				v66 = v60
			} else {
				v66 = int32(1)
			}
		}
	case 30, 62:
		v19 = F_pg_strncasecmp(m, l0, int32(238805), v3)
		mBase = m.M
		if v19 != 0 {
			v47 = int32(0)
			if l1 != 0 {
				v56 = v47
				v58 = v47
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v56)
				v60 = v58
			} else {
				v60 = v47
			}
			v66 = v60
		} else {
			if l1 != 0 {
				v42 = int32(1)
				v56 = int32(0)
				v58 = v42
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v56)
				v60 = v58
				v66 = v60
			} else {
				v66 = int32(1)
			}
		}
	case 31, 63:
		v23 = int32(2)
		if base.Ui32(v3) <= base.Ui32(v23) {
			v26 = v23
		} else {
			v26 = v3
		}
		v27 = F_pg_strncasecmp(m, l0, int32(270646), v26)
		mBase = m.M
		if v27 == int32(0) {
			v51 = int32(1)
			if l1 == int32(0) {
				v60 = v51
			} else {
				v56 = v51
				v58 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v56)
				v60 = v58
			}
			v66 = v60
		} else {
			v31 = F_pg_strncasecmp(m, l0, int32(335658), v26)
			mBase = m.M
			if v31 != 0 {
				v47 = int32(0)
				if l1 != 0 {
					v56 = v47
					v58 = v47
					*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v56)
					v60 = v58
				} else {
					v60 = v47
				}
				v66 = v60
			} else {
				if l1 != 0 {
					v42 = int32(1)
					v56 = int32(0)
					v58 = v42
					*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v56)
					v60 = v58
					v66 = v60
				} else {
					v66 = int32(1)
				}
			}
		}
	case 36, 68:
		v9 = F_pg_strncasecmp(m, l0, int32(340937), v3)
		mBase = m.M
		if v9 != 0 {
			v47 = int32(0)
			if l1 != 0 {
				v56 = v47
				v58 = v47
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v56)
				v60 = v58
			} else {
				v60 = v47
			}
		} else {
			v51 = int32(1)
			if l1 == int32(0) {
				v60 = v51
			} else {
				v56 = v51
				v58 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v56)
				v60 = v58
			}
		}
		v66 = v60
	case 41, 73:
		v15 = F_pg_strncasecmp(m, l0, int32(155980), v3)
		mBase = m.M
		if v15 == int32(0) {
			v51 = int32(1)
			if l1 == int32(0) {
				v60 = v51
			} else {
				v56 = v51
				v58 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v56)
				v60 = v58
			}
		} else {
			v47 = int32(0)
			if l1 != 0 {
				v56 = v47
				v58 = v47
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v56)
				v60 = v58
			} else {
				v60 = v47
			}
		}
		v66 = v60
	}
	return v66
}
