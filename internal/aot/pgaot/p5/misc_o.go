package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_OutputFunctionCall(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	v3 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+13)) = v8
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v8
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l0
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+28)) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = l1
	v16 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+22)) = uint16(v16)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = m.T0[v20].(func(*base.Module, int32) int32)(m, v6+int32(4))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+20)))
		if v25 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v32
				F_errmsg_internal(m, int32(529142), v6)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(491985), int32(1143), int32(302317))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v6 + int32(32)
			return v21
		}
	}
}
func F_oidgt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return base.B2i32(base.Ui32(v3) < base.Ui32(v2))
}
func F_oidle(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return base.B2i32(base.Ui32(v2) <= base.Ui32(v3))
}
func F_oidvectorhashfast(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_DirectFunctionCall1Coll(m, int32(1588), int32(0), l0)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_operationPriority(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v5 = l0 - int32(4)
	if base.Ui32(v5) <= base.Ui32(int32(37)) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v5<<(uint(int32(2))%32))+uint32(_consts[1074])))
		v13 = v12
	} else {
		v13 = int32(6)
	}
	return v13
}
func F_overlaps_timestamp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	v11 = int32(1)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v18 == v11 {
		if v14&int32(1) == int32(0) {
			v37 = v17
			v39 = v17
			v40 = v11
			if v13&int32(1) != 0 {
				if v12&int32(1) != 0 {
					v107 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
					v117 = int32(0)
					return v117
				} else {
					v48 = F_DirectFunctionCall2Coll(m, int32(1513), int32(0), v39, v15)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						if v48 != 0 {
							v107 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
							v117 = int32(0)
							return v117
						} else {
							v68 = v15
							v69 = int32(1)
							v73 = F_DirectFunctionCall2Coll(m, int32(1514), int32(0), v39, v68)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								if v73 != 0 {
									if v40 != 0 {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
										return v117
									} else {
										v77 = F_DirectFunctionCall2Coll(m, int32(1514), int32(0), v68, v37)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											if v69&base.B2i32(v77 == int32(0)) != 0 {
												v107 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
												v117 = int32(0)
												return v117
											} else {
												return base.B2i32(v77 != int32(0))
											}
										}
									}
								} else {
									v85 = int32(1)
									if v69|v40 != v85 {
										v117 = v85
									} else {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
									}
									return v117
								}
							}
						}
					}
				}
			} else {
				if v12&int32(1) != 0 {
					v55 = F_DirectFunctionCall2Coll(m, int32(1513), int32(0), v39, v16)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						if v55 != 0 {
							v107 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
							v117 = int32(0)
							return v117
						} else {
							v68 = v16
							v69 = int32(1)
							v73 = F_DirectFunctionCall2Coll(m, int32(1514), int32(0), v39, v68)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								if v73 != 0 {
									if v40 != 0 {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
										return v117
									} else {
										v77 = F_DirectFunctionCall2Coll(m, int32(1514), int32(0), v68, v37)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											if v69&base.B2i32(v77 == int32(0)) != 0 {
												v107 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
												v117 = int32(0)
												return v117
											} else {
												return base.B2i32(v77 != int32(0))
											}
										}
									}
								} else {
									v85 = int32(1)
									if v69|v40 != v85 {
										v117 = v85
									} else {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
									}
									return v117
								}
							}
						}
					}
				} else {
					v58 = int32(0)
					v59 = int32(1513)
					v63 = F_DirectFunctionCall2Coll(m, v59, v58, v16, v15)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						if v63 != 0 {
							v65 = v15
						} else {
							v65 = v16
						}
						v66 = F_DirectFunctionCall2Coll(m, v59, v58, v39, v65)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							if v66 != 0 {
								if v63 != 0 {
									v91 = v16
								} else {
									v91 = v15
								}
								v92 = F_DirectFunctionCall2Coll(m, int32(1514), int32(0), v39, v91)
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int32(0)
								} else {
									if v40&base.B2i32(v92 == int32(0)) != 0 {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
										return v117
									} else {
										return base.B2i32(v92 != int32(0))
									}
								}
							} else {
								v68 = v65
								v69 = v58
								v73 = F_DirectFunctionCall2Coll(m, int32(1514), int32(0), v39, v68)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									if v73 != 0 {
										if v40 != 0 {
											v107 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
											v117 = int32(0)
											return v117
										} else {
											v77 = F_DirectFunctionCall2Coll(m, int32(1514), int32(0), v68, v37)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												if v69&base.B2i32(v77 == int32(0)) != 0 {
													v107 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
													v117 = int32(0)
													return v117
												} else {
													return base.B2i32(v77 != int32(0))
												}
											}
										}
									} else {
										v85 = int32(1)
										if v69|v40 != v85 {
											v117 = v85
										} else {
											v107 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
											v117 = int32(0)
										}
										return v117
									}
								}
							}
						}
					}
				}
			}
		} else {
			v107 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
			v117 = int32(0)
			return v117
		}
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v14&int32(1) != 0 {
			v37 = v17
			v39 = v25
			v40 = v11
			if v13&int32(1) != 0 {
				if v12&int32(1) != 0 {
					v107 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
					v117 = int32(0)
					return v117
				} else {
					v48 = F_DirectFunctionCall2Coll(m, int32(1513), int32(0), v39, v15)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						if v48 != 0 {
							v107 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
							v117 = int32(0)
							return v117
						} else {
							v68 = v15
							v69 = int32(1)
							v73 = F_DirectFunctionCall2Coll(m, int32(1514), int32(0), v39, v68)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								if v73 != 0 {
									if v40 != 0 {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
										return v117
									} else {
										v77 = F_DirectFunctionCall2Coll(m, int32(1514), int32(0), v68, v37)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											if v69&base.B2i32(v77 == int32(0)) != 0 {
												v107 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
												v117 = int32(0)
												return v117
											} else {
												return base.B2i32(v77 != int32(0))
											}
										}
									}
								} else {
									v85 = int32(1)
									if v69|v40 != v85 {
										v117 = v85
									} else {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
									}
									return v117
								}
							}
						}
					}
				}
			} else {
				if v12&int32(1) != 0 {
					v55 = F_DirectFunctionCall2Coll(m, int32(1513), int32(0), v39, v16)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						if v55 != 0 {
							v107 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
							v117 = int32(0)
							return v117
						} else {
							v68 = v16
							v69 = int32(1)
							v73 = F_DirectFunctionCall2Coll(m, int32(1514), int32(0), v39, v68)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								if v73 != 0 {
									if v40 != 0 {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
										return v117
									} else {
										v77 = F_DirectFunctionCall2Coll(m, int32(1514), int32(0), v68, v37)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											if v69&base.B2i32(v77 == int32(0)) != 0 {
												v107 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
												v117 = int32(0)
												return v117
											} else {
												return base.B2i32(v77 != int32(0))
											}
										}
									}
								} else {
									v85 = int32(1)
									if v69|v40 != v85 {
										v117 = v85
									} else {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
									}
									return v117
								}
							}
						}
					}
				} else {
					v58 = int32(0)
					v59 = int32(1513)
					v63 = F_DirectFunctionCall2Coll(m, v59, v58, v16, v15)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						if v63 != 0 {
							v65 = v15
						} else {
							v65 = v16
						}
						v66 = F_DirectFunctionCall2Coll(m, v59, v58, v39, v65)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							if v66 != 0 {
								if v63 != 0 {
									v91 = v16
								} else {
									v91 = v15
								}
								v92 = F_DirectFunctionCall2Coll(m, int32(1514), int32(0), v39, v91)
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int32(0)
								} else {
									if v40&base.B2i32(v92 == int32(0)) != 0 {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
										return v117
									} else {
										return base.B2i32(v92 != int32(0))
									}
								}
							} else {
								v68 = v65
								v69 = v58
								v73 = F_DirectFunctionCall2Coll(m, int32(1514), int32(0), v39, v68)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									if v73 != 0 {
										if v40 != 0 {
											v107 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
											v117 = int32(0)
											return v117
										} else {
											v77 = F_DirectFunctionCall2Coll(m, int32(1514), int32(0), v68, v37)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												if v69&base.B2i32(v77 == int32(0)) != 0 {
													v107 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
													v117 = int32(0)
													return v117
												} else {
													return base.B2i32(v77 != int32(0))
												}
											}
										}
									} else {
										v85 = int32(1)
										if v69|v40 != v85 {
											v117 = v85
										} else {
											v107 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
											v117 = int32(0)
										}
										return v117
									}
								}
							}
						}
					}
				}
			}
		} else {
			v28 = int32(0)
			v31 = F_DirectFunctionCall2Coll(m, int32(1513), v28, v25, v17)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				if v31 != 0 {
					v35 = v17
				} else {
					v35 = v25
				}
				if v31 != 0 {
					v36 = v25
				} else {
					v36 = v17
				}
				v37 = v36
				v39 = v35
				v40 = v28
				if v13&int32(1) != 0 {
					if v12&int32(1) != 0 {
						v107 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
						v117 = int32(0)
						return v117
					} else {
						v48 = F_DirectFunctionCall2Coll(m, int32(1513), int32(0), v39, v15)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							if v48 != 0 {
								v107 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
								v117 = int32(0)
								return v117
							} else {
								v68 = v15
								v69 = int32(1)
								v73 = F_DirectFunctionCall2Coll(m, int32(1514), int32(0), v39, v68)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									if v73 != 0 {
										if v40 != 0 {
											v107 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
											v117 = int32(0)
											return v117
										} else {
											v77 = F_DirectFunctionCall2Coll(m, int32(1514), int32(0), v68, v37)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												if v69&base.B2i32(v77 == int32(0)) != 0 {
													v107 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
													v117 = int32(0)
													return v117
												} else {
													return base.B2i32(v77 != int32(0))
												}
											}
										}
									} else {
										v85 = int32(1)
										if v69|v40 != v85 {
											v117 = v85
										} else {
											v107 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
											v117 = int32(0)
										}
										return v117
									}
								}
							}
						}
					}
				} else {
					if v12&int32(1) != 0 {
						v55 = F_DirectFunctionCall2Coll(m, int32(1513), int32(0), v39, v16)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							if v55 != 0 {
								v107 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
								v117 = int32(0)
								return v117
							} else {
								v68 = v16
								v69 = int32(1)
								v73 = F_DirectFunctionCall2Coll(m, int32(1514), int32(0), v39, v68)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									if v73 != 0 {
										if v40 != 0 {
											v107 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
											v117 = int32(0)
											return v117
										} else {
											v77 = F_DirectFunctionCall2Coll(m, int32(1514), int32(0), v68, v37)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												if v69&base.B2i32(v77 == int32(0)) != 0 {
													v107 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
													v117 = int32(0)
													return v117
												} else {
													return base.B2i32(v77 != int32(0))
												}
											}
										}
									} else {
										v85 = int32(1)
										if v69|v40 != v85 {
											v117 = v85
										} else {
											v107 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
											v117 = int32(0)
										}
										return v117
									}
								}
							}
						}
					} else {
						v58 = int32(0)
						v59 = int32(1513)
						v63 = F_DirectFunctionCall2Coll(m, v59, v58, v16, v15)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							if v63 != 0 {
								v65 = v15
							} else {
								v65 = v16
							}
							v66 = F_DirectFunctionCall2Coll(m, v59, v58, v39, v65)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								if v66 != 0 {
									if v63 != 0 {
										v91 = v16
									} else {
										v91 = v15
									}
									v92 = F_DirectFunctionCall2Coll(m, int32(1514), int32(0), v39, v91)
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return int32(0)
									} else {
										if v40&base.B2i32(v92 == int32(0)) != 0 {
											v107 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
											v117 = int32(0)
											return v117
										} else {
											return base.B2i32(v92 != int32(0))
										}
									}
								} else {
									v68 = v65
									v69 = v58
									v73 = F_DirectFunctionCall2Coll(m, int32(1514), int32(0), v39, v68)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										if v73 != 0 {
											if v40 != 0 {
												v107 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
												v117 = int32(0)
												return v117
											} else {
												v77 = F_DirectFunctionCall2Coll(m, int32(1514), int32(0), v68, v37)
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return int32(0)
												} else {
													if v69&base.B2i32(v77 == int32(0)) != 0 {
														v107 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
														v117 = int32(0)
														return v117
													} else {
														return base.B2i32(v77 != int32(0))
													}
												}
											}
										} else {
											v85 = int32(1)
											if v69|v40 != v85 {
												v117 = v85
											} else {
												v107 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
												v117 = int32(0)
											}
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
